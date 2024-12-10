package store

import (
	"bytes"
	"context"
	"errors"

	cosmosstore "cosmossdk.io/core/store"
	"github.com/vitwit/cosmos-zk-gov/x/zkgov/types"
)

func InitUsers(ctx context.Context, store cosmosstore.KVStore, proposalID uint64) error {
	usersKey := types.UsersStoreKey(proposalID)

	usersBytes := []byte{}

	return store.Set(usersKey, usersBytes)

}

func StoreUser(ctx context.Context, store cosmosstore.KVStore, proposalID uint64, user string) error {
	usersKey := types.UsersStoreKey(proposalID)
	usersBytes, err := store.Get(usersKey)
	if err != nil {
		return err
	}

	userBytes := []byte(user)

	// if user already stored, throw error
	for i := 0; i < len(usersBytes); i += types.USER_SIZE {
		storedUser := usersBytes[i : i+types.USER_SIZE]
		if bytes.Equal(storedUser, userBytes) {
			return errors.New("user is already registered")
		}
	}

	usersBytes = append(usersBytes, userBytes...)
	store.Set(usersKey, usersBytes)

	return nil
}

func GetUsers(ctx context.Context, store cosmosstore.KVStore, proposalID uint64) ([]string, error) {
	userStoreKey := types.UsersStoreKey(proposalID)
	users := make([]string, 0)

	storedUsers, err := store.Get(userStoreKey)
	if err != nil {
		return users, err
	}

	for i := 0; i < len(storedUsers); i += types.USER_SIZE {
		user := storedUsers[i : i+types.USER_SIZE]
		users = append(users, string(user))
	}

	return users, nil
}
