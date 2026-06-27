package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var freqDomainSlotHSNAconfigurationListItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "slotIndex"},
		{Name: "hSNADownlink", Optional: true},
		{Name: "hSNAUplink", Optional: true},
		{Name: "hSNAFlexible", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type FreqDomainSlotHSNAconfigurationListItem struct {
	SlotIndex    int64
	HSNADownlink *HSNADownlink
	HSNAUplink   *HSNAUplink
	HSNAFlexible *HSNAFlexible
	IEExtensions []byte
}

func (ie *FreqDomainSlotHSNAconfigurationListItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(freqDomainSlotHSNAconfigurationListItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.HSNADownlink != nil, ie.HSNAUplink != nil, ie.HSNAFlexible != nil, false}); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.SlotIndex, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(1)),
		UpperBound: common.Ptr(int64(common.MaxnoofHSNASlots)),
	}); err != nil {
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

func (ie *FreqDomainSlotHSNAconfigurationListItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(freqDomainSlotHSNAconfigurationListItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(1)),
			UpperBound: common.Ptr(int64(common.MaxnoofHSNASlots)),
		})
		if err != nil {
			return err
		}
		ie.SlotIndex = val
	}
	if seq.IsComponentPresent(1) {
		ie.HSNADownlink = new(HSNADownlink)
		if err := ie.HSNADownlink.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.HSNAUplink = new(HSNAUplink)
		if err := ie.HSNAUplink.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
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
