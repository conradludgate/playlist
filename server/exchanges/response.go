package exchanges

import (
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

type FFJSONMarshaler interface {
	MarshalJSONBuf(buf fflib.EncodingBuffer) error
}

// ffjson: skip
type Response struct {
	Status  int
	Success bool
	Value   FFJSONMarshaler
	Error   string
}

// MarshalJSON marshal bytes to json - template
func (r *Response) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if r == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := r.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (r *Response) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if r == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"status":`)
	fflib.FormatBits2(buf, uint64(r.Status), 10, r.Status < 0)
	if r.Success {
		buf.WriteString(`,"success":true`)
	} else {
		buf.WriteString(`,"success":false`)
	}

	buf.WriteString(`,"value":`)

	if r.Value != nil {
		err = r.Value.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}
	} else {
		buf.WriteString("null")
	}

	buf.WriteString(`,"error":`)
	fflib.WriteJsonString(buf, string(r.Error))
	buf.WriteByte('}')
	return nil
}
