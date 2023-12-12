package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

type CaptiveJson struct {
	Captive       bool   `json:"captive"`
	UserPortalUrl string `json:"user-portal-url"`
}

var captive CaptiveJson = CaptiveJson{
	Captive:       true,
	UserPortalUrl: "http://192.168.2.2/",
}

const linuxCmd = "nmcli"

func WifiName() string {
	platform := runtime.GOOS

	if platform == "linux" {
		return forLinux()
	}

	return ""
}

func forLinux() string {
	cmd := exec.Command(linuxCmd, `-f`, `SSID`, `-t`, `dev`, `wifi`)
	log.Printf("cmd", cmd)
	output, _ := cmd.CombinedOutput()

	// if err != nil {
	// 	panic(err)
	// }

	return string(output)
}

func main() {

	http.HandleFunc("/captive", func(w http.ResponseWriter, r *http.Request) {
		log.Printf(r.Host, r.URL.Path, r.Body, r.Header, r.Method)
		j, err := json.Marshal(captive)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/captive+json")
		w.Write(j)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf(r.Host, r.URL.Path, r.Body, r.Header, r.Method)

		http.ServeFile(w, r, "./static/index.html")

	})

	http.HandleFunc("/wifilist", func(w http.ResponseWriter, r *http.Request) {
		log.Printf(r.Host, r.URL.Path, r.Body, r.Header, r.Method)
		wifis := WifiName()
		log.Printf("wifis", wifis)
		w.Write([]byte(wifis))
	})

	log.Fatal(http.ListenAndServe(":50052", nil))

}
