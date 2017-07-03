package common

import (
	"time"
	"strconv"
	"log"
)

func Tick(times string, duration int64, string string, fun func() error) {
	var ticker *time.Ticker
	switch times {
	case "freedom":
		ticker = time.NewTicker(time.Millisecond * time.Duration(duration))
	case "minute":
		ticker = time.NewTicker(time.Minute)
	case "hour":
		ticker = time.NewTicker(time.Minute * 60)
	case "day":
		ticker = time.NewTicker(time.Minute * 60 * 24)
	}

	i := 0
	for range ticker.C {
		i = i + 1
		err := fun()
		if err != nil {
			LogErr(err)
		}
		if Debug1 == true {
			log.Printf("%s\t%s\t%s\t%s\t", string, "ticked at ", strconv.Itoa(i), time.Now())
		}

	}

}
