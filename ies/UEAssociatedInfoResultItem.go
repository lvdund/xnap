package ies

import (
	"github.com/lvdund/asn1go/per"
)

var uEAssociatedInfoResultItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "uEAssistantIdentifier"},
		{Name: "uEPerformance", Optional: true},
		{Name: "measuredUETrajectory", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type UEAssociatedInfoResultItem struct {
	UEAssistantIdentifier NGRANnodeUEXnAPID
	UEPerformance         *UEPerformance
	MeasuredUETrajectory  *MeasuredUETrajectory
	IEExtensions          []byte
}

func (ie *UEAssociatedInfoResultItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEAssociatedInfoResultItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.UEPerformance != nil, ie.MeasuredUETrajectory != nil, false}); err != nil {
		return err
	}
	if err := ie.UEAssistantIdentifier.Encode(e); err != nil {
		return err
	}
	if ie.UEPerformance != nil {
		if err := ie.UEPerformance.Encode(e); err != nil {
			return err
		}
	}
	if ie.MeasuredUETrajectory != nil {
		if err := ie.MeasuredUETrajectory.Encode(e); err != nil {
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

func (ie *UEAssociatedInfoResultItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEAssociatedInfoResultItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.UEAssistantIdentifier.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.UEPerformance = new(UEPerformance)
		if err := ie.UEPerformance.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.MeasuredUETrajectory = new(MeasuredUETrajectory)
		if err := ie.MeasuredUETrajectory.Decode(d); err != nil {
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
