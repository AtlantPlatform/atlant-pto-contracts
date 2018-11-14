all:

compile:
	solc --abi contracts/KYC.sol

lint:
	solhint "contracts/**/*.sol"
	# disabled:solium -d contracts/
