package pkg

import (
	"time"
)

func CalculateExpiration(expired int64) time.Time {
	return time.Unix(expired, 0)
	//return expiration.Format("2006-01-02 15:04:05")
	//expiration := time.Unix(expired, 0)
	//return expiration
}
