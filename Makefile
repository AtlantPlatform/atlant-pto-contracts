all:

gen:
	cd contracts/ && abigen -sol KYC.sol -pkg contracts -out ../kyc_gen.go

compile:
	solc --abi contracts/KYC.sol

init:
	mkdir -p var
	geth --datadir=var/chain/ init var/genesis.json

node:
	geth --datadir=var/chain/ --rpc

attach:
	geth attach var/chain/geth.ipc --preload "scripts/allbalances.js"

estimate:
	go run scripts/run_estimate.go
