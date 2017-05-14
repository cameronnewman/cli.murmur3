package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/spaolacci/murmur3"
)

func main() {

	key := flag.String("key", "foo", "api Key")

	flag.Parse()
	hash := strconv.FormatUint(murmur3.Sum64([]byte(*key)), 16)

	fmt.Println("Key: " + *key)
	fmt.Println("Hash: " + hash)

}
