package client

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/types"
)

type RelayerClient struct {
	Address string
}

func NewRelayerClient(address string) *RelayerClient {
	return &RelayerClient{
		Address: address,
	}
}

func (relayerClient *RelayerClient) BroadCastTx(msg types.MsgVoteProposal) error {
	// Marshal the message to JSON
	msgBytes, err := msg.Marshal()
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	// Create a new POST request
	url := relayerClient.Address + "/broadCastTransaction"
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
