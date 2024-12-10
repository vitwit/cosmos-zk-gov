package circuit

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/vitwit/cosmos-zk-gov/x/zkgov/types"
)

type VoterInfo struct {
	ProposalID    uint64 `json:"proposal_id"`
	Commitment    string `json:"commitment"`
	Nullifier     string `json:"nullifier"`
	RandomSecret1 uint64 `json:"random_secret1"`
	RandomSecret2 uint64 `json:"random_secret2"`
	VoteOption    uint64 `json:"vote_option"`
}

// SaveInfo saves the voter info as JSON
func SaveInfo(
	proposalID uint64,
	commitment []byte,
	nullifier []byte,
	voteOption uint64,
	randomSecret1 uint64,
	randomSecret2 uint64,
	sender string,
) error {
	commitmentString := types.BytesToHexString(commitment)
	nullifierString := types.BytesToHexString(nullifier)
	info := VoterInfo{
		ProposalID:    proposalID,
		Commitment:    commitmentString,
		Nullifier:     nullifierString,
		VoteOption:    voteOption,
		RandomSecret1: randomSecret1,
		RandomSecret2: randomSecret2,
	}
	data, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join("commitments", fmt.Sprintf("%s-%d.json", sender, proposalID)), data, 0666)
}

// FetchInfo fetches the voter info from JSON
func FetchInfo(proposalID string, sender string) (VoterInfo, error) {
	var info VoterInfo
	data, err := os.ReadFile(filepath.Join("commitments", fmt.Sprintf("%s-%s.json", sender, proposalID)))
	if err != nil {
		return info, err
	}
	err = json.Unmarshal(data, &info)

	return info, err
}
