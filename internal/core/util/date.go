package util

import (
	"time"
)

const timeZone = "America/Sao_Paulo"

func TimeNowInSaoPaulo() time.Time {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return time.Now()
	}

	return time.Now().In(loc)
}
