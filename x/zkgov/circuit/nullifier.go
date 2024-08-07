package circuit

import (
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
)

func CreateNullifier(secret2, voteOption int64) []byte {
	hashFunc := mimc.NewMiMC()
	hashFunc.Reset()

	randomSecret2 := *big.NewInt(secret2)
	vote := *big.NewInt(voteOption)

	hashFunc.Write(randomSecret2.Bytes())
	hashFunc.Write(vote.Bytes())

	return hashFunc.Sum(nil)
}
