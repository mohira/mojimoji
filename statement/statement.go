package statement

import (
	"errors"
	"strings"
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
	return utf8.RuneCountInString(s.body) - CharacterLimit
}

type NgStatements []NgStatement

func NewNgStatements(line string) (NgStatements, error) {
	var result NgStatements

	for _, statement := range strings.Split(line, "。") {
		statement += "。"
		count := utf8.RuneCountInString(statement)
		if count > CharacterLimit {
			s, err := NewNgStatement(statement)
			if err != nil {
				return nil, err
			}
			result = append(result, *s)
		}
	}

	return result, nil
}
