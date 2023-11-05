package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/miguelmota/go-ethereum-hdwallet"
)

func main() {
	rpcURL := "https://sepolia.infura.io/v3/339aeb3ae4f948a39f86ff263d8c77d7"

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Secure seed for generating the HD wallet
	mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"

	// Create a new HD wallet from the mnemonic
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatalf("Failed to create HD wallet: %v", err)
	}

	// Derive a child Ethereum address for the user
	userAddress, err := wallet.Derive(hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0"), false)
	if err != nil {
		log.Fatalf("Failed to derive Ethereum address: %v", err)
	}

	fmt.Printf("User Ethereum Address: %s\n", userAddress.Address.Hex())

	// Start monitoring for incoming ETH payments to the user's address
	monitorIncomingETH(client, userAddress.Address)
}

func monitorIncomingETH(client *ethclient.Client, address common.Address) {
	ctx := context.Background()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				balance, err := client.BalanceAt(ctx, address, nil)
				if err != nil {
					log.Fatalf("Failed to get balance: %v", err)
				}

				fmt.Printf("Balance of address %s: %s ETH\n", address.Hex(), balance)

				if balance.Cmp(big.NewInt(5)) >= 0 {
					fmt.Println("Received at least 5 ETH. Payment complete.")
					return
				}

				time.Sleep(30 * time.Second) // Check balance every 30 seconds
			}
		}
	}()

	// Keep the monitoring routine running
	select {}
}
