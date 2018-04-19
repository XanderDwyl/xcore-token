pragma solidity ^0.4.18;


contract Bbs {
    string value;

    function Bbs(string _value) public {
      value = _value;
    }

    function set(string _value) public {
        value = _value;
    }

    function get() public constant returns (string) {
        return value;
    }
} 