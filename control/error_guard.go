package control

import "log"

func Safe(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[error]", r)
		}
	}()
	fn()
}
