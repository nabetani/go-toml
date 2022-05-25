package toml_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/pelletier/go-toml/v2"
)

var simpleTomlI = []byte(`Foo=123`)
var simpleTomlF = []byte(`Foo=123.0`)

type Float64Values struct {
	Foo float64
}

type Float64PtrValues struct {
	Foo *float64
}

type Float32Values struct {
	Foo float64
}

type Float32PtrValues struct {
	Foo *float64
}

type Int64Values struct {
	Foo int64
}

type Uint64Values struct {
	Foo uint64
}

type Int32Values struct {
	Foo int32
}

type Uint32Values struct {
	Foo uint32
}

type Int16Values struct {
	Foo int16
}

type Uint16Values struct {
	Foo uint16
}
type Int8Values struct {
	Foo int8
}

type Uint8Values struct {
	Foo uint8
}

type Int64PtrValues struct {
	Foo *int64
}

func TestReadAsFloat(t *testing.T) {
	v := Float64Values{}
	err := toml.Unmarshal(simpleTomlI, &v)
	if nil == err {
		t.Error("err should not be nil")
	}
}

func TestReadAsFloatWithLaxNumericType(t *testing.T) {
	v := Float64Values{}
	opts := toml.DecorderOpts{LaxNumericType: true}
	err := toml.UnmarshalWithOpts(simpleTomlI, &v, opts)
	if nil != err {
		t.Errorf("err is %q want nil", err)
	}
	if v.Foo != 123 {
		t.Errorf("v.Foo=%v, want 123", v.Foo)
	}
}

func TestReadAsFloatPtrWithLaxNumericType(t *testing.T) {
	v := Float64PtrValues{}
	opts := toml.DecorderOpts{LaxNumericType: true}
	err := toml.UnmarshalWithOpts(simpleTomlI, &v, opts)
	if nil != err {
		t.Errorf("err is %q want nil", err)
	}
	if *v.Foo != 123 {
		t.Errorf("v.Foo=%v, want 123", *v.Foo)
	}
}

func TestReadAsFloat32WithLaxNumericType(t *testing.T) {
	v := Float32Values{}
	opts := toml.DecorderOpts{LaxNumericType: true}
	err := toml.UnmarshalWithOpts(simpleTomlI, &v, opts)
	if nil != err {
		t.Errorf("err is %q want nil", err)
	}
	if v.Foo != 123 {
		t.Errorf("v.Foo=%v, want 123", v.Foo)
	}
}

func TestReadAsFloat32PtrWithLaxNumericType(t *testing.T) {
	v := Float32PtrValues{}
	opts := toml.DecorderOpts{LaxNumericType: true}
	err := toml.UnmarshalWithOpts(simpleTomlI, &v, opts)
	if nil != err {
		t.Errorf("err is %q want nil", err)
	}
	if *v.Foo != 123 {
		t.Errorf("v.Foo=%v, want 123", *v.Foo)
	}
}

func TestReadAsInt64WithLaxNumericType(t *testing.T) {
	v := Int64Values{}
	opts := toml.DecorderOpts{LaxNumericType: true}
	err := toml.UnmarshalWithOpts(simpleTomlF, &v, opts)
	if nil != err {
		t.Errorf("err is %q want nil", err)
	}
	if v.Foo != 123 {
		t.Errorf("v.Foo=%v, want 123", v.Foo)
	}
}

func TestReadAsInt64PtrWithLaxNumericType(t *testing.T) {
	v := Int64PtrValues{}
	opts := toml.DecorderOpts{LaxNumericType: true}
	err := toml.UnmarshalWithOpts(simpleTomlI, &v, opts)
	if nil != err {
		t.Errorf("err is %q want nil", err)
	}
	if *v.Foo != 123 {
		t.Errorf("v.Foo=%v, want 123", *v.Foo)
	}
}

var viciousToml = []byte(`
I64=-64.0
U64=64.0
I32=-32.0
U32=32.0
I16=-16.0
U16=16.0
I8=-8.0
U8=8.0
F64=12
F32=34
`)

type ManyNumTypes struct {
	I64 int64
	U64 uint64
	I32 int32
	U32 uint32
	I16 int16
	U16 uint16
	I8  int8
	U8  uint8
	F64 float64
	F32 float32
}

