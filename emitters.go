package main

import (
	"bytes"
	"os/exec"
	"time"
)

/*
// Periodically generate some data with the 'date' command
*/
func broadcaster() {
	var buf bytes.Buffer

	for {
		cmd := exec.Command("date")
		cmd.Stdout = &buf
		cmd.Stderr = &buf
		cmd.Run()

		m := NewMessage(buf.String())
		h.broadcast <- m
		buf.Reset()

		time.Sleep(2 * time.Second)
	}
}
