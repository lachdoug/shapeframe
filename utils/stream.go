package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Stream struct {
	Identifier string
	File       string
}

type StreamMessage struct {
	Type string
	Text string
}

func StreamCreate() (s *Stream) {
	s = &Stream{Identifier: uuid.New().String()}
	s.setFile()
	MakeFile(s.File)
	return
}

func StreamLoad(identifier string) (s *Stream) {
	s = &Stream{Identifier: identifier}
	s.setFile()
	return
}

func (s *Stream) setFile() {
	s.File = s.file()
}

func (s *Stream) directory() (d string) {
	d = TempDir(filepath.Join("streams", s.Identifier))
	return
}

func (s *Stream) file() (f string) {
	f = filepath.Join(s.directory(), "out")
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
	AppendFile(s.File, []byte{4})
}

func (s *Stream) ReadOut(Out *os.File, Err *os.File) (err error) {
	ch := make(chan []byte)
	go TailFile(s.File, ch)
	for b := range ch {
		m := &StreamMessage{}
		JsonUnmarshal(b, m)
		if m.Type == "error" {
			err = fmt.Errorf(m.Text)
			return
		}
		if _, err = fmt.Fprint(Out, m.Text); err != nil {
			panic(err)
		}
	}
	return
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
	j := JsonMarshal(m)
	AppendFile(s.File, s.TransmissionBlock(j))
}
