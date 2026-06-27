package ies

import (
	"github.com/lvdund/asn1go/per"
)

var resetResponsePartialReleaseItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ng-ran-node1UEXnAPID", Optional: true},
		{Name: "ng-ran-node2UEXnAPID", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ResetResponsePartialReleaseItem struct {
	NgRanNode1UEXnAPID *NGRANnodeUEXnAPID
	NgRanNode2UEXnAPID *NGRANnodeUEXnAPID
	IEExtensions       []byte
}

func (ie *ResetResponsePartialReleaseItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(resetResponsePartialReleaseItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.NgRanNode1UEXnAPID != nil, ie.NgRanNode2UEXnAPID != nil, false}); err != nil {
		return err
	}
	if ie.NgRanNode1UEXnAPID != nil {
		if err := ie.NgRanNode1UEXnAPID.Encode(e); err != nil {
			return err
		}
	}
	if ie.NgRanNode2UEXnAPID != nil {
		if err := ie.NgRanNode2UEXnAPID.Encode(e); err != nil {
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

func (ie *ResetResponsePartialReleaseItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(resetResponsePartialReleaseItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.NgRanNode1UEXnAPID = new(NGRANnodeUEXnAPID)
		if err := ie.NgRanNode1UEXnAPID.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.NgRanNode2UEXnAPID = new(NGRANnodeUEXnAPID)
		if err := ie.NgRanNode2UEXnAPID.Decode(d); err != nil {
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
