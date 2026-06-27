package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cPCTargetSNRequiredListItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "target-S-NG-RANnodeID"},
		{Name: "cpc-indicator"},
		{Name: "max-no-of-pscells"},
		{Name: "cpac-EstimatedArrivalProbability", Optional: true},
		{Name: "sN-to-MN-Container"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CPCTargetSNRequiredListItem struct {
	TargetSNGRANnodeID              GlobalNGRANNodeID
	CpcIndicator                    CPCindicator
	MaxNoOfPscells                  int64
	CpacEstimatedArrivalProbability *CHOProbability
	SNToMNContainer                 []byte
	IEExtensions                    []byte
}

func (ie *CPCTargetSNRequiredListItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cPCTargetSNRequiredListItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.CpacEstimatedArrivalProbability != nil, false}); err != nil {
		return err
	}
	if err := ie.TargetSNGRANnodeID.Encode(e); err != nil {
		return err
	}
	if err := ie.CpcIndicator.Encode(e); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.MaxNoOfPscells, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(1)),
		UpperBound: common.Ptr(int64(common.MaxnoofPSCellCandidates)),
	}); err != nil {
		return err
	}
	if ie.CpacEstimatedArrivalProbability != nil {
		if err := ie.CpacEstimatedArrivalProbability.Encode(e); err != nil {
			return err
		}
	}
	if err := e.EncodeOctetString(ie.SNToMNContainer, per.SizeConstraints{
		Extensible: false,
		Min:        nil,
		Max:        nil,
	}); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *CPCTargetSNRequiredListItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cPCTargetSNRequiredListItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.TargetSNGRANnodeID.Decode(d); err != nil {
		return err
	}
	if err := ie.CpcIndicator.Decode(d); err != nil {
		return err
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(1)),
			UpperBound: common.Ptr(int64(common.MaxnoofPSCellCandidates)),
		})
		if err != nil {
			return err
		}
		ie.MaxNoOfPscells = val
	}
	if seq.IsComponentPresent(3) {
		ie.CpacEstimatedArrivalProbability = new(CHOProbability)
		if err := ie.CpacEstimatedArrivalProbability.Decode(d); err != nil {
			return err
		}
	}
	{
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.SNToMNContainer = val
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
