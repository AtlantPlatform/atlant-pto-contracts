all:

compile:
	solc --abi contracts/KYC.sol

lint:
	solhint "contracts/**/*.sol"
	solium -d contracts/
