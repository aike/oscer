// oscer_test.go by aike
// licenced under MIT License. 

package osc_test

import (
	"testing"
	"bytes"
	"../src/osc"
)

func Test_Ok1(t *testing.T) {
	arg := []string{"oscer", "localhost", "12000", "/test/test"}
	err := osc.CheckArg(arg)
	if err != nil {
		t.Error("legal arguments NG")
	}
}

func Test_Ok2(t *testing.T) {
	arg := []string{"oscer", "192.168.0.1", "12000", "/test/test"}
	err := osc.CheckArg(arg)
	if err != nil {
		t.Error("legal arguments NG")
	}
}

///////////////////////////////////////


func Test_NgHost(t *testing.T) {
	arg := []string{"oscer", "", "12000", "/test/test"}
	err := osc.CheckArg(arg)
	if err == nil {
		t.Error("illegal host test NG")
	}
}

func Test_NgPort1(t *testing.T) {
	arg := []string{"oscer", "localhost", "-1", "test/test"}
	err := osc.CheckArg(arg)
	if err == nil {
		t.Error("illegal port number test NG")
	}
}

func Test_NgPort2(t *testing.T) {
	arg := []string{"oscer", "localhost", "65536", "test/test"}
	err := osc.CheckArg(arg)
	if err == nil {
		t.Error("illegal port number test NG")
	}
}

func Test_NgAdsr1(t *testing.T) {
	arg := []string{"oscer", "localhost", "12000", "test/test"}
	err := osc.CheckArg(arg)
	if err == nil {
		t.Error("illegal address test NG")
	}
}

func Test_NgAdsr2(t *testing.T) {
	arg := []string{"oscer", "localhost", "12000", "/test/test/"}
	err := osc.CheckArg(arg)
	if err == nil {
		t.Error("illegal address test NG")
	}
}

func Test_OkArg1(t *testing.T) {
	arg := []string{"oscer", "localhost", "12000", "/test/test", "100"}
	err := osc.CheckArg(arg)
	if err != nil {
		t.Error("legal arguments NG")
	}
}

func Test_OkArg2(t *testing.T) {
	arg := []string{"oscer", "localhost", "12000", "/test/test", "-100"}
	err := osc.CheckArg(arg)
	if err != nil {
		t.Error("legal arguments NG")
	}
}

func Test_OkArg3(t *testing.T) {
	arg := []string{"oscer", "localhost", "12000", "/test/test", "-100.5"}
	err := osc.CheckArg(arg)
	if err != nil {
		t.Error("legal arguments NG")
	}
}

func Test_OkArg4(t *testing.T) {
	arg := []string{"oscer", "localhost", "12000", "/test/test", "-100.5"}
	err := osc.CheckArg(arg)
	if err != nil {
		t.Error("legal arguments NG")
	}
}

func Test_NgArg1(t *testing.T) {
	arg := []string{"oscer", "localhost", "12000", "/test/test", "100abc"}
	err := osc.CheckArg(arg)
	if err == nil {
		t.Error("illegal arguments test NG")
	}
}

func Test_NgArg2(t *testing.T) {
	arg := []string{"oscer", "localhost", "12000", "/test/test", "1.1.1"}
	err := osc.CheckArg(arg)
	if err == nil {
		t.Error("illegal arguments test NG")
	}
}

///////////////////////////////////////

func Test_MessageArgNone(t *testing.T) {
	arg := []string{"oscer", "localhost", "12000", "/test"}
	expected := []byte {0x2f, 0x74, 0x65, 0x73, 0x74, 0x00, 0x00, 0x00,
						0x2c, 0x00, 0x00, 0x00 }
	_ = osc.CheckArg(arg)
	if bytes.Compare(expected, osc.GetData()) != 0 {
		t.Error("message test none NG")
	}
}

func Test_MessageArgInt(t *testing.T) {
	arg := []string{"oscer", "localhost", "12000", "/test", "100"}
	expected := []byte {0x2f, 0x74, 0x65, 0x73, 0x74, 0x00, 0x00, 0x00,
						0x2c, 0x69, 0x00, 0x00, 0x00, 0x00, 0x00, 0x64 }
	_ = osc.CheckArg(arg)	
	if bytes.Compare(expected, osc.GetData()) != 0 {
		t.Error("message test int NG")
	}
}

func Test_MessageArgFloat(t *testing.T) {
	arg := []string{"oscer", "localhost", "12000", "/test", "10.5"}
	expected := []byte {0x2f, 0x74, 0x65, 0x73, 0x74, 0x00, 0x00, 0x00,
						0x2c, 0x66, 0x00, 0x00, 0x41, 0x28, 0x00, 0x00 }
	_ = osc.CheckArg(arg)	
	if bytes.Compare(expected, osc.GetData()) != 0 {
		t.Error("message test float NG")
	}
}

func Test_MessageArgIntFloat(t *testing.T) {
	arg := []string{"oscer", "localhost", "12000", "/test", "100", "10.5"}
	expected := []byte {0x2f, 0x74, 0x65, 0x73, 0x74, 0x00, 0x00, 0x00,
						0x2c, 0x69, 0x66, 0x00, 0x00, 0x00, 0x00, 0x64,
						0x41, 0x28, 0x00, 0x00 }
	_ = osc.CheckArg(arg)	
	if bytes.Compare(expected, osc.GetData()) != 0 {
		t.Error("message test int float NG")
	}
}

