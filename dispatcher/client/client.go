package client

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/vitwit/cosmos-zk-gov/x/zkgov/types"
)

type DispatcherClient struct {
	Address string
}

func NewdispatcherClient(address string) *DispatcherClient {
	return &DispatcherClient{
		Address: address,
	}
}

func (dispatcherClient *DispatcherClient) BroadCastTx(msg types.MsgVoteProposal) error {
	// Marshal the message to JSON
	msgBytes, err := msg.Marshal()
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	// Create a new POST request
	url := dispatcherClient.Address + "/broadCastTransaction"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(msgBytes))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	client := &http.Client{}
	client.Timeout = 50 * time.Second

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to perform HTTP request: %w", err)
	}

	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK HTTP status: %s", resp.Status)
	}

	return nil
}
