package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/spaolacci/murmur3"
)

const (
	hash128 string = "128"
	hash64  string = "64"
	base    int    = 36
)

func main() {

	value := flag.String("value", "", "Value to hash")
	hash := flag.String("hash", hash64, "Hash Type (defaults to 64)")

	flag.Parse()

	if *value == "" {
		fmt.Println("Value not supplied")
		return
	}

	var hashResult string

	if *hash == hash128 {
		h1, h2 := murmur3.Sum128([]byte(*value))
		hashResult = strconv.FormatUint(h1+h2, base)
	}

	if *hash == hash64 {
		h1 := murmur3.Sum64([]byte(*value))
		hashResult = strconv.FormatUint(h1, base)
	}

	fmt.Println(hashResult)
}
