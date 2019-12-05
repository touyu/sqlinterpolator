package sqlinterpolator

import "errors"

var (
	ErrInvalidArgsCount = errors.New("Invalid args count")
	ErrInvalidBytes     = errors.New("Invalid []byte")
)
