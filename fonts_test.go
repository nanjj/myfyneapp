package main

import "testing"

func TestFcList(t *testing.T) {
	fonts := fcList()
	if len(fonts) == 0 {
		t.Fatal(fonts)
	}
}
