package ies

import (
	"github.com/lvdund/asn1go/per"
)

var cellBasedMDTEUTRAConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cellIdListforMDT-EUTRA"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CellBasedMDTEUTRA struct {
	CellIdListforMDTEUTRA CellIdListforMDTEUTRA
	IEExtensions          []byte
}

func (ie *CellBasedMDTEUTRA) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cellBasedMDTEUTRAConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.CellIdListforMDTEUTRA.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *CellBasedMDTEUTRA) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cellBasedMDTEUTRAConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.CellIdListforMDTEUTRA.Decode(d); err != nil {
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
