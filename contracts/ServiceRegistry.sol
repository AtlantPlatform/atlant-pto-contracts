pragma solidity ^0.4.20; // nice

import "./lib/Owned.sol";

contract ServiceRegistry is Owned {
  address public service;

  event ReplaceService(address oldService, address newService);

  // https://github.com/Dexaran/ERC223-token-standard/blob/Recommended/ERC223_Token.sol#L107-L114
  modifier withContract(address _addr) {
    uint length;
    assembly { length := extcodesize(_addr) }
    require(length > 0);
    _;
  }

  function ServiceRegistry(address _service) public {
    service = _service;
  }

  function replaceService(address _service) onlyOwner withContract(_service) public {
    address oldService = service;
    service = _service;
    ReplaceService(oldService, service);
  }
}
