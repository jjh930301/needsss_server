package logs

import (
	"time"

	"github.com/jjh930301/market/common/constants"
	"github.com/jjh930301/market/common/database"
	"github.com/jjh930301/market/common/models"
)

func Log(ip string, method int, endpoint string) {
	now := time.Now().In(constants.KrTime())
	database.DB.Create(&models.LogModel{
		Ip:       ip,
		Method:   method,
		DateTime: now,
		EndPoint: endpoint,
	})
}
