package v1

import "context"

type StoreDomain struct {
	Id       string
	Name     string
	Address  string
	Contacts string
}

type StoreRepository interface {
	Create(ctx context.Context, inDom *StoreDomain) (err error)
	FindAll(ctx context.Context) (outDom []StoreDomain)
	FindById(ctx context.Context, id string) (outDom StoreDomain, err error)
	UpdateById(ctx context.Context, inDom *StoreDomain) (err error)
	DeleteById(ctx context.Context, id string) (err error)
}
