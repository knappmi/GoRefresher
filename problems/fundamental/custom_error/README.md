# Custom Error

Implement sentinel, wrapped, and typed errors. Use errors.Is / errors.As.

Components:
- var ErrNotFound
- type ParseError struct { Line int; Msg string }
- func (e *ParseError) Error() string
- Function that returns wrapped errors with fmt.Errorf("...: %w", err)

Test detection using errors.Is / errors.As.
