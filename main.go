package main

import (
	"ice/common"
	"strconv"
	"sync"
)

func main() {
	common.LoadConfig()
	scanPorts := common.ParsePort(common.Ports)

	sw := sync.WaitGroup{}
	for _, port := range scanPorts {
		common.StartTcpServer(strconv.Itoa(port), &sw)
		sw.Add(1)
	}
	sw.Wait()
}
