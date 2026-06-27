package ies

import (
	"github.com/lvdund/asn1go/per"
)

var cellAndCapacityAssistanceInfoEUTRAConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "maximumCellListSize", Optional: true},
		{Name: "cellAssistanceInfo-EUTRA", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CellAndCapacityAssistanceInfoEUTRA struct {
	MaximumCellListSize     *MaximumCellListSize
	CellAssistanceInfoEUTRA *CellAssistanceInfoEUTRA
	IEExtensions            []byte
}

func (ie *CellAndCapacityAssistanceInfoEUTRA) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cellAndCapacityAssistanceInfoEUTRAConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.MaximumCellListSize != nil, ie.CellAssistanceInfoEUTRA != nil, false}); err != nil {
		return err
	}
	if ie.MaximumCellListSize != nil {
		if err := ie.MaximumCellListSize.Encode(e); err != nil {
			return err
		}
	}
	if ie.CellAssistanceInfoEUTRA != nil {
		if err := ie.CellAssistanceInfoEUTRA.Encode(e); err != nil {
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

func (ie *CellAndCapacityAssistanceInfoEUTRA) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cellAndCapacityAssistanceInfoEUTRAConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.MaximumCellListSize = new(MaximumCellListSize)
		if err := ie.MaximumCellListSize.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.CellAssistanceInfoEUTRA = new(CellAssistanceInfoEUTRA)
		if err := ie.CellAssistanceInfoEUTRA.Decode(d); err != nil {
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
