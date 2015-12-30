package osc

import (
	"fmt"
//	"net"
	"os"
	"regexp"
	"strconv"
	"errors"
//	"time"
)

var bytes []byte

/*
func main() {
	if len(os.Args) < 4 {
		// snowosc localhost 12000 /test/test 120 30 25
		usage()
	}

	serverIP := "localhost"
	serverPort := "12000"

	message := os.Args[1]


	udpAddr, err := net.ResolveUDPAddr("udp", serverIP+":"+serverPort)
	checkError(err)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	defer conn.Close()

//	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.Write([]byte(message))

}
*/

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: error: %s", err.Error())
		os.Exit(1)
	}
}

func CheckArg(arr []string) error {
//	var arg []string
	if len(arr) < 4 {
		return errors.New("args error")
	}

	host := arr[1]
	if !match(`^[A-Za-z0-9\-\.]+$`, host) {
		return errors.New("hostname error")
	}

	port, err := strconv.Atoi(arr[2])
	if err != nil {
		return errors.New("port number error")		
	}
	if port < 0 || port > 65535 {
		return errors.New("port number error")		
	}

	adsr := arr[3]
	if !match(`^/.+[^/]$`, adsr) {
		return errors.New("osc address error")
	}
	pushByteStr(adsr)

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
			pushByteF32(num_f32)
		} else {
			pushByteI32(num_i32)
		}
		idx++
	}
	return nil
}

func match(reg, str string) bool {
    return regexp.MustCompile(reg).Match([]byte(str))
}

func pushByteStr(str string) {

}

func pushByteI32(num int32) {

}

func pushByteF32(num float32) {

}