func TestReadManyNumTypesWithLaxNumericType(t *testing.T) {
	v := ManyNumTypes{}
	opts := toml.DecorderOpts{LaxNumericType: true}
	err := toml.UnmarshalWithOpts(viciousToml, &v, opts)
	if nil != err {
		t.Errorf("err is %q want nil", err)
	}
	jbytes, err := json.Marshal(v)
	jstr := string(jbytes)

	expect := "{\"I64\":-64,\"U64\":64,\"I32\":-32,\"U32\":32,\"I16\":-16,\"U16\":16,\"I8\":-8,\"U8\":8,\"F64\":12,\"F32\":34}"
	if jstr != expect {
		t.Errorf("jst=%q, want %q", jstr, expect)
	}
}

var complexToml = []byte(`
Foo=12
[Bar]
Grault=23
Baz.Qux=34
Baz.Quux=45.0
Baz.Corge=[0x56, 78e2]
[Bar.Garply]
Waldo=123
Fred=234.0
[Bar.Plugh]
Xyzzy=345
Thud=456.0
`)

type BazF struct {
	Qux, Quux float64
	Corge     []float64
}
type BarF struct {
	Grault *float64
	Baz    BazF
	Garply map[string]float64
	Plugh  map[string]*float64
}

type ComplexF struct {
	Foo float64
	Bar BarF
}

type BazI struct {
	Qux, Quux int64
	Corge     []int64
}
type BarI struct {
	Grault *int64
	Baz    BazI
	Garply map[string]int64
	Plugh  map[string]*int64
}

type ComplexI struct {
	Foo int64
	Bar BarI
}

func TestReadComplexValueAsFloat64WithLaxNumericType(t *testing.T) {
	v := ComplexF{}
	opts := toml.DecorderOpts{LaxNumericType: true}
	err := toml.UnmarshalWithOpts(complexToml, &v, opts)
	if nil != err {
		t.Errorf("err is %q want nil", err)
	}
	jbytes, err := json.Marshal(v)
	jstr := string(jbytes)
	expect := "{\"Foo\":12,\"Bar\":{\"Grault\":23,\"Baz\":{\"Qux\":34,\"Quux\":45,\"Corge\":[86,7800]},\"Garply\":{\"Fred\":234,\"Waldo\":123},\"Plugh\":{\"Thud\":456,\"Xyzzy\":345}}}"
	if jstr != expect {
		t.Errorf("jst=%q, want %q", jstr, expect)
	}
}
func TestReadComplexValueAsFloat64WithoutLaxNumericType(t *testing.T) {
	v := ComplexF{}
	opts := toml.DecorderOpts{}
	err := toml.UnmarshalWithOpts(complexToml, &v, opts)
	if nil == err {
		t.Errorf("err should not be nil")
	}
}

func TestReadComplexValueAsInt64WithLaxNumericType(t *testing.T) {
	v := ComplexI{}
	opts := toml.DecorderOpts{LaxNumericType: true}
	err := toml.UnmarshalWithOpts(complexToml, &v, opts)
	if nil != err {
		t.Errorf("err is %q want nil", err)
	}
	jbytes, err := json.Marshal(v)
	jstr := string(jbytes)
	expect := "{\"Foo\":12,\"Bar\":{\"Grault\":23,\"Baz\":{\"Qux\":34,\"Quux\":45,\"Corge\":[86,7800]},\"Garply\":{\"Fred\":234,\"Waldo\":123},\"Plugh\":{\"Thud\":456,\"Xyzzy\":345}}}"
	if jstr != expect {
		t.Errorf("jst=%q, want %q", jstr, expect)
	}
}

func TestReadComplexValueAsInt64WithoutLaxNumericType(t *testing.T) {
	v := ComplexI{}
	opts := toml.DecorderOpts{}
	err := toml.UnmarshalWithOpts(complexToml, &v, opts)
	if nil == err {
		t.Errorf("err should not be nil")
	}
}

func TestReadHugeIntAsInt64(t *testing.T) {
	val := int64(0x5555_5555_5555_5555)
	tomlStr := []byte(fmt.Sprintf("Foo=%d.0", val))
	v := Int64Values{}
	err := toml.UnmarshalWithOpts(tomlStr, &v, toml.DecorderOpts{LaxNumericType: true})
	if err != nil {
		t.Errorf("err=%q, want nil", err)
	}
	if v.Foo != val {
		t.Errorf("v.Foo=%v, want %v (toml is %q)", v.Foo, val, tomlStr)
	}
}
func TestReadHugeIntAsInt64ExpFormat(t *testing.T) {
	val := int64(-8608480567731124087)
	tomlStr := []byte("Foo=-8.608480567731124087000e+18")
	v := Int64Values{}
	err := toml.UnmarshalWithOpts(tomlStr, &v, toml.DecorderOpts{LaxNumericType: true})
	if err != nil {
		t.Errorf("err=%q, want nil", err)
	}
	if v.Foo != val {
		t.Errorf("v.Foo=%v, want %v (toml is %q)", v.Foo, val, tomlStr)
	}
}

