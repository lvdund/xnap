package ies

import (
	"github.com/lvdund/asn1go/per"
)

var criticalityDiagnosticsIEItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "iECriticality"},
		{Name: "iE-ID"},
		{Name: "typeOfError"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CriticalityDiagnosticsIEItem struct {
	IECriticality Criticality
	IEID          ProtocolIEID
	TypeOfError   TypeOfError
	IEExtensions  []byte
}

func (ie *CriticalityDiagnosticsIEItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(criticalityDiagnosticsIEItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.IECriticality.Encode(e); err != nil {
		return err
	}
	if err := ie.IEID.Encode(e); err != nil {
		return err
	}
	if err := ie.TypeOfError.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *CriticalityDiagnosticsIEItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(criticalityDiagnosticsIEItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.IECriticality.Decode(d); err != nil {
		return err
	}
	if err := ie.IEID.Decode(d); err != nil {
		return err
	}
	if err := ie.TypeOfError.Decode(d); err != nil {
		return err
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
