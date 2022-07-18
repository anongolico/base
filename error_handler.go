package base

import "log"

func handle(err error, customMsg string) {
	if err != nil {
		if customMsg == "" {
			customMsg = err.Error()
		}
		log.Fatalf("%s", customMsg)
	}
}
