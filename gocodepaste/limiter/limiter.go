package limiter

import (
	"context"
	"io/ioutil"
)

func readLuaScript(lua string) string {
	bs, err := ioutil.ReadFile(lua)
	if err != nil {
		panic(err)
	}
	return string(bs)
}

type Limiter interface {
	Allow(ctx context.Context, key string) bool
}
