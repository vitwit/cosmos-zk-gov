package main

import (
	"bytes"
	"fmt"
	"path/filepath"

	"log/slog"
	"os"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	zkcircuit "github.com/vitwit/cosmos-zk-gov/x/zkgov/circuit"
)

func GenerateZKKeys(merkleProofSize int) {

	slog.Info(
		fmt.Sprintf(
			"Generating prover key, verifier key and constraint system for circuit with merkle proof size: %d", merkleProofSize,
		),
	)
	var circuit zkcircuit.PrivateVotingCircuit
	circuit.MerkleProof.Path = make([]frontend.Variable, merkleProofSize)
	ccs, _ := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	pk, vk, _ := groth16.Setup(ccs)

	vkBuf, pkBuf, ccsBuf := new(bytes.Buffer), new(bytes.Buffer), new(bytes.Buffer)
	pk.WriteTo(pkBuf)
	vk.WriteTo(vkBuf)
	ccs.WriteTo(ccsBuf)

	proverKeyfileName := fmt.Sprintf("./keys/prover-%d", merkleProofSize)
	WriteToFile(proverKeyfileName, pkBuf)

	verifierfileName := fmt.Sprintf("./keys/verifier-%d", merkleProofSize)
	WriteToFile(verifierfileName, vkBuf)

	ccsfilename := fmt.Sprintf("./keys/ccs-%d", merkleProofSize)
	WriteToFile(ccsfilename, ccsBuf)

	slog.Info("Keys are successfully generated and stored in {keys} folder")

}

func WriteToFile(filename string, dataBuf *bytes.Buffer) {
	dir := filepath.Dir(filename)

	// Create the directory if it doesn't exist
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		slog.Error("Failed to create directory for storing keys")
	}
	err = os.WriteFile(filename, dataBuf.Bytes(), 0666)
	if err != nil {
		slog.Error("Failed to create directory for storing keys")
	}
}
