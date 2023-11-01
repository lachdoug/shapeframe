package models

import (
	"path/filepath"
	"sf/utils"

	"github.com/google/uuid"
)

type Stream struct {
	Identifier string
}

type StreamMessage struct {
	Type string
	Text string
}

func StreamCreate() (s *Stream) {
	s = &Stream{Identifier: uuid.New().String()}
	utils.StreamMakeFile(s.directory())
	return
}

func StreamLoad(identifier string) (s *Stream) {
	s = &Stream{Identifier: identifier}
	return
}

func (s *Stream) directory() (d string) {
	d = utils.TempDir(filepath.Join("streams", s.Identifier))
	return
}

func (s *Stream) Write(p []byte) (i int, err error) {
	m := &StreamMessage{
		Type: "output",
		Text: string(p),
	}
	j := utils.JsonMarshal(m)
	utils.StreamAppend(s.directory(), s.TransmissionBlock(j))
	return
}

func (s *Stream) Error(err error) {
	m := &StreamMessage{
		Type: "error",
		Text: string(err.Error()),
	}
	j := utils.JsonMarshal(m)
	utils.StreamAppend(s.directory(), s.TransmissionBlock(j))
}

func (s *Stream) Close() {
	utils.StreamAppend(s.directory(), []byte{4})
}

func (s *Stream) Read(ch chan []byte) {
	utils.StreamTail(s.directory(), ch)
}

func (s *Stream) TransmissionBlock(j []byte) (p []byte) {
	p = append(j, byte(27))
	return
}
