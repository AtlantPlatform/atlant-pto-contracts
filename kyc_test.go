package contracts

import (
	"log"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterProvider(t *testing.T) {
	require := require.New(t)
	tx, err := kyc.RegisterProvider(opts, providers[0])
	if precheckFailed(err) {
		log.Println("WARN: retrying RegisterProvider with RemoveProvider first")
		tx, err = kyc.RemoveProvider(opts, providers[0])
		require.NoError(err)
		log.Println("RemoveProvider Tx:", tx.Hash().String())
		time.Sleep(20 * time.Second)
		tx, err = kyc.RegisterProvider(opts, providers[0])
	}
	require.NoError(err)
	log.Println("RegisterProvider Tx:", tx.Hash().String())
	time.Sleep(20 * time.Second)
}

func TestRemoveProvider(t *testing.T) {
	require := require.New(t)
	tx, err := kyc.RegisterProvider(opts, providers[1])
	if precheckFailed(err) {
		// just exists already
	} else {
		require.NoError(err)
		log.Println("RegisterProvider Tx:", tx.Hash().String())
		time.Sleep(20 * time.Second)
	}
	tx, err = kyc.RemoveProvider(opts, providers[1])
	require.NoError(err)
	log.Println("RemoveProvider Tx:", tx.Hash().String())
	time.Sleep(20 * time.Second)
}

func TestOwnerKYC(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	ro := &bind.CallOpts{}
	status, err := kyc.GetStatus(ro, KYC_CUSTOMER)
	require.NoError(err)
	assert.EqualValues(0, status)

	tx, err := kyc.ApproveAddr(opts, KYC_CUSTOMER)
	require.NoError(err)
	log.Println("ApproveAddr Tx:", tx.Hash().String())
	time.Sleep(20 * time.Second)
	status, err = kyc.GetStatus(ro, KYC_CUSTOMER)
	require.NoError(err)
	assert.EqualValues(1, status)

	tx, err = kyc.SuspendAddr(opts, KYC_CUSTOMER)
	require.NoError(err)
	log.Println("SuspendAddr Tx:", tx.Hash().String())
	time.Sleep(20 * time.Second)
	status, err = kyc.GetStatus(ro, KYC_CUSTOMER)
	require.NoError(err)
	assert.EqualValues(2, status)

	tx, err = kyc.ApproveAddr(opts, KYC_CUSTOMER)
	require.NoError(err)
	log.Println("ApproveAddr Tx:", tx.Hash().String())
	time.Sleep(20 * time.Second)
	status, err = kyc.GetStatus(ro, KYC_CUSTOMER)
	require.NoError(err)
	assert.EqualValues(1, status)

	status, err = kyc.GetStatus(ro, KYC_CUSTOMER0)
	require.NoError(err)
	assert.EqualValues(0, status)
}

func TestProviderKYC(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	tx, err := kyc.RegisterProvider(opts, providers[1])
	if precheckFailed(err) {
		// just exists already
	} else {
		require.NoError(err)
		log.Println("RegisterProvider Tx:", tx.Hash().String())
		time.Sleep(20 * time.Second)
	}
	providerOpts := transactOpts(providers[1], PROVIDER_PASS)

	ro := &bind.CallOpts{}
	status, err := kyc.GetStatus(ro, KYC_CUSTOMER2)
	require.NoError(err)
	assert.EqualValues(0, status)

	tx, err = kyc.ApproveAddr(providerOpts, KYC_CUSTOMER2)
	require.NoError(err)
	log.Println("ApproveAddr Tx:", tx.Hash().String())
	time.Sleep(20 * time.Second)
	status, err = kyc.GetStatus(ro, KYC_CUSTOMER2)
	require.NoError(err)
	assert.EqualValues(1, status)

	tx, err = kyc.SuspendAddr(providerOpts, KYC_CUSTOMER2)
	require.NoError(err)
	log.Println("SuspendAddr Tx:", tx.Hash().String())
	time.Sleep(20 * time.Second)
	status, err = kyc.GetStatus(ro, KYC_CUSTOMER2)
	require.NoError(err)
	assert.EqualValues(2, status)

	tx, err = kyc.ApproveAddr(providerOpts, KYC_CUSTOMER2)
	require.NoError(err)
	log.Println("ApproveAddr Tx:", tx.Hash().String())
	time.Sleep(20 * time.Second)
	status, err = kyc.GetStatus(ro, KYC_CUSTOMER2)
	require.NoError(err)
	assert.EqualValues(1, status)

	status, err = kyc.GetStatus(ro, KYC_CUSTOMER0)
	require.NoError(err)
	assert.EqualValues(0, status)
}
