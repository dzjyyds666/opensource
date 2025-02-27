package common

import "testing"

type TestStruct struct {
	A int
	B *string
	C float64
	D *test1
}

type test1 struct {
	E int
	F *string
	G float64
}

func TestInterfaceToString(t *testing.T) {
	str := "test"
	testStruct := &TestStruct{
		A: 1,
		B: &str,
		C: 1.1,
	}

	test := &test1{
		E: 1,
		F: &str,
		G: 1.1,
	}

	testStruct.D = test

	withoutError := ToStringWithoutError(testStruct)

	t.Log(withoutError)
}
