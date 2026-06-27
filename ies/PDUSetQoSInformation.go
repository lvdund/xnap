package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	PDUSetQoSInformationPduSetIntegratedHandlingInformationTrue  int64 = 0
	PDUSetQoSInformationPduSetIntegratedHandlingInformationFalse int64 = 1
)

var pDUSetQoSInformationPduSetIntegratedHandlingInformationConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type PDUSetQoSInformationPduSetIntegratedHandlingInformation struct {
	Value int64
}

func (ie *PDUSetQoSInformationPduSetIntegratedHandlingInformation) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, pDUSetQoSInformationPduSetIntegratedHandlingInformationConstraints)
}

func (ie *PDUSetQoSInformationPduSetIntegratedHandlingInformation) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(pDUSetQoSInformationPduSetIntegratedHandlingInformationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var pDUSetQoSInformationConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pduSetDelayBudget", Optional: true},
		{Name: "pduSetErrorRate", Optional: true},
		{Name: "pduSetIntegratedHandlingInformation", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: nil,
}

type PDUSetQoSInformation struct {
	PduSetDelayBudget                   *ExtendedPacketDelayBudget
	PduSetErrorRate                     *PacketErrorRate
	PduSetIntegratedHandlingInformation *PDUSetQoSInformationPduSetIntegratedHandlingInformation
	IEExtensions                        []byte
}

func (ie *PDUSetQoSInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSetQoSInformationConstraints)
	if err := seq.EncodePreamble([]bool{ie.PduSetDelayBudget != nil, ie.PduSetErrorRate != nil, ie.PduSetIntegratedHandlingInformation != nil, false}); err != nil {
		return err
	}
	if ie.PduSetDelayBudget != nil {
		if err := ie.PduSetDelayBudget.Encode(e); err != nil {
			return err
		}
	}
	if ie.PduSetErrorRate != nil {
		if err := ie.PduSetErrorRate.Encode(e); err != nil {
			return err
		}
	}
	if ie.PduSetIntegratedHandlingInformation != nil {
		if err := ie.PduSetIntegratedHandlingInformation.Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *PDUSetQoSInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSetQoSInformationConstraints)
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.PduSetDelayBudget = new(ExtendedPacketDelayBudget)
		if err := ie.PduSetDelayBudget.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.PduSetErrorRate = new(PacketErrorRate)
		if err := ie.PduSetErrorRate.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.PduSetIntegratedHandlingInformation = new(PDUSetQoSInformationPduSetIntegratedHandlingInformation)
		if err := ie.PduSetIntegratedHandlingInformation.Decode(d); err != nil {
			return err
		}
	}
	return nil
}
