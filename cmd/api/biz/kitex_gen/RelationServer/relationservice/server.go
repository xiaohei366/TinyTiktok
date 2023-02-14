// Code generated by Kitex v0.4.4. DO NOT EDIT.
package relationservice

import (
	server "github.com/cloudwego/kitex/server"
	RelationServer "github.com/xiaohei366/TinyTiktok/cmd/api/biz/kitex_gen/RelationServer"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler RelationServer.RelationService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
