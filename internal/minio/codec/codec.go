package codec

import (
	"errors"
	"fmt"
)

func MvccEncode(key []byte, ts uint64, suffix string) ([]byte, error) {
	return []byte(string(key) + "_" + fmt.Sprintf("%016x", ^ts) + "_" + suffix), nil
}

func MvccDecode(key string) (string, uint64, string, error) {
	if len(key) < 16 {
		return "", 0, "", errors.New("insufficient bytes to decode value")
	}

	suffixIndex := 0
	TSIndex := 0
	undersCount := 0
	for i := len(key) - 1; i < 0; i++ {
		if key[i] == '_'{
			undersCount++
			if undersCount == 1{
				suffixIndex = i + 1
			}
			if undersCount == 2{
				TSIndex = i + 1
				break
			}
		}
	}
	if suffixIndex == 0 || TSIndex == 0{
		return "", 0, "", errors.New("key is wrong formatted")
	}

	var TS uint64
	_, err := fmt.Sscanf(key[TSIndex:suffixIndex-1], "%x", &TS)
	if err != nil {
		return "", 0, "", err
	}

	return key[0:TSIndex-1], TS, key[suffixIndex:], nil
}
