package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ignoxx/sl3/checkout/eth"
	"github.com/ignoxx/sl3/checkout/types"
	"github.com/ignoxx/sl3/checkout/utils"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	env := getEnvironment()

	fmt.Printf("Starting checkout server with build version %s\n", env.BuildVersion())

	// store := storage.NewMemoryStorage()
	ethClient, err := eth.NewEth(env.RpcURL, env.Mnemonic)
	if err != nil {
		log.Fatalf("Failed to create Ethereum client: %v", err)
	}
	// api := api.NewApi(env.WebServerPort, store)
	// go api.Start()

	account, err := ethClient.GetAccountForOrder(0)
	if err != nil {
		log.Fatalf("Failed to get account for order: %v", err)
	}

	fmt.Printf("User Ethereum Address: %s\n", account.Address.Hex())

	// Start monitoring for incoming ETH payments to the user's address
	monitorIncomingETH(ethClient, account.Address)
}

func monitorIncomingETH(client *eth.EthClient, address common.Address) {
	ctx := context.Background()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				balance, err := client.GetBalanceInETH(ctx, address)
				if err != nil {
					log.Fatalf("Failed to get balance: %v", err)
				}

				fmt.Printf("Balance of address %s: %f ETH\n", address.Hex(), balance)

                if balance >= .5 {
					fmt.Println("Received at least 0.5 ETH. Payment complete.")
					return
				}

				time.Sleep(15 * time.Second) // Check balance every 30 seconds
			}
		}
	}()

	// Keep the monitoring routine running
	select {}
}

func getEnvironment() types.Environment {
	webServerPort, ok := os.LookupEnv("PORT")
	if !ok {
		panic("PORT environment variable not set")
	}

	env, ok := os.LookupEnv("CE_ENV")
	if !ok {
		log.Println("CE_ENV environment variable not set. Defaulting to dev.")
		env = "dev"
	}

	revision, time, modified, err := utils.GetBuildInfo()
	if err != nil {
		panic("Failed to get build info: " + err.Error())
	}

	rpcURL, ok := os.LookupEnv("RPC_URL")
	if !ok {
		panic("RPC_URL environment variable not set")
	}

	mnemonic, ok := os.LookupEnv("MNEMONIC")
	if !ok {
		panic("MNEMONIC environment variable not set")
	}

	return types.Environment{
		WebServerPort: webServerPort,
		RpcURL:        rpcURL,
		Mnemonic:      mnemonic,
		Env:           env,
		CommitHash:    revision,
		CommitTime:    time,
		Modified:      modified == "true",
	}
}
