package utils

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"time"
)

func TailFile(filePath string, ch chan []byte) {
	f, err := os.Open(filePath)
	checkErr(err)

	defer func() {
		err := f.Close()
		checkErr(err)
	}()

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
		if bytes.Equal(b, []byte{4}) { // Check for EOT
			break
		} else if !bytes.Equal(b, []byte{}) { // Skip empty bytes
			ch <- bytes.Trim(b, string([]byte{27}))
		}
	}
	close(ch)
}
