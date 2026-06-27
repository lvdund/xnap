package ies

import (
	"github.com/lvdund/asn1go/per"
)

var sCPACRequestInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "s-CPAC-Security-Config-List"},
		{Name: "s-CPAC-MultiTargetSN-List", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SCPACRequestInfo struct {
	SCPACSecurityConfigList SCPACSecurityConfigList
	SCPACMultiTargetSNList  *SCPACMultiTargetSNList
	IEExtensions            []byte
}

func (ie *SCPACRequestInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sCPACRequestInfoConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.SCPACMultiTargetSNList != nil, false}); err != nil {
		return err
	}
	if err := ie.SCPACSecurityConfigList.Encode(e); err != nil {
		return err
	}
	if ie.SCPACMultiTargetSNList != nil {
		if err := ie.SCPACMultiTargetSNList.Encode(e); err != nil {
			return err
		}
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SCPACRequestInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sCPACRequestInfoConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.SCPACSecurityConfigList.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.SCPACMultiTargetSNList = new(SCPACMultiTargetSNList)
		if err := ie.SCPACMultiTargetSNList.Decode(d); err != nil {
			return err
		}
	}
	extBytes, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	if len(extBytes) > 0 {
		ie.IEExtensions = extBytes[0]
	}
	return nil
}
