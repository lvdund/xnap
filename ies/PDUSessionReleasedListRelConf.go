package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionReleasedListRelConfConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pduSessionsReleasedList-SNterminated", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionReleasedListRelConf struct {
	PduSessionsReleasedListSNterminated *PDUSessionListWithDataForwardingFromTarget
	IEExtensions                        []byte
}

func (ie *PDUSessionReleasedListRelConf) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionReleasedListRelConfConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.PduSessionsReleasedListSNterminated != nil, false}); err != nil {
		return err
	}
	if ie.PduSessionsReleasedListSNterminated != nil {
		if err := ie.PduSessionsReleasedListSNterminated.Encode(e); err != nil {
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

func (ie *PDUSessionReleasedListRelConf) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionReleasedListRelConfConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.PduSessionsReleasedListSNterminated = new(PDUSessionListWithDataForwardingFromTarget)
		if err := ie.PduSessionsReleasedListSNterminated.Decode(d); err != nil {
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
