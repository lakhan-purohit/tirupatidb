package router

import (
	"log"
	"os"
	"tirupatiBE/dal/dbModel"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type Rtr struct {
	db         dbModel.RouterFunc
	srv        *fasthttp.Server
	httpServer HttpServer
}

func CreateServer(dbFunc dbModel.RouterFunc, htp HttpServer) *Rtr {

	router := router.New()

	server := &fasthttp.Server{
		ReadTimeout:          htp.ReadTimeOut,
		WriteTimeout:         htp.WriteTimeOut,
		MaxConnsPerIP:        htp.MaxConnPerIp,
		MaxRequestsPerConn:   htp.MaxRequestPerIp,
		MaxKeepaliveDuration: htp.KeepAliveConnDuration,
		Handler:              router.Handler,
	}

	r := &Rtr{
		db:         dbFunc,
		httpServer: htp,
		srv:        server,
	}

	router.POST("/api/v2/createUser", r.createUser)
	router.GET("/api/v2/LoginUser", r.loginUser)
	router.GET("/api/v2/userList", r.userList)
	router.PUT("/api/v2/userDetailsUpdate", r.userDetailsUpdate)
	router.POST("/api/v2/addProduct", r.addProduct)
	router.POST("/api/v2/saveUnSave", r.saveUnSaveProduct)
	router.GET("/api/v2/getProductList/{role}", r.getProduct)
	router.GET("/api/v2/servFile/{fileName}", r.serveFile)
	router.POST("/api/v2/saveImage/{id}", r.saveImages)
	router.DELETE("/api/v2/deleteImage/{id}", r.deleteImages)
	router.DELETE("/api/v2/deleteProduct/{id}", r.deleteProduct)
	router.PUT("/api/v2/updateProduct", r.updateProduct)
	router.POST("/api/v2/makeOrder", r.makeOrder)
	router.GET("/api/v2/getOrder/{id}", r.getOrder)
	router.PUT("/api/v2/updateOrder/{id}/{price}/{qty}", r.upateOrder)
	router.PUT("/api/v2/updateStatus/{id}/{status}", r.upateOrderStatus)
	router.DELETE("/api/v2/deleteOrder/{id}", r.deleteOrder)
	router.DELETE("/api/v2/deleteOrder1/{id}", r.deleteOrder)

	return r
}

func (r *Rtr) Begin() {

	if err := fasthttp.ListenAndServe(r.httpServer.Listen+os.Getenv("PORT"), r.srv.Handler); err != nil {

		panic(err)
	} else {
		log.Fatal("create DB")
	}
}
