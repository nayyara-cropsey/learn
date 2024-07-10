package main

import (
	"context"
	"log"

	"github.com/auth0/go-auth0/management"

	"auth0/internal/users"
)

type user struct {
	UserID      string                 `json:"user_id"`
	AppMetadata map[string]interface{} `json:"app_metadata"`
}

const (
	gnqlKey = "gnqlKey"
)

func main() {
	ctx := context.Background()
	m, err := users.NewManagement(ctx, users.Config{
		Domain:       "greynoise2.auth0.com",
		ClientID:     "G6pNjo4UMAtPorMqx1GtbtoLnjvwHqvO",
		ClientSecret: "ZkEV-uhny6BZit_Uyz-xwBhAa1Z7HH8PH9A-_ZHhYbNbHzMfntUKrsIJ5QNta5KI",
	})
	if err != nil {
		log.Fatalf("failed to initialize the user management client: %+v", err)
	}

	total, err := m.UpdateAll(ctx, func(user *management.User) *management.User {
		a := make(map[string]interface{})
		if user.AppMetadata != nil {
			a = *user.AppMetadata
			if a[gnqlKey] != nil {
				a[gnqlKey] = nil
			}
		}

		log.Println("Updating user app metadata", *user.ID, *user.Email, a)
		return &management.User{
			AppMetadata: &a,
		}
	})
	if err != nil {
		log.Fatalf("failed to update users: %+v", err)
	}

	log.Printf("Total: %v", total)
}
