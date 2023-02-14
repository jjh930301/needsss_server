package responses

import "github.com/jjh930301/needsss/gin/dto"

type InterestListResponse struct {
	Ticker dto.InterestList             `json:"ticker"`
	Recent dto.InterestTickerChartModel `json:"recent"`
}
