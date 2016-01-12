// osc.go by aike
// licenced under MIT License. 

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
)

var serverIP string
var serverPort string
var senddata []byte
var oscarg []byte
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

	conn.Write(senddata)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: error: %s", err.Error())
		os.Exit(1)
	}
}

func CheckArg(arr []string) error {
	senddata = []byte{}
	oscarg = []byte{}

	if len(arr) < 4 {
		return errors.New("args error")
	}

	host := arr[1]
	if match(`^[A-Za-z0-9\-\.]+$`, host) {
		// IPv4
		serverIP = host
	} else if match(`^[:%A-Za-z0-9]+$`, host) {
		// IPv6
		serverIP = "[" + host + "]"
	} else {
		return errors.New("hostname error")
	}

	port, err := strconv.Atoi(arr[2])
	if err != nil {
		return errors.New("port number error")		
	}
	if port < 0 || port > 65535 {
		return errors.New("port number error")		
	}
	serverPort = arr[2]

	adsr := arr[3]
	if !match(`^/.*[^/]$`, adsr) {
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

	senddata = append(senddata, 0)
	fill4byte()
	senddata = append(senddata, oscarg...)

	initdata = true
	return nil
}

func match(reg, str string) bool {
    return regexp.MustCompile(reg).Match([]byte(str))
}

func pushDataStr(str string) {
	senddata = append(senddata, []byte(str)...)
	senddata = append(senddata, 0)
	fill4byte()
	senddata = append(senddata, 0x2c)
}

func fill4byte() {
	for datalen := len(senddata); datalen % 4 != 0; datalen++ {
		senddata = append(senddata, 0)
	}
}

func pushDataI32(num int32) {
	senddata = append(senddata, 'i')
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, num)
	oscarg = append(oscarg, buf.Bytes()...)
}

func pushDataF32(num float32) {
	senddata = append(senddata, 'f')
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, num)
	oscarg = append(oscarg, buf.Bytes()...)
}

func GetData() []byte {
	return senddata
}

