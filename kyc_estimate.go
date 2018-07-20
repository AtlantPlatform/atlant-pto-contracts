// Copyright 2017, 2018 Tensigma Ltd. All rights reserved.
// Use of this source code is governed by Microsoft Reference Source
// License (MS-RSL) that can be found in the LICENSE file.

package contracts

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/AtlantPlatform/ethfw"
)

var (
	GETH_IPC      = "var/chain/geth.ipc"
	KYC_ADDRESS   = common.HexToAddress("0x51fdfcc9309a7e4f3eccacb1d58e24ec80ecf296")
	OWNER_ADDRESS = common.HexToAddress("0xafcc2ccb0f20a9d46ccde16960df93fac2c8c577")
	OWNER_PASS    = "1234"
	PROVIDER_PASS = "1234"
	OWNER_KEY     = "var/chain/keystore/UTC--2018-04-19T10-19-35.897354000Z--afcc2ccb0f20a9d46ccde16960df93fac2c8c577"
	KYC_CUSTOMER  = common.HexToAddress("0x80ea25bc4aac5d4e4a82255a117e526946857aa2")
	KYC_CUSTOMER2 = common.HexToAddress("0xfd10c8b38aab823ce2364f3d45cceb399bbf9587")
	KYC_CUSTOMER0 = common.HexToAddress("0x0000000000000000000000000000000000000000")
)

var providers = []common.Address{
	common.HexToAddress("0x8dc12760659ac3b327b7c2bd37c25f5f206d1ff3"),
	common.HexToAddress("0x8b5201d35d5e8dcecd6686d5b5731168d19381d2"),
}

var (
	kyc      *KYC
	client   *ethclient.Client
	keycache ethfw.KeyCache
	opts     *bind.TransactOpts
)

func init() {
	cli, err := ethclient.Dial(GETH_IPC)
	orPanic(err)
	client = cli
	kyc, err = NewKYC(KYC_ADDRESS, cli)
	orPanic(err)
	keycache = ethfw.NewKeyCache()
	keycache.SetPath(OWNER_ADDRESS, OWNER_KEY)
	keycache.SetPath(providers[0], "var/chain/keystore/UTC--2018-04-19T10-19-49.026595000Z--8dc12760659ac3b327b7c2bd37c25f5f206d1ff3")
	keycache.SetPath(providers[1], "var/chain/keystore/UTC--2018-04-19T10-20-39.056144000Z--8b5201d35d5e8dcecd6686d5b5731168d19381d2")
	opts = transactOpts(OWNER_ADDRESS, OWNER_PASS)
}

func RunEstimate() {
	// tx, err := kyc.RegisterProvider(opts, providers[1])
	// if precheckFailed(err) {
	// 	// just exists already
	// } else {
	// 	orPanic(err)
	// 	log.Println("RegisterProvider Tx:", tx.Hash().String())
	// 	time.Sleep(10 * time.Second)
	// }
	providerOpts := transactOpts(providers[1], PROVIDER_PASS)

	target := randomAddr()
	log.Println("Approving", target.Hex())
	tx, err := kyc.ApproveAddr(providerOpts, target)
	orPanic(err)

	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	for err != nil && strings.HasSuffix(err.Error(), "not found") {
		time.Sleep(time.Second)
		receipt, err = client.TransactionReceipt(context.Background(), tx.Hash())
	}
	if err != nil {
		log.Fatalln("Failed to get receipt:", err)
	} else {
		fmt.Printf("Status: %v Gas used: %d  CumulativeGasUsed: %d\n", receipt.Status,
			receipt.GasUsed, receipt.CumulativeGasUsed)
	}

}

func orPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func transactOpts(account common.Address, pass string) *bind.TransactOpts {
	signerFn := keycache.SignerFn(account, pass)
	if signerFn == nil {
		panic(fmt.Sprintf("auth failed: %s", account.Hex()))
	}
	return &bind.TransactOpts{
		From:   account,
		Signer: signerFn,
	}
}

func precheckFailed(err error) bool {
	if err == nil {
		return false
	}
	if err.Error() == "failed to estimate gas needed: gas required exceeds allowance or always failing transaction" {
		return true
	}
	return false
}

func randomAddr() common.Address {
	buf := make([]byte, 20)
	rand.Read(buf)
	return common.HexToAddress("0x" + hex.EncodeToString(buf))
}
