// Copyright 2017-2021 Digital Asset Exchange Limited. All rights reserved.
// Use of this source code is governed by Microsoft Reference Source
// License (MS-RSL) that can be found in the LICENSE file.

pragma solidity ^0.4.24;

import "./lib/Owned.sol";
import "./lib/AddrSet.sol";


// KYC implements "Know Your Customer" storage for identity approvals by KYC providers.
contract KYC is Owned {

    // Status corresponding to the state of approvement:
    // * Unknown when an address has not been processed yet;
    // * Approved when an address has been approved by contract owner or 3rd party KYC provider;
    // * Suspended means a temporary or permanent suspension of all operations, any KYC providers may
    // set this status when account needs to be re-verified due to legal events or blocked because of fraud.
    enum Status {
        unknown,
        approved,
        suspended
    }

    // Events emitted by this contract
    event ProviderAdded(address indexed addr);
    event ProviderRemoved(address indexed addr);
    event AddrApproved(address indexed addr, address indexed by);
    event AddrSuspended(address indexed addr, address indexed by);

    // Contract state
    AddrSet.Data private kycProviders;
    mapping(address => Status) public kycStatus;

    // registerProvider adds a new 3rd-party provider that is authorized to perform KYC.
    function registerProvider(address addr) public onlyOwner {
        require(AddrSet.insert(kycProviders, addr));
        emit ProviderAdded(addr);
    }

    // removeProvider removes a 3rd-party provider that was authorized to perform KYC.
    function removeProvider(address addr) public onlyOwner {
        require(AddrSet.remove(kycProviders, addr));
        emit ProviderRemoved(addr);
    }

    // isProvider returns true if the given address was authorized to perform KYC.
    function isProvider(address addr) public view returns (bool) {
        return addr == owner || AddrSet.contains(kycProviders, addr);
    }

    // getStatus returns the KYC status for a given address.
    function getStatus(address addr) public view returns (Status) {
        return kycStatus[addr];
    }

    // approveAddr sets the address status to Approved, see Status for details.
    // Can be invoked by owner or authorized KYC providers only.
    function approveAddr(address addr) public onlyAuthorized {
        Status status = kycStatus[addr];
        require(status != Status.approved);
        kycStatus[addr] = Status.approved;
        emit AddrApproved(addr, msg.sender);
    }

    // suspendAddr sets the address status to Suspended, see Status for details.
    // Can be invoked by owner or authorized KYC providers only.
    function suspendAddr(address addr) public onlyAuthorized {
        Status status = kycStatus[addr];
        require(status != Status.suspended);
        kycStatus[addr] = Status.suspended;
        emit AddrSuspended(addr, msg.sender);
    }

    // onlyAuthorized modifier restricts write access to contract owner and authorized KYC providers.
    modifier onlyAuthorized() {
        require(msg.sender == owner || AddrSet.contains(kycProviders, msg.sender));
        _;
    }
}
