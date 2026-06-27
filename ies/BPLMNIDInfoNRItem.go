package ies

import (
	"github.com/lvdund/asn1go/per"
)

var bPLMNIDInfoNRItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "broadcastPLMNs"},
		{Name: "tac"},
		{Name: "nr-CI"},
		{Name: "ranac", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type BPLMNIDInfoNRItem struct {
	BroadcastPLMNs BroadcastPLMNs
	Tac            TAC
	NrCI           NRCellIdentity
	Ranac          *RANAC
	IEExtensions   []byte
}

func (ie *BPLMNIDInfoNRItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bPLMNIDInfoNRItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.Ranac != nil, false}); err != nil {
		return err
	}
	if err := ie.BroadcastPLMNs.Encode(e); err != nil {
		return err
	}
	if err := ie.Tac.Encode(e); err != nil {
		return err
	}
	if err := ie.NrCI.Encode(e); err != nil {
		return err
	}
	if ie.Ranac != nil {
		if err := ie.Ranac.Encode(e); err != nil {
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

func (ie *BPLMNIDInfoNRItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bPLMNIDInfoNRItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.BroadcastPLMNs.Decode(d); err != nil {
		return err
	}
	if err := ie.Tac.Decode(d); err != nil {
		return err
	}
	if err := ie.NrCI.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(3) {
		ie.Ranac = new(RANAC)
		if err := ie.Ranac.Decode(d); err != nil {
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
