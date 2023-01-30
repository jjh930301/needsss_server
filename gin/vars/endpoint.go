package vars

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
	Token         string = "/token"
	GoogleIDToken string = "/google/token"
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
