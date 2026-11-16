package common

import (
	"github.com/leekchan/accounting"
)

func FormatMoney(cents int) string {
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	return ac.FormatMoney(cents)
}
