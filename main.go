package main

import (
	"log"

	"gokeepasspoc/input"
	"gokeepasspoc/keystore"
)

func main() {
	kbdxfile := input.Env("PASSDB")
	if kbdxfile == "" {
		log.Panic("define PASSDB env to your kdbx file")
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
	input := input.Text("Enter key name:")
	pw := keystore.Lookup(input)
	println("Lookup: ", pw)
}
