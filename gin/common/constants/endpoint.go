package constants

const (
	GET    int = 0
	POST   int = 1
	PUT    int = 2
	DELETE int = 3
)

const (
	AuthGroup     string = "/auth"
	TickerGroup   string = "/ticker"
	CommentGroup  string = "/comment"
	InterestGroup string = "/interest"
	UserGroup     string = "/user"
)

const (
	Default string = ""
	//auth
	Token          string = "/token"
	Regist         string = "/regist"
	Login          string = "/login"
	GoogleLogin    string = "/google/login"
	GoogleCallback string = "/google/callback"
	//tiocker
	GetTicker    string = "/:ticker"
	GetChart     string = "/chart/:ticker"
	TickerSearch string = "/search"
	//interest
	SaleInterestTicker string = "/sale"
	//user
	SetNickname string = "/nickname"
	//comment
	GetTickerComments string = "/:ticker"
)
