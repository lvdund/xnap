package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SplitSRBRRCTransferSrbTypeSrb1 int64 = 0
	SplitSRBRRCTransferSrbTypeSrb2 int64 = 1
)

var splitSRBRRCTransferSrbTypeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type SplitSRBRRCTransferSrbType struct {
	Value int64
}

func (ie *SplitSRBRRCTransferSrbType) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, splitSRBRRCTransferSrbTypeConstraints)
}

func (ie *SplitSRBRRCTransferSrbType) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(splitSRBRRCTransferSrbTypeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var splitSRBRRCTransferConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "rrcContainer", Optional: true},
		{Name: "srbType"},
		{Name: "deliveryStatus", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SplitSRBRRCTransfer struct {
	RrcContainer   []byte
	SrbType        SplitSRBRRCTransferSrbType
	DeliveryStatus *DeliveryStatus
	IEExtensions   []byte
}

func (ie *SplitSRBRRCTransfer) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(splitSRBRRCTransferConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{len(ie.RrcContainer) > 0, ie.DeliveryStatus != nil, false}); err != nil {
		return err
	}
	if len(ie.RrcContainer) > 0 {
		if err := e.EncodeOctetString(ie.RrcContainer, per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		}); err != nil {
			return err
		}
	}
	if err := ie.SrbType.Encode(e); err != nil {
		return err
	}
	if ie.DeliveryStatus != nil {
		if err := ie.DeliveryStatus.Encode(e); err != nil {
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

func (ie *SplitSRBRRCTransfer) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(splitSRBRRCTransferConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.RrcContainer = val
	}
	if err := ie.SrbType.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(2) {
		ie.DeliveryStatus = new(DeliveryStatus)
		if err := ie.DeliveryStatus.Decode(d); err != nil {
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
