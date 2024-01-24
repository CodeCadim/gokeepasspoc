package keystore

import (
	"os"

	"github.com/tobischo/gokeepasslib/v3"
)

type keystore struct {
	db   *gokeepasslib.Database
	file *os.File
}

var ks keystore

func Open(filename, pass string) (err error) {
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

func Close() error {
	return ks.file.Close()
}

func Lookup(key string) (pass string) {
	for _, entry := range ks.db.Content.Root.Groups[0].Entries {
		if key == entry.GetTitle() {
			return entry.GetPassword()
		}
	}
	return ""
}

func List() (res []string) {
	for _, entry := range ks.db.Content.Root.Groups[0].Entries {
		res = append(res, entry.GetTitle())
	}
	return res
}
