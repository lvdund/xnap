package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// Xn Message IE
type IE interface {
	Encode(w *per.Encoder) error
	Decode(r *per.Decoder) error
}

type XnAPMessageIE struct {
	ID          ProtocolIEID
	Criticality Criticality
	Value       IE
}

func (ie *XnAPMessageIE) Encode(e *per.Encoder) error {
	if err := e.EncodeInteger(ie.ID.Value, per.Constrained(0, 65535)); err != nil {
		return fmt.Errorf("encode IE ID failed: %w", err)
	}
	if err := ie.Criticality.Encode(e); err != nil {
		return fmt.Errorf("encode IE Criticality failed: %w", err)
	}
	valEncoder := per.NewEncoder(per.APER)
	if err := ie.Value.Encode(valEncoder); err != nil {
		return fmt.Errorf("encode IE Value failed: %w", err)
	}
	if err := e.EncodeOpenType(valEncoder.Bytes()); err != nil {
		return fmt.Errorf("encode IE Value as OpenType failed: %w", err)
	}
	return nil
}

func (ie *XnAPMessageIE) Decode(d *per.Decoder) error {
	return fmt.Errorf("E1APMessageIE.Decode should not be called directly")
}

var choiceTypeXnMsg []per.AlternativeInfo = []per.AlternativeInfo{
	{Name: "initiatingMessage", Tag: int(common.InitiatingMessage)},
	{Name: "successfulOutcome", Tag: int(common.SuccessfulOutcome)},
	{Name: "unsuccessfulOutcome", Tag: int(common.UnsuccessfulOutcome)},
}

func choiceXnMsg(e *per.Encoder, c int64) error {
	choice := e.NewChoiceEncoder(per.ChoiceConstraints{
		Extensible:       true,
		RootAlternatives: choiceTypeXnMsg,
	})
	return choice.EncodeChoice(c, false, nil)
}

func procedureCodeXnMsg(e *per.Encoder, c int64) error {
	procedureCode := ProcedureCode{Value: c}
	return procedureCode.Encode(e)
}

func criticalityXnMsg(e *per.Encoder, c int64) error {
	criticality := Criticality{Value: c}
	return criticality.Encode(e)
}

type xnMsgWithIEs interface {
	toIEs() []XnAPMessageIE
}

func encodeBody[T xnMsgWithIEs](msg T, e *per.Encoder) error {
	seq := e.NewSequenceEncoder(per.SequenceConstraints{Extensible: true})
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	ies := msg.toIEs()
	seqOf := e.NewSequenceOfEncoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	if err := seqOf.EncodeLength(int64(len(ies))); err != nil {
		return err
	}
	for i := range ies {
		if err := ies[i].Encode(e); err != nil {
			return err
		}
	}

	return nil
}

/////////////////////////

type integerIE struct {
	Value int64
	c     per.IntegerConstraints
}

func NewINTEGER(v int64, c per.IntegerConstraints) integerIE {
	return integerIE{Value: v, c: c}
}
func (t *integerIE) Encode(e *per.Encoder) (err error) { return e.EncodeInteger(t.Value, t.c) }
func (t *integerIE) Decode(d *per.Decoder) (err error) {
	v, err := d.DecodeInteger(t.c)
	if err != nil {
		return
	}
	t.Value = v
	return
}

type octetStringIE struct {
	Value []byte
	c     per.SizeConstraints
}

func NewOCTETSTRING(v []byte, c per.SizeConstraints) octetStringIE {
	return octetStringIE{Value: v, c: c}
}

func (t *octetStringIE) Encode(e *per.Encoder) (err error) { return e.EncodeOctetString(t.Value, t.c) }

func (t *octetStringIE) Decode(d *per.Decoder) (err error) {
	v, err := d.DecodeOctetString(t.c)
	if err != nil {
		return
	}
	t.Value = v
	return
}

type bitStringIE struct {
	Value per.BitString
	c     per.SizeConstraints
}

func NewBITSTRING(v per.BitString, c per.SizeConstraints) bitStringIE {
	return bitStringIE{Value: v, c: c}
}
func (t *bitStringIE) Encode(e *per.Encoder) (err error) { return e.EncodeBitString(t.Value, t.c) }
func (t *bitStringIE) Decode(d *per.Decoder) (err error) {
	v, err := d.DecodeBitString(t.c)
	if err != nil {
		return
	}
	t.Value = v
	return
}

