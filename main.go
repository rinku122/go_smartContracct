package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	MyContract "test/gen"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://goerli.infura.io/v3/d90e669987d4483590a00ad4908191f6"

// Contract Address = 0x5Ef2492f42E3b5Ab72797C0516486f68E90b0799
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
	privateKey, err := crypto.HexToECDSA("PRIVATE KEY")
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
	cAdd := common.HexToAddress("0x5Ef2492f42E3b5Ab72797C0516486f68E90b0799")
	contract, err := MyContract.NewMyContract(cAdd, client)
	if err != nil {
		log.Fatalf("Contract instance error:%v", err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Error in chainID  key:%v", err)

	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Error in gasPrice  key:%v", err)

	}

	tx, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Tx err :%v", err)

	}

	tx.GasLimit = 3000000
	tx.GasPrice = gasPrice
	txHash, err := contract.ProduceEvents(tx, big.NewInt(1))

	if err != nil {
		log.Fatalf("Calling Contract Faild:%v", err)

	}
	fmt.Println("Transaction Hash", txHash)

	eventCounter, err := contract.EventCounter(&bind.CallOpts{
		From: userAddress,
	})

	if err != nil {
		log.Fatalf("Getting new Value Error:%v", err)

	}

	fmt.Println("New Event Counter : ", eventCounter)

}
