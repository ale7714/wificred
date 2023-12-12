package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

type CaptiveJson struct {
	Captive       bool   `json:"captive"`
	UserPortalUrl string `json:"user-portal-url"`
}

var captive CaptiveJson = CaptiveJson{
	Captive:       true,
	UserPortalUrl: "http://192.168.2.2/",
}

type WiFis struct {
	Ssids []string
}

type Credentials struct {
	Ssid     string
	Password string
}

var credentials Credentials
var mu sync.Mutex

func WifiNames() WiFis {
	var wifis WiFis

	if runtime.GOOS == "linux" {
		cmd := exec.Command("nmcli", `-f`, `SSID`, `-t`, `dev`, `wifi`)

		output, err := cmd.CombinedOutput()
		log.Printf("out", output)
		if err == nil {
			log.Printf("split", strings.Split(string(output), "\n"))
			wifis.Ssids = strings.Split(string(output), "\n")
		}

	}

	return wifis
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/captive", serveCaptive)
	mux.HandleFunc("/wifilist", getWifiList)
	mux.HandleFunc("/save", saveWifi)

	err := http.ListenAndServe(":50052", mux)
	log.Fatal(err)
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.Host, r.URL.Path, r.Body, r.Header, r.Method)

	t, _ := template.ParseFiles("templates/base.html", "templates/index.html")
	err := t.Execute(w, credentials)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func serveCaptive(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.Host, r.URL.Path, r.Body, r.Header, r.Method)
	j, err := json.Marshal(captive)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/captive+json")
	w.Write(j)
}

func getWifiList(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.Host, r.URL.Path, r.Body, r.Header, r.Method)
	wifis := WifiNames()

	t, _ := template.ParseFiles("templates/base.html", "templates/wifiform.html")
	err := t.Execute(w, wifis)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func saveWifi(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.Host, r.URL.Path, r.Body, r.Header, r.Method)

	if r.Method == "POST" {

		mu.Lock()
		credentials.Ssid = r.FormValue("ssid")
		credentials.Password = r.FormValue("password")
		log.Printf("saving credentials for %s", credentials.Ssid)
		mu.Unlock()
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
