package types

import (
	"encoding/binary"
)

const (
	// ModuleName defines the name of the nft module
	ModuleName = "zk-gov"

	// StoreKey is the default store key for nft
	StoreKey = ModuleName

	// RouterKey is the message route for nft
	RouterKey = ModuleName
)

var (
	CommitmentsKey     = []byte{0x01}
	UsersKey           = []byte{0x02}
	MerkleRootKey      = []byte{0x03}
	NullifiersKey      = []byte{0x04}
	ProposalCounterKey = []byte{0x05}
	ProposalInfoKey    = []byte{0x06}
	ProposalResultKey  = []byte{0x07}
	VotesKey           = []byte{0x08}
)

func ProposalInfoStoreKey(proposalID uint64) []byte {
	proposalIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(proposalIDBytes, proposalID)

	key := append(ProposalInfoKey, proposalIDBytes...)

	return key
}

func ProposalResultStoreKey(proposalID uint64) []byte {
	proposalIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(proposalIDBytes, proposalID)

	key := append(ProposalResultKey, proposalIDBytes...)

	return key
}

func CommitmentsStoreKey(proposalID uint64) []byte {
	proposalIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(proposalIDBytes, proposalID)

	key := append(CommitmentsKey, proposalIDBytes...)

	return key
}

func NullifiersStoreKey(proposalID uint64) []byte {
	proposalIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(proposalIDBytes, proposalID)

	key := append(NullifiersKey, proposalIDBytes...)

	return key
}

func UsersStoreKey(proposalID uint64) []byte {

	proposalIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(proposalIDBytes, proposalID)

	key := append(UsersKey, proposalIDBytes...)

	return key
}

func MerkleRootStoreKey(proposalID uint64) []byte {
	proposalIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(proposalIDBytes, proposalID)
	key := append(MerkleRootKey, proposalIDBytes...)

	return key
}

func VotesStoreKey(proposalID uint64) []byte {
	proposalIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(proposalIDBytes, proposalID)
	key := append(VotesKey, proposalIDBytes...)

	return key
}
