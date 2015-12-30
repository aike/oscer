package osc

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"bytes"
	"encoding/binary"
	"errors"
//	"time"
)

var serverIP string
var serverPort string
var data []byte
var initdata bool = false

func Send() {
	if !initdata {
		return
	}

	udpAddr, err := net.ResolveUDPAddr("udp", serverIP + ":" + serverPort)
	checkError(err)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	defer conn.Close()

//	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.Write(data)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: error: %s", err.Error())
		os.Exit(1)
	}
}

func CheckArg(arr []string) error {
	data = []byte{}

	if len(arr) < 4 {
		return errors.New("args error")
	}

	host := arr[1]
	if !match(`^[A-Za-z0-9\-\.]+$`, host) {
		return errors.New("hostname error")
	}
	serverIP = host

	port, err := strconv.Atoi(arr[2])
	if err != nil {
		return errors.New("port number error")		
	}
	if port < 0 || port > 65535 {
		return errors.New("port number error")		
	}
	serverPort = arr[2]

	adsr := arr[3]
	if !match(`^/.+[^/]$`, adsr) {
		return errors.New("osc address error")
	}
	pushDataStr(adsr)

	idx := 0
	for i := 4; i < len(arr); i++ {
		num_i64, err := strconv.ParseInt(arr[i], 10, 32)
		num_i32 := int32(num_i64)
		if err != nil {
			num_f64, err := strconv.ParseFloat(arr[i], 32)
			num_f32 := float32(num_f64)
			if err != nil {
				return errors.New("osc args error")
			}
			pushDataF32(num_f32)
		} else {
			pushDataI32(num_i32)
		}
		idx++
	}

	initdata = true
	return nil
}

func match(reg, str string) bool {
    return regexp.MustCompile(reg).Match([]byte(str))
}

func pushDataStr(str string) {
	data = append(data, []byte(str)...)
	for datalen := len(str); datalen % 4 != 0; datalen++ {
		data = append(data, 0)
	}
}

func pushDataI32(num int32) {
	buf := bytes.NewBuffer([]byte{})
	data = append(data, 0x2c, 0x69, 0, 0)
	binary.Write(buf, binary.BigEndian, num)
	data = append(data, buf.Bytes()...)
}

func pushDataF32(num float32) {
	buf := bytes.NewBuffer([]byte{})
	data = append(data, 0x2c, 0x66, 0, 0)
	binary.Write(buf, binary.BigEndian, num)
	data = append(data, buf.Bytes()...)
}



