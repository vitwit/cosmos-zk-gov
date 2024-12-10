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
	"github.com/vitwit/cosmos-zk-gov/x/zkgov/types"
)

type DispatcherServer struct {
	ClientCtx *client.Context
	Cmd       *cobra.Command
}

func NewDispatcher(ctx *client.Context, cmd *cobra.Command) DispatcherServer {
	return DispatcherServer{
		ClientCtx: ctx,
		Cmd:       cmd,
	}
}

func (dispatcher *DispatcherServer) broadCastTransactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		dispatcherAddress := dispatcher.ClientCtx.GetFromAddress().String()

		var msg types.MsgVoteProposal
		err = msg.Unmarshal(body)
		msg.Sender = dispatcherAddress

		if err != nil {
			slog.Error("unable to unmarshal transaction")
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Process the transaction (this is just an example)
		slog.Info(
			fmt.Sprintf("Received vote transaction: proposal_ID=%d, vote_option=%d\n", msg.ProposalId, msg.VoteOption),
		)

		err = tx.GenerateOrBroadcastTxCLI(*dispatcher.ClientCtx, dispatcher.Cmd.Flags(), &msg)
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

func (dispatcher *DispatcherServer) Run(port string) error {
	http.HandleFunc("/broadCastTransaction", dispatcher.broadCastTransactionHandler)
	slog.Info("Starting dispatcher server on :" + port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		return fmt.Errorf("Could not start dispatcher: %s\n", err)
	}
	return nil
}
