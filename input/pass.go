package input

import (
	"bufio"
	"log"
	"os"
	"syscall"

	"golang.org/x/term"
)

func Env(key string) string {
	res, _ := os.LookupEnv(key)
	return res
}

func Pass(prompt string) string {
	print(prompt)
	bytepw, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		log.Printf("%s", err)
		return ""
	}
	return string(bytepw)
}

func Text(prompt string) string {
	print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Printf("%s", err)
		return ""
	}
	return scanner.Text()
}
