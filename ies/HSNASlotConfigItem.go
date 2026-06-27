package ies

import (
	"github.com/lvdund/asn1go/per"
)

var hSNASlotConfigItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "hSNADownlink", Optional: true},
		{Name: "hSNAUplink", Optional: true},
		{Name: "hSNAFlexible", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type HSNASlotConfigItem struct {
	HSNADownlink *HSNADownlink
	HSNAUplink   *HSNAUplink
	HSNAFlexible *HSNAFlexible
	IEExtensions []byte
}

func (ie *HSNASlotConfigItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(hSNASlotConfigItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.HSNADownlink != nil, ie.HSNAUplink != nil, ie.HSNAFlexible != nil, false}); err != nil {
		return err
	}
	if ie.HSNADownlink != nil {
		if err := ie.HSNADownlink.Encode(e); err != nil {
			return err
		}
	}
	if ie.HSNAUplink != nil {
		if err := ie.HSNAUplink.Encode(e); err != nil {
			return err
		}
	}
	if ie.HSNAFlexible != nil {
		if err := ie.HSNAFlexible.Encode(e); err != nil {
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

func (ie *HSNASlotConfigItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(hSNASlotConfigItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.HSNADownlink = new(HSNADownlink)
		if err := ie.HSNADownlink.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.HSNAUplink = new(HSNAUplink)
		if err := ie.HSNAUplink.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.HSNAFlexible = new(HSNAFlexible)
		if err := ie.HSNAFlexible.Decode(d); err != nil {
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
