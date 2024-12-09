package circuit

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
	"testing"

	"hash"

	"github.com/consensys/gnark-crypto/accumulator/merkletree"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/test"
	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/store"
)

func GetRandomCommitmentList(size uint64, hFunc hash.Hash) []byte {

	commitments := []byte{}
	for i := uint64(0); i < size; i++ {
		hFunc.Reset()
		numBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(numBytes, i)
		hFunc.Write(numBytes)
		numHash := hFunc.Sum(nil)
		commitments = append(commitments, numHash...)
	}

	return commitments
}
func TestPrivateVotingCircuit(t *testing.T) {
	assert := test.NewAssert(t)

	curve := ecc.BN254

	// the size of bytes should be of power of 2
	size := 7
	commitmentIndex := 7
	dataSize := 32

	hashFunc := mimc.NewMiMC()

	hashFunc.Reset()
	randomSecret1 := *big.NewInt(10)
	randomSecret2 := *big.NewInt(20)
	voteOption := *big.NewInt(2)

	hashFunc.Write(randomSecret1.Bytes())
	hashFunc.Write(randomSecret2.Bytes())
	hashFunc.Write(voteOption.Bytes())

	commitment := hashFunc.Sum(nil)
	commitments := GetRandomCommitmentList(uint64(size), hashFunc)
	commitments = append(commitments, commitment...)

	fmt.Println("data length", len(commitments))

	var buf bytes.Buffer
	buf.Write(commitments)

	hashFunc.Reset()

	hashFunc.Write(randomSecret2.Bytes())
	hashFunc.Write(voteOption.Bytes())

	nullifier := hashFunc.Sum(nil)

	fmt.Println("before")
	hashFunc.Reset()
	merkleRootBytes, merkleProofBytes, numLeaves, err := merkletree.BuildReaderProof(&buf, hashFunc, dataSize, uint64(commitmentIndex))
	assert.NoError(err)

	fmt.Println("after")
	merkleProof := store.GetMerkleProofFromBytes(merkleRootBytes, merkleProofBytes)
	// index should be 4
	hashFunc.Reset()
	verified := merkletree.VerifyProof(hashFunc, merkleRootBytes, merkleProofBytes, 0, numLeaves)
	assert.Equal(verified, false)
	hashFunc.Reset()
	verified = merkletree.VerifyProof(hashFunc, merkleRootBytes, merkleProofBytes, uint64(commitmentIndex), numLeaves)
	assert.Equal(verified, true)

	circuit := PrivateVotingCircuit{}
	circuit.MerkleProof.Path = make([]frontend.Variable, len(merkleProofBytes))
	witness := PrivateVotingCircuit{
		SecretUniqueId1: randomSecret1,
		SecretUniqueId2: randomSecret2,
		VoteOption:      voteOption,
		Commitment:      commitment,
		Nullifier:       nullifier,
		CommitmentIndex: uint64(commitmentIndex),
		MerkleProof:     merkleProof,
		MerkleRoot:      merkleRootBytes,
	}

	// err = test.IsSolved(&circuit, &witness, ecc.BN254.ScalarField())
	assert.ProverSucceeded(&circuit, &witness, test.WithCurves(curve))

}
