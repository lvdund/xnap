package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionAggregateMaximumBitRateConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "downlink-session-AMBR"},
		{Name: "uplink-session-AMBR"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionAggregateMaximumBitRate struct {
	DownlinkSessionAMBR BitRate
	UplinkSessionAMBR   BitRate
	IEExtensions        []byte
}

func (ie *PDUSessionAggregateMaximumBitRate) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionAggregateMaximumBitRateConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.DownlinkSessionAMBR.Encode(e); err != nil {
		return err
	}
	if err := ie.UplinkSessionAMBR.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *PDUSessionAggregateMaximumBitRate) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionAggregateMaximumBitRateConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.DownlinkSessionAMBR.Decode(d); err != nil {
		return err
	}
	if err := ie.UplinkSessionAMBR.Decode(d); err != nil {
		return err
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
