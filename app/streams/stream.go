package streams

import (
	"fmt"
	"path/filepath"
	"sf/app/errors"
	"sf/app/io"
	"sf/utils"
	"strings"

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
	d = utils.TempDir(filepath.Join("streams", st.Identifier))
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
	go utils.TailFile(st.File, ch)
	for b := range ch {
		m := &Message{}
		utils.JsonUnmarshal(b, m)
		if m.Type == "error" {
			err = errors.Error(m.Text)
			return
		} else {
			out(m.Text)
		}
	}
	st.Complete = true
	return
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
