package utils

import (
	"log"
	"time"
)

var HCMLocationTime *time.Location

func init() {
	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		log.Panic(err)
	}

	HCMLocationTime = location
}
