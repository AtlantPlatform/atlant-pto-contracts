atlant-pto-contracts
--------------------

## Dependencies

* Geth (if need to test locally): https://geth.ethereum.org/downloads/ 
* Ethereum Playbook (deploying): https://github.com/AtlantPlatform/ethereum-playbook
* Solhint (linting): https://github.com/protofire/solhint

## Deploying using Ethereum Playbook

Endpoint to see the Tx results if using `testnet` inventory group: https://ropsten.etherscan.io/

```
$ ethereum-playbook -f playbook.testnet.yml -g testnet -h

Usage: ethereum-playbook [OPTIONS] COMMAND [arg...]

Ethereum contracts deployment and management tool.

Options:
  -f                    Custom path to playbook.yml spec file. (default "playbook.testnet.yml")
  -s                    Name or path of Solidity compiler (solc, not solcjs). (default "solc")
  -g                    Inventory group name, corresponding to Geth nodes. (default "testnet")
  -h                    Print help. (default true)
  -l, --log-level       Sets the log level (default: info) (default 4)

Commands:
  balances              Generic CALL command, accepts 0 args
  txInfo                Generic CALL command, accepts 1 args
  txReceipt             Generic CALL command, accepts 1 args

  kycCheckProviders     Checks what addresses from wallets are authorized to perform KYC
  kycGetStatus          Gets the KYC status of a target address (0 = unknown, 1 = Approved, 2 = Suspended)
  kycIsProvider         Checks if the provided address is authorized to perform KYC
  kycOwner              Gets the owner address of a KYC contract
  kycApproveAddr        Sets the address status to 1 (= Approved). Can be invoked by owner or authorized KYC providers only
  kycDeploy             Deploys a KYC contract
  kycRegisterProvider   Adds a new 3rd-party provider that is authorized to perform KYC
  kycRemoveProvider     Removes a 3rd-party provider that was authorized to perform KYC
  kycSuspendAddr        Sets the address status to 2 (= Suspended). Can be invoked by owner or authorized KYC providers only

Run 'ethereum-playbook COMMAND --help' for more information on a command.
```

To deploy KYC:

```
$ ethereum-playbook -f playbook.testnet.yml -g testnet kycDeploy

INFO[0001] located keyfile by address       address=0xee84ebaaccfd55dca9b773f5d0db7e7e5d5e04e2 section=Wallets wallet=kycOwner
INFO[0003] contract deployed        address=0x406ffb87bf17727242240b9c446ed92cd9c9b06c contract=KYC
"0xa6031c5f18a7794ca8086c8bacdddbd926ed86718c603fc3b5d295e62e7b63c5"
```

To register a KYC provider:

```
$ ethereum-playbook -f playbook.testnet.yml -g testnet kycRegisterProvider 0xf86ef48734f7209054cfb563b551a11179d84e50

"0x4f2f2a18a8365f2581316f2ab1bd53fa18e8be018aa00d778ffcbb8b9daab0c1"

$ ethereum-playbook -f playbook.testnet.yml -g testnet kycCheckProviders

0xee84ebaaccfd55dca9b773f5d0db7e7e5d5e04e2 (@kycOwner): true
0xf86ef48734f7209054cfb563b551a11179d84e50 (@kycProvider1): true
```

To alter the KYC status:

```
$ ethereum-playbook -f playbook.testnet.yml -g testnet kycApproveAddr 0x5130df55dc8db83dd86832c947e41d1db4aa21e1

"0x81b54b0d3447d1a7d00590c0d249205910be113aca5b4b643fae0abe2dcec1f2"

$ ethereum-playbook -f playbook.testnet.yml -g testnet kycGetStatus 0x5130df55dc8db83dd86832c947e41d1db4aa21e1

1
```
