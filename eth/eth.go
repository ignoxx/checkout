package eth

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

const (
    ONE_ETH_IN_WEI = 1_000_000_000_000_000_000
)

type EthClient struct {
	client *ethclient.Client
	wallet *hdwallet.Wallet
}

func NewEth(rpcURL string, mnemonic string) (*EthClient, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}

	return &EthClient{
		client: client,
		wallet: wallet,
	}, nil
}

func (e *EthClient) GetBalance(ctx context.Context, address common.Address) (float64, error) {
	balance, err := e.client.BalanceAt(ctx, address, nil)
	if err != nil {
		return 0, err
	}

	fbalance, _ := balance.Float64()
	return fbalance, nil
}

func (e *EthClient) GetBalanceInETH(ctx context.Context, address common.Address) (float64, error) {
    balance, err := e.GetBalance(ctx, address)
    if err != nil {
        return 0, err
    }

    return balance / ONE_ETH_IN_WEI, nil
}

func (e *EthClient) GetAccountForOrder(orderID int) (*accounts.Account, error) {
	// Derive a child Ethereum address for the order
	path := fmt.Sprintf("m/44'/60'/0'/0/%d", orderID)
	derivationPath, err := hdwallet.ParseDerivationPath(path)
	if err != nil {
		return nil, errors.Newf("failed to parse derivation path for path %s and orderID %d: %v", path, orderID, err)
	}

	orderAddress, err := e.wallet.Derive(derivationPath, false)

	if err != nil {
		return nil, err
	}

	return &orderAddress, nil
}
