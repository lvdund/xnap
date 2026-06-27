package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	CauseProtocolTransferSyntaxError                          int64 = 0
	CauseProtocolAbstractSyntaxErrorReject                    int64 = 1
	CauseProtocolAbstractSyntaxErrorIgnoreAndNotify           int64 = 2
	CauseProtocolMessageNotCompatibleWithReceiverState        int64 = 3
	CauseProtocolSemanticError                                int64 = 4
	CauseProtocolAbstractSyntaxErrorFalselyConstructedMessage int64 = 5
	CauseProtocolUnspecified                                  int64 = 6
)

var causeProtocolConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6},
	ExtValues:  nil,
}

type CauseProtocol struct {
	Value int64
}

func (ie *CauseProtocol) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, causeProtocolConstraints)
}

func (ie *CauseProtocol) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(causeProtocolConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
