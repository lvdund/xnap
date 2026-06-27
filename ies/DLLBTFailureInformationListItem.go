package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var dLLBTFailureInformationListItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "uEAssistantIdentifier"},
		{Name: "numberOfDLLBTFailures"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type DLLBTFailureInformationListItem struct {
	UEAssistantIdentifier NGRANnodeUEXnAPID
	NumberOfDLLBTFailures int64
	IEExtensions          []byte
}

func (ie *DLLBTFailureInformationListItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dLLBTFailureInformationListItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.UEAssistantIdentifier.Encode(e); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.NumberOfDLLBTFailures, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(1)),
		UpperBound: common.Ptr(int64(1000)),
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

func (ie *DLLBTFailureInformationListItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dLLBTFailureInformationListItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.UEAssistantIdentifier.Decode(d); err != nil {
		return err
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(1)),
			UpperBound: common.Ptr(int64(1000)),
		})
		if err != nil {
			return err
		}
		ie.NumberOfDLLBTFailures = val
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
