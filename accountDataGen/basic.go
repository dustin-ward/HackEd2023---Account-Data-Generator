package accountDataGen

import (
	"dustin-ward/genAccountData/user"
	"fmt"
	"time"
)

func (d *AccountData) gen_memberSince() {
    yearsOver17 := d.UserData.Age - 17
    d.MemberSince = time.Date(2023-rng.Intn(yearsOver17),
                    time.Month(rng.Intn(12)+1),
                    rng.Intn(31)+1,
                    0,
                    0,
                    0,
                    0,
                    time.UTC,
                ).Format("2006-02-01")
}

func (d *AccountData) gen_numAccounts() {
    d.NumAccounts = rng.Intn(4)
    d.NumAccounts += int(0.05 * float32(d.UserData.Age))
}

func GenerateData_Basic(size int, split float64) []AccountData {
	activeAccounts := int(float64(size) * split)

	data := make([]AccountData, 0)

	for i := 0; i < size; i++ {
		var d AccountData
		d.UserData = user.GenUserData()
		d.AccountNumber = int(rng.Int31())
        d.gen_memberSince()
        d.Income = fmt.Sprintf("%.2f", rng.Float64()*999999)
        d.NetWorth = fmt.Sprintf("%.2f", rng.Float32()*2000000)
        d.CreditScore = rng.Intn(999-300) + 300
		d.gen_numAccounts()
		d.AvgAccountValue = fmt.Sprintf("%.2f", rng.Float32()*1000000)
		d.AvgCreditRate = fmt.Sprintf("%.2f", rng.Float32()*50)
		d.AvgInestmentRate = fmt.Sprintf("%.2f", rng.Float32()*50)
		d.TotalLiabilities = fmt.Sprintf("%.2f", rng.Float32()*200000)
		if i < activeAccounts {
			d.AccountActive = "1.0"
		} else {
			d.AccountActive = "0.0"
		}
		data = append(data, d)
	}

	return data
}
