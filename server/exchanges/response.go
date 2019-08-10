package exchanges

import (
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// FFJSONMarshaler is an interface to any type that is quickly converted to json
type FFJSONMarshaler interface {
	MarshalJSON() ([]byte, error)
	MarshalJSONBuf(fflib.EncodingBuffer) error
}

// Response is the type that all http requests get sent back
type Response struct {
	Status  int
	Success bool
	Value   FFJSONMarshaler
	Error   error
}

// MarshalJSON turns a response object into it's JSON form as bytes
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

// MarshalJSONBuf writes the object to a buffer as json
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
	if r.Error != nil {
		fflib.WriteJsonString(buf, string(r.Error.Error()))
	} else {
		buf.WriteString("null")
	}

	buf.WriteByte('}')
	return nil
}
