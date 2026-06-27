package ies

import (
	"github.com/lvdund/asn1go/per"
)

var globalCellIDConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-id"},
		{Name: "cell-type"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type GlobalCellID struct {
	PlmnId       PLMNIdentity
	CellType     CellTypeChoice
	IEExtensions []byte
}

func (ie *GlobalCellID) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(globalCellIDConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.PlmnId.Encode(e); err != nil {
		return err
	}
	if err := ie.CellType.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *GlobalCellID) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(globalCellIDConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PlmnId.Decode(d); err != nil {
		return err
	}
	if err := ie.CellType.Decode(d); err != nil {
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
