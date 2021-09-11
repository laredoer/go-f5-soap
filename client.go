package soap

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"github.com/FishGoddess/logit"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type Option func(*options)

func WithRequestTimeout(t time.Duration) Option {
	return func(o *options) {
		o.conTimeout = t
	}
}

func WithBasicAuth(login, password string) Option {
	return func(o *options) {
		o.auth = &basicAuth{Login: login, Password: password}
	}
}

func WithTLS(tls *tls.Config) Option {
	return func(o *options) {
		o.tlsCfg = tls
	}
}

func WithTimeout(t time.Duration) Option {
	return func(o *options) {
		o.timeout = t
	}
}

func WithHTTPHeaders(headers map[string]string) Option {
	return func(o *options) {
		o.httpHeaders = headers
	}
}

func WithDebug() Option {
	return func(o *options) {
		o.debug = true
	}
}

type options struct {
	tlsCfg           *tls.Config
	auth             *basicAuth
	client           HTTPClient
	timeout          time.Duration
	conTimeout       time.Duration
	tlsHShakeTimeout time.Duration
	httpHeaders      map[string]string
	debug            bool
}
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type basicAuth struct {
	Login    string
	Password string
}

type HTTPError struct {
	StatusCode   int
	ResponseBody []byte
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP Status %d: %s", e.StatusCode, string(e.ResponseBody))
}

var defaultOptions = options{
	timeout:          time.Duration(30 * time.Second),
	conTimeout:       time.Duration(90 * time.Second),
	tlsHShakeTimeout: time.Duration(15 * time.Second),
}

type Client struct {
	url     string
	opts    *options
	headers []interface{}
}

func NewClient(url string, opt ...Option) *Client {

	opts := defaultOptions

	for _, o := range opt {
		o(&opts)
	}

	return &Client{
		url:  url,
		opts: &opts,
	}
}

func (c *Client) Call(ctx context.Context, request interface{}) (response []byte, err error) {

	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)

	if err := encoder.Encode(request); err != nil {
		return nil, err
	}

	if err := encoder.Flush(); err != nil {
		return nil, err
	}

	if c.opts.debug {
		logit.Debug("==== request body ===>>\n", buffer.String())
	}

	req, err := http.NewRequest("POST", c.url, buffer)
	if err != nil {
		return nil, err
	}
	if c.opts.auth != nil {
		req.SetBasicAuth(c.opts.auth.Login, c.opts.auth.Password)
	}

	req = req.WithContext(ctx)
	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	req.Header.Set("User-Agent", "go-f5-soap/0.1")
	if c.opts.httpHeaders != nil {
		for k, v := range c.opts.httpHeaders {
			req.Header.Set(k, v)
		}
	}
	req.Close = true

	client := c.opts.client
	if client == nil {
		tr := &http.Transport{
			TLSClientConfig: c.opts.tlsCfg,
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				d := net.Dialer{Timeout: c.opts.timeout}
				return d.DialContext(ctx, network, addr)
			},
			TLSHandshakeTimeout: c.opts.tlsHShakeTimeout,
		}
		client = &http.Client{Timeout: c.opts.conTimeout, Transport: tr}
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		body, _ := ioutil.ReadAll(res.Body)
		return nil, &HTTPError{
			StatusCode:   res.StatusCode,
			ResponseBody: body,
		}
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if c.opts.debug {
		logit.Debug("==== response body ===>>\n", string(body))
	}

	return body, nil
}
