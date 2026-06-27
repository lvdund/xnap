package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var forbiddenAreaItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-Identity"},
		{Name: "forbidden-TACs"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ForbiddenAreaItem struct {
	PlmnIdentity  PLMNIdentity
	ForbiddenTACs []*TAC
	IEExtensions  []byte
}

func (ie *ForbiddenAreaItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(forbiddenAreaItemConstraints)
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
	soEnc := e.NewSequenceOfEncoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(1)),
		Max:        common.Ptr(int64(common.MaxnoofForbiddenTACs)),
	})
	if err := soEnc.EncodeLength(int64(len(ie.ForbiddenTACs))); err != nil {
		return err
	}
	for _, item := range ie.ForbiddenTACs {
		if err := item.Encode(e); err != nil {
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

func (ie *ForbiddenAreaItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(forbiddenAreaItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PlmnIdentity.Decode(d); err != nil {
		return err
	}
	{
		soDec := d.NewSequenceOfDecoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofForbiddenTACs)),
		})
		n, err := soDec.DecodeLength()
		if err != nil {
			return err
		}
		ie.ForbiddenTACs = make([]*TAC, n)
		for i := range ie.ForbiddenTACs {
			ie.ForbiddenTACs[i] = new(TAC)
			if err := ie.ForbiddenTACs[i].Decode(d); err != nil {
				return err
			}
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
