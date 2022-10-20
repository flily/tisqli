package syntax

import (
	"regexp"
	"strconv"
	"strings"
)

type ParserError struct {
	Line   int
	Column int
	SQL    string
	Err    error
}

func parseInt(str string) int {
	if str == "" {
		return 0
	}

	result, _ := strconv.Atoi(str)
	return result
}

var regexParserError = regexp.MustCompile(`line (\d+) column (\d+)`)

func NewParserError(sql string, err error) error {
	results := regexParserError.FindStringSubmatch(err.Error())
	if results == nil {
		return nil
	}

	return &ParserError{
		Line:   parseInt(results[1]),
		Column: parseInt(results[2]),
		SQL:    sql,
		Err:    err,
	}
}

func (e *ParserError) Hint() string {
	lines := strings.Split(e.SQL, "\n")
	results := make([]string, 0, len(lines)+3)

	for i, line := range lines {
		results = append(results, line)
		if i+1 == e.Line {
			results = append(results,
				strings.Repeat(" ", e.Column-1)+"^",
				strings.Repeat(" ", e.Column-1)+"|",
				strings.Repeat(" ", e.Column-1)+"+-- "+e.Err.Error(),
			)
		}
	}

	return strings.Join(results, "\n")
}

func (e *ParserError) Error() string {
	return e.Hint()
}

func (e *ParserError) Unwrap() error {
	return e.Err
}
