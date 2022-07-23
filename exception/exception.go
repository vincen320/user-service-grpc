package exception

import "errors"

var (
	ErrParseId = errors.New("unable to parse and get ID")
	ErrDecode  = errors.New("errors when trying decode data")
)

//kalau dibutuhkan buat la versi structny
