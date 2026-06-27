package ies

import (
	"github.com/lvdund/asn1go/per"
)

var trafficToBeModifiedItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "trafficIndex"},
		{Name: "trafficProfile", Optional: true},
		{Name: "f1-TerminatingTopologyBHInformation", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type TrafficToBeModifiedItem struct {
	TrafficIndex                       TrafficIndex
	TrafficProfile                     *TrafficProfile
	F1TerminatingTopologyBHInformation *F1TerminatingTopologyBHInformation
	IEExtensions                       []byte
}

func (ie *TrafficToBeModifiedItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(trafficToBeModifiedItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.TrafficProfile != nil, ie.F1TerminatingTopologyBHInformation != nil, false}); err != nil {
		return err
	}
	if err := ie.TrafficIndex.Encode(e); err != nil {
		return err
	}
	if ie.TrafficProfile != nil {
		if err := ie.TrafficProfile.Encode(e); err != nil {
			return err
		}
	}
	if ie.F1TerminatingTopologyBHInformation != nil {
		if err := ie.F1TerminatingTopologyBHInformation.Encode(e); err != nil {
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

func (ie *TrafficToBeModifiedItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(trafficToBeModifiedItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.TrafficIndex.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.TrafficProfile = new(TrafficProfile)
		if err := ie.TrafficProfile.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.F1TerminatingTopologyBHInformation = new(F1TerminatingTopologyBHInformation)
		if err := ie.F1TerminatingTopologyBHInformation.Decode(d); err != nil {
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
