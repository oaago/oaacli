package ipstring

import (
	"errors"
	"fmt"
	"net"
)

func GetIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println("ip:", ipnet.IP.String())
				return ipnet.IP.String(), nil
			}

		}
	}
	return "", errors.New("无法提取ip")
}
