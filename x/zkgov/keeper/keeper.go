package keeper

import (
	"context"

	cosmosstore "cosmossdk.io/core/store"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/vitwit/cosmos-zk-gov/x/zkgov/circuit"
	storeImpl "github.com/vitwit/cosmos-zk-gov/x/zkgov/store"
	"github.com/vitwit/cosmos-zk-gov/x/zkgov/types"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

type Keeper struct {
	storeKey cosmosstore.KVStoreService
	cdc      codec.BinaryCodec
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey cosmosstore.KVStoreService,
) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}

func (k *Keeper) Vote(ctx context.Context, votePropal types.MsgVoteProposal) error {

	nullifier := votePropal.Nullifier
	proposalID := votePropal.ProposalId
	voteOption := votePropal.VoteOption

	zkProofBytes := votePropal.ZkProof

	store := k.storeKey.OpenKVStore(ctx)

	merkleRoot, err := storeImpl.GetMerkleRoot(ctx, store, proposalID)
	if err != nil {
		return err
	}

	err = storeImpl.StoreNullifier(ctx, store, proposalID, nullifier)
	if err != nil {
		return err
	}

	err = storeImpl.StoreVote(ctx, store, proposalID, voteOption)
	if err != nil {
		return err
	}

	publicWitness := circuit.PreparePublicWitness(nullifier, types.VoteOptionToInt(voteOption), merkleRoot)
	zkProof, err := circuit.UnMarshalZkProof(zkProofBytes[:])
	if err != nil {
		return err
	}

	// verifier key should be initialized at genesis
	vkey, err := circuit.FetchVerifier(int(votePropal.MerkleproofSize))
	if err != nil {
		return err
	}

	err = groth16.Verify(zkProof, vkey, publicWitness)
	if err != nil {
		return err
	}

	return nil
}

func (k *Keeper) RegisterUser(ctx context.Context, commitment string, user string, proposalID uint64) error {
	store := k.storeKey.OpenKVStore(ctx)
	if err := storeImpl.StoreUser(ctx, store, proposalID, user); err != nil {
		return err
	}

	if err := storeImpl.StoreCommitment(ctx, store, proposalID, commitment); err != nil {
		return err
	}

	return nil
}

func (k *Keeper) CreatePropsal(ctx context.Context, proposal types.MsgCreateProposal) (uint64, error) {
	store := k.storeKey.OpenKVStore(ctx)

	proposalID, err := storeImpl.StoreProposal(ctx, store, proposal)
	if err != nil {
		return proposalID, err
	}

	err = storeImpl.InitCommitments(ctx, store, proposalID)
	if err != nil {
		return proposalID, err
	}

	err = storeImpl.InitMerkleRoot(ctx, store, proposalID)
	if err != nil {
		return proposalID, err
	}

	err = storeImpl.InitNullifiers(ctx, store, proposalID)
	if err != nil {
		return proposalID, err
	}

	err = storeImpl.InitUsers(ctx, store, proposalID)
	if err != nil {
		return proposalID, err
	}

	err = storeImpl.InitVotes(ctx, store, proposalID)

	return proposalID, err

}

/* ------------------------- Queries ------------------------------*/

func (k *Keeper) MerkleProof(ctx context.Context, req *types.QueryCommitmentMerkleProofRequest) (*types.QueryCommitmentMerkleProofResponse, error) {
	store := k.storeKey.OpenKVStore(ctx)
	return storeImpl.GetMerkleProof(ctx, store, req)
}

func (k *Keeper) GetProposalAllInfo(ctx context.Context, req *types.QueryProposalAllInfoRequest) (*types.QueryProposalAllInfoResponse, error) {
	store := k.storeKey.OpenKVStore(ctx)

	// query the store votes
	votes, err := storeImpl.GetVotes(ctx, store, req.ProposalId)
	if err != nil {
		return nil, err
	}

	// query the nullifier
	nullifiers, err := storeImpl.GetNullifiers(ctx, store, req.ProposalId)
	if err != nil {
		return nil, err
	}

	votesInfo := GetVotesInfo(nullifiers, votes)

	return &types.QueryProposalAllInfoResponse{
		Votes: votesInfo,
	}, nil
}

func GetVotesInfo(nullifiers []string, votes []types.VoteOption) []*types.VoteInfo {
	VotesInfo := make([]*types.VoteInfo, len(nullifiers))
	for i, vote := range votes {
		voteInfo := &types.VoteInfo{
			Nullifer:   nullifiers[i],
			VoteOption: vote,
		}
		VotesInfo[i] = voteInfo
	}

	return VotesInfo
}

func GetUsersInfo(commitments []string, users []string) []*types.UserInfo {
	usersInfo := make([]*types.UserInfo, len(users))
	for i, user := range users {
		userInfo := &types.UserInfo{
			Commitment:  commitments[i],
			UserAddress: user,
		}
		usersInfo[i] = userInfo

	}

	return usersInfo
}

func (k *Keeper) GetAllProposals(ctx context.Context, req *types.GetProposalRequest) (*types.GetProposalsResponse, error) {
	store := runtime.KVStoreAdapter(k.storeKey.OpenKVStore(ctx))
	var lists []*types.Proposal
	iter := storetypes.KVStorePrefixIterator(store, types.ProposalInfoKey)
	for ; iter.Valid(); iter.Next() {
		var list types.Proposal
		k.cdc.Unmarshal(iter.Value(), &list)
		lists = append(lists, &list)
	}
	return &types.GetProposalsResponse{Proposals: lists}, nil
}
