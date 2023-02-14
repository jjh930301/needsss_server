package dto

import "time"

type HolidayDto struct {
	Name     string    `json:"name"` // 휴일 이름
	Date     string    `json:"date"` // 휴일 날짜
	TDate    time.Time `json:"-"`
	OpenedAt string    `json:"opened_at"` // 시간 [수능일 같은 경우 10시에 개장]
	TOpendAt time.Time `json:"-"`
	Type     int       `json:"type"` // 1인 경우 opened_at 이후에 가능
}
