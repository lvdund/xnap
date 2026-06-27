package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var freqDomainHSNAconfigurationListItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "rBsetIndex"},
		{Name: "freqDomainSlotHSNAconfiguration-List"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type FreqDomainHSNAconfigurationListItem struct {
	RBsetIndex                          int64
	FreqDomainSlotHSNAconfigurationList FreqDomainSlotHSNAconfigurationList
	IEExtensions                        []byte
}

func (ie *FreqDomainHSNAconfigurationListItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(freqDomainHSNAconfigurationListItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.RBsetIndex, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(0)),
		UpperBound: common.Ptr(int64(common.MaxnoofRBsetsPerCell1)),
	}); err != nil {
		return err
	}
	if err := ie.FreqDomainSlotHSNAconfigurationList.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *FreqDomainHSNAconfigurationListItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(freqDomainHSNAconfigurationListItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(common.MaxnoofRBsetsPerCell1)),
		})
		if err != nil {
			return err
		}
		ie.RBsetIndex = val
	}
	if err := ie.FreqDomainSlotHSNAconfigurationList.Decode(d); err != nil {
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
