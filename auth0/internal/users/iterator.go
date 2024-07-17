package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/auth0/go-auth0/management"
)

var (
	ErrNoItemsLeft = errors.New("no items left")
)

type Iterator struct {
	mgr      manager
	page     int
	pageSize int
	hasNext  bool
}

func (i *Iterator) HasNext() bool {
	return i.hasNext
}

func (i *Iterator) Next(ctx context.Context) ([]*management.User, error) {
	if !i.HasNext() {
		return nil, ErrNoItemsLeft
	}

	usrs, err := i.mgr.List(ctx, management.Page(i.page), management.PerPage(i.pageSize))
	if err != nil {
		return nil, fmt.Errorf("next: %w", err)
	}

	i.page++
	fmt.Println("page", i.page)
	i.hasNext = usrs.HasNext()

	return usrs.Users, nil
}
