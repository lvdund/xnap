package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	CNTypeRestrictionsForEquivalentItemCnTypeEpcForbidden    int64 = 0
	CNTypeRestrictionsForEquivalentItemCnTypeFiveGCForbidden int64 = 1
)

var cNTypeRestrictionsForEquivalentItemCnTypeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type CNTypeRestrictionsForEquivalentItemCnType struct {
	Value int64
}

func (ie *CNTypeRestrictionsForEquivalentItemCnType) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, cNTypeRestrictionsForEquivalentItemCnTypeConstraints)
}

func (ie *CNTypeRestrictionsForEquivalentItemCnType) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(cNTypeRestrictionsForEquivalentItemCnTypeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var cNTypeRestrictionsForEquivalentItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-Identity"},
		{Name: "cn-Type"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CNTypeRestrictionsForEquivalentItem struct {
	PlmnIdentity PLMNIdentity
	CnType       CNTypeRestrictionsForEquivalentItemCnType
	IEExtensions []byte
}

func (ie *CNTypeRestrictionsForEquivalentItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cNTypeRestrictionsForEquivalentItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.PlmnIdentity.Encode(e); err != nil {
		return err
	}
	if err := ie.CnType.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *CNTypeRestrictionsForEquivalentItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cNTypeRestrictionsForEquivalentItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PlmnIdentity.Decode(d); err != nil {
		return err
	}
	if err := ie.CnType.Decode(d); err != nil {
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
