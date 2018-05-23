# reddit-go

go get github.com/Jleagle/reddit-go/reddit

##### Create a client
```go
var client = reddit.GetClient(
	os.Getenv("API_CLIENT"),
	os.Getenv("API_SECRET"),
	os.Getenv("AUTH_CALLBACK_URL"),
	os.Getenv("UNIQUE_USER_AGENT"),
)
```

##### Send the user to a login page
```go
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	u, state := client.Login([]reddit.AuthScope{reddit.ScopeRead}, false, "")

	setToSession("state", state)

	http.Redirect(w, r, u, 302)
}
```

##### Retrieve a code and swap it for a token to store
```go
func LoginCallbackHandler(w http.ResponseWriter, r *http.Request) {

	errStr := r.URL.Query().Get("error")
	if errStr != "" {
		fmt.Println(errStr)
	}

	state := getFromSession("state")

	if state != r.URL.Query().Get("state") {
		fmt.Println(errors.New("invalid state"))
	}

	tok, err := client.GetToken(r)
	if err != nil {
		fmt.Println(err)
	}

	setToSession("token", tok)

	http.Redirect(w, r, "/", 302)
}
```

#### Make an API call
```go
func ListingHandler(w http.ResponseWriter, r *http.Request) {

	tok := getFromSession("tok")

	client.SetToken(tok)

	posts, err := client.GetListing(reddit.ListingOptions{})
	if err != nil {
		fmt.Println(err)
	}
}
```
