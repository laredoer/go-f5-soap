package utils

import (
	"log"
	"runtime"
)

// Go open a goroutine
func Go(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, 64<<10)
				buf = buf[:runtime.Stack(buf, false)]
				log.Printf("go func: panic recovered: %s\n%s", err, buf)
			}
		}()

		f()
	}()
}
