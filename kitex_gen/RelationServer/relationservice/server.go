// Code generated by Kitex v0.4.4. DO NOT EDIT.
<<<<<<<< HEAD:kitex_gen/VideoServer/videosrv/server.go
package videosrv

import (
	server "github.com/cloudwego/kitex/server"
	VideoServer "github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler VideoServer.VideoSrv, opts ...server.Option) server.Server {
========
package relationservice

import (
	server "github.com/cloudwego/kitex/server"
	RelationServer "github.com/xiaohei366/TinyTiktok/kitex_gen/RelationServer"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler RelationServer.RelationService, opts ...server.Option) server.Server {
>>>>>>>> origin/user&follow:kitex_gen/RelationServer/relationservice/server.go
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}