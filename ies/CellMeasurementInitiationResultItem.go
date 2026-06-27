package ies

import (
	"github.com/lvdund/asn1go/per"
)

var cellMeasurementInitiationResultItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cellID"},
		{Name: "cellMeasurementFailureCause-List", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CellMeasurementInitiationResultItem struct {
	CellID                          GlobalNGRANCellID
	CellMeasurementFailureCauseList *CellMeasurementFailureCauseList
	IEExtensions                    []byte
}

func (ie *CellMeasurementInitiationResultItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cellMeasurementInitiationResultItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.CellMeasurementFailureCauseList != nil, false}); err != nil {
		return err
	}
	if err := ie.CellID.Encode(e); err != nil {
		return err
	}
	if ie.CellMeasurementFailureCauseList != nil {
		if err := ie.CellMeasurementFailureCauseList.Encode(e); err != nil {
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

func (ie *CellMeasurementInitiationResultItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cellMeasurementInitiationResultItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.CellID.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.CellMeasurementFailureCauseList = new(CellMeasurementFailureCauseList)
		if err := ie.CellMeasurementFailureCauseList.Decode(d); err != nil {
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
