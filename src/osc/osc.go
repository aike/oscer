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
	pushAdrsString(adsr)

	idx := 0
	for i := 4; i < len(arr); i++ {
		if match(`^[+-]?[0-9]+$`, arr[i]) {
			// Int32
			num_i64, err := strconv.ParseInt(arr[i], 10, 32)
			num_i32 := int32(num_i64)
			if err != nil {
				return errors.New("osc args error")
			}
			pushDataI32(num_i32)

		} else if match(`^[+-]?[0-9.]+$`, arr[i]) {
			// Float32
			num_f64, err := strconv.ParseFloat(arr[i], 32)
			num_f32 := float32(num_f64)
			if err != nil {
				return errors.New("osc args error")
			}
			pushDataF32(num_f32)

		} else {
			// String
			pushDataString(arr[i])
		}
		idx++
	}

	senddata = append(senddata, 0)
	fill4byte()
	senddata = append(senddata, oscarg...)

	initdata = true
	return nil
}


func IsServer(arr []string) bool {
	if len(arr) != 3 {
		return false
	} else if arr[1] != "receive" {
		return false
	}
	return true
}

func CreateServer(portstr string) error {
	port, err := strconv.Atoi(portstr)
	if err != nil {
		return errors.New("port number error")		
	}
	if port < 0 || port > 65535 {
		return errors.New("port number error")		
	}

    addr, err := net.ResolveUDPAddr("udp", "localhost:" + portstr)
	if err != nil {
		return errors.New("server resolve error")		
	}

    conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return errors.New("server listen error")		
	}

    defer conn.Close()
    buf := make([]byte, 4096)
    for {
        len, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			return errors.New("server data read error")		
		}

        parse(buf, len)
    }

    return nil;
}


func match(reg, str string) bool {
    return regexp.MustCompile(reg).Match([]byte(str))
}

func pushAdrsString(str string) {
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

func pushDataString(str string) {
	senddata = append(senddata, 's')
	buf := bytes.NewBuffer([]byte(str))
	oscarg = append(oscarg, buf.Bytes()...)

	oscarg = append(oscarg, 0)
	for datalen := len(oscarg); datalen % 4 != 0; datalen++ {
		oscarg = append(oscarg, 0)
	}
}

func GetData() []byte {
	return senddata
}

func parse(arr []byte, datalen int) {
	//	for i := 0; i < datalen; i++ {
	//		fmt.Printf("%02x ", arr[i])
	//		if i % 16 == 15 {
	//			fmt.Printf("\n")
	//		}
	//	}
	//	fmt.Printf("\n==================\n")

	var pos int
	var i32 int32
	var f32 float32
	var str string

	adrs, pos := getString(arr, 0)
	fmt.Printf("%s", adrs)

	flags, pos := getString(arr, pos)
	flags = flags[1:]

	for i := 0; i < len(flags); i++ {
		switch flags[i] {
			case 'i':
				i32, pos = getInt32(arr, pos)
				fmt.Printf(" %v", i32)
			case 'f':
				f32, pos = getFloat32(arr, pos)
				fmt.Printf(" %v", f32)
			case 's':
				str, pos = getString(arr, pos)
				fmt.Printf(` "%s"`, str)
		}
	}
	fmt.Printf("\n")
}

func getString(arr []byte, start int) (string, int) {
	pos := start
	for ; arr[pos] != 0 && pos < len(arr); pos++ {}

	rest := 4 - (pos % 4)
	pos += rest

	return string(arr[start:pos]), pos
}

func getInt32(arr []byte, start int) (int32, int) {
	var n int32
	buf := bytes.NewBuffer(arr[start:start + 4])
	binary.Read(buf, binary.BigEndian, &n)
	return n, start + 4
}

func getFloat32(arr []byte, start int) (float32, int) {
	var f float32
	buf := bytes.NewBuffer(arr[start:start + 4])
	binary.Read(buf, binary.BigEndian, &f)
	return f, start + 4	
}

