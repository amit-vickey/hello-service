package test

import pb "github.com/amit/hello-service/rpc"

var (
	requestList = [] pb.HelloReq{
		{
			Subject: "sample",
		},
		{
			Subject: "amit",
		},
	}
)
