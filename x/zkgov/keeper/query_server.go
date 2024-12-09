package keeper

import (
	"context"

	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) CommitmentMerkleProof(ctx context.Context, req *types.QueryCommitmentMerkleProofRequest) (*types.QueryCommitmentMerkleProofResponse, error) {
	return k.MerkleProof(ctx, req)
}

func (k Keeper) ProposalAllInfo(ctx context.Context, req *types.QueryProposalAllInfoRequest) (*types.QueryProposalAllInfoResponse, error) {
	return k.GetProposalAllInfo(ctx, req)
}

func (k Keeper) GetProposals(ctx context.Context, req *types.GetProposalRequest) (*types.GetProposalsResponse, error) {
	return k.GetAllProposals(ctx, req)
}
