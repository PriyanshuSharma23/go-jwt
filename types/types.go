package types

type ErrorStr string

func (e ErrorStr) Error() string {
	return string(e)
}
