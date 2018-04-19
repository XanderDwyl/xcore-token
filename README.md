


# how to start to dev this

```
go get github.com/pilu/fresh
```

```
cd ~/src/github.com/a-fis/xcore-token
fresh
```

# using dep instead of glide 

- https://golang.github.io/dep/docs/new-project.html
- https://github.com/golang/dep

## install dep 

```
 curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```

## dep init
```
dep init 
```

## add package 

```
dep ensure -add XXXXX
```

## update package 
```
dep ensure -update
```

-------
# spec 
## Get XCore Token Info
- [ ] Get total raised amount
- [ ] Display Token Info
  - Price
  - Total Supply
  - Start Date
  - End Date
  _ First week bonus
  - Max Cap / investor

- [ ] Buy token / send token

#NOTES:

## Generate Solidity Contract to ABI
Make sure to install `solc` and `abigen`

## Install solidity
Reference: http://solidity.readthedocs.io/en/v0.4.21/installing-solidity.html
Homebrew is missing pre-built bottles at the time of writing, following a Jenkins to TravisCI migration, but Homebrew should still work just fine as a means to build-from-source. We will re-add the pre-built bottles soon.

```
brew update
brew upgrade
brew tap ethereum/ethereum
brew install solidity
brew linkapps solidity
```

### Building Go-Ethereum from source
Building Geth (command line client)
Clone the repository to a directory of your choosing:

```
git clone https://github.com/ethereum/go-ethereum

ln -s $GOPATH/src/github.com/ethereum/go-ethereum /usr/local/bin/abigen
```

## How to compile solidity contracts
```
solc --abi contracts/<xxxx.sol>
```

## How to convert Ethereum contract definitions into compile-time type-safe Go packages.
```
abigen --abi contracts/<filename.abi> --pkg <contract package> --out contracts/<filename.go>

or simply run the command below:

./compile.sh
```