func TestReadHugeIntAsUint64(t *testing.T) {
	val := uint64(0xF555_5555_5555_5555)
	tomlStr := []byte(fmt.Sprintf("Foo=%d.0", val))
	v := Uint64Values{}
	err := toml.UnmarshalWithOpts(tomlStr, &v, toml.DecorderOpts{LaxNumericType: true})
	if err != nil {
		t.Errorf("err=%q, want nil", err)
	}
	if v.Foo != val {
		t.Errorf("v.Foo=%v, want %v (toml is %q)", v.Foo, val, tomlStr)
	}
}

func TestReadHugeIntAsInt32(t *testing.T) {
	val := int32(0x5555_5555)
	tomlStr := []byte(fmt.Sprintf("Foo=%d.0", val))
	v := Int32Values{}
	err := toml.UnmarshalWithOpts(tomlStr, &v, toml.DecorderOpts{LaxNumericType: true})
	if err != nil {
		t.Errorf("err=%q, want nil", err)
	}
	if v.Foo != val {
		t.Errorf("v.Foo=%v, want %v (toml is %q)", v.Foo, val, tomlStr)
	}
}
func TestReadHugeIntAsUint32(t *testing.T) {
	val := uint32(0xFAFA_CAFE)
	tomlStr := []byte(fmt.Sprintf("Foo=%d.0", val))
	v := Uint32Values{}
	err := toml.UnmarshalWithOpts(tomlStr, &v, toml.DecorderOpts{LaxNumericType: true})
	if err != nil {
		t.Errorf("err=%q, want nil", err)
	}
	if v.Foo != val {
		t.Errorf("v.Foo=%v, want %v (toml is %q)", v.Foo, val, tomlStr)
	}
}

func TestReadHugeIntAsInt16(t *testing.T) {
	val := int16(0x5555)
	tomlStr := []byte(fmt.Sprintf("Foo=%d.0", val))
	v := Int16Values{}
	err := toml.UnmarshalWithOpts(tomlStr, &v, toml.DecorderOpts{LaxNumericType: true})
	if err != nil {
		t.Errorf("err=%q, want nil", err)
	}
	if v.Foo != val {
		t.Errorf("v.Foo=%v, want %v (toml is %q)", v.Foo, val, tomlStr)
	}
}
func TestReadHugeIntAsUint16(t *testing.T) {
	val := uint16(0xCAFE)
	tomlStr := []byte(fmt.Sprintf("Foo=%d.0", val))
	v := Uint16Values{}
	err := toml.UnmarshalWithOpts(tomlStr, &v, toml.DecorderOpts{LaxNumericType: true})
	if err != nil {
		t.Errorf("err=%q, want nil", err)
	}
	if v.Foo != val {
		t.Errorf("v.Foo=%v, want %v (toml is %q)", v.Foo, val, tomlStr)
	}
}

func TestReadHugeIntAsInt8(t *testing.T) {
	val := int8(0x55)
	tomlStr := []byte(fmt.Sprintf("Foo=%d.0", val))
	v := Int8Values{}
	err := toml.UnmarshalWithOpts(tomlStr, &v, toml.DecorderOpts{LaxNumericType: true})
	if err != nil {
		t.Errorf("err=%q, want nil", err)
	}
	if v.Foo != val {
		t.Errorf("v.Foo=%v, want %v (toml is %q)", v.Foo, val, tomlStr)
	}
}
func TestReadHugeIntAsUint8(t *testing.T) {
	val := uint8(0xCA)
	tomlStr := []byte(fmt.Sprintf("Foo=%d.0", val))
	v := Uint8Values{}
	err := toml.UnmarshalWithOpts(tomlStr, &v, toml.DecorderOpts{LaxNumericType: true})
	if err != nil {
		t.Errorf("err=%q, want nil", err)
	}
	if v.Foo != val {
		t.Errorf("v.Foo=%v, want %v (toml is %q)", v.Foo, val, tomlStr)
	}
}
