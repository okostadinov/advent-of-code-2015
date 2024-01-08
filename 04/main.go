package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

const input = "iwrupvqb"

func main() {
	prefix := "000000"
	i := 1

	for {
		hash := getMD5Hash(input + strconv.Itoa(i))

		if strings.HasPrefix(hash, prefix) {
			fmt.Println(i)
			break
		}
		i++
	}
}

func getMD5Hash(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}
