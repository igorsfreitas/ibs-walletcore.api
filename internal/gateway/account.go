package gateway

import "github.com/igorsfreitas/ibs-walletcore.api/internal/entity"

type AccountGateway interface {
	GetByID(id string) (*entity.Account, error)
	Save(account *entity.Account) error
}
