package config

import (
	"io"
	"log"
	"os"
	"strings"
)

func NewLog(path string) {
	if !strings.HasSuffix(path, ".txt") {
		panic(".txt is not suffixed at logName Env")
	} else {
		if f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err != nil {
			if os.IsNotExist(err) {
				if f, err = os.Create(path); err != nil {
					panic(err)
				}
			} else {
				panic(err)
			}
		} else {

			// For log 초기화
			if f, err = os.Create(path); err != nil {
				panic(err)
			}

			logWriter := io.MultiWriter(f, os.Stdout)
			log.SetOutput(logWriter)

			// set Log Setting
			log.SetPrefix("Log : ")
			log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

		}
	}
}
