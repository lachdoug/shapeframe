package streams

import (
	"bufio"
	"bytes"
	"fmt"
	fileio "io"
	"os"
	"path/filepath"
	"sf/app/dirs"
	"sf/app/errors"
	"sf/app/io"
	"sf/utils"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Stream struct {
	Identifier string
	File       string
	Writer     *Writer
	Complete   bool
}

// Construction

func StreamCreate() (st *Stream) {
	st = &Stream{Identifier: uuid.New().String()}
	st.Writer = NewWriter(st, "output")
	st.setFile()
	utils.MakeFile(st.File)
	return
}

func StreamLoad(identifier string) (st *Stream) {
	st = &Stream{Identifier: identifier}
	st.setFile()
	return
}

func (st *Stream) setFile() {
	st.File = st.file()
}

// Location

func (st *Stream) directory() (d string) {
	d = dirs.TempDir(filepath.Join("streams", st.Identifier))
	return
}

func (st *Stream) file() (f string) {
	f = filepath.Join(st.directory(), "out")
	return
}

// Writing

func (st *Stream) Write(s string) {
	st.saveNewLines("output", s)
}

func (st *Stream) Writef(format string, a ...any) {
	st.Write(fmt.Sprintf(format, a...))
}

func (st *Stream) Writeln(a ...any) {
	st.Write(fmt.Sprintln(a...))
}

func (st *Stream) Error(s string) {
	st.save("error", s)
}

func (st *Stream) Errorf(format string, a ...any) {
	st.Error(fmt.Sprintf(format, a...))
}

func (st *Stream) Errorln(a ...any) {
	st.Error(fmt.Sprintln(a...))
}

func (st *Stream) ErrorWrapf(err error, format string, a ...any) {
	err = errors.ErrorWrapf(err, format, a...)
	st.Errorf(err.Error())
}

func (st *Stream) Close() {
	utils.AppendFile(st.File, []byte{4})
}

// Styled writing

func (st *Stream) Heading(format string, a ...any) {
	st.Writef(io.LightYellowColor+format+io.ResetText+"\n", a...)
}

func (st *Stream) SubHeading(format string, a ...any) {
	st.Writef(io.LightBlueColor+format+io.ResetText+"\n", a...)
}

func (st *Stream) DotPoint(format string, a ...any) {
	st.Writef(io.GrayColor+"• "+format+io.ResetText+"\n", a...)
}

func (st *Stream) ScriptSuccess(isOutput bool) {
	suffix := ""
	if !isOutput {
		suffix = ": no output"
	}
	st.Writeln(io.GrayColor + "╰ success" + suffix + io.ResetText)
}

func (st *Stream) ScriptError(err error) {
	st.Writeln(io.RedColor + err.Error())
	st.Writeln(io.GrayColor + "╰ failed" + io.ResetText)
}

func (st *Stream) ScriptMissing() {
	st.Writeln(io.GrayColor + "╰ does not exist" + io.ResetText)
}

// Reading

func (st *Stream) Print() (err error) {
	io.Print(io.HideCursor)
	defer io.Print(io.ShowCursor)
	st.Read(io.Print)
	return
}

func (st *Stream) Read(out func(...any)) (err error) {
	ch := make(chan []byte)
	go st.tailFile(ch)
	for b := range ch {
		m := &Message{}
		utils.JsonUnmarshal(b, m)
		if m.Type == "error" {
			err = errors.Error(m.Text)
			st.Complete = true
			return
		} else {
			out(m.Text)
		}
	}
	st.Complete = true
	return
}

func (st *Stream) tailFile(ch chan []byte) {
	var f *os.File
	var err error

	if f, err = os.Open(st.File); err != nil {
		panic(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			panic(err)
		}
	}()

	r := bufio.NewReader(f)

	for {
		b, err := r.ReadBytes(27)
		if err != nil {
			if err == fileio.EOF {
				time.Sleep(100 * time.Millisecond)
			} else {
				break
			}
		}
		if bytes.Equal(b, []byte{4}) { // Check for EOT
			break
		} else if !bytes.Equal(b, []byte{}) { // Skip empty bytes
			ch <- bytes.Trim(b, string([]byte{27}))
		}
	}
	close(ch)
}

// Saving

func (st *Stream) TransmissionBlock(j []byte) (p []byte) {
	p = append(j, byte(27))
	return
}

func (st *Stream) saveNewLines(kind string, text string) {
	nlines := strings.Split(text, "\n")
	last := len(nlines) - 1
	for i, nline := range nlines {
		if i < last {
			nline = nline + "\n"
		}
		st.saveLine(kind, nline)
	}
}

func (st *Stream) saveLine(kind string, nline string) {
	rlines := strings.Split(string(nline), "\r")
	last := len(rlines) - 1
	for i, rline := range rlines {
		if i < last {
			rline = rline + "\r"
		}
		st.save(kind, rline)
	}
}

func (st *Stream) save(kind string, text string) {
	if text == "" {
		return // Skip blank strings
	}
	m := &Message{
		Type: kind,
		Text: text,
	}
	j := utils.JsonMarshal(m)
	utils.AppendFile(st.File, st.TransmissionBlock(j))
}
