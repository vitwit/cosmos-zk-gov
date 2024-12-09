package store

import (
	"bytes"
	"context"

	cosmosstore "cosmossdk.io/core/store"
	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/types"
)

func StoreCommitment(ctx context.Context, store cosmosstore.KVStore, proposalID uint64, commitment string) error {

	commitmentsKey := types.CommitmentsStoreKey(proposalID)
	commitmentsBytes, err := store.Get(commitmentsKey)
	if err != nil {
		return err
	}

	commitmentsBytes = RemoveDefaultCommitments(commitmentsBytes)

	commitmentBytes, err := types.HexStringToBytes(commitment)
	if err != nil {
		return err
	}

	commitmentsBytes = append(commitmentsBytes, commitmentBytes...)
	commitmentsBytes = ExtendTillPowerof2(commitmentsBytes)
	store.Set(commitmentsKey, commitmentsBytes)

	UpdateMerkleRoot(ctx, store, proposalID, commitmentsBytes)
	return nil
}

func GetCommitments(ctx context.Context, store cosmosstore.KVStore, proposalID uint64) ([]string, error) {
	commitmentStoreKey := types.CommitmentsStoreKey(proposalID)
	commitments := make([]string, 0)
	StoredCommitments, err := store.Get(commitmentStoreKey)
	if err != nil {
		return commitments, err
	}

	for i := 0; i < len(StoredCommitments); i += types.COMMITMENT_SIZE {
		commitmentBytes := StoredCommitments[i : i+types.COMMITMENT_SIZE]
		commitmentString := types.BytesToHexString(commitmentBytes)
		commitments = append(commitments, commitmentString)
	}
	return commitments, nil
}

func InitCommitments(ctx context.Context, store cosmosstore.KVStore, proposalID uint64) error {
	commitmentsKey := types.CommitmentsStoreKey(proposalID)

	commitmenstBytes := []byte{}

	return store.Set(commitmentsKey, commitmenstBytes)

}

func RemoveDefaultCommitments(commitmentsBytes []byte) []byte {
	filteredCommitmentBytes := []byte{}
	for i := 0; i < len(commitmentsBytes); i += types.COMMITMENT_SIZE {
		commitment := commitmentsBytes[i : i+types.COMMITMENT_SIZE]
		if bytes.Equal(commitment, DefaultCommitment()) {
			break
		}
		filteredCommitmentBytes = append(filteredCommitmentBytes, commitment...)
	}

	return filteredCommitmentBytes
}

func ExtendTillPowerof2(commitmentsBytes []byte) []byte {
	TotalSize := len(commitmentsBytes)
	commitmentsCount := TotalSize / types.COMMITMENT_SIZE

	for commitmentsCount < 2 || (commitmentsCount&(commitmentsCount-1)) > 0 {
		commitmentsBytes = append(commitmentsBytes, DefaultCommitment()...)
		TotalSize = len(commitmentsBytes)
		commitmentsCount = TotalSize / types.COMMITMENT_SIZE
	}

	return commitmentsBytes
}

func DefaultCommitment() []byte {
	defaultCommitment := make([]byte, 32)
	return defaultCommitment
}
