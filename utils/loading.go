package utils

import (
	"fmt"
	"time"

	"github.com/tj/go-spin"
)

func Loading() {
	s := spin.New()
	for {
		fmt.Printf("\r %s ", s.Next())
		time.Sleep(100 * time.Millisecond)
	}
}
