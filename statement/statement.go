package statement

import (
	"errors"
	"strings"
	"unicode"
	"unicode/utf8"
)

const CharacterLimit int = 35

type NgStatement struct {
	body string
}

func NewNgStatement(body string) (*NgStatement, error) {
	if utf8.RuneCountInString(body) <= CharacterLimit {
		return nil, errors.New("文字数制限の違反をしていません")
	}
	return &NgStatement{body: body}, nil
}

func (s *NgStatement) String() string {
	return s.body
}

func (s *NgStatement) OverCount() int {
	return Count(s.body) - CharacterLimit
}

type NgStatements []NgStatement

// Count 文字数を数える
// ただし、英数字やスペースは0.5文字換算、小数点は切り上げる。
func Count(s string) int {
	count := 0

	for _, c := range s {
		isSymbol := c == '(' || c == ')' || c == '_'
		if unicode.IsDigit(c) ||
			unicode.IsLower(c) ||
			unicode.IsUpper(c) ||
			unicode.IsSpace(c) ||
			isSymbol {
			count++
		}
	}

	return utf8.RuneCountInString(s) - count/2
}

func NewNgStatements(line string) (NgStatements, error) {
	var result NgStatements

	for _, statement := range strings.Split(line, "。") {
		statement += "。"

		if Count(statement) > CharacterLimit {
			s, err := NewNgStatement(statement)
			if err != nil {
				return nil, err
			}
			result = append(result, *s)
		}
	}

	return result, nil
}
