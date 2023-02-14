package endpoint

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
	HolidayGroup  string = "/holiday"
)

const (
	Default string = ""
	//auth
	Token         string = "/token"
	GoogleIDToken string = "/google/token"
	//tiocker
	GetTicker    string = "/:symbol"
	GetChart     string = "/chart/:symbol"
	TickerSearch string = "/search"
	//interest
	SaleInterestTicker string = "/sale"
	//user
	SetNickname string = "/nickname"
	//comment
	GetTickerComments string = "/:symbol"
)
