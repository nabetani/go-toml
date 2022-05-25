package toml

import (
	"encoding/json"
	"testing"
)

var simpleToml = []byte(`Foo=123`)

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

type Int64PtrValues struct {
	Foo *int64
}

func TestReadAsFloat(t *testing.T) {
	v := Float64Values{}
	err := Unmarshal(simpleToml, &v)
	if nil == err {
		t.Error("err should not be nil")
	}
}

func TestReadAsFloatWithLaxNumericType(t *testing.T) {
	v := Float64Values{}
	opts := DecorderOpts{LaxNumericType: true}
	err := UnmarshalWithOpts(simpleToml, &v, opts)
	if nil != err {
		t.Errorf("err is %q want nil", err)
	}
	if v.Foo != 123 {
		t.Errorf("v.Foo=%v, want 123", v.Foo)
	}
}

func TestReadAsFloatPtrWithLaxNumericType(t *testing.T) {
	v := Float64PtrValues{}
	opts := DecorderOpts{LaxNumericType: true}
	err := UnmarshalWithOpts(simpleToml, &v, opts)
	if nil != err {
		t.Errorf("err is %q want nil", err)
	}
	if *v.Foo != 123 {
		t.Errorf("v.Foo=%v, want 123", *v.Foo)
	}
}

func TestReadAsFloat32WithLaxNumericType(t *testing.T) {
	v := Float32Values{}
	opts := DecorderOpts{LaxNumericType: true}
	err := UnmarshalWithOpts(simpleToml, &v, opts)
	if nil != err {
		t.Errorf("err is %q want nil", err)
	}
	if v.Foo != 123 {
		t.Errorf("v.Foo=%v, want 123", v.Foo)
	}
}

func TestReadAsFloat32PtrWithLaxNumericType(t *testing.T) {
	v := Float32PtrValues{}
	opts := DecorderOpts{LaxNumericType: true}
	err := UnmarshalWithOpts(simpleToml, &v, opts)
	if nil != err {
		t.Errorf("err is %q want nil", err)
	}
	if *v.Foo != 123 {
		t.Errorf("v.Foo=%v, want 123", *v.Foo)
	}
}

func TestReadAsInt64WithLaxNumericType(t *testing.T) {
	v := Int64Values{}
	opts := DecorderOpts{LaxNumericType: true}
	err := UnmarshalWithOpts(simpleToml, &v, opts)
	if nil != err {
		t.Errorf("err is %q want nil", err)
	}
	if v.Foo != 123 {
		t.Errorf("v.Foo=%v, want 123", v.Foo)
	}
}

func TestReadAsInt64PtrWithLaxNumericType(t *testing.T) {
	v := Int64PtrValues{}
	opts := DecorderOpts{LaxNumericType: true}
	err := UnmarshalWithOpts(simpleToml, &v, opts)
	if nil != err {
		t.Errorf("err is %q want nil", err)
	}
	if *v.Foo != 123 {
		t.Errorf("v.Foo=%v, want 123", *v.Foo)
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
	opts := DecorderOpts{LaxNumericType: true}
	err := UnmarshalWithOpts(complexToml, &v, opts)
	if nil != err {
		t.Errorf("err is %q want nil", err)
	}
	jbytes, err := json.Marshal(v)
	jstr := string(jbytes)
	expects := "{\"Foo\":12,\"Bar\":{\"Grault\":23,\"Baz\":{\"Qux\":34,\"Quux\":45,\"Corge\":[86,7800]},\"Garply\":{\"Fred\":234,\"Waldo\":123},\"Plugh\":{\"Thud\":456,\"Xyzzy\":345}}}"
	if jstr != expects {
		t.Errorf("jst=%q, want %q", jstr, expects)
	}
}
func TestReadComplexValueAsFloat64WithoutLaxNumericType(t *testing.T) {
	v := ComplexF{}
	opts := DecorderOpts{}
	err := UnmarshalWithOpts(complexToml, &v, opts)
	if nil == err {
		t.Errorf("err should not be nil")
	}
}

func TestReadComplexValueAsInt64WithLaxNumericType(t *testing.T) {
	v := ComplexI{}
	opts := DecorderOpts{LaxNumericType: true}
	err := UnmarshalWithOpts(complexToml, &v, opts)
	if nil != err {
		t.Errorf("err is %q want nil", err)
	}
	jbytes, err := json.Marshal(v)
	jstr := string(jbytes)
	expects := "{\"Foo\":12,\"Bar\":{\"Grault\":23,\"Baz\":{\"Qux\":34,\"Quux\":45,\"Corge\":[86,7800]},\"Garply\":{\"Fred\":234,\"Waldo\":123},\"Plugh\":{\"Thud\":456,\"Xyzzy\":345}}}"
	if jstr != expects {
		t.Errorf("jst=%q, want %q", jstr, expects)
	}
}

func TestReadComplexValueAsInt64WithoutLaxNumericType(t *testing.T) {
	v := ComplexI{}
	opts := DecorderOpts{}
	err := UnmarshalWithOpts(complexToml, &v, opts)
	if nil == err {
		t.Errorf("err should not be nil")
	}
}
