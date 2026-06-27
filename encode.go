package xn

import "bytes"

func XnEncode(msg XnApMessageEncoder) (wire []byte, err error) {
	var buf bytes.Buffer
	if err = msg.Encode(&buf); err == nil {
		wire = buf.Bytes()
	}
	return
}
