package mobile

import (
	"github.com/OpenIoTHub/gateway-go/v2/client"
	"github.com/OpenIoTHub/gateway-go/v2/register"
	"log"
	"net"
)

//只可以由用户开启
//func init() {
//	runGatewayGO()
//}

func runGatewayGO() {
	log.Println("Start gateway-go service")
	register.RegisterService("localhost",
		"gateway-go",
		"_http._tcp",
		"localhost",
		5244,
		[]string{},
		0,
		[]net.IP{net.ParseIP("127.0.0.1")},
		[]net.IP{},
	)
	client.Run()
}
