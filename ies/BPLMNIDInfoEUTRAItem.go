package ies

import (
	"github.com/lvdund/asn1go/per"
)

var bPLMNIDInfoEUTRAItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "broadcastPLMNs"},
		{Name: "tac"},
		{Name: "e-utraCI"},
		{Name: "ranac", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type BPLMNIDInfoEUTRAItem struct {
	BroadcastPLMNs BroadcastEUTRAPLMNs
	Tac            TAC
	EUtraCI        EUTRACellIdentity
	Ranac          *RANAC
	IEExtensions   []byte
}

func (ie *BPLMNIDInfoEUTRAItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bPLMNIDInfoEUTRAItemConstraints)
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
	if err := ie.EUtraCI.Encode(e); err != nil {
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

func (ie *BPLMNIDInfoEUTRAItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bPLMNIDInfoEUTRAItemConstraints)
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
	if err := ie.EUtraCI.Decode(d); err != nil {
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
