package ies

import (
	"github.com/lvdund/asn1go/per"
)

var nRTransmissionBandwidthConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nRSCS"},
		{Name: "nRNRB"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NRTransmissionBandwidth struct {
	NRSCS        NRSCS
	NRNRB        NRNRB
	IEExtensions []byte
}

func (ie *NRTransmissionBandwidth) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRTransmissionBandwidthConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.NRSCS.Encode(e); err != nil {
		return err
	}
	if err := ie.NRNRB.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *NRTransmissionBandwidth) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRTransmissionBandwidthConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.NRSCS.Decode(d); err != nil {
		return err
	}
	if err := ie.NRNRB.Decode(d); err != nil {
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
