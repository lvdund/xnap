package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cPAInformationModReqConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "max-no-of-pscells"},
		{Name: "cpac-EstimatedArrivalProbability", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CPAInformationModReq struct {
	MaxNoOfPscells                  int64
	CpacEstimatedArrivalProbability *CHOProbability
	IEExtensions                    []byte
}

func (ie *CPAInformationModReq) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cPAInformationModReqConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.CpacEstimatedArrivalProbability != nil, false}); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.MaxNoOfPscells, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(1)),
		UpperBound: common.Ptr(int64(8)),
	}); err != nil {
		return err
	}
	if ie.CpacEstimatedArrivalProbability != nil {
		if err := ie.CpacEstimatedArrivalProbability.Encode(e); err != nil {
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

func (ie *CPAInformationModReq) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cPAInformationModReqConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(1)),
			UpperBound: common.Ptr(int64(8)),
		})
		if err != nil {
			return err
		}
		ie.MaxNoOfPscells = val
	}
	if seq.IsComponentPresent(1) {
		ie.CpacEstimatedArrivalProbability = new(CHOProbability)
		if err := ie.CpacEstimatedArrivalProbability.Decode(d); err != nil {
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
