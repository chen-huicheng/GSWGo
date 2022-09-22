package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	rawDecodedText, err := base64.StdEncoding.DecodeString("aGVsbG8gZnJvbSBnb3NhbXBsZXMuZGV2IGJhc2U2NCBlbmNvZGluZyBleGFtcGxlIQ==")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decoded text: %s\n", rawDecodedText)
}
