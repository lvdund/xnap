package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var nRPPaPositioningInformationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "routingID"},
		{Name: "nRPPaTransactionID"},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NRPPaPositioningInformation struct {
	RoutingID          RoutingID
	NRPPaTransactionID int64
	IEExtensions       []byte
}

func (ie *NRPPaPositioningInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRPPaPositioningInformationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.RoutingID.Encode(e); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.NRPPaTransactionID, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(0)),
		UpperBound: common.Ptr(int64(32767)),
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

func (ie *NRPPaPositioningInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRPPaPositioningInformationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.RoutingID.Decode(d); err != nil {
		return err
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(32767)),
		})
		if err != nil {
			return err
		}
		ie.NRPPaTransactionID = val
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
