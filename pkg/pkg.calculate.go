package pkg

import (
	"time"
)

func CalculateExpiration(expired int64) string {
	expiration := time.Unix(expired, 0)
	return expiration.Format("2006-01-02 15:04:05")
	//expiration := time.Unix(expired, 0)
	//return expiration
}
