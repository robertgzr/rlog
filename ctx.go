package rlog

import (
	"bytes"
	"fmt"
	"strconv"
)

const (
	defaultCtxOpener    = "("
	defaultCtxCloser    = ")"
	defaultCtxDelimiter = " "
)

type Ctx struct {
	values      []interface{}
	op, cl, del string
}

func newCtx() Ctx {
	return Ctx{
		values: make([]interface{}, 0),
		op:     defaultCtxOpener,
		cl:     defaultCtxCloser,
		del:    defaultCtxDelimiter,
	}
}

// Add takes a list in the form of ( (string, interface{}), ... )
func (c *Ctx) Add(ctx []interface{}) {
	if len(ctx) == 0 {
		return
	}
	if len(ctx)%2 != 0 {
		panic("cannot append odd number of context elements")
	}
	c.values = append(c.values, ctx...)

}

func (c *Ctx) String() string {
	if len(c.values) == 0 {
		return ""
	}
	var buf bytes.Buffer
	buf.WriteString(c.op)

	for i := 0; i < len(c.values); i += 2 {
		k, ok := c.values[i].(string)
		v := c.values[i+1]
		if !ok {
			panic(fmt.Sprintf("unable to parse \"%#+v\" (type '%[1]T') as ctx key, want type 'string'", c.values[i]))
		}

		buf.WriteString(fmt.Sprintf("%s=%q", k, formatCtxValue(v)))
		if i != len(c.values)-2 {
			buf.WriteString(c.del)
		}
	}

	buf.WriteString(c.cl)
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
