package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	PDCPChangeIndicationFromSNGRANNodeSNgRanNodeKeyUpdateRequired int64 = 0
	PDCPChangeIndicationFromSNGRANNodePdcpDataRecoveryRequired    int64 = 1
)

var pDCPChangeIndicationFromSNGRANNodeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type PDCPChangeIndicationFromSNGRANNode struct {
	Value int64
}

func (ie *PDCPChangeIndicationFromSNGRANNode) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, pDCPChangeIndicationFromSNGRANNodeConstraints)
}

func (ie *PDCPChangeIndicationFromSNGRANNode) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(pDCPChangeIndicationFromSNGRANNodeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	PDCPChangeIndicationFromMNGRANNodePdcpDataRecoveryRequired int64 = 0
)

var pDCPChangeIndicationFromMNGRANNodeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type PDCPChangeIndicationFromMNGRANNode struct {
	Value int64
}

func (ie *PDCPChangeIndicationFromMNGRANNode) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, pDCPChangeIndicationFromMNGRANNodeConstraints)
}

func (ie *PDCPChangeIndicationFromMNGRANNode) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(pDCPChangeIndicationFromMNGRANNodeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	PDCPChangeIndicationChFromSNGRANNode  = 0
	PDCPChangeIndicationChFromMNGRANNode  = 1
	PDCPChangeIndicationChChoiceExtension = 2
)

var pDCPChangeIndicationConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "from-S-NG-RAN-node"},
		{Name: "from-M-NG-RAN-node"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type PDCPChangeIndication struct {
	Choice          int
	FromSNGRANNode  *PDCPChangeIndicationFromSNGRANNode
	FromMNGRANNode  *PDCPChangeIndicationFromMNGRANNode
	ChoiceExtension []byte
}

func (ie *PDCPChangeIndication) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(pDCPChangeIndicationConstraints)
	switch ie.Choice {
	case 0: // from-S-NG-RAN-node
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.FromSNGRANNode.Encode(e); err != nil {
			return err
		}
	case 1: // from-M-NG-RAN-node
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.FromMNGRANNode.Encode(e); err != nil {
			return err
		}
	case 2: // choice-extension
		if err := choice.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		// TODO encode field ChoiceExtension (kind=ext)
	}
	return nil
}

func (ie *PDCPChangeIndication) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(pDCPChangeIndicationConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // from-S-NG-RAN-node
		ie.FromSNGRANNode = new(PDCPChangeIndicationFromSNGRANNode)
		if err := ie.FromSNGRANNode.Decode(d); err != nil {
			return err
		}
	case 1: // from-M-NG-RAN-node
		ie.FromMNGRANNode = new(PDCPChangeIndicationFromMNGRANNode)
		if err := ie.FromMNGRANNode.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
