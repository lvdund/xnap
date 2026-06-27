package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SecurityIndicationIntegrityProtectionIndicationRequired  int64 = 0
	SecurityIndicationIntegrityProtectionIndicationPreferred int64 = 1
	SecurityIndicationIntegrityProtectionIndicationNotNeeded int64 = 2
)

var securityIndicationIntegrityProtectionIndicationConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type SecurityIndicationIntegrityProtectionIndication struct {
	Value int64
}

func (ie *SecurityIndicationIntegrityProtectionIndication) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, securityIndicationIntegrityProtectionIndicationConstraints)
}

func (ie *SecurityIndicationIntegrityProtectionIndication) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(securityIndicationIntegrityProtectionIndicationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	SecurityIndicationConfidentialityProtectionIndicationRequired  int64 = 0
	SecurityIndicationConfidentialityProtectionIndicationPreferred int64 = 1
	SecurityIndicationConfidentialityProtectionIndicationNotNeeded int64 = 2
)

var securityIndicationConfidentialityProtectionIndicationConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type SecurityIndicationConfidentialityProtectionIndication struct {
	Value int64
}

func (ie *SecurityIndicationConfidentialityProtectionIndication) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, securityIndicationConfidentialityProtectionIndicationConstraints)
}

func (ie *SecurityIndicationConfidentialityProtectionIndication) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(securityIndicationConfidentialityProtectionIndicationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var securityIndicationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "integrityProtectionIndication"},
		{Name: "confidentialityProtectionIndication"},
		{Name: "maximumIPdatarate", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SecurityIndication struct {
	IntegrityProtectionIndication       SecurityIndicationIntegrityProtectionIndication
	ConfidentialityProtectionIndication SecurityIndicationConfidentialityProtectionIndication
	MaximumIPdatarate                   *MaximumIPdatarate
	IEExtensions                        []byte
}

func (ie *SecurityIndication) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(securityIndicationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.MaximumIPdatarate != nil, false}); err != nil {
		return err
	}
	if err := ie.IntegrityProtectionIndication.Encode(e); err != nil {
		return err
	}
	if err := ie.ConfidentialityProtectionIndication.Encode(e); err != nil {
		return err
	}
	if ie.MaximumIPdatarate != nil {
		if err := ie.MaximumIPdatarate.Encode(e); err != nil {
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

func (ie *SecurityIndication) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(securityIndicationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.IntegrityProtectionIndication.Decode(d); err != nil {
		return err
	}
	if err := ie.ConfidentialityProtectionIndication.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(2) {
		ie.MaximumIPdatarate = new(MaximumIPdatarate)
		if err := ie.MaximumIPdatarate.Decode(d); err != nil {
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
