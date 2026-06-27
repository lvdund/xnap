package ies

import (
	"github.com/lvdund/asn1go/per"
)

var dLNonF1TerminatingBHInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ingressBAPRoutingID"},
		{Name: "ingressBHRLCCHID"},
		{Name: "priorhopBAPAddress"},
		{Name: "iabqosMappingInformation"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type DLNonF1TerminatingBHInfo struct {
	IngressBAPRoutingID      BAPRoutingID
	IngressBHRLCCHID         BHRLCChannelID
	PriorhopBAPAddress       BAPAddress
	IabqosMappingInformation IABQoSMappingInformation
	IEExtensions             []byte
}

func (ie *DLNonF1TerminatingBHInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dLNonF1TerminatingBHInfoConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.IngressBAPRoutingID.Encode(e); err != nil {
		return err
	}
	if err := ie.IngressBHRLCCHID.Encode(e); err != nil {
		return err
	}
	if err := ie.PriorhopBAPAddress.Encode(e); err != nil {
		return err
	}
	if err := ie.IabqosMappingInformation.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *DLNonF1TerminatingBHInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dLNonF1TerminatingBHInfoConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.IngressBAPRoutingID.Decode(d); err != nil {
		return err
	}
	if err := ie.IngressBHRLCCHID.Decode(d); err != nil {
		return err
	}
	if err := ie.PriorhopBAPAddress.Decode(d); err != nil {
		return err
	}
	if err := ie.IabqosMappingInformation.Decode(d); err != nil {
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
