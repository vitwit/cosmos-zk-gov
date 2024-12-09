package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/types"
)

type RelayerServer struct {
	ClientCtx *client.Context
	Cmd       *cobra.Command
}

func NewRelayer(ctx *client.Context, cmd *cobra.Command) RelayerServer {
	return RelayerServer{
		ClientCtx: ctx,
		Cmd:       cmd,
	}
}

func (relayer *RelayerServer) broadCastTransactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		relayerAddress := relayer.ClientCtx.GetFromAddress().String()

		var msg types.MsgVoteProposal
		err = msg.Unmarshal(body)
		msg.Sender = relayerAddress

		if err != nil {
			slog.Error("unable to unmarshal transaction")
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Process the transaction (this is just an example)
		slog.Info(
			fmt.Sprintf("Received vote transaction: proposal_ID=%d, vote_option=%d\n", msg.ProposalId, msg.VoteOption),
		)

		err = tx.GenerateOrBroadcastTxCLI(*relayer.ClientCtx, relayer.Cmd.Flags(), &msg)
		if err != nil {
			slog.Error("Error while broadcasting vote proposal", err.Error())
			http.Error(w, "Broadcast error ", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"status": "success"}
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			slog.Error("Error while writing response", err.Error())
			http.Error(w, "Error while writing response ", http.StatusInternalServerError)
			return
		}

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}

func (relayer *RelayerServer) Run(port string) error {
	http.HandleFunc("/broadCastTransaction", relayer.broadCastTransactionHandler)
	slog.Info("Starting Relayer server on :" + port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		return fmt.Errorf("Could not start relayer: %s\n", err)
	}
	return nil
}
