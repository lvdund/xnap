package ies

import (
	"github.com/lvdund/asn1go/per"
)

var cellMeasurementResultItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cell-ID"},
		{Name: "radioResourceStatus", Optional: true},
		{Name: "tNLCapacityIndicator", Optional: true},
		{Name: "compositeAvailableCapacityGroup", Optional: true},
		{Name: "sliceAvailableCapacity", Optional: true},
		{Name: "numberofActiveUEs", Optional: true},
		{Name: "rRCConnections", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CellMeasurementResultItem struct {
	CellID                          GlobalNGRANCellID
	RadioResourceStatus             *RadioResourceStatus
	TNLCapacityIndicator            *TNLCapacityIndicator
	CompositeAvailableCapacityGroup *CompositeAvailableCapacityGroup
	SliceAvailableCapacity          *SliceAvailableCapacity
	NumberofActiveUEs               *NumberofActiveUEs
	RRCConnections                  *RRCConnections
	IEExtensions                    []byte
}

func (ie *CellMeasurementResultItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cellMeasurementResultItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.RadioResourceStatus != nil, ie.TNLCapacityIndicator != nil, ie.CompositeAvailableCapacityGroup != nil, ie.SliceAvailableCapacity != nil, ie.NumberofActiveUEs != nil, ie.RRCConnections != nil, false}); err != nil {
		return err
	}
	if err := ie.CellID.Encode(e); err != nil {
		return err
	}
	if ie.RadioResourceStatus != nil {
		if err := ie.RadioResourceStatus.Encode(e); err != nil {
			return err
		}
	}
	if ie.TNLCapacityIndicator != nil {
		if err := ie.TNLCapacityIndicator.Encode(e); err != nil {
			return err
		}
	}
	if ie.CompositeAvailableCapacityGroup != nil {
		if err := ie.CompositeAvailableCapacityGroup.Encode(e); err != nil {
			return err
		}
	}
	if ie.SliceAvailableCapacity != nil {
		if err := ie.SliceAvailableCapacity.Encode(e); err != nil {
			return err
		}
	}
	if ie.NumberofActiveUEs != nil {
		if err := ie.NumberofActiveUEs.Encode(e); err != nil {
			return err
		}
	}
	if ie.RRCConnections != nil {
		if err := ie.RRCConnections.Encode(e); err != nil {
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

func (ie *CellMeasurementResultItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cellMeasurementResultItemConstraints)
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
		ie.RadioResourceStatus = new(RadioResourceStatus)
		if err := ie.RadioResourceStatus.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.TNLCapacityIndicator = new(TNLCapacityIndicator)
		if err := ie.TNLCapacityIndicator.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.CompositeAvailableCapacityGroup = new(CompositeAvailableCapacityGroup)
		if err := ie.CompositeAvailableCapacityGroup.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(4) {
		ie.SliceAvailableCapacity = new(SliceAvailableCapacity)
		if err := ie.SliceAvailableCapacity.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(5) {
		ie.NumberofActiveUEs = new(NumberofActiveUEs)
		if err := ie.NumberofActiveUEs.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(6) {
		ie.RRCConnections = new(RRCConnections)
		if err := ie.RRCConnections.Decode(d); err != nil {
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
