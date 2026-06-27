package ies

import (
	"math"
	"math/big"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var volumeTimedReportItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "startTimeStamp"},
		{Name: "endTimeStamp"},
		{Name: "usageCountUL"},
		{Name: "usageCountDL"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

// usageCountConstraints models INTEGER (0..18446744073709551615) i.e. 0..2^64-1,
// an unsigned 64-bit range that exceeds int64. Use the BigInt encode/decode path.
var usageCountConstraints = per.ConstrainedBig(big.NewInt(0),
	new(big.Int).SetUint64(math.MaxUint64))

type VolumeTimedReportItem struct {
	StartTimeStamp []byte
	EndTimeStamp   []byte
	UsageCountUL   *big.Int
	UsageCountDL   *big.Int
	IEExtensions   []byte
}

func (ie *VolumeTimedReportItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(volumeTimedReportItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := e.EncodeOctetString(ie.StartTimeStamp, per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(4)),
		Max:        common.Ptr(int64(4)),
	}); err != nil {
		return err
	}
	if err := e.EncodeOctetString(ie.EndTimeStamp, per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(4)),
		Max:        common.Ptr(int64(4)),
	}); err != nil {
		return err
	}
	if err := e.EncodeBigInteger(ie.UsageCountUL, usageCountConstraints); err != nil {
		return err
	}
	if err := e.EncodeBigInteger(ie.UsageCountDL, usageCountConstraints); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *VolumeTimedReportItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(volumeTimedReportItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	{
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(4)),
			Max:        common.Ptr(int64(4)),
		})
		if err != nil {
			return err
		}
		ie.StartTimeStamp = val
	}
	{
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(4)),
			Max:        common.Ptr(int64(4)),
		})
		if err != nil {
			return err
		}
		ie.EndTimeStamp = val
	}
	{
		val, err := d.DecodeBigInteger(usageCountConstraints)
		if err != nil {
			return err
		}
		ie.UsageCountUL = val
	}
	{
		val, err := d.DecodeBigInteger(usageCountConstraints)
		if err != nil {
			return err
		}
		ie.UsageCountDL = val
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
