package ies

import (
	"github.com/lvdund/asn1go/per"
)

var posPartialUEContextInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "requestedSRSTransmissionCharacteristics", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PosPartialUEContextInfo struct {
	RequestedSRSTransmissionCharacteristics *RequestedSRSTransmissionCharacteristics
	IEExtensions                            []byte
}

func (ie *PosPartialUEContextInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(posPartialUEContextInfoConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.RequestedSRSTransmissionCharacteristics != nil, false}); err != nil {
		return err
	}
	if ie.RequestedSRSTransmissionCharacteristics != nil {
		if err := ie.RequestedSRSTransmissionCharacteristics.Encode(e); err != nil {
			return err
		}
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *PosPartialUEContextInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(posPartialUEContextInfoConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.RequestedSRSTransmissionCharacteristics = new(RequestedSRSTransmissionCharacteristics)
		if err := ie.RequestedSRSTransmissionCharacteristics.Decode(d); err != nil {
			return err
		}
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
