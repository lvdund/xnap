package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var bluetoothNameConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(248)),
}

type BluetoothName struct {
	Value []byte
}

func (ie *BluetoothName) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, bluetoothNameConstraints)
}

func (ie *BluetoothName) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(bluetoothNameConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
