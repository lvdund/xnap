package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	DRBBStatusTransferChoiceChPdcpSn12bits    = 0
	DRBBStatusTransferChoiceChPdcpSn18bits    = 1
	DRBBStatusTransferChoiceChChoiceExtension = 2
)

var dRBBStatusTransferChoiceConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "pdcp-sn-12bits"},
		{Name: "pdcp-sn-18bits"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type DRBBStatusTransferChoice struct {
	Choice          int
	PdcpSn12bits    *DRBBStatusTransfer12bitsSN
	PdcpSn18bits    *DRBBStatusTransfer18bitsSN
	ChoiceExtension []byte
}

func (ie *DRBBStatusTransferChoice) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(dRBBStatusTransferChoiceConstraints)
	switch ie.Choice {
	case 0: // pdcp-sn-12bits
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.PdcpSn12bits.Encode(e); err != nil {
			return err
		}
	case 1: // pdcp-sn-18bits
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.PdcpSn18bits.Encode(e); err != nil {
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

func (ie *DRBBStatusTransferChoice) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(dRBBStatusTransferChoiceConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // pdcp-sn-12bits
		ie.PdcpSn12bits = new(DRBBStatusTransfer12bitsSN)
		if err := ie.PdcpSn12bits.Decode(d); err != nil {
			return err
		}
	case 1: // pdcp-sn-18bits
		ie.PdcpSn18bits = new(DRBBStatusTransfer18bitsSN)
		if err := ie.PdcpSn18bits.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
