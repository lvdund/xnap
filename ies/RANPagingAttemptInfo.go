package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	RANPagingAttemptInfoNextPagingAreaScopeSame    int64 = 0
	RANPagingAttemptInfoNextPagingAreaScopeChanged int64 = 1
)

var rANPagingAttemptInfoNextPagingAreaScopeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type RANPagingAttemptInfoNextPagingAreaScope struct {
	Value int64
}

func (ie *RANPagingAttemptInfoNextPagingAreaScope) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, rANPagingAttemptInfoNextPagingAreaScopeConstraints)
}

func (ie *RANPagingAttemptInfoNextPagingAreaScope) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(rANPagingAttemptInfoNextPagingAreaScopeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var rANPagingAttemptInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pagingAttemptCount"},
		{Name: "intendedNumberOfPagingAttempts"},
		{Name: "nextPagingAreaScope", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type RANPagingAttemptInfo struct {
	PagingAttemptCount             int64
	IntendedNumberOfPagingAttempts int64
	NextPagingAreaScope            *RANPagingAttemptInfoNextPagingAreaScope
	IEExtensions                   []byte
}

func (ie *RANPagingAttemptInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rANPagingAttemptInfoConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.NextPagingAreaScope != nil, false}); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.PagingAttemptCount, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(1)),
		UpperBound: common.Ptr(int64(16)),
	}); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.IntendedNumberOfPagingAttempts, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(1)),
		UpperBound: common.Ptr(int64(16)),
	}); err != nil {
		return err
	}
	if ie.NextPagingAreaScope != nil {
		if err := ie.NextPagingAreaScope.Encode(e); err != nil {
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

func (ie *RANPagingAttemptInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rANPagingAttemptInfoConstraints)
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
			UpperBound: common.Ptr(int64(16)),
		})
		if err != nil {
			return err
		}
		ie.PagingAttemptCount = val
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(1)),
			UpperBound: common.Ptr(int64(16)),
		})
		if err != nil {
			return err
		}
		ie.IntendedNumberOfPagingAttempts = val
	}
	if seq.IsComponentPresent(2) {
		ie.NextPagingAreaScope = new(RANPagingAttemptInfoNextPagingAreaScope)
		if err := ie.NextPagingAreaScope.Decode(d); err != nil {
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
