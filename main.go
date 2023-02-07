// package main

// import (
// 	"context"
// 	"crypto/ecdsa"
// 	"fmt"
// 	"log"
// 	"math/big"

// 	MyContract "test/gen"

// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	"github.com/ethereum/go-ethereum/ethclient"
// )

// var infuraURL = "https://goerli.infura.io/v3/d90e669987d4483590a00ad4908191f6"

// func main() {
// 	client, err := ethclient.DialContext(context.Background(), infuraURL)

// 	if err != nil {
// 		log.Fatalf("Error to create a ether client:%v", err)
// 	}

// 	defer client.Close()

// 	block, err := client.BlockByNumber(context.Background(), nil)
// 	if err != nil {
// 		log.Fatalf("Error to get the block:%v", err)
// 	}

// 	fmt.Println("Block Numer : ", block.Number())

// 	cAdd := common.HexToAddress("0x5Ef2492f42E3b5Ab72797C0516486f68E90b0799")
// 	contract, err := MyContract.NewMyContract(cAdd, client)
// 	if err != nil {
// 		log.Fatalf("Contract instance error:%v", err)
// 	}

// 	privateKey, err := crypto.HexToECDSA("0d03a89dbe214d61bdfd8afec6ca5c46cf4ccb205820bdd10a7834b5e7cd4db7")

// 	if err != nil {
// 		log.Fatalf("Error in private key:%v", err)
// 	}

// 	c := make(chan string)

// 	for i := 1; i <= 5; i++ {
// 		go callContract(client, contract, privateKey, c)
// 		fmt.Println(<-c, i, "Times")
// 	}

// }

// func callContract(client *ethclient.Client, contract *MyContract.MyContract, privateKey *ecdsa.PrivateKey, c chan string) {

// 	chainID, err := client.NetworkID(context.Background())
// 	if err != nil {
// 		log.Fatalf("Error in chainID  key:%v", err)
// 		c <- `Produced event Error in ChainID`

// 	}

// 	gasPrice, err := client.SuggestGasPrice(context.Background())
// 	if err != nil {
// 		log.Fatalf("Error in gasPrice  key:%v", err)
// 		c <- `Produced event Error in gasPrice`

// 	}

// 	tx, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
// 	if err != nil {
// 		log.Fatalf("Tx err :%v", err)
// 		c <- `Produced event Error in tx`

// 	}

// 	tx.GasLimit = 3000000
// 	tx.GasPrice = gasPrice
// 	txHash, err := contract.ProduceEvents(tx, big.NewInt(20))

// 	if err != nil {
// 		log.Fatalf("Calling Contract Faild:%v", err)
// 		c <- `Produced event Error`

// 	}
// 	fmt.Println("Transaction Hash", txHash)

// 	c <- `Produced event`

// }
package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	MySmartContract "test/gen"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "http://127.0.0.8545"
var Pki = "Private Key"
var contractAddress = "0xCAFBb81820fD028dd83ef148815b4d1519aa16F6"
var numberOfEvents = big.NewInt(2)

func main() {
	client, err := ethclient.DialContext(context.Background(), infuraURL)

	if err != nil {
		log.Fatalf("Error to create a ether client:%v", err)
	}

	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error to get the block:%v", err)
	}

	fmt.Println("Block Numer : ", block.Number())
	privateKey, err := crypto.HexToECDSA(Pki)
	if err != nil {
		log.Fatalf("Error in private key:%v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("Error in public  key:%v", err)
	}

	userAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("userAddress : ", userAddress)

	cAdd := common.HexToAddress(contractAddress)
	contract, err := MySmartContract.NewMySmartContract(cAdd, client)

	if err != nil {
		log.Fatalf("Contract instance error:%v", err)
	}

	c := make(chan string)

	for i := 1; i <= 100; i++ {
		go callContract(client, contract, privateKey, userAddress, c)
		fmt.Println(<-c, i, "Times")
	}

}

func callContract(client *ethclient.Client, contract *MySmartContract.MySmartContract, privateKey *ecdsa.PrivateKey, userAddress common.Address, c chan string) {
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Error in chainID  key:%v", err)
		c <- `Error in ChainID`

	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Error in gasPrice  key:%v", err)
		c <- `Error in gasPrice`

	}

	nonce, err := client.PendingNonceAt(context.Background(), userAddress)

	if err != nil {
		c <- `Error in nounce`
		panic(err)

	}

	tx, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Tx err :%v", err)
		c <- `Error in Tx`

	}

	tx.GasLimit = 3000000
	tx.GasPrice = gasPrice
	tx.Nonce = big.NewInt(int64(nonce))
	tx.Value = big.NewInt(0)
	txHash, err := contract.ProduceEvents(tx, numberOfEvents)

	if err != nil {
		log.Fatalf("Calling Contract Faild:%v", err)
		c <- `Error in Producing Event`

	}
	fmt.Println("Transaction Hash", txHash.Hash().Hex())

	eventCounter, err := contract.EventCounter(&bind.CallOpts{
		From: userAddress,
	})

	if err != nil {
		log.Fatalf("Getting new Value Error:%v", err)

	}

	fmt.Println("New Event Counter : ", eventCounter)
	c <- `Produced event`
}
