package http

import (
	"net/http"
)

// Transport 传输，纯粹是为了解决名冲突而创建
// 在某些情况下，有可能会同时引入net.http包和本包，而Go语言对包冲突的解决办法实在是太操蛋
type Transport = http.Transport
