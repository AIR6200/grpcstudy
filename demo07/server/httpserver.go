package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"grpcstudy/demo07/server/services"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	gwmux := runtime.NewServeMux()
	defer cancel()
	/*	cert, _ := tls.LoadX509KeyPair("cert/client.pem","cert/client.key")
		certPool := x509.NewCertPool()
		ca, _ := ioutil.ReadFile("cert/ca.pem")
		certPool.AppendCertsFromPEM(ca)

		creds := credentials.NewTLS(&tls.Config{
			Certificates: []tls.Certificate{cert},//服务端证书
			ServerName: "localhost",
			RootCAs: certPool,
		})*/
	opt := []grpc.DialOption{grpc.WithInsecure()}
	//opt := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	//8099是对应的grpc的监听端口，必须与grpc开放的端口一致,所以启动前务必启动grpc服务
	err := services.RegisterProdServiceHandlerFromEndpoint(ctx, gwmux, "localhost:8099", opt)
	if err != nil {
		log.Fatal(err)
	}
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}
	httpServer.ListenAndServe()
}
