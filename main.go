package main

import (
	log "ShareKnows/logs"
	_ "ShareKnows/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	var port int = 8080
	if httpport, err := beego.AppConfig.Int("http.port"); nil == err {
		port = httpport
	}
	log.Infof("Prometheus_oper listen on port:%d", port)
	beego.Run(fmt.Sprintf(":%d", port))
}
