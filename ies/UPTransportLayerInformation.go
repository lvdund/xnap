package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	UPTransportLayerInformationChGtpTunnel       = 0
	UPTransportLayerInformationChChoiceExtension = 1
)

var uPTransportLayerInformationConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "gtpTunnel"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type UPTransportLayerInformation struct {
	Choice          int
	GtpTunnel       *GTPtunnelTransportLayerInformation
	ChoiceExtension []byte
}

func (ie *UPTransportLayerInformation) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(uPTransportLayerInformationConstraints)
	switch ie.Choice {
	case 0: // gtpTunnel
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.GtpTunnel.Encode(e); err != nil {
			return err
		}
	case 1: // choice-extension
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		// TODO encode field ChoiceExtension (kind=ext)
	}
	return nil
}

func (ie *UPTransportLayerInformation) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(uPTransportLayerInformationConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // gtpTunnel
		ie.GtpTunnel = new(GTPtunnelTransportLayerInformation)
		if err := ie.GtpTunnel.Decode(d); err != nil {
			return err
		}
	case 1: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
