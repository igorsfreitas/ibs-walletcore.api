package gateway

import "github.com/igorsfreitas/ibs-walletcore.api/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
