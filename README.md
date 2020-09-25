# gorle
[Run length encoder/decoder](https://en.wikipedia.org/wiki/Run-length_encoding) implementation in Golang.

### Usage

```
package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/jasonh-ca/gorle"
)

func main() {
	value,err := hex.DecodeString("00000000000000000000000000000000000000000000000000800000000000000000000000010000000800000000000002010000200000000000004000000000000000000000000000000000002000000000008000000000000000000000000000000000080005000000000000000000000000000000000002000001000000000000000000000000100000000000000000000010000000000200000000000040000002000000000000000000000000000010200000000000000000020000000000000008000000000000000000800000000000000000000000000000000000000000000000010400000000000000000000000000400000000000080020000000")
	if err != nil {
		fmt.Print("err = ", err)
		return
	}
	encoded := gorle.Encode(value)
	fmt.Printf("original len=%d, encoded len=%d, result: %s\n", len(value), len(encoded), hex.EncodeToString(encoded))

	decoded := gorle.Decode(encoded)
	fmt.Printf("decoded len=%d, result: %s\n", len(decoded), hex.EncodeToString(decoded))

	if bytes.Compare(value, decoded) == 0 {
		fmt.Println("Encode and decode successfully!")
	}
}

```

