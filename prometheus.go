package common

import (
	"net"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func PrometheusBoot(port int) {
	http.Handle("/metrics", promhttp.Handler())
	addr := net.JoinHostPort("0.0.0.0", strconv.Itoa(port))
	go func() {
		err := http.ListenAndServe(
			addr,
			nil)

		if err != nil {
			logger.Fatal("启动失败")
		}

		logger.Infof("监控启动,端口为 %s", addr)
	}()
}
