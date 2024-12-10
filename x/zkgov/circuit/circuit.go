package circuit

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/witness"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/accumulator/merkle"
	"github.com/consensys/gnark/std/hash/mimc"
	"github.com/vitwit/cosmos-zk-gov/x/zkgov/types"
)

// Define the circuit
type PrivateVotingCircuit struct {
	SecretUniqueId1 frontend.Variable // randomly generated
	SecretUniqueId2 frontend.Variable // randomly generated

	Commitment frontend.Variable //  hash(secret1 + secret2 + voteOption)
	Nullifier  frontend.Variable `gnark:",public"` // hash(secret2 + voteOption)
	VoteOption frontend.Variable `gnark:",public"`

	MerkleProof     merkle.MerkleProof
	CommitmentIndex frontend.Variable
	MerkleRoot      frontend.Variable `gnark:",public"`
}

func (circuit *PrivateVotingCircuit) Define(api frontend.API) error {

	mimc, _ := mimc.NewMiMC(api)

	circuit.checkCommitment(api, mimc)
	circuit.checkNullifier(api, mimc)
	circuit.checkMerkleProof(api, mimc)

	return nil
}

func (circuit *PrivateVotingCircuit) checkCommitment(api frontend.API, hFunc mimc.MiMC) error {

	hFunc.Reset()
	hFunc.Write(circuit.SecretUniqueId1)
	hFunc.Write(circuit.SecretUniqueId2)
	hFunc.Write(circuit.VoteOption)

	computedCommitment := hFunc.Sum()
	api.AssertIsEqual(circuit.Commitment, computedCommitment)

	return nil
}

func (circuit *PrivateVotingCircuit) checkMerkleProof(api frontend.API, hFunc mimc.MiMC) error {
	hFunc.Reset()
	api.AssertIsEqual(circuit.Commitment, circuit.MerkleProof.Path[0])
	api.AssertIsEqual(circuit.MerkleRoot, circuit.MerkleProof.RootHash)
	circuit.MerkleProof.VerifyProof(api, &hFunc, circuit.CommitmentIndex)

	return nil
}

func (circuit *PrivateVotingCircuit) checkNullifier(api frontend.API, hFunc mimc.MiMC) error {

	hFunc.Reset()
	hFunc.Write(circuit.SecretUniqueId2)
	hFunc.Write(circuit.VoteOption)

	computedNullifier := hFunc.Sum()
	api.AssertIsEqual(circuit.Nullifier, computedNullifier)

	return nil
}

func PreparePublicWitness(nullifier string, voteOptionInt int64, merkleRoot string) witness.Witness {
	nullifierBytes, _ := types.HexStringToBytes(nullifier)
	merkleRootBytes, _ := types.HexStringToBytes(merkleRoot)

	voteOption := *big.NewInt(voteOptionInt)

	field := ecc.BN254.ScalarField()
	publicWitness, err := ConstructWitness(field, nullifierBytes, voteOption.Bytes(), merkleRootBytes)
	if err != nil {
		panic(err)
	}

	return publicWitness

}

// constructs new public witness using assignment's public inputs
func ConstructWitness(field *big.Int, nullifierBytes []byte, voteOption []byte, merkleRootBytes []byte) (witness.Witness, error) {
	newWitness, err := witness.New(field)
	if err != nil {
		return nil, err
	}

	witnessChan := make(chan any)
	go passPubInputs(&witnessChan, nullifierBytes, voteOption, merkleRootBytes)
	newWitness.Fill(3, 0, witnessChan)

	return newWitness, nil
}

// close the channel after passing the values to end the for loop over channel values
func passPubInputs(witnessChan *chan any, nullifierBytes []byte, voteOption []byte, merkleRootBytes []byte) {
	*witnessChan <- nullifierBytes
	*witnessChan <- voteOption
	*witnessChan <- merkleRootBytes

	fmt.Println("pulbic values sent via channel for witness construction...")
	close(*witnessChan)
}

func Sha256Hash(message []byte) []byte {
	shaFunc := sha256.New()
	shaFunc.Write(message)
	return shaFunc.Sum(nil)
}
