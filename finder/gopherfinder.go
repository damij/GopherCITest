package finder

import (
	"errors"
)

func IsItGopher(word string) (bool, error) {
	switch(word) {
	case "gopher":
		return true, nil
	case "weasel":
		err := errors.New("I AM WEASEL!")
		return false, err
	}
	return false, nil
}
