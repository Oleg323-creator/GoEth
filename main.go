package main

//module github.com/Oleg323-creator

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"log"
	"math/big"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	client, err := ethclient.Dial("https://sepolia.infura.io/v3/8ab0a7925db44ab094fa1b6546409a6b")
	if err != nil {
		log.Fatalf("Error connecting to Sepolia node: %v", err)
	}
	fmt.Println("Connection success!")

	fromAddress := common.HexToAddress("0xAa8ff5ed1dA7832b6b361F90b9bA6D7b384Ea5E9")
	toAddress := common.HexToAddress("0x886577048713f65d6e26e61e82597A523887645B")
	amount := new(big.Int)
	amount.SetString("100000000000000000", 10) //WEI

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Error getting gas price: %v", err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Error getting nonce: %v", err)
	}

	gasLimit := uint64(21000)

	tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, nil)

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("Error reading private key: %v", err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP2930Signer(big.NewInt(11155111)), privateKey)
	if err != nil {
		log.Fatalf("Error singing tx: %v", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Error sending tx: %v", err)
	}

	log.Printf("Tx has sent! Hash: %s\n", signedTx.Hash().Hex())
}

//	//GENERATING PRIVAT KEY
//	privatKey, err := crypto.GenerateKey()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	privateKeyBytes := crypto.FromECDSA(privatKey)
//	fmt.Println("Private key: ", hexutil.Encode(privateKeyBytes)[2:]) //JUST TO SEE THAT IT WORKS
//
//	//GENERATING PUBLIC KEY
//	publicAddres := crypto.PubkeyToAddress(privatKey.PublicKey).Hex()
//	fmt.Printf("Public address: %s", publicAddres)
