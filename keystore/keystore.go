package keystore

import (
	"fmt"
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

func Index(idx int) (pass string) {
	return ks.db.Content.Root.Groups[0].Entries[idx].GetPassword()
}

func Lookup(key string) (pass string) {
	for _, entry := range ks.db.Content.Root.Groups[0].Entries {
		if key == entry.GetTitle() {
			username := entry.GetContent("UserName")
			password := entry.GetContent("Password")
			return fmt.Sprintf("%s\t%s\n", username, password)
		}
	}
	return ""
}

func List() (res []string) {
	for idx, entry := range ks.db.Content.Root.Groups[0].Entries {
		res = append(res, fmt.Sprintf("%d: %s", idx, entry.GetTitle()))
	}
	return res
}
