package dunapi

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/xchacha20-poly1305/dun/dunbox"
)

func CreateProxyHttpClient(box *dunbox.Box) *http.Client {
	transport := &http.Transport{
		TLSHandshakeTimeout:   time.Second * 3,
		ResponseHeaderTimeout: time.Second * 3,
	}

	if box != nil {
		transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return DialContext(ctx, box, network, addr)
		}
	}

	client := &http.Client{
		Transport: transport,
	}

	return client
}
