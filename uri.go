package gox

const (
	// URISchemeSocksV4 Socks协议
	URISchemeSocksV4 URIScheme = "socks4"
	// URISchemeSocksV5 Socks协议
	URISchemeSocksV5 URIScheme = "socks4"
	// URISchemeHttp Socks协议
	URISchemeHttp URIScheme = "http"
	// URISchemeHttps Socks协议
	URISchemeHttps URIScheme = "https"
)

// URIScheme 协议
type URIScheme string
