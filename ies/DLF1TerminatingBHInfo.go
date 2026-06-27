package ies

import (
	"github.com/lvdund/asn1go/per"
)

var dLF1TerminatingBHInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "egressBAPRoutingID"},
		{Name: "egressBHRLCCHID"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type DLF1TerminatingBHInfo struct {
	EgressBAPRoutingID BAPRoutingID
	EgressBHRLCCHID    BHRLCChannelID
	IEExtensions       []byte
}

func (ie *DLF1TerminatingBHInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dLF1TerminatingBHInfoConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.EgressBAPRoutingID.Encode(e); err != nil {
		return err
	}
	if err := ie.EgressBHRLCCHID.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *DLF1TerminatingBHInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dLF1TerminatingBHInfoConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.EgressBAPRoutingID.Decode(d); err != nil {
		return err
	}
	if err := ie.EgressBHRLCCHID.Decode(d); err != nil {
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
