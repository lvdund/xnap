package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	CauseChRadioNetwork    = 0
	CauseChTransport       = 1
	CauseChProtocol        = 2
	CauseChMisc            = 3
	CauseChChoiceExtension = 4
)

var causeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "radioNetwork"},
		{Name: "transport"},
		{Name: "protocol"},
		{Name: "misc"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type Cause struct {
	Choice          int
	RadioNetwork    *CauseRadioNetworkLayer
	Transport       *CauseTransportLayer
	Protocol        *CauseProtocol
	Misc            *CauseMisc
	ChoiceExtension []byte
}

func (ie *Cause) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(causeConstraints)
	switch ie.Choice {
	case 0: // radioNetwork
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.RadioNetwork.Encode(e); err != nil {
			return err
		}
	case 1: // transport
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.Transport.Encode(e); err != nil {
			return err
		}
	case 2: // protocol
		if err := choice.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		if err := ie.Protocol.Encode(e); err != nil {
			return err
		}
	case 3: // misc
		if err := choice.EncodeChoice(3, false, nil); err != nil {
			return err
		}
		if err := ie.Misc.Encode(e); err != nil {
			return err
		}
	case 4: // choice-extension
		if err := choice.EncodeChoice(4, false, nil); err != nil {
			return err
		}
		// TODO encode field ChoiceExtension (kind=ext)
	}
	return nil
}

func (ie *Cause) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(causeConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // radioNetwork
		ie.RadioNetwork = new(CauseRadioNetworkLayer)
		if err := ie.RadioNetwork.Decode(d); err != nil {
			return err
		}
	case 1: // transport
		ie.Transport = new(CauseTransportLayer)
		if err := ie.Transport.Decode(d); err != nil {
			return err
		}
	case 2: // protocol
		ie.Protocol = new(CauseProtocol)
		if err := ie.Protocol.Decode(d); err != nil {
			return err
		}
	case 3: // misc
		ie.Misc = new(CauseMisc)
		if err := ie.Misc.Decode(d); err != nil {
			return err
		}
	case 4: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
