基本信息
========
    - 团队名称：`1' or '1' = '1`
    - 作者：Flily Hsu
    - 项目进展：构思中


项目介绍
========
通过`pincap/tidb/parser`对SQL的解析能力，进行SQL注入对检测和防范。


背景&动机
==========
对于Web接口潜在的SQL注入进行检测、拦截和防范。可进一步扩展用于：
- 片段SQL语句的注入检测，可用于WAF等网关设备，或者应用内的前置过滤器；
- 完整SQL语句的注入检测，可用于数据库内置的SQL检测，或者应用内的数据库中间件；

项目设计
=========

针对片段SQL语句的注入检测
-----------------------
在该应用场景下，检测程序运行于WAF之类的网关设备上，或者应用本身的前置过滤器，具备能力读取到参数主体内容，但是缺乏使用的具体SQL的上下文。通过将提交的参数按照常用参数插入点进行SQL语句拼接，通过parser参数语法树之后，与预期输入值拼接的SQL产生的语法树进行对比，如果语法树不一致，则说明存在SQL注入风险。

例如，简单SQL语句
```sql
SELECT * FROM users WHERE id = ?
```

未具备相应安全能力的开发者，可能会使用如下方式构造SQL
```python
user_id = request.args.get("id")
sql = "SELECT * FROM users WHERE `id` = {};".format(user_id)
```

对于WAF设备，或者应用的前置过滤器而言，可能收到从用户发来的`user_id`的可能取值有遗下几类情况：
    - `42`，单纯整型数值，有机会成为合法id；
    - `3.1415926`，单纯浮点数数值，无机会成为合法id；
    - `lorem`，字符串，无机会拼接到SQL中；
    - `1 or 1 = 1`，简单注入类型，会改变语法树结构；
    > 以上仅为几种简单情况的说明，并非包含真实世界中的所有可能性。

大部分情况下不知道开发者会如何使用参数拼接到SQL的哪个位置，也不知道拼接位置的参数类型，因此需要对参数进行多种拼接方式的尝试，例如猜测可能的拼接位置是可能是一个整形或者字符串形，那拼接的语句可能如下
```python
user_id = request.args.get("id")
sql = "SELECT * FROM users WHERE `id` = {};".format(user_id)

user_id = request.args.get("id")
sql = "SELECT * FROM users WHERE `id` = '{}';".format(user_id)
```

我们尝试将用户输入按照以上方法拼接到SQL中，并同时假设正确情况下整形使用`13`，字符串使用`ipsum`进行拼接，那么我们可以得到如下的语法树

### 整形参数的拼接
```sql
-- input: 42
    SELECT * FROM users WHERE `id` = 42;
    SELECT * FROM users WHERE `id` = 13;
--                            ^^^^ ^ ^^
--                            |    | |
--                            |    | +---- ValueExpr(bigint): 42/13
--                            |    +------ BinaryOperationExpr: <eq>
--                            +----------- ColumnNameExpr: id

    SELECT * FROM users WHERE `id` = '42';
    SELECT * FROM users WHERE `id` = 'ipsum';
--                            ^^^^ ^ ^^^^^^^
--                            |    | |
--                            |    | +---- ValueExpr(var_string): '42'/'ipsum'
--                            |    +------ BinaryOperationExpr: <eq>
--                            +----------- ColumnNameExpr: id
```
在整形输入的情况下，两种拼接得到的SQL能够参数同样的语法树，因为我们可以判定不存在SQL注入的风险；


### 浮点型参数的拼接
```sql
-- input: 3.1415926
    SELECT * FROM users WHERE `id` = 3.1415926;
--                            ^^^^ ^ ^^^^^^^^^
--                            |    | |
--                            |    | +---- ValueExpr(decimal): 3.141592
--                            |    +------ BinaryOperationExpr: <eq>
--                            +----------- ColumnNameExpr: id
    SELECT * FROM users WHERE `id` = 13;
--                            ^^^^ ^ ^^
--                            |    | |
--                            |    | +---- ValueExpr(bigint): 13
--                            |    +------ BinaryOperationExpr: <eq>
--                            +----------- ColumnNameExpr: id

    SELECT * FROM users WHERE `id` = '3.1415926';
    SELECT * FROM users WHERE `id` = 'ipsum';
--                            ^^^^ ^ ^^^^^^^
--                            |    | |
--                            |    | +---- ValueExpr(var_string): '3.1415926'/'ipsum'
--                            |    +------ BinaryOperationExpr: <eq>
--                            +----------- ColumnNameExpr: id
```
在浮点型输入的情况下，出现两种情况：
- 在整形的拼接位置，语法树出现细微差别，虽然`pingcap/tidb/parser`解析的结果中最后一项均为`ValueExpr`，但是具体数据类型分别是`bigint`和`decimal`，但是却可以解释为无差异，不会产生注入可能；
- 在字符串形的拼接位置，语法树完全相同；

