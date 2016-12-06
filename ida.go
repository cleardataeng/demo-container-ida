package main

import "os"
import "net/http"
import "encoding/json"
import "io/ioutil"
import "log"

func slashHandler(w http.ResponseWriter, r *http.Request) {
	// Get dactyl's URL from the environment
	url := os.Getenv("DACTYL_URL")
	if len(url) < 1 {
		http.Error(w, "No DACTYL_URL is set in the environment",
			http.StatusInternalServerError)
		return
	}

	// GET / from dactyl
	res, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Load JSON response
	type DactylInfo struct {
		Hostname   string
		RemoteAddr string
	}
	var dinfo DactylInfo
	err = json.Unmarshal(response, &dinfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Find my hostname
	name, err := os.Hostname()
	if err != nil {
		name = "(failed to load hostname)"
	}

	// Ida's response include the info from dactyl
	type IdaInfo struct {
		Hostname   string
		RemoteAddr string
		DactylInfo DactylInfo
	}
	info := IdaInfo{
		name,
		r.RemoteAddr,
		dinfo,
	}

	// Return
	js, err := json.Marshal(info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func logHandler(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
	handler.ServeHTTP(w, r)
    })
}

func main() {
	http.HandleFunc("/", slashHandler)
	http.ListenAndServe(":8081", logHandler(http.DefaultServeMux))
}
