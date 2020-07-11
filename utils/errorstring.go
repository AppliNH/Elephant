package utils

type ErrorString struct {
	S string
}

func (m *ErrorString) Error() string {
	return m.S
}
