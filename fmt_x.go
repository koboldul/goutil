package util

import "fmt"

func Printlnf(msg string, p ...interface{}) (int, error) {
	return fmt.Println(fmt.Sprintf(msg, p...))
}
