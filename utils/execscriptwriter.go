package utils

// type ExecScriptWriter struct {
// 	Stream *Stream
// 	Length int
// }

// func NewExecScriptWriter(st *Stream) (stw *ExecScriptWriter) {
// 	stw = &ExecScriptWriter{Stream: st}
// 	return
// }

// func (stw *ExecScriptWriter) Write(p []byte) (i int, err error) {
// 	stw.Stream.Write(string(p))
// 	i = len(p)
// 	stw.Length = stw.Length + i
// 	return
// }

// func (stw *ExecScriptWriter) isUsed() (is bool) {
// 	if stw.Length > 0 {
// 		is = true
// 	}
// 	return
// }
