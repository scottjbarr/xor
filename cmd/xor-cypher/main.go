package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/scottjbarr/xor"
)

func main() {
	key := []byte(os.Getenv("KEY"))
	if len(key) == 0 {
		panic("KEY environment variable not set")
	}

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	b := xor.EncryptDecrypt(data, key)
	fmt.Printf("%s", b)
}
