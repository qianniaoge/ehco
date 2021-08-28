package constant

import "time"

var (
	DefaultDeadline = 30 * time.Second
)

const (
	Version = "1.0.7"

	MaxMWSSStreamCnt = 100
	DialTimeOut      = 3 * time.Second

	Listen_RAW  = "raw"
	Listen_WS   = "ws"
	Listen_WSS  = "wss"
	Listen_MWSS = "mwss"

	Transport_RAW  = "raw"
	Transport_WS   = "ws"
	Transport_WSS  = "wss"
	Transport_MWSS = "mwss"

	BUFFER_POOL_SIZE = 128      // 128 * 4kb
	BUFFER_SIZE      = 4 * 1024 // 4kb

	IndexHTMLTMPL = `<!doctype html>
<html>
<head>
	<meta charset="UTF-8">
</head>
<body>
	<h2>Ehco(Version ` + Version + `)</h2>
	<h3>ehco is a network relay tool and a typo :)</h3>
	<p><a href="https://github.com/Ehco1996/ehco">More information here</a></p>

	<p><a href="/metrics/">Metrics</a></p>
	<p><a href="/debug/pprof/">Debug</a></p>
</body>
</html>
`
)
