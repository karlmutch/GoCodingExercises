// stub.go
package main

import "errors"

func Strawman(word string) (bool, error) {
	switch word {
	case "strawman":
		return true, nil
	default:
		err := errors.New("There was no strawman")
		return false, err
	}
	return false, nil
}
