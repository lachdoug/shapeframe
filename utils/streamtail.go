package utils

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"time"
)

func StreamTail(dirPath string, ch chan []byte) {
	filePath := filepath.Join(dirPath, "out")

	f, err := os.Open(filePath)
	checkErr(err)

	defer f.Close()

	r := bufio.NewReader(f)

	for {
		b, err := r.ReadBytes(27)
		if err != nil {
			if err == io.EOF {
				time.Sleep(100 * time.Millisecond)
			} else {
				break
			}
		}
		// Check for EOT
		if bytes.Equal(b, []byte{4}) {
			break
		} else if !bytes.Equal(b, []byte{}) {
			ch <- bytes.Trim(b, string([]byte{27}))
		}
	}
	close(ch)
}
