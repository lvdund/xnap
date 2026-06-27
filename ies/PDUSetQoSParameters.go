package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSetQoSParametersConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ulPDUSetQoSInformation", Optional: true},
		{Name: "dlPDUSetQoSInformation", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: nil,
}

type PDUSetQoSParameters struct {
	UlPDUSetQoSInformation *PDUSetQoSInformation
	DlPDUSetQoSInformation *PDUSetQoSInformation
	IEExtensions           []byte
}

func (ie *PDUSetQoSParameters) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSetQoSParametersConstraints)
	if err := seq.EncodePreamble([]bool{ie.UlPDUSetQoSInformation != nil, ie.DlPDUSetQoSInformation != nil, false}); err != nil {
		return err
	}
	if ie.UlPDUSetQoSInformation != nil {
		if err := ie.UlPDUSetQoSInformation.Encode(e); err != nil {
			return err
		}
	}
	if ie.DlPDUSetQoSInformation != nil {
		if err := ie.DlPDUSetQoSInformation.Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *PDUSetQoSParameters) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSetQoSParametersConstraints)
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.UlPDUSetQoSInformation = new(PDUSetQoSInformation)
		if err := ie.UlPDUSetQoSInformation.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.DlPDUSetQoSInformation = new(PDUSetQoSInformation)
		if err := ie.DlPDUSetQoSInformation.Decode(d); err != nil {
			return err
		}
	}
	return nil
}
