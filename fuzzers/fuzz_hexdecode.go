package myfuzz
import (
	"github.com/binance-chain/go-sdk/types/msg"
)

func Fuzz(data []byte) int {
	msg.HexDecode(string(data))
	return 1
}
