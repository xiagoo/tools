package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	tc := flag.String("u2s", "", "Unicode2String")
	flag.Usage = usage
	flag.Parse()
	if len(*tc) > 0 {
		Unicode2String(*tc)
	}
}

func Unicode2String(tc string) {
	textQuoted := strconv.QuoteToASCII(tc)
	textUnquoted := textQuoted[1 : len(textQuoted)-1]
	unicode := strings.Split(textUnquoted, "u")
	var s string
	for _, v := range unicode {
		if len(v) < 1 {
			continue
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			panic(err)
		}
		s += fmt.Sprintf("%c", temp)
	}
	fmt.Println(s)
}

func usage() {
	fmt.Fprintf(os.Stderr, `transcoding versin: 1.0
Options:
`)
	flag.PrintDefaults()
}
