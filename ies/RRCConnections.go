package ies

import (
	"github.com/lvdund/asn1go/per"
)

var rRCConnectionsConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "noofRRCConnections"},
		{Name: "availableRRCConnectionCapacityValue"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type RRCConnections struct {
	NoofRRCConnections                  NoofRRCConnections
	AvailableRRCConnectionCapacityValue AvailableRRCConnectionCapacityValue
	IEExtensions                        []byte
}

func (ie *RRCConnections) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCConnectionsConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.NoofRRCConnections.Encode(e); err != nil {
		return err
	}
	if err := ie.AvailableRRCConnectionCapacityValue.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *RRCConnections) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCConnectionsConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.NoofRRCConnections.Decode(d); err != nil {
		return err
	}
	if err := ie.AvailableRRCConnectionCapacityValue.Decode(d); err != nil {
		return err
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
