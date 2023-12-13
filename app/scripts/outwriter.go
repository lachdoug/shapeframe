package scripts

import "sf/app/streams"

type OutWriter struct {
	Stream *streams.Stream
	Length int
}

func NewOutWriter(st *streams.Stream) (ow *OutWriter) {
	ow = &OutWriter{Stream: st}
	return
}

func (ow *OutWriter) Write(p []byte) (i int, err error) {
	ow.Stream.Writer.Write(p)
	i = len(p)
	ow.Length = ow.Length + i
	return
}

func (ow *OutWriter) length() (i int) {
	i = ow.Length
	return
}
