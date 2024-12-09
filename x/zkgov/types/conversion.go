package types

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

func HexStringToBytes(data string) ([]byte, error) {
	bytes, err := hex.DecodeString(data)
	if err != nil {
		return nil, fmt.Errorf("invalid hex string: %v", err)
	}
	return bytes, nil
}

func BytesToHexString(data []byte) string {
	return hex.EncodeToString(data)
}

func MarshalVoteOption(voteOption VoteOption) []byte {
	vote := uint64(OPTION_NO)
	if voteOption == VoteOption_VOTE_OPTION_YES {
		vote = uint64(OPTION_YES)
	}

	voteOptionBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(voteOptionBytes, vote)

	return voteOptionBytes
}

func UnMarshalVoteOption(voteOptionBytes []byte) VoteOption {
	vote := binary.BigEndian.Uint64(voteOptionBytes)
	voteOption := VoteOption_VOTE_OPTION_NO
	if vote == OPTION_YES {
		voteOption = VoteOption_VOTE_OPTION_YES
	}
	return voteOption
}

func VoteOptionToInt(voteOption VoteOption) int64 {
	if voteOption == VoteOption_VOTE_OPTION_YES {
		return OPTION_YES
	}
	return OPTION_NO
}

func IntToVoteOption(voteOption int64) VoteOption {
	if voteOption == OPTION_YES {
		return VoteOption_VOTE_OPTION_YES
	}

	return VoteOption_VOTE_OPTION_NO
}

func StringToVoteOption(voteOptionString string) (int64, error) {
	if voteOptionString == "YES" {
		return OPTION_YES, nil
	}
	if voteOptionString == "NO" {
		return OPTION_NO, nil
	}

	return -1, fmt.Errorf("invalid vote option, choose YES or NO!")
}
