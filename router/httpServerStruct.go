package router

import "time"

type HttpServer struct {
	ReadTimeOut           time.Duration `json:"ReadTimeOut"`
	WriteTimeOut          time.Duration `json"WriteTimeOut"`
	Listen                string        `json:"Listen"`
	KeepAliveConnDuration time.Duration `json:"MaxKeepAliveDuration"`
	MaxConnPerIp          int           `json:"MaxConnPerIp"`
	MaxRequestPerIp       int           `json:"MaxRequestPerConn"`
}
