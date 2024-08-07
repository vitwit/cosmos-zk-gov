package circuit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	bn254 "github.com/consensys/gnark/backend/groth16/bn254"
	"github.com/consensys/gnark/backend/witness"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

// bytes to zk proof
func UnMarshalZkProof(zkProofBytes []byte) (groth16.Proof, error) {

	// unmarshal sig into proof
	zkProof := new(bn254.Proof)
	err := json.Unmarshal(zkProofBytes, zkProof)
	if err != nil {
		return nil, err
	}

	return zkProof, nil
}

// zkproof to bytes
func MarshalZkProof(zkproof groth16.Proof) ([]byte, error) {
	proofbn254 := zkproof.(*bn254.Proof)
	proofBytes, err := json.Marshal(proofbn254)
	if err != nil {
		return nil, err
	}

	return proofBytes, nil
}

// public witness from assignment
func GetPublicWitness(assignment *PrivateVotingCircuit) (witness.Witness, error) {
	witness, _ := frontend.NewWitness(assignment, ecc.BN254.ScalarField())
	return witness.Public()
}

// generate for for the assignment. Pk, Cs are fetching from predefined files
func GenerateProof(assignment *PrivateVotingCircuit) ([]byte, error) {

	merkleproofSize := len(assignment.MerkleProof.Path)

	pk, err := FetchProver(merkleproofSize)
	if err != nil {
		return nil, err
	}

	cs, err := FetchCs(merkleproofSize)
	if err != nil {
		return nil, err
	}

	witness, _ := frontend.NewWitness(assignment, ecc.BN254.ScalarField())

	zkproof, err := groth16.Prove(cs, pk, witness)

	if err != nil {
		return nil, err
	}
	return MarshalZkProof(zkproof)
}

// FetchProver fetches prover key from stored files
func FetchProver(size int) (*bn254.ProvingKey, error) {
	provingKeyBytes, err := os.ReadFile(filepath.Join("keys", fmt.Sprintf("prover-%d", size)))
	if err != nil {
		return nil, err
	}

	provingKey := new(bn254.ProvingKey)
	_, err = provingKey.ReadFrom(bytes.NewBuffer(provingKeyBytes))
	if err != nil {
		return nil, err
	}

	return provingKey, nil
}

// to fetch contraint system
func CompileCircuit(size int) constraint.ConstraintSystem {
	var circuit PrivateVotingCircuit
	circuit.MerkleProof.Path = make([]frontend.Variable, size)

	ccs, _ := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	return ccs
}

// instant setup for testing
func FetchKeys(size int) (groth16.ProvingKey, groth16.VerifyingKey, constraint.ConstraintSystem) {
	cs := CompileCircuit(size)
	pk, vk, err := groth16.Setup(cs)
	if err != nil {
		panic("error while setup: " + err.Error())
	}
	return pk, vk, cs

}

// FetchCs fetches constraint system from stored files
func FetchCs(size int) (constraint.ConstraintSystem, error) {
	csBytes, err := os.ReadFile(filepath.Join("keys", fmt.Sprintf("ccs-%d", size)))
	if err != nil {
		return nil, err
	}

	cs := groth16.NewCS(ecc.BN254)
	_, err = cs.ReadFrom(bytes.NewBuffer(csBytes))
	if err != nil {
		return nil, err
	}

	return cs, nil
}

// fetch verifier key from store file
func FetchVerifier(size int) (*bn254.VerifyingKey, error) {
	verifierKeyBytes, err := os.ReadFile(filepath.Join("keys", fmt.Sprintf("verifier-%d", size)))
	if err != nil {
		return nil, err
	}

	verifierKey := new(bn254.VerifyingKey)
	_, err = verifierKey.ReadFrom(bytes.NewBuffer(verifierKeyBytes))
	if err != nil {
		return nil, err
	}

	return verifierKey, nil
}

// complete test for an assignment before broadcasting the tx
func TestZKProof(assignment *PrivateVotingCircuit) {
	size := len(assignment.MerkleProof.Path)

	zkproofBytes, err := GenerateProof(assignment)
	if err != nil {
		panic("error while generating zk proo" + err.Error())
	}

	zkproof, err := UnMarshalZkProof(zkproofBytes)
	if err != nil {
		panic("error while unmarshalling zkproof" + err.Error())
	}

	pubWitness, err := GetPublicWitness(assignment)
	if err != nil {
		panic(err)
	}

	vk, err := FetchVerifier(size)
	if err != nil {
		panic(err)
	}

	groth16.Verify(zkproof, vk, pubWitness)
}
