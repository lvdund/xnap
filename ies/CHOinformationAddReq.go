package ies

import (
	"github.com/lvdund/asn1go/per"
)

var cHOinformationAddReqConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "source-M-NGRAN-node-ID"},
		{Name: "source-M-NGRAN-node-UE-XnAP-ID"},
		{Name: "cHO-EstimatedArrivalProbability", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CHOinformationAddReq struct {
	SourceMNGRANNodeID             GlobalNGRANNodeID
	SourceMNGRANNodeUEXnAPID       NGRANnodeUEXnAPID
	CHOEstimatedArrivalProbability *CHOProbability
	IEExtensions                   []byte
}

func (ie *CHOinformationAddReq) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cHOinformationAddReqConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.CHOEstimatedArrivalProbability != nil, false}); err != nil {
		return err
	}
	if err := ie.SourceMNGRANNodeID.Encode(e); err != nil {
		return err
	}
	if err := ie.SourceMNGRANNodeUEXnAPID.Encode(e); err != nil {
		return err
	}
	if ie.CHOEstimatedArrivalProbability != nil {
		if err := ie.CHOEstimatedArrivalProbability.Encode(e); err != nil {
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

func (ie *CHOinformationAddReq) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cHOinformationAddReqConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.SourceMNGRANNodeID.Decode(d); err != nil {
		return err
	}
	if err := ie.SourceMNGRANNodeUEXnAPID.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(2) {
		ie.CHOEstimatedArrivalProbability = new(CHOProbability)
		if err := ie.CHOEstimatedArrivalProbability.Decode(d); err != nil {
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
