package store

import (
	"context"
	"encoding/binary"

	cosmosstore "cosmossdk.io/core/store"
	"github.com/vitwit/cosmos-zk-gov/x/zkgov/types"
)

func StoreProposal(ctx context.Context, store cosmosstore.KVStore, req types.MsgCreateProposal) (uint64, error) {

	proposalCounter := GetProposalCounter(ctx, store)
	proposalCounter++
	StoreProposalCounter(ctx, store, proposalCounter)

	proposalInfoStoreKey := types.ProposalInfoStoreKey(proposalCounter)
	proposal := types.Proposal{
		ProposalId:           proposalCounter,
		Title:                req.Title,
		Description:          req.Description,
		RegistrationDeadline: req.RegistrationDeadline,
		VotingDeadline:       req.VotingDeadline,
	}

	proposalBytes, err := proposal.Marshal()
	if err != nil {
		return 0, err
	}
	err = store.Set(proposalInfoStoreKey, proposalBytes)
	if err != nil {
		return 0, err
	}

	return proposalCounter, nil
}

func GetProposalCounter(ctx context.Context, store cosmosstore.KVStore) uint64 {
	proposalCounterKey := types.ProposalCounterKey
	var proposalCounterBytes []byte
	if found, err := store.Has(proposalCounterKey); !found || err != nil {
		zero := 0
		zeroBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(zeroBytes, uint64(zero))
		proposalCounterBytes = zeroBytes
		store.Set(proposalCounterKey, proposalCounterBytes)
	}

	proposalCounterBytes, _ = store.Get(proposalCounterKey)

	proposalCounter := binary.BigEndian.Uint64(proposalCounterBytes)

	return proposalCounter
}

func StoreProposalCounter(ctx context.Context, store cosmosstore.KVStore, proposalCounter uint64) error {
	proposalCounterKey := types.ProposalCounterKey

	proposalCounterBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(proposalCounterBytes, proposalCounter)

	return store.Set(proposalCounterKey, proposalCounterBytes)
}

// Get the stored proposals
func GetProposal(ctx context.Context, store cosmosstore.KVStore) (*types.MsgCreateProposal, error) {
	var proposal types.MsgCreateProposal
	proposalInfo, err := store.Get(types.ProposalInfoKey)
	if err != nil {
		return &types.MsgCreateProposal{}, err
	}
	err = proposal.Unmarshal(proposalInfo)
	if err != nil {
		return &types.MsgCreateProposal{}, err
	}

	return &proposal, nil
}