### 字符串参数的拼接
```sql
-- input: lorem
    SELECT * FROM users WHERE `id` = lorem;
--                            ^^^^ ^ ^^^^^
--                            |    | |
--                            |    | +---- ColumnNameExpr: lorem
--                            |    +------ BinaryOperationExpr: <eq>
--                            +----------- ColumnNameExpr: id
    SELECT * FROM users WHERE `id` = 13;
--                            ^^^^ ^ ^^
--                            |    | |
--                            |    | +---- ValueExpr(bigint): 13
--                            |    +------ BinaryOperationExpr: <eq>
--                            +----------- ColumnNameExpr: id

    SELECT * FROM users WHERE `id` = 'lorem';
    SELECT * FROM users WHERE `id` = 'ipsum';
--                            ^^^^ ^ ^^^^^^^
--                            |    | |
--                            |    | +---- ValueExpr(var_string): 'lorem'/'ipsum'
--                            |    +------ BinaryOperationExpr: <eq>
--                            +----------- ColumnNameExpr: id
```
在字符串形输入的情况下，出现两种情况：
- 在整形的拼接位置，语法树出现不同，最后的节点从`ValueExpr`变为`ColumnNameExpr`，但是语法树并没有插入新的语法元素，并不会出现注入的情况，该情况可以通过正则表达式的解析进行排除；
- 在字符串形的拼接位置，语法树完全相同；

### 简单注入类型
```sql
-- input: 1 or 1 = 1
    SELECT * FROM users WHERE `id` = 1 or 1 = 1;
--                            ^^^^ ^ ^ ^^ ^ ^ ^
--                            |    | | |  | | |
--                            |    | | |  | | +---- ValueExpr(bigint): 1
--                            |    | | |  | +------ BinaryOperationExpr: <eq>
--                            |    | | |  +-------- ValueExpr(bigint): 1
--                            |    | | +----------- BinaryOperationExpr: <or>
--                            |    | +------------- ColumnNameExpr: lorem
--                            |    +--------------- BinaryOperationExpr: <eq>
--                            +-------------------- ColumnNameExpr: id
    SELECT * FROM users WHERE `id` = 13;
--                            ^^^^ ^ ^^
--                            |    | |
--                            |    | +---- ValueExpr(bigint): 13
--                            |    +------ BinaryOperationExpr: <eq>
--                            +----------- ColumnNameExpr: id

    SELECT * FROM users WHERE `id` = '1 or 1 = 1';
    SELECT * FROM users WHERE `id` = 'ipsum';
--                            ^^^^ ^ ^^^^^^^
--                            |    | |
--                            |    | +---- ValueExpr(var_string): 'lorem'/'ipsum'
--                            |    +------ BinaryOperationExpr: <eq>
--                            +----------- ColumnNameExpr: id
```
在简单注入类型输入的情况下，出现两种情况：
- 在整形的拼接位置，语法树出现明显不同，增加了多个语法元素，可以判定为注入；
- 在字符串形的拼接位置，语法树完全相同；

针对完整SQL语句的注入检测
-----------------------
在该应用场景下，推测可能的场景为：
1. 检测程序运行于数据库内的检测模块，具备能力读取到完整的SQL语句；
2. 检测程序运行于应用本身的数据库查询中间件，具备能力读取到完整的SQL语句；
    1. 如果不能获取SQL语句在程序中执行的位置，能力和运行于数据库中相同，没有额外检测能力；
    2. 如果具备获取SQL语句在程序中运行位置的能力，可以对执行的SQL语句以及对应位置进行保存，后续执行查询时对比与之前查询的语法树是否发生变化，可以更准确的判断SQL注入的发生；

对于获取SQL执行位置判断前后语法树是否发生变化的情形，在不同的编程语言中可能为如下操作：
1. C语言之类的静态语言，如果可以使用宏函数改造已有的查询函数，或者是新实现的查询函数，可以通过`__FILE__`和`__LINE__`之类的预定义宏获取执行位置；
2. Python之类的动态语言，可以通过反射机制获取到函数运行到位置，从而对SQL语句出现位置和语法树解析结果进行保存和对比，该方法适合在RASP之类的侵入式的运行时检测技术中应用；

