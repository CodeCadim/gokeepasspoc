package main

import (
	"bufio"
	"os"
	"syscall"

	"github.com/tobischo/gokeepasslib/v3"
	"golang.org/x/term"
)

func main() {
	kbdxfile, ok := os.LookupEnv("PASSDB")
	if !ok {
		println("Define a PASSDB env pointing to your databse kdbx file")
		os.Exit(-1)
	}

	print("\nEnter password: ")
	pass := readpass()

	passmngr := keystore{}
	err := passmngr.open(kbdxfile, pass)
	if err != nil {
		panic(err)
	}
	defer passmngr.close()

	for _, item := range passmngr.list() {
		println(item)
	}

	print("\nEnter key name: ")
	input := readtext()
	pw := passmngr.lookup(input)
	println("Lookup: ", pw)
}

type keystore struct {
	db   *gokeepasslib.Database
	file *os.File
}

func (ks *keystore) open(filename, pass string) (err error) {
	ks.file, err = os.Open(filename)
	if err != nil {
		return err
	}
	ks.db = gokeepasslib.NewDatabase()
	ks.db.Credentials = gokeepasslib.NewPasswordCredentials(pass)
	err = gokeepasslib.NewDecoder(ks.file).Decode(ks.db)
	if err != nil {
		return err
	}
	return ks.db.UnlockProtectedEntries()
}

func (ks *keystore) close() error {
	return ks.file.Close()
}

func (ks *keystore) lookup(key string) (pass string) {
	for _, entry := range ks.db.Content.Root.Groups[0].Entries {
		if key == entry.GetTitle() {
			return entry.GetPassword()
		}
	}
	return ""
}

func (ks *keystore) list() (res []string) {
	for _, entry := range ks.db.Content.Root.Groups[0].Entries {
		res = append(res, entry.GetTitle())
	}
	return res
}

func readpass() string {
	bytepw, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		panic(err)
	}
	return string(bytepw)
}

func readtext() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return scanner.Text()
}
