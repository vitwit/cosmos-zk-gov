package circuit

import (
	"bytes"
	"fmt"
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/accumulator/merkletree"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/test"
	"github.com/vitwit/cosmos-zk-gov/x/zkgov/store"
)

func TestReadingKeys(t *testing.T) {
	assert := test.NewAssert(t)

	curve := ecc.BN254

	// the size of bytes should be of power of 2
	size := 3
	commitmentIndex := 3
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

	hashFunc.Reset()
	merkleRootBytes, merkleProofBytes, numLeaves, err := merkletree.BuildReaderProof(&buf, hashFunc, dataSize, uint64(commitmentIndex))
	assert.NoError(err)

	merkleProof := store.GetMerkleProofFromBytes(merkleRootBytes, merkleProofBytes)
	// index should be 4
	hashFunc.Reset()
	verified := merkletree.VerifyProof(hashFunc, merkleRootBytes, merkleProofBytes, 0, numLeaves)
	assert.Equal(verified, false)
	hashFunc.Reset()
	verified = merkletree.VerifyProof(hashFunc, merkleRootBytes, merkleProofBytes, uint64(commitmentIndex), numLeaves)
	assert.Equal(verified, true)

	circuit := PrivateVotingCircuit{}
	merkleproofSize := len(merkleProofBytes)
	GenerateZKKeys(merkleproofSize)

	circuit.MerkleProof.Path = make([]frontend.Variable, merkleproofSize)
	assignment := PrivateVotingCircuit{
		SecretUniqueId1: randomSecret1,
		SecretUniqueId2: randomSecret2,
		VoteOption:      voteOption,
		Commitment:      commitment,
		Nullifier:       nullifier,
		CommitmentIndex: uint64(commitmentIndex),
		MerkleProof:     merkleProof,
		MerkleRoot:      merkleRootBytes,
	}

	witness, _ := frontend.NewWitness(&assignment, curve.ScalarField())
	pubwitness, _ := witness.Public()

	pk, err := FetchProver(merkleproofSize)
	assert.NoError(err)

	vk, err := FetchVerifier(merkleproofSize)
	assert.NoError(err)

	cs, err := FetchCs(merkleproofSize)
	assert.NoError(err)

	proof, err := groth16.Prove(cs, pk, witness)
	groth16.Verify(proof, vk, pubwitness)
}
