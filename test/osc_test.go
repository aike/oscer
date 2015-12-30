package osc_test

import (
	"testing"
	"../src/osc"
)

func Test_Ok1(t *testing.T) {
	arg := []string{"snowosc", "localhost", "12000", "/test/test"}
	err := osc.CheckArg(arg)
	if err != nil {
		t.Error("正常引数テスト　NG")
	}
}

func Test_Ok2(t *testing.T) {
	arg := []string{"snowosc", "192.168.0.1", "12000", "/test/test"}
	err := osc.CheckArg(arg)
	if err != nil {
		t.Error("正常引数テスト　NG")
	}
}

func Test_NgHost(t *testing.T) {
	arg := []string{"snowosc", "", "12000", "/test/test"}
	err := osc.CheckArg(arg)
	if err == nil {
		t.Error("異常ホスト名テスト　NG")
	}
}

func Test_NgPort1(t *testing.T) {
	arg := []string{"snowosc", "localhost", "-1", "test/test"}
	err := osc.CheckArg(arg)
	if err == nil {
		t.Error("異常ポート番号テスト　NG")
	}
}

func Test_NgPort2(t *testing.T) {
	arg := []string{"snowosc", "localhost", "65536", "test/test"}
	err := osc.CheckArg(arg)
	if err == nil {
		t.Error("異常ポート番号テスト　NG")
	}
}

func Test_NgAdsr1(t *testing.T) {
	arg := []string{"snowosc", "localhost", "12000", "test/test"}
	err := osc.CheckArg(arg)
	if err == nil {
		t.Error("異常アドレステスト　NG")
	}
}

func Test_NgAdsr2(t *testing.T) {
	arg := []string{"snowosc", "localhost", "12000", "/test/test/"}
	err := osc.CheckArg(arg)
	if err == nil {
		t.Error("異常アドレステスト　NG")
	}
}

func Test_OkArg1(t *testing.T) {
	arg := []string{"snowosc", "localhost", "12000", "/test/test", "100"}
	err := osc.CheckArg(arg)
	if err != nil {
		t.Error("正常引数テスト　NG")
	}
}

func Test_OkArg2(t *testing.T) {
	arg := []string{"snowosc", "localhost", "12000", "/test/test", "-100"}
	err := osc.CheckArg(arg)
	if err != nil {
		t.Error("正常引数テスト　NG")
	}
}

func Test_OkArg3(t *testing.T) {
	arg := []string{"snowosc", "localhost", "12000", "/test/test", "-100.5"}
	err := osc.CheckArg(arg)
	if err != nil {
		t.Error("正常引数テスト　NG")
	}
}

func Test_OkArg4(t *testing.T) {
	arg := []string{"snowosc", "localhost", "12000", "/test/test", "-100.5"}
	err := osc.CheckArg(arg)
	if err != nil {
		t.Error("正常引数テスト　NG")
	}
}

func Test_NgArg1(t *testing.T) {
	arg := []string{"snowosc", "localhost", "12000", "/test/test", "100abc"}
	err := osc.CheckArg(arg)
	if err == nil {
		t.Error("異常引数テスト　NG")
	}
}

func Test_NgArg2(t *testing.T) {
	arg := []string{"snowosc", "localhost", "12000", "/test/test", "1.1.1"}
	err := osc.CheckArg(arg)
	if err == nil {
		t.Error("異常引数テスト　NG")
	}
}
