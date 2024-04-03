package main

import (
	cli "github.com/coinbase-samples/commerce-cli/cmd"

	"github.com/coinbase-samples/commerce-cli/sdk"
)

func main() {
	sdk.InitClient()
	cli.Execute()
}
