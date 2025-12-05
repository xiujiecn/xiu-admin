package model

import (
	"encoding/base64"
	"encoding/json"
	"errors"
)

type MyBytes []byte

func (b MyBytes) MarshalJSON() ([]byte, error) {
	base64Str := base64.StdEncoding.EncodeToString(b)
	return json.Marshal(base64Str)
}

func (b *MyBytes) UnmarshalValue(data interface{}) error {
	str, ok := data.(string)
	if !ok {
		return errors.New("data is not a string")
	}
	bytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return err
	}
	*b = bytes
	return nil
}

func (b *MyBytes) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return (*b).UnmarshalString(s)
}

func (b *MyBytes) UnmarshalString(s string) error {
	bytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return err
	}
	*b = bytes
	return nil
}

func (b *MyBytes) String() string {
	return base64.StdEncoding.EncodeToString(*b)
}

func (b *MyBytes) Bytes() []byte {
	return *b
}
