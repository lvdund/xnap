package ies

import (
	"github.com/lvdund/asn1go/per"
)

var mBSSessionIDConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "tMGI"},
		{Name: "nID", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type MBSSessionID struct {
	TMGI         TMGI
	NID          *NID
	IEExtensions []byte
}

func (ie *MBSSessionID) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mBSSessionIDConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.NID != nil, false}); err != nil {
		return err
	}
	if err := ie.TMGI.Encode(e); err != nil {
		return err
	}
	if ie.NID != nil {
		if err := ie.NID.Encode(e); err != nil {
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

func (ie *MBSSessionID) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mBSSessionIDConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.TMGI.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.NID = new(NID)
		if err := ie.NID.Decode(d); err != nil {
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
