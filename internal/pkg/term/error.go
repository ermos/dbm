package term

import (
	"fmt"
)

func ErrorHandler() {
	if err := recover(); err != nil {
		fmt.Printf("dbm: %v.\n", err)
	}
}
