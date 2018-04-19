#!/bin/sh

rm -rf contracts/bbs.go
# rm -rf contracts/bbs.abi

# solc --abi contracts/bbs.sol| grep '^[[a-zA-Z]' | grep -vE "(Contract JSON ABI)" >> contracts/bbs.abi

abigen --sol=contracts/bbs.sol --pkg=bbs --out=contracts/bbs.go
