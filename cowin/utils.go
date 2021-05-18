package cowin

import (
	"time"
)

func today() string {
	return time.Now().Local().Format("02-01-2006")
}
