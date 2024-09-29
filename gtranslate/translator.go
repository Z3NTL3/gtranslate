/*
Written by @z3ntl3 under GNU license

License: see LICENSE file
*/
package gtranslate

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	qs "github.com/google/go-querystring/query"
)

const (
	API_URL = "https://translate.google.com/translate_a/single"
	DstTarget = "t"
)

type Lang = string

/*
 Comes into handy when you want to implement your own GoogleTranslator
 instead of the current internal one
*/
type GoogleTranslator[T any] interface {
	SetProxy(string) (*T, error)
	ClearProxy() *T 
	SetTimeout(time.Duration) *T
	Translate(TranslationOptions) (string, error)
}

/* 
Translation options
Gets encoded in query string
*/
 type TranslationOptions struct {
	Client string `url:"client"` // Client, should be 'gtx'
	SourceLang string `url:"sl"` // Source language
	TargetLang string `url:"tl"` // Target language
	DstTarget string `url:"dt"` // Destination target, should be 't'
	Query string `url:"q"` // Query (text to translate from source lang to target lang)
}

/*
	Starts the translation handshake.
	Works with every GoogleTranslator instance that comforts the interface
	
	Returns translation string in targetLang or error if any
*/
func Translate[T any, Trans GoogleTranslator[T]](instance Trans, options TranslationOptions) (string, error) {
	return instance.Translate(options)
}

// GoogleTranslator client
type Client struct {
	*http.Client
}

// Creates a new google translator instance
func New() *Client {
	return &Client{&http.Client{}}
}

// Sets proxy to use for google translator instance
func (c *Client) SetProxy(proxyURI string) (*Client, error) {
	uri, err := url.Parse(proxyURI)
	if err != nil {
		return nil, err
	}

	c.Transport = &http.Transport{
		Proxy: http.ProxyURL(uri),
	}

	return c, nil
}

// Clears the proxy so that it is not used on the google translator instance
// when translating
func (c *Client) ClearProxy() *Client {
	c.Transport = &http.Transport{}
	return c
}

// Sets timeout
func (c *Client) SetTimeout(d time.Duration) *Client {
	c.Timeout = d
	return c
}

func (c *Client) Translate(opts TranslationOptions) (string, error) {
	// qs: options
	
	if opts.Client == "" {
		opts.Client = "gtx"
	}

	if opts.DstTarget == "" {
		opts.DstTarget = "t"
	}

	if opts.SourceLang == "" {
		return "", fmt.Errorf("sourceLang must not be empty")
	}

	if opts.TargetLang == "" {
		return "", fmt.Errorf("targetLang must not be empty")
	}
	
	if opts.Query == "" {
		return "", fmt.Errorf("query must not be empty")
	}

	v, err := qs.Values(opts)
	if err != nil {
		return "", err
	}

	// setup the request 
	req, err := http.NewRequest(
		http.MethodGet, 
		fmt.Sprintf("%s?%s", API_URL, v.Encode()),
		nil,
	)

	v = nil // gc

	// some headers to use
	{
		req.Header.Add("Referer", "https://translate.google.com/")
		req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Mobile Safari/537.36")
	}

	// perform the request
	res, err := c.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	// translation failure
	if res.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("%d %s", res.StatusCode, res.Status))
	}

	rawBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if !(len(string(rawBody)) > 0) {
		return "", errors.New("Google Translator could not translate given input query")
	}


	return strings.Split(string(rawBody), "\"")[1], nil
}