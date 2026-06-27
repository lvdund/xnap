package utils

import (
	"encoding/hex"

	"github.com/lvdund/asn1go/per"
)

func BitStringToHex(bitString *per.BitString) (hexString string) {
	hexString = hex.EncodeToString(bitString.Bits)
	hexLen := (bitString.Length + 3) / 4
	hexString = hexString[:hexLen]
	return
}

func HexToBitString(hexString string, bitLength int) (bitString per.BitString) {
	hexLen := len(hexString)
	if hexLen != (bitLength+3)/4 {
		return
	}
	if hexLen%2 == 1 {
		hexString += "0"
	}
	if byteTmp, err := hex.DecodeString(hexString); err != nil {
		//TODO: print warning
	} else {
		bitString.Bits = byteTmp
	}
	bitString.Length = bitLength
	mask := byte(0xff)
	mask = mask << uint(8-bitLength%8)
	if mask != 0 {
		bitString.Bits[len(bitString.Bits)-1] &= mask
	}
	return
}

func ByteToBitString(byteArray []byte, bitLength int) (bitString per.BitString) {
	byteLen := (bitLength + 7) / 8
	if byteLen > len(byteArray) {
		return
	}
	bitString.Bits = byteArray
	bitString.Length = bitLength
	return
}