通过完整SQL语句，且没有运行位置上下文的大概可判断几种类型的注入：
1. 赘述型SQL注入；
2. 扩展语句型SQL注入；
3. 注释型SQL注入；

### 赘述型SQL注入
赘述型通过增加SQL语句的语法元素，改变判定条件的结果，制造与预期不同的查询结果。

例如查询用户名和密码
```python
username, password = get_from_request()
sql = "SELECT * FROM users WHERE user = '{}' AND pswd = '{}'".format(username, password)
```

对比口令处正常输入`ipsum`和注入输入`ipsum' OR '1' = '1`的语法树：
```sql
    SELECT * FROM users WHERE user = 'ipsum' AND pswd = 'ipsum';
--                            ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
--     +-+ *ast.BinaryOperationExpr (26)  '<and>': 
--       +-+ *ast.BinaryOperationExpr (26)  '<eq>': 
--       | +-+ *ast.ColumnNameExpr (26)  '': 
--       | | +-+ *ast.ColumnName (0)  'user': 
--       | +-+ *driver.ValueExpr (33)  'var_string(5) <ipsum>': const
--       +-+ *ast.BinaryOperationExpr (45)  '<eq>': 
--         +-+ *ast.ColumnNameExpr (45)  '': 
--         | +-+ *ast.ColumnName (0)  'pswd': 
--         +-+ *driver.ValueExpr (52)  'var_string(5) <ipsum>': const

    SELECT * FROM users WHERE user = 'ipsum' AND pswd = 'ipsum' OR '1' = '1';
--                            ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
--     +-+ *ast.BinaryOperationExpr (26)  '<or>': 
--           +-+ *ast.BinaryOperationExpr (26)  '<and>': 
--           | +-+ *ast.BinaryOperationExpr (26)  '<eq>': 
--           | | +-+ *ast.ColumnNameExpr (26)  '': 
--           | | | +-+ *ast.ColumnName (0)  'user': 
--           | | +-+ *driver.ValueExpr (33)  'var_string(5) <ipsum>': const
--           | +-+ *ast.BinaryOperationExpr (45)  '<eq>': 
--           |   +-+ *ast.ColumnNameExpr (45)  '': 
--           |   | +-+ *ast.ColumnName (0)  'pswd': 
--           |   +-+ *driver.ValueExpr (52)  'var_string(5) <ipsum>': const
--           +-+ *ast.BinaryOperationExpr (63)  '<eq>': const
--             +-+ *driver.ValueExpr (63)  'var_string(1) <1>': const
--             +-+ *driver.ValueExpr (69)  'var_string(1) <1>': const
```
通过注入，WHERE子句中的最顶层语法元素从原来的`ast.BinaryOperationExpr<and>`变成了`ast.BinaryOperationExpr<or>`，同时加入了一个恒真的判定条件`'1' = '1'`，使得其他条件完全失效。这一类型SQL注入的检测，考虑通过对SQL语句中已经出现的常数进行预先计算，找到出现恒等式并绕过SQL条件的情况，从而判定存在SQL注入的可能性。

### 扩展语句型SQL注入
该类型主要是通过`UNION`语句，获取新的查询结果，从而绕过原有的查询条件。

