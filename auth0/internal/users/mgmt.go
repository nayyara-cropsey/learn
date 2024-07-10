package users

import (
	"context"
	"fmt"

	"github.com/auth0/go-auth0/management"
)

const (
	defaultPageSize = 100
)

type UpdateFunc func(user *management.User) *management.User

type Management struct {
	mgr manager
}

func NewManagement(ctx context.Context, c Config) (*Management, error) {
	auth0API, err := management.New(
		c.Domain,
		management.WithClientCredentials(ctx, c.ClientID, c.ClientSecret),
	)

	if err != nil {
		return nil, fmt.Errorf("auth0 connection: %w", err)
	}

	return &Management{
		mgr: auth0API.User,
	}, nil
}

func (m *Management) Iterator(pageSize int) *Iterator {
	return &Iterator{
		mgr:      m.mgr,
		page:     0,
		pageSize: pageSize,
		hasNext:  true,
	}
}

func (m *Management) UpdateAll(ctx context.Context, updateFunc UpdateFunc) (int, error) {
	iterator := m.Iterator(defaultPageSize)

	var total int
	for iterator.HasNext() {
		usrs, err := iterator.Next(ctx)
		if err != nil {
			return total, fmt.Errorf("next: %w", err)
		}

		for _, usr := range usrs {
			u := updateFunc(usr)

			if err := m.mgr.Update(ctx, *usr.ID, u); err != nil {
				return total, fmt.Errorf("update: %w", err)
			}

			total++
		}
	}

	return total, nil
}
