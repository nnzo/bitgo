# BitGo
A golang wrapper for the Bitcoin RPC API


## Usage

Install:
`go get github.com/nnzo/bitgo`

```go
package main

import (
	"fmt"
	bitgo "github.com/nnzo/bitgo"
)

func main() {
	client := bitcoinrpc.NewClient("http://localhost:8332", "yourUsername", "yourPassword")

	blockCount, err := client.GetBlockCount()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Block Count:", blockCount)
}
```