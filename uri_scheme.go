package gox

const (
	// URISchemeSocksV4 Socks协议
	URISchemeSocksV4 URIScheme = `socks4`
	// URISchemeSocksV5 Socks协议
	URISchemeSocksV5 URIScheme = `socks4`
	// URISchemeHttp Http协议
	URISchemeHttp URIScheme = `http`
	// URISchemeHttps Https协议
	URISchemeHttps URIScheme = `https`
	// URISchemeRtmp Rtmp协议
	URISchemeRtmp URIScheme = `rtmp`
	// URISchemeMqtt MQTT协议
	URISchemeMqtt URIScheme = `mqtt`
	// URISchemeMqtts MQTT安全协议
	URISchemeMqtts URIScheme = `mqtts`
	// URISchemeWs Websocket协议
	URISchemeWs URIScheme = `ws`
	// URISchemeWss Websocket安全协议
	URISchemeWss URIScheme = `wss`
)

// URIScheme 协议
type URIScheme string
