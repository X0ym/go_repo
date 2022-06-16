package main

import (
	"fmt"
	"regexp"
)

func main() {
	compile, err := regexp.Compile("[0-9]+")
	if err != nil {
		return
	}
	str := "proxytest_header99"
	// Match reports whether the byte slice b
	// contains any match of the regular expression re.
	match := compile.Match([]byte(str))
	fmt.Println(match)

	re := regexp.MustCompile(`a(x*)b`)
	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("T")))
	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("$1")))
	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("$1W")))
	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("${1}W")))

	mustCompile := regexp.MustCompile(`/test([0-9]+)`)
	fmt.Printf("%s\n", mustCompile.ReplaceAll([]byte("/test9"), []byte("/test_new$1")))

	//Output:
	//
	//-T-T-
	//--xx-
	//---
	//-W-xxW-
}
