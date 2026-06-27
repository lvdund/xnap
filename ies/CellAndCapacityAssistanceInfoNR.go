package ies

import (
	"github.com/lvdund/asn1go/per"
)

var cellAndCapacityAssistanceInfoNRConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "maximumCellListSize", Optional: true},
		{Name: "cellAssistanceInfo-NR", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CellAndCapacityAssistanceInfoNR struct {
	MaximumCellListSize  *MaximumCellListSize
	CellAssistanceInfoNR *CellAssistanceInfoNR
	IEExtensions         []byte
}

func (ie *CellAndCapacityAssistanceInfoNR) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cellAndCapacityAssistanceInfoNRConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.MaximumCellListSize != nil, ie.CellAssistanceInfoNR != nil, false}); err != nil {
		return err
	}
	if ie.MaximumCellListSize != nil {
		if err := ie.MaximumCellListSize.Encode(e); err != nil {
			return err
		}
	}
	if ie.CellAssistanceInfoNR != nil {
		if err := ie.CellAssistanceInfoNR.Encode(e); err != nil {
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

func (ie *CellAndCapacityAssistanceInfoNR) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cellAndCapacityAssistanceInfoNRConstraints)
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
		ie.CellAssistanceInfoNR = new(CellAssistanceInfoNR)
		if err := ie.CellAssistanceInfoNR.Decode(d); err != nil {
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
