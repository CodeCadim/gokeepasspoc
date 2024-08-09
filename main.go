package main

import (
	"errors"
	"log"
	"os"
	"strconv"

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
	for _, item := range keystore.List() {
		println(item)
	}
	input := input.Text("Enter key index:")
	idx, err := strconv.Atoi(input)
	if err != nil {
		log.Panic(err)
	}
	pw := keystore.Index(idx)
	println("Lookup by index", idx, pw)
}
