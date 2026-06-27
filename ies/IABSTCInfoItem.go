package ies

import (
	"github.com/lvdund/asn1go/per"
)

var iABSTCInfoItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sSB-freqInfo"},
		{Name: "sSB-subcarrierSpacing"},
		{Name: "sSB-transmissionPeriodicity"},
		{Name: "sSB-transmissionTimingOffset"},
		{Name: "sSB-transmissionBitmap"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type IABSTCInfoItem struct {
	SSBFreqInfo                 SSBFreqInfo
	SSBSubcarrierSpacing        SSBSubcarrierSpacing
	SSBTransmissionPeriodicity  SSBTransmissionPeriodicity
	SSBTransmissionTimingOffset SSBTransmissionTimingOffset
	SSBTransmissionBitmap       SSBTransmissionBitmap
	IEExtensions                []byte
}

func (ie *IABSTCInfoItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iABSTCInfoItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.SSBFreqInfo.Encode(e); err != nil {
		return err
	}
	if err := ie.SSBSubcarrierSpacing.Encode(e); err != nil {
		return err
	}
	if err := ie.SSBTransmissionPeriodicity.Encode(e); err != nil {
		return err
	}
	if err := ie.SSBTransmissionTimingOffset.Encode(e); err != nil {
		return err
	}
	if err := ie.SSBTransmissionBitmap.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *IABSTCInfoItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iABSTCInfoItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.SSBFreqInfo.Decode(d); err != nil {
		return err
	}
	if err := ie.SSBSubcarrierSpacing.Decode(d); err != nil {
		return err
	}
	if err := ie.SSBTransmissionPeriodicity.Decode(d); err != nil {
		return err
	}
	if err := ie.SSBTransmissionTimingOffset.Decode(d); err != nil {
		return err
	}
	if err := ie.SSBTransmissionBitmap.Decode(d); err != nil {
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
