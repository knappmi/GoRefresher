package customerror

import "errors"

var ErrNotFound = errors.New("not found") // TODO refine message

type ParseError struct { Line int; Msg string }
func (e *ParseError) Error() string { // TODO implement formatted error
	return "parse error" }

func Lookup(k string) (string, error) {
	// TODO implement lookup logic returning wrapped errors
	return "", ErrNotFound
}

func Parse(line string, ln int) error {
	// TODO implement line parsing & error creation
	return nil
}
