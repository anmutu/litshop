package context

import (
	"net/http"
	"sync"
	"time"
)

type Context struct {
	mu        sync.RWMutex
	RequestId string

	Request *http.Request
	Writer  http.ResponseWriter

	App        string
	Opts       map[string]interface{}
	HttpHeader http.Header

	Session Session
}

func (*Context) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*Context) Done() <-chan struct{} {
	return nil
}

func (*Context) Err() error {
	return nil
}

func (*Context) Value(key interface{}) interface{} {
	return nil
}

type setter map[string][]string

func (c *Context) SetHeader(header http.Header) {
	c.HttpHeader = header
}

func (c *Context) GetHeaders() http.Header {
	return c.HttpHeader
}

func (c *Context) GetHeader(key string) string {
	return c.HttpHeader.Get(key)
}

func (c *Context) Set(key string, value interface{}) {
	c.mu.Lock()
	if c.Opts == nil {
		c.Opts = make(map[string]interface{})
	}

	c.Opts[key] = value
	c.mu.Unlock()
}

func (c *Context) Get(key string) (value interface{}, exists bool) {
	c.mu.RLock()
	value, exists = c.Opts[key]
	c.mu.RUnlock()
	return
}

func (c *Context) GetString(key string) (s string) {
	if val, ok := c.Get(key); ok && val != nil {
		s, _ = val.(string)
	}
	return
}

func (c *Context) GetBool(key string) (b bool) {
	if val, ok := c.Get(key); ok && val != nil {
		b, _ = val.(bool)
	}
	return
}

func (c *Context) GetInt(key string) (i int) {
	if val, ok := c.Get(key); ok && val != nil {
		i, _ = val.(int)
	}
	return
}

func (c *Context) GetInt64(key string) (i64 int64) {
	if val, ok := c.Get(key); ok && val != nil {
		i64, _ = val.(int64)
	}
	return
}

func (c *Context) GetFloat64(key string) (f64 float64) {
	if val, ok := c.Get(key); ok && val != nil {
		f64, _ = val.(float64)
	}
	return
}

func (c *Context) GetTime(key string) (t time.Time) {
	if val, ok := c.Get(key); ok && val != nil {
		t, _ = val.(time.Time)
	}
	return
}

func (c *Context) GetDuration(key string) (d time.Duration) {
	if val, ok := c.Get(key); ok && val != nil {
		d, _ = val.(time.Duration)
	}
	return
}

func (c *Context) GetStringSlice(key string) (ss []string) {
	if val, ok := c.Get(key); ok && val != nil {
		ss, _ = val.([]string)
	}
	return
}

func (c *Context) GetStringMap(key string) (sm map[string]interface{}) {
	if val, ok := c.Get(key); ok && val != nil {
		sm, _ = val.(map[string]interface{})
	}
	return
}

func (c *Context) GetStringMapString(key string) (sms map[string]string) {
	if val, ok := c.Get(key); ok && val != nil {
		sms, _ = val.(map[string]string)
	}
	return
}

func (c *Context) GetStringMapStringSlice(key string) (smss map[string][]string) {
	if val, ok := c.Get(key); ok && val != nil {
		smss, _ = val.(map[string][]string)
	}

	return
}

func New() *Context {
	return new(Context)
}
