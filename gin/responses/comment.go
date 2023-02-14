package responses

import (
	"time"

	"github.com/jjh930301/needsss/gin/dto"
	uuid "github.com/satori/go.uuid"
)

type TickerCommentsResponse struct {
	UserId    uuid.UUID       `gorm:"column:user_id;type:varchar(36)" json:"-"`
	User      dto.CommentUser `gorm:"foreignKey:UserId;references:ID" json:"user"`
	CommentId string          `gorm:"column:id" json:"comment_id"`         // 코멘트 uuid 수정이나 삭제시 필요할 수도 있습니다.
	Comment   string          `gorm:"column:comment" json:"comment"`       // 코멘트
	Symbol    string          `gorm:"column:symbol" json:"symbol"`         // 메세지를 뿌려줄 symbol
	CreatedAt time.Time       `gorm:"column:created_at" json:"created_at"` // 코멘트 등록일시
}
