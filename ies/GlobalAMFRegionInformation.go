package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var globalAMFRegionInformationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-ID"},
		{Name: "amf-region-id"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type GlobalAMFRegionInformation struct {
	PlmnID       PLMNIdentity
	AmfRegionId  per.BitString
	IEExtensions []byte
}

func (ie *GlobalAMFRegionInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(globalAMFRegionInformationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.PlmnID.Encode(e); err != nil {
		return err
	}
	if err := e.EncodeBitString(ie.AmfRegionId, per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(8)),
		Max:        common.Ptr(int64(8)),
	}); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *GlobalAMFRegionInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(globalAMFRegionInformationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PlmnID.Decode(d); err != nil {
		return err
	}
	{
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(8)),
			Max:        common.Ptr(int64(8)),
		})
		if err != nil {
			return err
		}
		ie.AmfRegionId = val
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
