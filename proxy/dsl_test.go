package proxy_test

import (
	"encoding/json"
	"testing"

	"github.com/Meduzz/dsl/proxy"
)

func TestDSL(t *testing.T) {
	p := proxy.ProxyConfig("test.com", "/test")

	p.WithMiddlewares(proxy.StripPrefix("strip-test", "/test"), proxy.ForwardAuth("forward-auth", "auth-service"), proxy.Chain("strip-test", "forward-auth"))

	bs, _ := json.Marshal(p)

	println(string(bs))
}
