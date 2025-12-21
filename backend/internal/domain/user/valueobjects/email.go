package user

import (
	"regexp"
	"strings"

	"github.com/cockroachdb/errors"
)

type Email struct {
	value string
}

const (
	EmailInitialValue = ""
)

// よくあるエラー定義
var (
	ErrInvalidEmail = errors.New("メールアドレスtの形式が不正です")
	ErrEmailTooLong = errors.New("メールアドレスが長すぎます")
)

func NewEmail(email string) (Email, error) {
	if email == EmailInitialValue {
		return Email{value: EmailInitialValue}, nil
	}

	// トリミングして小文字に変換
	trimmed := strings.TrimSpace(strings.ToLower((email)))

	if !isValidEmailFormat(trimmed) {
		return Email{}, ErrInvalidEmail
	}
	
	return Email{value: trimmed}, nil
}

func (e Email) Value() string {
	return e.value
}

// メールアドレスのフォーマット検証
func isValidEmailFormat(email string) bool {
	pattern := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	// patternが静的なのでerr判定なし
	matched, _ := regexp.MatchString(pattern, email)

	return matched
}