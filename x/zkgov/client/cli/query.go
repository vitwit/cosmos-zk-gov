package cli

import (
	"strconv"

	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   types.ModuleName,
		Short: "Querying commands for the zkgov module",
		RunE:  client.ValidateCmd,
	}

	cmd.AddCommand(
		GetProposalAllInfo(),
		GetAllProposal(),
	)

	return cmd
}

func GetProposalAllInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-proposal-info [proposal_id]",
		Short: "Show all the info related to the proposal including votes",
		Long: `Show all the info related to the proposal including votes,
		`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			proposalID, _ := strconv.Atoi(args[0])

			req := &types.QueryProposalAllInfoRequest{ProposalId: uint64(proposalID)}
			res, _ := queryClient.ProposalAllInfo(cmd.Context(), req)

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetAllProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-proposals [proposal_id]",
		Short: "Show all proposals",
		Long: `Show all proposals,
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			req := &types.GetProposalRequest{}
			res, _ := queryClient.GetProposals(cmd.Context(), req)

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
