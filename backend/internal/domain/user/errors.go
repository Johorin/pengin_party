package user

import (
	"fmt"

	"github.com/cockroachdb/errors"
)

// ErrUserAlreadyExists は既に同じemailのユーザーが存在する場合のエラー
type ErrUserAlreadyExists struct {
	Email string
}

func (e *ErrUserAlreadyExists) Error() string {
	return fmt.Sprintf("既に同じemailのユーザーが存在します: %s", e.Email)
}

// IsErrUserAlreadyExists はエラーがErrUserAlreadyExistsかどうかを判定する
func IsErrUserAlreadyExists(err error) bool {
	var errUserAlreadyExists *ErrUserAlreadyExists
	return errors.As(err, &errUserAlreadyExists)
}
