package ies

import (
	"github.com/lvdund/asn1go/per"
)

var requestedSRSTransmissionCharacteristicsConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type RequestedSRSTransmissionCharacteristics struct {
	Value []byte
}

func (ie *RequestedSRSTransmissionCharacteristics) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, requestedSRSTransmissionCharacteristicsConstraints)
}

func (ie *RequestedSRSTransmissionCharacteristics) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(requestedSRSTransmissionCharacteristicsConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
