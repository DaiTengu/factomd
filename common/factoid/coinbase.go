// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package factoid

import (
	"github.com/FactomProject/factomd/common/interfaces"
	"time"
)

var adrs []interfaces.IAddress
var amount uint64 = 5000000000 // One Factoid (remember, fixed point math!
var addressCnt int = 0         // No coinbase payments until Milestone 3

// Allows the amount paid in the coinbase to be modified.   This is
// NOT allowed in production!  That's why it is here in Test!
func UpdateAmount(amt uint64) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidUpdateAmount.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	amount = amt
}

// This routine generates the Coinbase.  This is a fixed amount to be
// paid to the federated servers.
//
// Currently we are paying just a few fixed addresses.
//
func GetCoinbase(ftime interfaces.Timestamp) interfaces.ITransaction {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidGetCoinbase.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	coinbase := new(Transaction)
	coinbase.SetTimestamp(ftime)

	for _, adr := range adrs {
		coinbase.AddOutput(adr, amount) // add specified amount
	}

	return coinbase
}
