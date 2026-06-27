package utils

import (
	"encoding/hex"

	"github.com/lvdund/xnap/ies"
)

type Snssai struct {
	Sst int32
	Sd  string
}

func SNssaiToModels(ngapSnssai ies.SNSSAI) (modelsSnssai Snssai) {
	modelsSnssai.Sst = int32(ngapSnssai.Sst[0])
	if ngapSnssai.Sd != nil {
		modelsSnssai.Sd = hex.EncodeToString(ngapSnssai.Sd)
	}
	return
}

func SNssaiToNgap(modelsSnssai Snssai) ies.SNSSAI {
	var ngapSnssai ies.SNSSAI
	ngapSnssai.Sst = []byte{byte(modelsSnssai.Sst)}

	if modelsSnssai.Sd != "" {
		if sdTmp, err := hex.DecodeString(modelsSnssai.Sd); err != nil {
			return ngapSnssai
		} else {
			ngapSnssai.Sd = sdTmp
		}
	}
	return ngapSnssai
}
