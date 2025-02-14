package main

import (
	"ledger-system/db"
	"os"
)

func main() {
	db.Init(os.Args[1:])
}
