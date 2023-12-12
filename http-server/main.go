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

func WifiNames() WiFis {
	var wifis WiFis

	if runtime.GOOS == "linux" {
		cmd := exec.Command("nmcli", `-f`, `SSID`, `-t`, `dev`, `wifi`)
		log.Printf("cmd", cmd)
		output, err := cmd.CombinedOutput()

		if err != nil {
			panic(err)
		}

		wifis.Ssids = strings.Split(string(output), "\n")
	}

	return wifis
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/captive", serveCaptive)
	mux.HandleFunc("/wifilist", getWifiList)

	err := http.ListenAndServe(":50052", mux)
	log.Fatal(err)
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.Host, r.URL.Path, r.Body, r.Header, r.Method)

	t, _ := template.ParseFiles("templates/base.html", "templates/index.html")
	err := t.Execute(w, nil)

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
