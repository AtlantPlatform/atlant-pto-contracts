// Copyright 2017, 2018 Tensigma Ltd. All rights reserved.
// Use of this source code is governed by Microsoft Reference Source
// License (MS-RSL) that can be found in the LICENSE file.

pragma solidity ^0.4.24;

import "./lib/Owned.sol";


contract ServiceRegistry is Owned {
    address public service;

    event ServiceReplaced(address indexed oldService, address indexed newService);

    // https://github.com/Dexaran/ERC223-token-standard/blob/Recommended/ERC223_Token.sol#L107-L114
    modifier withContract(address _addr) {
        uint length;
        assembly { length := extcodesize(_addr) }
        require(length > 0);
        _;
    }

    constructor(address _service) public {
        service = _service;
    }

    function replaceService(address _service) public onlyOwner withContract(_service) {
        address oldService = service;
        service = _service;
        emit ServiceReplaced(oldService, service);
    }
}
