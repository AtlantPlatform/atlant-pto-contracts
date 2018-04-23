pragma solidity ^0.4.20; // nice

import "./lib/token/StandardToken.sol";
import "./ServiceRegistry.sol";
import "./KYC.sol";

contract PropertyToken is StandardToken, Owned {

  uint public limit = 50 * 1e6 * 1e18;
  string public name = "Atlant Property Token 0000";
  string public symbol = "ATPT0000";
  uint8 public decimals = 18;

  ServiceRegistry public kycRegistry;

  function PropertyToken(ServiceRegistry _kycRegistry, string _name, string _symbol, uint _limit) public {
    require(_kycRegistry != address(0));
    require(_limit != 0);
    kycRegistry = _kycRegistry;

    name = _name;
    symbol = _symbol;
    limit = _limit;
  }

  function mint(address _holder, uint _value) external onlyOwner {
    require(_value != 0);
    require(totalSupply + _value <= limit);
    _checkKYC(_holder, _holder);

    balances[_holder] += _value;
    totalSupply += _value;
    Transfer(0x0, _holder, _value);
  }

  function transfer(address _to, uint _value) public {
    _checkKYC(msg.sender, _to);
    return super.transfer(_to, _value);
  }

  function transferFrom(address _from, address _to, uint _value) public {
    _checkKYC(_from, _to);
    return super.transferFrom(_from, _to, _value);
  }

  function _checkKYC(address _from, address _to) private {
    require(_service().getStatus(_from) == KYC.Status.approved);
    require(_service().getStatus(_to) == KYC.Status.approved);
  }

  function _service() constant public returns (KYC) {
    return KYC(kycRegistry.service());
  }
}
