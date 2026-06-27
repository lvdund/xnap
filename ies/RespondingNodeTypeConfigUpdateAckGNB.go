package ies

import (
	"github.com/lvdund/asn1go/per"
)

var respondingNodeTypeConfigUpdateAckGNBConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "served-NR-Cells", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type RespondingNodeTypeConfigUpdateAckGNB struct {
	ServedNRCells *ServedCellsNR
	IEExtensions  []byte
}

func (ie *RespondingNodeTypeConfigUpdateAckGNB) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(respondingNodeTypeConfigUpdateAckGNBConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.ServedNRCells != nil, false}); err != nil {
		return err
	}
	if ie.ServedNRCells != nil {
		if err := ie.ServedNRCells.Encode(e); err != nil {
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

func (ie *RespondingNodeTypeConfigUpdateAckGNB) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(respondingNodeTypeConfigUpdateAckGNBConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.ServedNRCells = new(ServedCellsNR)
		if err := ie.ServedNRCells.Decode(d); err != nil {
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
