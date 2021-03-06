package common

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

var Ports string
var Outputfile string
var IsSaveLog bool

func Getpath() string {
	file, _ := exec.LookPath(os.Args[0])
	path1, _ := filepath.Abs(file)
	filename := filepath.Dir(path1)
	var path string
	if strings.Contains(filename, "/") {
		tmp := strings.Split(filename, `/`)
		tmp[len(tmp)-1] = ``
		path = strings.Join(tmp, `/`)
	} else if strings.Contains(filename, `\`) {
		tmp := strings.Split(filename, `\`)
		tmp[len(tmp)-1] = ``
		path = strings.Join(tmp, `\`)
	}
	return path
}

func ParsePort(ports string) (scanPorts []int) {
	if ports == "" {
		return
	}
	slices := strings.Split(ports, ",")
	for _, port := range slices {
		port = strings.Trim(port, " ")
		upper := port
		if strings.Contains(port, "-") {
			ranges := strings.Split(port, "-")
			if len(ranges) < 2 {
				continue
			}

			startPort, _ := strconv.Atoi(ranges[0])
			endPort, _ := strconv.Atoi(ranges[1])
			if startPort < endPort {
				port = ranges[0]
				upper = ranges[1]
			} else {
				port = ranges[1]
				upper = ranges[0]
			}

		}
		start, _ := strconv.Atoi(port)
		end, _ := strconv.Atoi(upper)
		for i := start; i <= end; i++ {
			scanPorts = append(scanPorts, i)
		}
	}
	scanPorts = removeDuplicate(scanPorts)
	return scanPorts
}

func removeDuplicate(old []int) []int {
	result := []int{}
	temp := map[int]struct{}{}
	for _, item := range old {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
