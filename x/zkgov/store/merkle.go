package store

import (
	"bytes"
	"context"
	"errors"

	cosmosstore "cosmossdk.io/core/store"
	"github.com/consensys/gnark-crypto/accumulator/merkletree"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/accumulator/merkle"
	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/types"
)

func InitMerkleRoot(ctx context.Context, store cosmosstore.KVStore, proposalID uint64) error {
	merklerootKey := types.MerkleRootStoreKey(proposalID)

	merklerootBytes := []byte{}

	return store.Set(merklerootKey, merklerootBytes)

}

func UpdateMerkleRoot(ctx context.Context, store cosmosstore.KVStore, proposalID uint64, commitments []byte) error {
	merklerootKey := types.MerkleRootStoreKey(proposalID)

	var buf bytes.Buffer
	buf.Write(commitments)

	hFunc := mimc.NewMiMC()

	// TODO: find root directly
	merkleroot, _, _, err := merkletree.BuildReaderProof(&buf, hFunc, types.COMMITMENT_SIZE, uint64(0))
	if err != nil {
		return err
	}

	store.Set(merklerootKey, merkleroot)
	return nil
}

func GetMerkleProof(ctx context.Context, store cosmosstore.KVStore, req *types.QueryCommitmentMerkleProofRequest) (*types.QueryCommitmentMerkleProofResponse, error) {
	commitmentsKey := types.CommitmentsStoreKey(req.ProposalId)
	commitmentsBytes, err := store.Get(commitmentsKey)
	if err != nil {
		return nil, err
	}

	commitmentBytes, err := types.HexStringToBytes(req.Commitment)
	if err != nil {
		return nil, err
	}
	commitmentIndex := -1
	for i := 0; i < len(commitmentsBytes); i += types.COMMITMENT_SIZE {
		curCommitmentBytes := commitmentsBytes[i : i+types.COMMITMENT_SIZE]
		if bytes.Equal(curCommitmentBytes, commitmentBytes) {
			commitmentIndex = i / types.COMMITMENT_SIZE
			break
		}
	}

	if commitmentIndex == -1 {
		return nil, errors.New("Commitment is not registered")
	}

	var buf bytes.Buffer
	buf.Write(commitmentsBytes)

	hFunc := mimc.NewMiMC()

	root, merkleproof, _, err := merkletree.BuildReaderProof(&buf, hFunc, types.COMMITMENT_SIZE, uint64(commitmentIndex))
	if err != nil {
		return nil, err
	}

	return &types.QueryCommitmentMerkleProofResponse{
		MerkleProof:     merkleproof,
		Root:            root,
		CommitmentIndex: uint64(commitmentIndex),
	}, nil
}

func GetMerkleRoot(ctx context.Context, store cosmosstore.KVStore, proposalID uint64) (string, error) {
	merkleKey := types.MerkleRootStoreKey(proposalID)
	merkleRootBytes, err := store.Get(merkleKey)
	if err != nil {
		return "", err
	}
	return types.BytesToHexString(merkleRootBytes), nil
}

func GetMerkleProofFromBytes(rootBytes []byte, proofBytes [][]byte) merkle.MerkleProof {
	var merkleProof merkle.MerkleProof
	merkleProof.RootHash = rootBytes
	merkleProof.Path = make([]frontend.Variable, len(proofBytes))
	for i := 0; i < len(proofBytes); i++ {
		merkleProof.Path[i] = proofBytes[i]
	}
	return merkleProof
}