例如对比用户名密码
```python
username, password = get_from_request()
sql = "SELECT * FROM users WHERE `name` = '{}'".format(username)

result = db.query(sql)
if result["password"] == password:
    return "Login success"
else:
    return "Login failed"
```
此时存在用户`tom`可以使用自己的口令登录用户`jerry`账户的操作方法，输入用户名为`jerry' AND 1 = 0 UNION SELECT * FROM users WHERE name = 'tom`，口令为用户`tom`的口令，构造的SQL语句如下：
```sql
    SELECT * FROM users WHERE `name` = 'jerry';
--   +-+ *ast.SelectStmt (0)  'SELECT * FROM users WHERE `name` = 'jerry';': 
--     +-+ *ast.FieldList (0)  '': 
--     | +-+ *ast.SelectField (0)  '': 
--     +-+ *ast.TableRefsClause (0)  '': 
--     | +-+ *ast.Join (0)  '': 
--     |   +-+ *ast.TableSource (0)  '': 
--     |     +-+ *ast.TableName (0)  '': 
--     +-+ *ast.BinaryOperationExpr (26)  '<eq>': 
--       +-+ *ast.ColumnNameExpr (26)  '': 
--       | +-+ *ast.ColumnName (0)  'name': 
--       +-+ *driver.ValueExpr (35)  'var_string(5) <jerry>': const

    SELECT * FROM users WHERE `name` = 'jerry' AND 1 = 0 UNION SELECT * FROM users WHERE name = 'tom';
--                                      ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
--   +-+ *ast.SetOprStmt (0)  'SELECT * FROM users WHERE `name` = 'jerry' AND 1 = 0 UNION SELECT * FROM users WHERE name = 'tom';': 
--     +-+ *ast.SetOprSelectList (0)  '': 
--       +-+ *ast.SelectStmt (0)  '': 
--       | +-+ *ast.FieldList (0)  '': 
--       | | +-+ *ast.SelectField (0)  '': 
--       | +-+ *ast.TableRefsClause (0)  '': 
--       | | +-+ *ast.Join (0)  '': 
--       | |   +-+ *ast.TableSource (0)  '': 
--       | |     +-+ *ast.TableName (0)  '': 
--       | +-+ *ast.BinaryOperationExpr (26)  '<and>': 
--       |   +-+ *ast.BinaryOperationExpr (26)  '<eq>': 
--       |   | +-+ *ast.ColumnNameExpr (26)  '': 
--       |   | | +-+ *ast.ColumnName (0)  'name': 
--       |   | +-+ *driver.ValueExpr (35)  'var_string(5) <jerry>': const
--       |   +-+ *ast.BinaryOperationExpr (47)  '<eq>': const
--       |     +-+ *driver.ValueExpr (47)  'bigint(1) BINARY <1>': const
--       |     +-+ *driver.ValueExpr (51)  'bigint(1) BINARY <0>': const
--       +-+ *ast.SelectStmt (0)  '': 
--         +-+ *ast.FieldList (0)  '': 
--         | +-+ *ast.SelectField (0)  '': 
--         +-+ *ast.TableRefsClause (0)  '': 
--         | +-+ *ast.Join (0)  '': 
--         |   +-+ *ast.TableSource (0)  '': 
--         |     +-+ *ast.TableName (0)  '': 
--         +-+ *ast.BinaryOperationExpr (85)  '<eq>': 
--           +-+ *ast.ColumnNameExpr (85)  '': 
--           | +-+ *ast.ColumnName (0)  'name': 
--           +-+ *driver.ValueExpr (92)  'var_string(3) <tom>': const
```
注入之后，整个SQL的语法树根节点从原来的`ast.SelectStmt`变为`ast.SetOprSelectList`，由于`UNION`语句的使用较难被排除风险，需要通过进一步的语义判断来识别注入可能，包括如下途径：
1. 通过对于UNION目标的查询条件，判断是否有可能替换或者覆盖原有条件，但可能精度较差；
2. 通过寻找恒等式方法，判断`UNION`子句对前文条件的绕过，该部分可服用赘述型的检测方法；

### 注释型SQL注入
通过注入注释，使得后续的SQL语句失效，从而达到绕过的目的，常见的注释符号有`--`和`#`，其中`--`是SQL标准的注释符号，`#`是MySQL的注释符号。对于TiDB的parser，注释内容都无法被解析为ast，因此需要对注释内容进行提取和判断，单独送入检测是否包含语法元素，或者撤销注释判断是否有新语法结构生成。

```python
username, password = get_from_request()
sql = "SELECT * FROM users WHERE user = '{}' AND pswd = '{}'".format(username, password)
```

输入例如`username`为`jerry' -- `
```sql
    SELECT * FROM users WHERE user = 'lorem' AND pswd = 'ipsum';
--                            ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    SELECT * FROM users WHERE user = 'jerry' -- ' AND pswd = '123456'
--                            ^^^^^^^^^^^^^^
```
该情况下，如果通过构造SQL前的参数，能够比较好判断注入的可能，但是对于完整SQL语句，还需对于注释去除后的SQL进行检测，但是难以判断注释是否是临时注释的问题，对于新开发系统可以避免误报，对于已有系统可能较难规避。


预计开发输出成果
================
1. 一个用于片段SQL进行注入检测的原型算法，和简单应用实例，包括：
    1. 一个纯算法的函数，以及配套的测试用例情况；
    2. 一个可被注入简单的Web接口，连接数据库后端，使用sqlmap扫描在该函数的开启和关闭状态下的检测情况；
2. 一个用于完整SQL注入风险监测的原型算法，尽量覆盖几种不同类型的可能性。