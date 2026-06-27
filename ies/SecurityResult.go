package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SecurityResultIntegrityProtectionResultPerformed    int64 = 0
	SecurityResultIntegrityProtectionResultNotPerformed int64 = 1
)

var securityResultIntegrityProtectionResultConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type SecurityResultIntegrityProtectionResult struct {
	Value int64
}

func (ie *SecurityResultIntegrityProtectionResult) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, securityResultIntegrityProtectionResultConstraints)
}

func (ie *SecurityResultIntegrityProtectionResult) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(securityResultIntegrityProtectionResultConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	SecurityResultConfidentialityProtectionResultPerformed    int64 = 0
	SecurityResultConfidentialityProtectionResultNotPerformed int64 = 1
)

var securityResultConfidentialityProtectionResultConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type SecurityResultConfidentialityProtectionResult struct {
	Value int64
}

func (ie *SecurityResultConfidentialityProtectionResult) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, securityResultConfidentialityProtectionResultConstraints)
}

func (ie *SecurityResultConfidentialityProtectionResult) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(securityResultConfidentialityProtectionResultConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var securityResultConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "integrityProtectionResult"},
		{Name: "confidentialityProtectionResult"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SecurityResult struct {
	IntegrityProtectionResult       SecurityResultIntegrityProtectionResult
	ConfidentialityProtectionResult SecurityResultConfidentialityProtectionResult
	IEExtensions                    []byte
}

func (ie *SecurityResult) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(securityResultConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.IntegrityProtectionResult.Encode(e); err != nil {
		return err
	}
	if err := ie.ConfidentialityProtectionResult.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SecurityResult) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(securityResultConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.IntegrityProtectionResult.Decode(d); err != nil {
		return err
	}
	if err := ie.ConfidentialityProtectionResult.Decode(d); err != nil {
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
