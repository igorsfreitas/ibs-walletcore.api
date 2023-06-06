package gateway

import "github.com/igorsfreitas/ibs-walletcore.api/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
