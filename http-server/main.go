package main

import (
	"encoding/json"
	"io"
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

const linuxCmd = "sudo nmcli"
const linuxArgs = "-f SSID dev wifi"

func WifiName() string {
	platform := runtime.GOOS
	log.Printf("platform", platform)
	if platform == "linux" {
		return forLinux()
	}

	return ""
}

func forLinux() string {
	cmd := exec.Command(linuxCmd, linuxArgs)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	// start the command after having set up the pipe
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	defer cmd.Wait()

	var str string

	if b, err := io.ReadAll(stdout); err == nil {
		log.Printf("cmd", b)
		str += (string(b) + "\n")
	}
	log.Printf("string", str)

	name := strings.Replace(str, "\n", "", -1)
	return name
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
