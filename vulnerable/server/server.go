package server

import (
	"database/sql"
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func CallerSignature() string {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		return fmt.Sprintf("%s:%d", file, line)
	}

	return ""
}

type SQLInjectionServer struct {
	Database *sql.DB
}

func (s *SQLInjectionServer) Init() error {
	_, err := s.Database.Exec(
		"CREATE TABLE IF NOT EXISTS users (" +
			"id BIGINT NOT NULL AUTO_INCREMENT, " +
			"name CHAR(64), " +
			"password CHAR(64), " +
			"age INT, " +
			"PRIMARY KEY (id)" +
			")",
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *SQLInjectionServer) RunSQL(base string, args ...interface{}) (*sql.Rows, error) {
	caller := CallerSignature()
	log.Printf("caller on %s", caller)

	sql := fmt.Sprintf(base, args...)
	log.Printf("SQL: %s", sql)
	return s.Database.Query(sql)
}

type User struct {
	ID       int
	Name     string
	Password string
	Age      int
}

type GetUserByID struct {
	ID string `uri:"id"`
}

func (s *SQLInjectionServer) GetUser(c *gin.Context) (*Response, error) {
	request := &GetUserByID{}
	err := c.BindUri(request)
	if err != nil {
		return nil, err
	}

	rows, err := s.RunSQL("SELECT * FROM users WHERE id = %s", request.ID)
	if err != nil {
		return nil, err
	}

	user := &User{}
	found := false
	for !found && rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Password, &user.Age)
		if err != nil {
			return nil, err
		}

		found = true
	}

	if !found {
		return NotFound(request.ID)
	}

	return Success(user)
}

func (s *SQLInjectionServer) AddUser(c *gin.Context) (*Response, error) {
	user := &User{}
	err := c.BindJSON(user)
	if err != nil {
		return nil, err
	}

	_, err = s.RunSQL("INSERT INTO users (name, password, age) VALUES ('%s', '%s', %d)",
		user.Name, user.Password, user.Age)
	if err != nil {
		return nil, err
	}

	return Success(user)
}

func (s *SQLInjectionServer) DeleteUser(c *gin.Context) (*Response, error) {
	request := &GetUserByID{}
	err := c.BindUri(request)
	if err != nil {
		return nil, err
	}

	_, err = s.RunSQL("DELETE FROM users WHERE id = %s", request.ID)
	if err != nil {
		return nil, err
	}

	return Success(nil)
}

func (s *SQLInjectionServer) ListUsers(c *gin.Context) (*Response, error) {
	rows, err := s.RunSQL("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	users := make([]*User, 0)
	for rows.Next() {
		user := &User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Password, &user.Age)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return Success(users)
}

func (s *SQLInjectionServer) UserQuery(c *gin.Context) (*Response, error) {
	base := ("SELECT * FROM users")
	conditions := []string{}

	if id := c.Query("id"); len(id) > 0 {
		conditions = append(conditions, fmt.Sprintf("id = %s", id))
	}

	if name := c.Query("name"); len(name) > 0 {
		conditions = append(conditions, fmt.Sprintf("name = '%s'", name))
	}

	if password := c.Query("password"); len(password) > 0 {
		conditions = append(conditions, fmt.Sprintf("password = '%s'", password))
	}

	if age := c.Query("age"); len(age) > 0 {
		conditions = append(conditions, fmt.Sprintf("age = %s", age))
	}

	if len(conditions) > 0 {
		base += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := s.RunSQL(base)
	if err != nil {
		return nil, err
	}

	users := make([]*User, 0)
	for rows.Next() {
		user := &User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Password, &user.Age)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return Success(users)
}

func (s *SQLInjectionServer) Serve(port int) error {
	engine := gin.Default()
	engine.Use(DefaultHandlerFatal)
	engine.NoRoute(DefaultHandler404)

	engine.GET("/user", MakeHandler(s.ListUsers))
	engine.POST("/user", MakeHandler(s.AddUser))
	engine.GET("/user/:id", MakeHandler(s.GetUser))
	engine.DELETE("/user/:id", MakeHandler(s.DeleteUser))

	engine.GET("/user-query", MakeHandler(s.UserQuery))

	address := fmt.Sprintf(":%d", port)
	fmt.Printf("server is listening on %s\n", address)
	return engine.Run(address)
}

func NewServer(dsn string) (*SQLInjectionServer, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	server := &SQLInjectionServer{
		Database: db,
	}

	return server, nil
}
