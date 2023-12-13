package streams

type Writer struct {
	Stream *Stream
	Kind   string
}

// Stream writer

func NewWriter(st *Stream, kind string) (w *Writer) {
	w = &Writer{Stream: st, Kind: kind}
	return
}

func (w *Writer) Write(p []byte) (i int, err error) {
	w.Stream.saveNewLines(w.Kind, string(p))
	i = len(p)
	return
}
