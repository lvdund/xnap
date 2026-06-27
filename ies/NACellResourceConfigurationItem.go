package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NACellResourceConfigurationItemNAdownlinkTrue  int64 = 0
	NACellResourceConfigurationItemNAdownlinkFalse int64 = 1
)

var nACellResourceConfigurationItemNAdownlinkConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type NACellResourceConfigurationItemNAdownlink struct {
	Value int64
}

func (ie *NACellResourceConfigurationItemNAdownlink) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, nACellResourceConfigurationItemNAdownlinkConstraints)
}

func (ie *NACellResourceConfigurationItemNAdownlink) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(nACellResourceConfigurationItemNAdownlinkConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	NACellResourceConfigurationItemNAuplinkTrue  int64 = 0
	NACellResourceConfigurationItemNAuplinkFalse int64 = 1
)

var nACellResourceConfigurationItemNAuplinkConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type NACellResourceConfigurationItemNAuplink struct {
	Value int64
}

func (ie *NACellResourceConfigurationItemNAuplink) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, nACellResourceConfigurationItemNAuplinkConstraints)
}

func (ie *NACellResourceConfigurationItemNAuplink) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(nACellResourceConfigurationItemNAuplinkConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	NACellResourceConfigurationItemNAflexibleTrue  int64 = 0
	NACellResourceConfigurationItemNAflexibleFalse int64 = 1
)

var nACellResourceConfigurationItemNAflexibleConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type NACellResourceConfigurationItemNAflexible struct {
	Value int64
}

func (ie *NACellResourceConfigurationItemNAflexible) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, nACellResourceConfigurationItemNAflexibleConstraints)
}

func (ie *NACellResourceConfigurationItemNAflexible) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(nACellResourceConfigurationItemNAflexibleConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var nACellResourceConfigurationItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nAdownlink", Optional: true},
		{Name: "nAuplink", Optional: true},
		{Name: "nAflexible", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NACellResourceConfigurationItem struct {
	NAdownlink   *NACellResourceConfigurationItemNAdownlink
	NAuplink     *NACellResourceConfigurationItemNAuplink
	NAflexible   *NACellResourceConfigurationItemNAflexible
	IEExtensions []byte
}

func (ie *NACellResourceConfigurationItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nACellResourceConfigurationItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.NAdownlink != nil, ie.NAuplink != nil, ie.NAflexible != nil, false}); err != nil {
		return err
	}
	if ie.NAdownlink != nil {
		if err := ie.NAdownlink.Encode(e); err != nil {
			return err
		}
	}
	if ie.NAuplink != nil {
		if err := ie.NAuplink.Encode(e); err != nil {
			return err
		}
	}
	if ie.NAflexible != nil {
		if err := ie.NAflexible.Encode(e); err != nil {
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

func (ie *NACellResourceConfigurationItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nACellResourceConfigurationItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.NAdownlink = new(NACellResourceConfigurationItemNAdownlink)
		if err := ie.NAdownlink.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.NAuplink = new(NACellResourceConfigurationItemNAuplink)
		if err := ie.NAuplink.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.NAflexible = new(NACellResourceConfigurationItemNAflexible)
		if err := ie.NAflexible.Decode(d); err != nil {
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
