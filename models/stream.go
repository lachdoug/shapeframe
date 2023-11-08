package models

import (
	"fmt"
	"path/filepath"
	"sf/utils"
	"strings"
	"time"

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

func (s *Stream) Writef(format string, a ...any) {
	_, _ = s.Write([]byte(fmt.Sprintf(format, a...)))
}

func (s *Stream) Write(p []byte) (i int, err error) {
	s.saveNewLines("output", string(p))
	return
}

func (s *Stream) Error(err error) {
	s.save("error", err.Error())
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

func (s *Stream) saveNewLines(kind string, text string) {
	nlines := strings.Split(text, "\n")
	last := len(nlines) - 1
	for i, nline := range nlines {
		if i < last {
			nline = nline + "\n"
		}
		s.saveLine(kind, nline)
	}
}

func (s *Stream) saveLine(kind string, nline string) {
	rlines := strings.Split(string(nline), "\r")
	last := len(rlines) - 1
	for i, rline := range rlines {
		if i < last {
			rline = rline + "\r"
		}
		time.Sleep(10 * time.Millisecond)
		s.save(kind, rline)
	}
}

func (s *Stream) save(kind string, text string) {
	if text == "" {
		return // Skip blank strings
	}
	m := &StreamMessage{
		Type: kind,
		Text: text,
	}
	j := utils.JsonMarshal(m)
	utils.StreamAppend(s.directory(), s.TransmissionBlock(j))
}
