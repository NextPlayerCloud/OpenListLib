package mobile

import (
	"fmt"
	"github.com/OpenIoTHub/gateway-go/v2/client"
	"github.com/OpenIoTHub/gateway-go/v2/register"
	uuid "github.com/satori/go.uuid"
	"log"
	"net"
)

//只可以由用户开启
//func init() {
//	runGatewayGO()
//}

func runGatewayGO() {
	log.Println("Start gateway-go service")
	register.RegisterService("localhost-gateway-go",
		"_http._tcp",
		"local",
		"localhost",
		34323,
		[]string{"name=gateway-go", fmt.Sprintf("id=gateway-go@%s", uuid.Must(uuid.NewV4()).String()), "home-page=https://github.com/OpenIoTHub/gateway-go"},
		0,
		[]net.IP{net.ParseIP("127.0.0.1")},
		[]net.IP{},
	)
	register.RegisterService("localhost-alist",
		"_http._tcp",
		"local",
		"localhost",
		5244,
		[]string{"name=OpenlistApp", fmt.Sprintf("id=alist@%s", uuid.Must(uuid.NewV4()).String()), "home-page=https://github.com/OpenlistApp/OpenlistApp"},
		0,
		[]net.IP{net.ParseIP("127.0.0.1")},
		[]net.IP{},
	)
	register.RegisterService("localhost-ddns-go",
		"_http._tcp",
		"local",
		"localhost",
		9876,
		[]string{"name=ddns-go", fmt.Sprintf("id=ddns-go@%s", uuid.Must(uuid.NewV4()).String()), "home-page=https://github.com/OpenIoTHub/ddns-go"},
		0,
		[]net.IP{net.ParseIP("127.0.0.1")},
		[]net.IP{},
	)
	client.Run()
}
