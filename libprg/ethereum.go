package libprg

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// ethereumBackend identifies transaction preparation failures and retries
// transient RPC reads. Submission is attempted once because an RPC failure
// does not establish whether the node accepted a transaction.
type ethereumBackend struct {
	*ethclient.Client
}

func (b ethereumBackend) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return ethereumRead(ctx, "eth_getCode (pending)", 3, func() ([]byte, error) {
		return b.Client.PendingCodeAt(ctx, account)
	})
}

func (b ethereumBackend) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return ethereumRead(ctx, "eth_estimateGas", 3, func() (uint64, error) {
		return b.Client.EstimateGas(ctx, msg)
	})
}

func (b ethereumBackend) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return ethereumRead(ctx, "eth_getTransactionCount (pending)", 3, func() (uint64, error) {
		return b.Client.PendingNonceAt(ctx, account)
	})
}

func (b ethereumBackend) SendTransaction(ctx context.Context, tx *ethTypes.Transaction) error {
	if err := b.Client.SendTransaction(ctx, tx); err != nil {
		return fmt.Errorf("eth_sendRawTransaction for %s (check transaction status before retrying): %w", tx.Hash(), err)
	}
	return nil
}

func ethereumRead[T any](ctx context.Context, method string, attempts int, read func() (T, error)) (T, error) {
	var value T
	var err error
	for attempt := 1; attempt <= attempts; attempt++ {
		value, err = read()
		if err == nil {
			return value, nil
		}
		if !transientEthereumError(err) || attempt == attempts {
			return value, fmt.Errorf("Ethereum RPC %s failed after %d attempt(s): %w", method, attempt, err)
		}
		timer := time.NewTimer(time.Duration(attempt) * 500 * time.Millisecond)
		select {
		case <-ctx.Done():
			timer.Stop()
			return value, fmt.Errorf("Ethereum RPC %s: %w", method, ctx.Err())
		case <-timer.C:
		}
	}
	return value, fmt.Errorf("Ethereum RPC %s requires at least one attempt", method)
}

func transientEthereumError(err error) bool {
	var rpcErr rpc.Error
	if errors.As(err, &rpcErr) {
		return rpcErr.ErrorCode() == -32603 || rpcErr.ErrorCode() == -32005
	}
	var httpErr rpc.HTTPError
	if errors.As(err, &httpErr) {
		switch httpErr.StatusCode {
		case http.StatusTooManyRequests, http.StatusBadGateway, http.StatusServiceUnavailable, http.StatusGatewayTimeout:
			return true
		}
	}
	return false
}
