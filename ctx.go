package rlog

import (
	"bytes"
	"fmt"
	"strconv"
	"sync"
)

type Ctx struct {
	m map[string]interface{}
	k []string
	x sync.Mutex
}

func newCtx() *Ctx {
	return &Ctx{m: make(map[string]interface{}), k: []string{}}
}

func (c *Ctx) Add(ctx []interface{}) {
	if len(ctx) == 0 {
		return
	}
	for i := 0; i < len(ctx); i += 2 {
		k, ok := ctx[i].(string)
		v := ctx[i+1]
		if !ok {
			panic(fmt.Sprintf("unable to parse \"%#+v\" (type '%[1]T') as ctx key, want type 'string'", ctx[i]))
		}
		c.Set(k, v)
	}
}

func (c *Ctx) Set(k string, v interface{}) {
	c.x.Lock()
	defer c.x.Unlock()
	c.m[k] = v
	c.k = append(c.k, k)
}
func (c *Ctx) Delete(k string) {
	c.x.Lock()
	defer c.x.Unlock()
	delete(c.m, k)
	for i, e := range c.k {
		if e == k {
			c.k = append(c.k[:i], c.k[i+1:]...)
			return
		}
	}
}

func (c *Ctx) String() string {
	if len(c.m) == 0 {
		return ""
	}
	var buf bytes.Buffer
	for _, k := range c.k {
		buf.WriteString(fmt.Sprintf(" %s=%q", k, formatCtxValue(c.m[k])))
	}
	buf.WriteString(" |")
	return buf.String()
}

func formatCtxValue(value interface{}) string {
	if value == nil {
		return "nil"
	}
	switch v := value.(type) {
	case func() string:
		return v()
	case bool:
		return strconv.FormatBool(v)
	default:
		return fmt.Sprintf("%+v", v)
	}
}
