package ies

import (
	"github.com/lvdund/asn1go/per"
)

var sharedResourceTypeULDLSharingDLResourcesChangedConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dl-resourceBitmap"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SharedResourceTypeULDLSharingDLResourcesChanged struct {
	DlResourceBitmap DataTrafficResources
	IEExtensions     []byte
}

func (ie *SharedResourceTypeULDLSharingDLResourcesChanged) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sharedResourceTypeULDLSharingDLResourcesChangedConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.DlResourceBitmap.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SharedResourceTypeULDLSharingDLResourcesChanged) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sharedResourceTypeULDLSharingDLResourcesChangedConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.DlResourceBitmap.Decode(d); err != nil {
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
