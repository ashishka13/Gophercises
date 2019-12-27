package main

import "gophercises/secret/cmd/cobra"

func main() {
	cobra.RootCmd.Execute()
}

/*
steps from secret folder itself
go run cmd/cli.go
go build -o ./secret cmd/cli.go
./secret set someKey someValue
./secret get someKey
*/
