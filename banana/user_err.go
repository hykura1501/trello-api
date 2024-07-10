package banana

import "errors"

var (
	ErrUserConflict = errors.New("user already exists")
)
