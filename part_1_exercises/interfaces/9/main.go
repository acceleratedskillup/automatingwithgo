package main

import (
	"fmt"
	"io"
)

type Diary struct {
	Entries []string
	Current int
}

func (d *Diary) Read(p []byte) (n int, err error) {
	if d.Current >= len(d.Entries) {
		return 0, io.EOF
	}
	entry := d.Entries[d.Current]
	d.Current++
	copy(p, entry)
	return len(entry), nil
}

func main() {
	diary := &Diary{
		Entries: []string{
			"Visited 1920s New York.",
			"Met Cleopatra in ancient Egypt.",
			"Witnessed the Big Bang.",
		},
	}

	buf := make([]byte, 50)
	for {
		n, err := diary.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Println(string(buf[:n]))
	}
}
