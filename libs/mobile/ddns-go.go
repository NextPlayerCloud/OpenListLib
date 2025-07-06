package mobile

import (
	"embed"
	"errors"
	"github.com/jeessy2/ddns-go/v6/config"
	"github.com/jeessy2/ddns-go/v6/dns"
	"github.com/jeessy2/ddns-go/v6/util"
	"github.com/jeessy2/ddns-go/v6/web"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

var listen = ":9876"

//go:embed static
var staticEmbeddedFiles embed.FS

//go:embed favicon.ico
var faviconEmbeddedFile embed.FS

func init() {
	go run()
}

func run() {
	// 兼容之前的配置文件
	conf, _ := config.GetConfigCached()
	conf.CompatibleConfig()
	// 初始化语言
	util.InitLogLang(conf.Lang)

	go func() {
		// 启动web服务
		err := runWebServer()
		if err != nil {
			log.Println(err)
			time.Sleep(time.Minute)
			os.Exit(1)
		}
	}()

	// 初始化备用DNS
	util.InitBackupDNS("", conf.Lang)

	// 等待网络连接
	util.WaitInternet(dns.Addresses)

	// 定时运行
	dns.RunTimer(time.Duration(300) * time.Second)
}

func staticFsFunc(writer http.ResponseWriter, request *http.Request) {
	http.FileServer(http.FS(staticEmbeddedFiles)).ServeHTTP(writer, request)
}

func faviconFsFunc(writer http.ResponseWriter, request *http.Request) {
	http.FileServer(http.FS(faviconEmbeddedFile)).ServeHTTP(writer, request)
}

func runWebServer() error {
	// 启动静态文件服务
	http.HandleFunc("/static/", web.AuthAssert(staticFsFunc))
	http.HandleFunc("/favicon.ico", web.AuthAssert(faviconFsFunc))

	http.HandleFunc("/login", web.AuthAssert(web.Login))
	http.HandleFunc("/loginFunc", web.AuthAssert(web.LoginFunc))

	http.HandleFunc("/", web.Auth(web.Writing))
	http.HandleFunc("/save", web.Auth(web.Save))
	http.HandleFunc("/logs", web.Auth(web.Logs))
	http.HandleFunc("/clearLog", web.Auth(web.ClearLog))
	http.HandleFunc("/webhookTest", web.Auth(web.WebhookTest))
	http.HandleFunc("/logout", web.Auth(web.Logout))

	util.Log("监听 %s", listen)

	l, err := net.Listen("tcp", listen)
	if err != nil {
		return errors.New(util.LogStr("监听端口发生异常, 请检查端口是否被占用! %s", err))
	}

	return http.Serve(l, nil)
}
