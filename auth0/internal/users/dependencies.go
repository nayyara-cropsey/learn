package users

import (
	"context"

	"github.com/auth0/go-auth0/management"
)

type manager interface {
	List(ctx context.Context, opts ...management.RequestOption) (*management.UserList, error)
	Update(ctx context.Context, id string, u *management.User, opts ...management.RequestOption) error
}