type enumeratedIE struct {
	Value int64
	c     per.EnumeratedConstraints
}

func NewENUMERATED(v int64, c per.EnumeratedConstraints) enumeratedIE {
	return enumeratedIE{Value: v, c: c}
}

func (t *enumeratedIE) Encode(e *per.Encoder) (err error) { return e.EncodeEnumerated(t.Value, t.c) }

func (t *enumeratedIE) Decode(d *per.Decoder) (err error) {
	v, err := d.DecodeEnumerated(t.c)
	if err != nil {
		return
	}
	t.Value = v
	return
}

///////////////////

const (
	MaxPrivateIEs         uint = 65535
	MaxProtocolExtensions uint = 65535
	MaxProtocolIEs        uint = 65535
)

// asn1: Criticality ::= ENUMERATED { reject, ignore, notify }
type Criticality struct {
	Value int64
}

const (
	CriticalityReject int64 = 0
	CriticalityIgnore int64 = 1
	CriticalityNotify int64 = 2
)

var criticalityConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

func (e *Criticality) Encode(w *per.Encoder) error {
	return w.EncodeEnumerated(e.Value, criticalityConstraints)
}

func (e *Criticality) Decode(r *per.Decoder) error {
	val, err := r.DecodeEnumerated(criticalityConstraints)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}

// asn1: Presence ::= ENUMERATED { optional, conditional, mandatory }
type Presence struct {
	Value int64
}

const (
	PresenceOptional    int64 = 0
	PresenceConditional int64 = 1
	PresenceMandatory   int64 = 2
)

var presenceConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{
		PresenceOptional,
		PresenceConditional,
		PresenceMandatory,
	},
	ExtValues: nil,
}

func (e *Presence) Encode(w *per.Encoder) error {
	return w.EncodeEnumerated(e.Value, presenceConstraints)
}

func (e *Presence) Decode(r *per.Decoder) error {
	val, err := r.DecodeEnumerated(presenceConstraints)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}

// asn1: ProcedureCode ::= INTEGER (0..255)
type ProcedureCode struct {
	Value int64 `asn1:"0..255"`
}

var procedureCodeConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(255)),
}

func (e *ProcedureCode) Encode(w *per.Encoder) error {
	return w.EncodeInteger(e.Value, procedureCodeConstraints)
}

func (e *ProcedureCode) Decode(r *per.Decoder) error {
	val, err := r.DecodeInteger(procedureCodeConstraints)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}

// asn1: ProtocolIE-ID ::= INTEGER (0..maxProtocolIEs)
type ProtocolIEID struct {
	Value int64 `asn1:"0..maxProtocolIEs"`
}

var protocolIEIDConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(MaxProtocolIEs)),
}

func (e *ProtocolIEID) Encode(w *per.Encoder) error {
	return w.EncodeInteger(e.Value, protocolIEIDConstraints)
}

func (e *ProtocolIEID) Decode(r *per.Decoder) error {
	val, err := r.DecodeInteger(protocolIEIDConstraints)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}

// asn1: TriggeringMessage ::= ENUMERATED { initiating-message, successful-outcome, unsuccessful-outcome }
type TriggeringMessage struct {
	Value int64
}

const (
	TriggeringMessageInitiatingMessage   int64 = 0
	TriggeringMessageSuccessfulOutcome   int64 = 1
	TriggeringMessageUnsuccessfulOutcome int64 = 2
)

var triggeringMessageConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{
		TriggeringMessageInitiatingMessage,
		TriggeringMessageSuccessfulOutcome,
		TriggeringMessageUnsuccessfulOutcome,
	},
	ExtValues: nil,
}

func (e *TriggeringMessage) Encode(w *per.Encoder) error {
	return w.EncodeEnumerated(e.Value, triggeringMessageConstraints)
}

func (e *TriggeringMessage) Decode(r *per.Decoder) error {
	val, err := r.DecodeEnumerated(triggeringMessageConstraints)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
