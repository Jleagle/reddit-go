package reddit

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/beefsack/go-rate"
	"golang.org/x/oauth2"
)

const (
	defaultUserAgent = "github.com/Jleagle/reddit-go"

	authURL        = "https://www.reddit.com/api/v1/authorize"
	authCompactURL = "https://www.reddit.com/api/v1/authorize.compact"
	tokenURL       = "https://www.reddit.com/api/v1/access_token"
	apiURL         = "https://oauth.reddit.com/"
)

var (
	errNoToken = errors.New("no token set")
	errNoCode  = errors.New("no code found")
)

type transport struct {
	http.RoundTripper // Interface
	useragent string
}

func (t *transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {

	req.Header.Set("User-Agent", t.useragent)
	return t.RoundTripper.RoundTrip(req)
}

type Reddit struct {
	oauthConfig oauth2.Config
	ctx         context.Context
	httpClient  *http.Client
	throttle    *rate.RateLimiter
}

func GetClient(client string, secret string, redirect string, userAgent string) (reddit Reddit) {

	if userAgent == "" {
		userAgent = defaultUserAgent
	}

	config := oauth2.Config{
		ClientID:     client,
		ClientSecret: secret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authURL,
			TokenURL: tokenURL,
		},
		RedirectURL: redirect,
	}

	reddit = Reddit{
		oauthConfig: config,
		ctx:         context.Background(),
	}

	// Use the custom HTTP client when requesting a token.
	httpClient := &http.Client{
		Timeout:   2 * time.Second,
		Transport: &transport{http.DefaultTransport, userAgent},
	}

	reddit.ctx = context.WithValue(reddit.ctx, oauth2.HTTPClient, httpClient)

	return reddit
}

func (r *Reddit) Throttle(duration time.Duration) {
	if duration == 0 {
		r.throttle = nil
	} else {
		r.throttle = rate.New(1, duration)
	}
}

func (r Reddit) Login(scopes []AuthScope, mobile bool, state string) (string, string) {

	// Set scopes
	r.oauthConfig.Scopes = []string{}
	for _, v := range scopes {
		r.oauthConfig.Scopes = append(r.oauthConfig.Scopes, string(v))
	}

	// Set auth URL
	if mobile {
		r.oauthConfig.Endpoint.AuthURL = authCompactURL
	}

	// Generate state
	if state == "" {
		state = strconv.Itoa(int(rand.Int31()))
	}

	// Generate login URL
	u := r.oauthConfig.AuthCodeURL(
		state,
		oauth2.SetAuthURLParam("response_type", "code"),
		oauth2.SetAuthURLParam("duration", "permanent"),
	)

	return u, state
}

func (r Reddit) GetToken(re *http.Request) (tok *oauth2.Token, err error) {

	code := re.URL.Query().Get("code")

	if code == "" {
		return tok, errNoCode
	}

	return r.oauthConfig.Exchange(r.ctx, code)
}

func (r *Reddit) SetToken(tok *oauth2.Token) {
	r.httpClient = r.oauthConfig.Client(r.ctx, tok)
}

func (r Reddit) fetchGet(u string, data url.Values, i interface{}) (err error) {

	u = apiURL + u

	encoded := data.Encode()
	if encoded != "" {
		u = u + "?" + encoded
	}

	if r.httpClient == nil {
		return errNoToken
	}

	if r.throttle != nil {
		r.throttle.Wait()
	}

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return err
	}

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, i)
	if err != nil {
		return err
	}

	return nil
}

func (r Reddit) fetchPost(u string, data url.Values, i interface{}) (err error) {

	u = apiURL + u

	if r.httpClient == nil {
		return errNoToken
	}

	if r.throttle != nil {
		r.throttle.Wait()
	}

	req, err := http.NewRequest("POST", u, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.PostForm = data

	err = r.fetch(req, i)
	if err != nil {
		return err
	}

	return nil
}

func (r Reddit) fetch(req *http.Request, i interface{}) (err error) {

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, i)
	if err != nil {
		return err
	}

	return nil
}
