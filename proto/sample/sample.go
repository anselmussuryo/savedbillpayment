package sample

import (
	"github.com/anselmussuryo/savedbillpayment/pb"
	"github.com/anselmussuryo/savedbillpayment/util"
)

// NewSavedBillPayment returns a new sample SavedBillPayment
func NewSavedBillPayment() *pb.SavedBillPayment {
	savedBillPayment := &pb.SavedBillPayment{
		MerchantID:   util.RandomInt(1, 100),
		CustomerID:   util.RandomInt(1, 100),
		SubsciberNo:  util.RandomInt(1, 100),
		Description:  util.RandomString(10),
		Isshowib:     util.RandomString(3),
		Isshowmobile: util.RandomString(3),
		Amount:       util.RandomFloat(0, 1000),
		Isfavorite:   util.RandomString(3),
	}

	return savedBillPayment
}
