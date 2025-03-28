package account

import (
	"context"
	"github.com/segmentio/ksuid"
)

type Service interface {
	PostAccounts(ctx context.Context, name string) (*Account, error)
	GetAccount(ctx context.Context, id string) (*Account, error)
	GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)
}

type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type accountService struct {
	repository Repository
}

func (aS accountService) PostAccounts(ctx context.Context, name string) (*Account, error) {
	a := &Account{
		Name: name,
		ID:   ksuid.New().String(),
	}
	if err := aS.repository.PutAccount(ctx, *a); err != nil {
		return nil, err
	}
	return a, nil
}

func (aS accountService) GetAccount(ctx context.Context, id string) (*Account, error) {
	accountById, err := aS.repository.GetAccountById(ctx, id)
	if err != nil {
		return nil, err
	}
	return accountById, nil

}

func (a accountService) GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}
	accounts, err := a.repository.ListAccounts(ctx, skip, take)
	if err != nil {
		return nil, err
	}
	return accounts, nil

}

func newService(r Repository) Service {
	return &accountService{r}
}
