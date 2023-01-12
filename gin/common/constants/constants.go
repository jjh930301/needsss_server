package constants

import "time"

func KrTime() *time.Location {
	loc, _ := time.LoadLocation("Asia/Seoul")
	return loc
}
