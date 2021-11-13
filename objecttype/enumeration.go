//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Any change will be overwritten                                                   *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package objecttype

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type ObjectType string
type ObjectTypes []ObjectType

func Slice(es ...ObjectType) ObjectTypes {
	var s ObjectTypes

	for _, e := range es {
		s = append(s, e)
	}

	return s
}

const (
	Array   ObjectType = "array"
	Boolean ObjectType = "boolean"
	Integer ObjectType = "integer"
	Null    ObjectType = "null"
	Number  ObjectType = "number"
	Object  ObjectType = "object"
	String  ObjectType = "string"
)

var (
	All = ObjectTypes{
		Array,
		Boolean,
		Integer,
		Null,
		Number,
		Object,
		String,
	}
)

func Is(v string) bool {
	return All.Is(v)
}

func Parse(v string) (ObjectType, error) {
	return All.Parse(v)
}

func Strings() []string {
	return All.Strings()
}

func (t ObjectType) String() string {
	return string(t)
}

var ErrObjectTypeInvalid = errors.New("invalid enumeration type")

func Error(vs ObjectTypes, v string) error {
	return fmt.Errorf(
		"%w '%v', must be one of %v",
		ErrObjectTypeInvalid, v, strings.Join(vs.Strings(), ","),
	)
}

func (t ObjectTypes) Strings() []string {
	var ss []string

	for _, v := range t {
		ss = append(ss, v.String())
	}

	return ss
}

func (t ObjectTypes) Parse(v string) (ObjectType, error) {
	var o ObjectType
	var f bool
	n := strings.ToLower(v)

	for _, e := range t {
		if strings.ToLower(e.String()) == n {
			o = e
			f = true
			break
		}
	}

	if !f {
		return o, Error(t, v)
	}

	return o, nil
}

func (t ObjectTypes) Is(v string) bool {
	var f bool

	for _, e := range t {
		if string(e) == v {
			f = true
			break
		}
	}

	return f
}

func (t ObjectTypes) Contains(v ObjectType) bool {
	for _, e := range t {
		if e == v {
			return true
		}
	}

	return false
}

func (t ObjectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *ObjectType) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	e, err := Parse(s)

	if err != nil {
		return err
	}

	*t = e

	return nil
}
