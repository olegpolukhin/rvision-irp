// Package ping is library seeking to emulate the unix "ping"
package ping

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	timeSliceLength = 2000
	trackerLength   = 1000
	seqNumberLength = 1
)

// NewPinger returns a new Pinger struct pointer
func NewPinger(server string) error {
	var seqNumber uint64 = seqNumberLength
	var timeout = time.Duration(timeSliceLength) * time.Millisecond
	var period = time.Duration(trackerLength) * time.Millisecond
	var network string
	var port int

	server = strings.TrimLeft(server, "http://")
	server = strings.TrimLeft(server, "https://")
	server = strings.TrimSpace(server)

	sArray := strings.Split(server, ":")
	if len(sArray) > 1 && len(sArray) <= 2 {
		server = sArray[0]
		portParse, err := strconv.Atoi(sArray[1])
		if err != nil {
			return fmt.Errorf("pinger error: not valid port")
		}

		port = portParse
	} else {
		server = sArray[0]
		port = 0
	}

	if len(server) == 0 && port == 0 {
		return fmt.Errorf("pinger error: not valid server or port")
	}

	if port == 0 {
		network = fmt.Sprintf("%s", server)
	} else {
		network = fmt.Sprintf("%s:%d", server, port)
	}

	ticker := time.NewTicker(period)
	quit := make(chan interface{})

	for ; ; seqNumber++ {
		select {
		case <-ticker.C:
			if err := tryPort(network, timeout); err != nil {
				return fmt.Errorf("pinger error: tryPort %v", err)
			}

			ticker.Stop()
			return nil
		case <-quit:
			ticker.Stop()
			return nil
		}
	}
}

func tryPort(network string, timeout time.Duration) error {
	//startTime := time.Now()
	conn, err := net.DialTimeout("tcp", network, timeout)
	//endTime := time.Now()
	if err != nil {
		return fmt.Errorf("connection failed")
	} else {
		defer conn.Close()
		//var t = float64(endTime.Sub(startTime)) / float64(time.Millisecond)
		//os.Stdout.Write([]byte(startTime.Format("[2006-01-02T15:04:05]:") + fmt.Sprintf(" addr=%s seq=%d time=%4.2fms\n", conn.RemoteAddr().String(), seq, t)))
	}

	return nil
}
