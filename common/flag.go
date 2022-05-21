package common

import "flag"

func LoadConfig() {
	flag.StringVar(&Ports, "p", "21,80-90,22,23,3389,6379,3306,1433,7001,1099", "tcp port")
	flag.StringVar(&Outputfile, "o", "iceLog.txt", "logfile")
	flag.BoolVar(&IsSaveLog, "nlog", false, "是否存储日志")
	flag.Parse()
}
