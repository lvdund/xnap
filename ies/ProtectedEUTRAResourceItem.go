package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ProtectedEUTRAResourceItemResourceTypeDownlinknonCRS int64 = 0
	ProtectedEUTRAResourceItemResourceTypeCRS            int64 = 1
	ProtectedEUTRAResourceItemResourceTypeUplink         int64 = 2
)

var protectedEUTRAResourceItemResourceTypeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type ProtectedEUTRAResourceItemResourceType struct {
	Value int64
}

func (ie *ProtectedEUTRAResourceItemResourceType) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, protectedEUTRAResourceItemResourceTypeConstraints)
}

func (ie *ProtectedEUTRAResourceItemResourceType) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(protectedEUTRAResourceItemResourceTypeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var protectedEUTRAResourceItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "resourceType"},
		{Name: "intra-PRBProtectedResourceFootprint"},
		{Name: "protectedFootprintFrequencyPattern"},
		{Name: "protectedFootprintTimePattern"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ProtectedEUTRAResourceItem struct {
	ResourceType                       ProtectedEUTRAResourceItemResourceType
	IntraPRBProtectedResourceFootprint per.BitString
	ProtectedFootprintFrequencyPattern per.BitString
	ProtectedFootprintTimePattern      ProtectedEUTRAFootprintTimePattern
	IEExtensions                       []byte
}

func (ie *ProtectedEUTRAResourceItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(protectedEUTRAResourceItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.ResourceType.Encode(e); err != nil {
		return err
	}
	if err := e.EncodeBitString(ie.IntraPRBProtectedResourceFootprint, per.SizeConstraints{
		Extensible: false,
		Min:        nil,
		Max:        nil,
	}); err != nil {
		return err
	}
	if err := e.EncodeBitString(ie.ProtectedFootprintFrequencyPattern, per.SizeConstraints{
		Extensible: false,
		Min:        nil,
		Max:        nil,
	}); err != nil {
		return err
	}
	if err := ie.ProtectedFootprintTimePattern.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *ProtectedEUTRAResourceItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(protectedEUTRAResourceItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.ResourceType.Decode(d); err != nil {
		return err
	}
	{
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.IntraPRBProtectedResourceFootprint = val
	}
	{
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.ProtectedFootprintFrequencyPattern = val
	}
	if err := ie.ProtectedFootprintTimePattern.Decode(d); err != nil {
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
