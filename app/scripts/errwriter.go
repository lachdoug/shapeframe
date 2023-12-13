package scripts

import (
	"strings"
)

type ErrWriter struct {
	Messages []string
}

func NewErrWriter() (ew *ErrWriter) {
	ew = &ErrWriter{}
	return
}

func (ew *ErrWriter) Write(p []byte) (i int, err error) {
	i = len(p)
	s := strings.TrimSpace(string(p))
	s = strings.TrimPrefix(s, "Error: ")
	ew.Messages = append(ew.Messages, s)
	return
}

func (ew *ErrWriter) Error() (s string) {
	if len(ew.Messages) > 0 {
		s = ": " + strings.Join(ew.Messages, ": ")
	}
	return
}
