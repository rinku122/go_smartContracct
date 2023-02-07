// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

contract MySmartContract {
    uint256 public eventCounter;
    event Event(uint256 Counter);

    function produceEvents(uint256 _event) public {
        for (uint256 i = eventCounter + 1; i <= eventCounter + _event; i++) {
            emit Event(i);
        }

        eventCounter += _event;
    }
}
