package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gokeepasspoc/input"
	"gokeepasspoc/keystore"
)

func main() {
	kbdxfile := input.Env("PASSDB")
	if kbdxfile == "" {
		log.Panic("define PASSDB env to your kdbx file")
	}
	if _, err := os.Stat(kbdxfile); errors.Is(err, os.ErrNotExist) {
		log.Fatalf("%s does not exist", kbdxfile)
	}
	pass := input.Pass("Enter password:")
	err := keystore.Open(kbdxfile, pass)
	if err != nil {
		log.Panic(err)
	}
	defer keystore.Close()

	result, err := fzfRun(keystore.List())
	if err != nil {
		log.Fatalf("running fzf: %s", err)
	}

	pw := keystore.Lookup(result)
	fmt.Print(pw)
}
