package main

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
)

func fcList() (fonts []string) {
	cmd := exec.Command("fc-list", ":lang=zh", "file")
	var (
		b   []byte
		err error
	)
	if b, err = cmd.Output(); err != nil {
		return
	}
	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), ": ")
		if strings.HasSuffix(line, ".ttf") {
			fonts = append(fonts, line)
		}
	}
	return
}
