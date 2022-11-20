package main

import "os"

func smaller(files []string) (small string) {
	sizeof := func(filename string) int64 {
		if fi, err := os.Stat(filename); err != nil {
			return 0
		} else {
			return fi.Size()
		}
	}
	l := len(files)
	switch l {
	case 0:
		return
	case 1:
		return files[0]
	}

	min := sizeof(files[0])
	for i := 1; i < l; i++ {
		size := sizeof(files[i])
		if size < min {
			min = size
			small = files[i]
		}
	}
	return
}

func init() {
	os.Setenv("FYNE_FONT", smaller(fcList()))
}
