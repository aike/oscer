// oscer_test.go by aike
// licenced under MIT License. 

package oscer_test

import (
	"testing"
	"bytes"
	"oscer/osc"
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

func Test_Ok3(t *testing.T) {
	arg := []string{"oscer", "::1", "12000", "/test/test"}
	err := osc.CheckArg(arg)
	if err != nil {
		t.Error("legal arguments NG")
	}
}

func Test_Ok4(t *testing.T) {
	arg := []string{"oscer", `fe80::aa66:7fff:fe22:12f8%en0`, "12000", "/test/test"}
	err := osc.CheckArg(arg)
	if err != nil {
		t.Error("legal arguments NG")
	}
}

func Test_Ok_Server1(t *testing.T) {
	arg := []string{"oscer", "receive", "12000"}
	ok := osc.IsServer(arg)
	if !ok {
		t.Error("server check NG")
	}
}

func Test_Ok_Server2(t *testing.T) {
	arg := []string{"oscer", "192.168.0.1", "12000", "/test/test"}
	ok := osc.IsServer(arg)
	if ok {
		t.Error("server check NG")
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
	arg := []string{"oscer", "localhost", "12000", "/test/test", "1.1.1"}
	err := osc.CheckArg(arg)
	if err == nil {
		t.Error("illegal arguments test NG")
	}
}

///////////////////////////////////////

func Test_MessageArgNone1(t *testing.T) {
	arg := []string{"oscer", "localhost", "12000", "/test"}
	expected := []byte {0x2f, 0x74, 0x65, 0x73, 0x74, 0x00, 0x00, 0x00,
						0x2c, 0x00, 0x00, 0x00 }
	_ = osc.CheckArg(arg)
	if bytes.Compare(expected, osc.GetData()) != 0 {
		t.Error("message test none NG")
	}
}

func Test_MessageArgNone2(t *testing.T) {
	arg := []string{"oscer", "localhost", "12000", "/testabc"}
	expected := []byte {0x2f, 0x74, 0x65, 0x73, 0x74, 0x61, 0x62, 0x63,
						0x00, 0x00, 0x00, 0x00, 0x2c, 0x00, 0x00, 0x00 }
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

func Test_Message3Args(t *testing.T) {
	arg := []string{"oscer", "localhost", "12000", "/test", "100", "100", "100"}
	expected := []byte {0x2f, 0x74, 0x65, 0x73, 0x74, 0x00, 0x00, 0x00,
						0x2c, 0x69, 0x69, 0x69, 0x00, 0x00, 0x00, 0x00,
						0x00, 0x00, 0x00, 0x64, 0x00, 0x00, 0x00, 0x64,
						0x00, 0x00, 0x00, 0x64 }
	_ = osc.CheckArg(arg)	
	if bytes.Compare(expected, osc.GetData()) != 0 {
		t.Error("message test 3 args NG")
	}
}

func Test_MessageString(t *testing.T) {
	arg := []string{"oscer", "localhost", "12000", "/hello", "world"}
	expected := []byte {0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x00, 0x00,
						0x2c, 0x73, 0x00, 0x00, 0x77, 0x6f, 0x72, 0x6c,
						0x64, 0x00, 0x00, 0x00}
	_ = osc.CheckArg(arg)	
	if bytes.Compare(expected, osc.GetData()) != 0 {
		t.Error("message test string NG")
	}
}

func Test_MessageIntStrFloatStr(t *testing.T) {
	arg := []string{"oscer", "localhost", "12000", "/hello", "10", "str1", "1.5", "str2"}
	expected := []byte {0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x00, 0x00,
						0x2c, 0x69, 0x73, 0x66, 0x73, 0x00, 0x00, 0x00,
						0x00, 0x00, 0x00, 0x0a, 0x73, 0x74, 0x72, 0x31,
						0x00, 0x00, 0x00, 0x00, 0x3f, 0xc0, 0x00, 0x00,
						0x73, 0x74, 0x72, 0x32, 0x00, 0x00, 0x00, 0x00 }
	_ = osc.CheckArg(arg)

//	b := osc.GetData()
//	for i := 0; i < len(b); i++ {
//		fmt.Printf("%02x ", b[i])
//	}
//	fmt.Println("")

	if bytes.Compare(expected, osc.GetData()) != 0 {
		t.Error("message test int str float str NG")
	}
}



