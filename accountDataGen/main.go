package accountDataGen

import (
	"dustin-ward/genAccountData/user"
	"math/rand"
	"time"
)

type AccountData struct {
	AccountNumber    int
	UserData         user.User
	NumAccounts      int
	AvgAccountValue  string
	AvgCreditRate    string
	AvgInestmentRate string
	NetWorth         string
	Income           string
	TotalLiabilities string
	CreditScore      int
	MemberSince      string
	AccountActive    string
}

var rng *rand.Rand

func Init() {
	src := rand.NewSource(time.Now().Unix())
	rng = rand.New(src)

	user.Init()
}
