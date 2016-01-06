package main

import (
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func main() {
	a := `abc
def
ghi`

	b := `abc
def
go-
ghi`

	lineDiff(a, b)
}

func lineDiff(src1, src2 string) []diffmatchpatch.Diff {
	dmp := diffmatchpatch.New()
	a, b, c := dmp.DiffLinesToChars(src1, src2)
	diffs := dmp.DiffMain(a, b, false)
	result := dmp.DiffCharsToLines(diffs, c)
	fmt.Println(result)
	return result
}
