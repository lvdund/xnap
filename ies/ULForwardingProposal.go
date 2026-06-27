package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ULForwardingProposalUlForwardingProposed int64 = 0
)

var uLForwardingProposalConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type ULForwardingProposal struct {
	Value int64
}

func (ie *ULForwardingProposal) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, uLForwardingProposalConstraints)
}

func (ie *ULForwardingProposal) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(uLForwardingProposalConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
