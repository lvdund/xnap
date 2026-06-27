package ies

import (
	"github.com/lvdund/asn1go/per"
)

var cellsinAoIItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pLMN-Identity"},
		{Name: "ng-ran-cell-id"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CellsinAoIItem struct {
	PLMNIdentity PLMNIdentity
	NgRanCellId  NGRANCellIdentity
	IEExtensions []byte
}

func (ie *CellsinAoIItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cellsinAoIItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.PLMNIdentity.Encode(e); err != nil {
		return err
	}
	if err := ie.NgRanCellId.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *CellsinAoIItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cellsinAoIItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PLMNIdentity.Decode(d); err != nil {
		return err
	}
	if err := ie.NgRanCellId.Decode(d); err != nil {
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
