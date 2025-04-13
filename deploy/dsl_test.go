package deploy_test

import (
	"encoding/json"
	"testing"

	. "github.com/Meduzz/dsl/deploy"
)

func TestDsl(t *testing.T) {
	var d = DeployConfig("wtf")
	d.WithOptions(WithDB(WithDialect("pg", "wtf", true), Env("DB_CONN")), WithRedis(Env("REDIS_URL")), WithNats(Env("NATS_URL")), WithTcpPort(8080, "http"), WithVolume("cache", "/cache"), WithCommand("./server"), WithArgument("start", ""))

	bs, _ := json.Marshal(d)

	println(string(bs))
}
