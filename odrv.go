package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	var port = getEnv("PORT", "8080")

	//http.HandleFunc("/", indexHandler)
	//http.HandleFunc("/query", queryHandler)
	http.HandleFunc("/resolve", resolveHandler)


	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func getEnv(envName string, envDef string) string {
	var env = os.Getenv(envName)
	if env == "" {
		env = envDef
		log.Printf("Defaulting %s to %s", envName, envDef)
	}
	return env
}

func coreFunc(w http.ResponseWriter, r *http.Request, path string) {

	var upstream = getEnv("UPSTREAM", "dns.google")

	if r.Method == "HEAD" {
		w.WriteHeader(http.StatusOK)
		fmt.Print(r.Method, " http://", r.Host, r.URL, "\n")
		return
	}

	err := r.ParseForm()
	rForm := r.Form.Encode()
	rPostForm := r.PostForm.Encode()

	var url url.URL
	url.Scheme = "https"
	url.Host = upstream
	url.Path = path
	url.RawQuery = rForm + rPostForm

	req := http.Request{
		Method: r.Method,
		Header: r.Header,
		URL:    &url,
	}

	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	c := &http.Client{
		//Transport: nil,
	}
	resp, err := c.Do(&req)
	defer resp.Body.Close()

	//var f interface{}
	//json.Unmarshal(body, &f)

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		defer reader.Close()
	default:
		reader = resp.Body
	}

	//w.Header().Add("Content-Type", "application/x-javascript; charset=UTF-8")
	w.WriteHeader(resp.StatusCode)

	body, err := ioutil.ReadAll(reader)
	w.Write(body)

	fmt.Fprint(w)

	fmt.Print(r.Method, " http://", r.Host, r.URL, "\n")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var index = getEnv("INDEX", "/")
	coreFunc(w, r, index)
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	var query = getEnv("QUERY", "/query")

	coreFunc(w, r, query)
}

func resolveHandler(w http.ResponseWriter, r *http.Request) {
	var resolve = getEnv("RESOLVE", "/resolve")

	coreFunc(w, r, resolve)
}
