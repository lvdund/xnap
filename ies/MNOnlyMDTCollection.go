package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	MNOnlyMDTCollectionMNOnly int64 = 0
)

var mNOnlyMDTCollectionConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type MNOnlyMDTCollection struct {
	Value int64
}

func (ie *MNOnlyMDTCollection) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, mNOnlyMDTCollectionConstraints)
}

func (ie *MNOnlyMDTCollection) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(mNOnlyMDTCollectionConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
