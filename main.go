package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/yard-turkey/cosi-prototype-interface/cosi"
	"google.golang.org/grpc"
)

func init(){
	flag.Parse()
}

func main() {
	glog.Info("starting s3 plugin services")
	rpcServer := grpc.NewServer()
	glog.Info("grpc server initialized")
	glog.Info("configuring provisioner server")
	cosi.RegisterProvisionerServer(rpcServer, newHandler())
	listener, err := configureHTTPListener()
	if err != nil {
		panic(err)
	}
	glog.Info("starting grpc server, waiting for requests")
	if err = rpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
