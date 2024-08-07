package types

import (
	"bytes"
	"fmt"
	"os"

	"github.com/consensys/gnark/backend/groth16"
	bn254 "github.com/consensys/gnark/backend/groth16/bn254"
)

func ReaderVerifier(size int) (groth16.VerifyingKey, error) {

	fileName := VerifierKeyName(size)
	vkBytes, err := os.ReadFile(fileName)

	vk := new(bn254.VerifyingKey)
	_, err = vk.ReadFrom(bytes.NewBuffer(vkBytes))
	if err != nil {
		return nil, err
	}

	return vk, nil
}

func VerifierKeyName(size int) string {
	return fmt.Sprintf("../client/zk/keys/verifier-%d", size)
}
