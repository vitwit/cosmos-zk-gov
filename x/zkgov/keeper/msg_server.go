package keeper

import (
	"context"

	"github.com/vitwit/cosmos-zk-gov/x/zkgov/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(k Keeper) types.MsgServer {
	return &msgServer{
		Keeper: k,
	}
}

// register users who wants to vote for a certain proposal
// a user can register only once for one proposal. Similar;y, he can vote only once.
func (k msgServer) RegisterUser(ctx context.Context, req *types.MsgRegisterUser) (*types.MsgRegisterUserResponse, error) {
	err := req.ValidateBasic()
	if err != nil {
		return nil, err
	}

	err = k.Keeper.RegisterUser(ctx, req.Commitment, req.Sender, req.ProposalId)
	if err != nil {
		return nil, err
	}
	return &types.MsgRegisterUserResponse{}, nil
}

// generate userId it should in seq
// generate the random number to get the nullifier
func (k msgServer) VoteProposal(ctx context.Context, req *types.MsgVoteProposal) (*types.MsgVoteProposalResponse, error) {

	if err := req.ValidateBasic(); err != nil {
		return nil, err
	}

	if err := k.Keeper.Vote(ctx, *req); err != nil {
		return nil, err
	}

	return &types.MsgVoteProposalResponse{}, nil
}

func (k msgServer) CreateProposal(ctx context.Context, req *types.MsgCreateProposal) (*types.MsgCreateProposalResponse, error) {

	proposalID, err := k.Keeper.CreatePropsal(ctx, *req)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateProposalResponse{
		ProposalId: proposalID,
	}, nil

}

func (k msgServer) ProcessProposal(ctx context.Context, req *types.MsgProcessProposal) (*types.MsgProcessProposalResponse, error) {

	// TODO: calculate the result

	return &types.MsgProcessProposalResponse{}, nil
}
