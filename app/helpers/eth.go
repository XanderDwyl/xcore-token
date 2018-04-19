package helpers

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	bbs "github.com/a-fis/xcore-token/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

type EthereumData struct {
	HashedData string
	Timestamp  int64
}

var key *ecdsa.PrivateKey
var auth *bind.TransactOpts
var alloc core.GenesisAlloc
var sim *backends.SimulatedBackend

var addr common.Address
var contract *bbs.Bbs
var bbsErr error

func init() {
	// Go ethereum imports
	key, _ = crypto.GenerateKey()
	auth = bind.NewKeyedTransactor(key)

	alloc = make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}
	sim = backends.NewSimulatedBackend(alloc)

	addr, _, contract, bbsErr = bbs.DeployBbs(auth, sim, "Hello")
	if bbsErr != nil {
		log.Fatal(bbsErr)
	}

	return
}

// PublishToEthNetwork ...
func PublishToEthNetwork(data string) {
	log.Println("Contract successfully deployed to: " + addr.String())
	contract.Set(auth, data)
	sim.Commit()
}

// GetDataFromEthNetwork ...
func GetDataFromEthNetwork() string {
	data, _ := contract.Get(nil)
	log.Println("Data successfully retrieved from: " + addr.String())
	log.Println(data)

	return data
}

// ToString ...
func (ed *EthereumData) ToString() string {
	return fmt.Sprintf("{hash: %s, timestamp: %d}", ed.HashedData, ed.Timestamp)
}
