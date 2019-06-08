package main

import (
	"sync"
	"flag"
	"time"
	"net/http"
	"log"
	"html/template"
	"expvar"
)

const changeURL = "https://code.google.com/p/go/source/detail?r=go1.5"

var (
	httpAdr = flag.String("http", "localhost:8080", "Liste address")
	pollPeriod   = flag.Duration("poll", 5*time.Second, "Poll period")

	hitCount		= expvar.NewInt("hitCount")
	pollCount		= expvar.NewInt("pollCount")
	pollError		= expvar.NewString("pollError")
	pollErrorCount	= expvar.NewInt("pollErrorCount")
)

var state struct {
	sync.RWMutex
	yes bool // whether Go 1.5 has been tagged.
}

var tmpl = template.Must(template.New("root").Parse(`
<!DOCTYPE html><html><body><center>
	<h2>Is Go 1.5 out yet?</h2>
	<h1>
	{{if .Yes}}
		<a href="{{.URL}}">YES!</a>
	{{else}}
		No.
	{{end}}
	</h1>
</center></body></html>
`))

func main() {
	flag.Parse()
	go poll(*pollPeriod)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(*httpAdr, nil))
}

func poll(period time.Duration) {
	for !isTagged() {
		time.Sleep(period)
	}
	state.Lock()
	state.yes = true
	state.Unlock()
}

func isTagged() bool {
	pollCount.Add(1)
	r, err := http.Head(changeURL)
	if err != nil {
		log.Print(err)
		pollError.Set(err.Error())
		pollErrorCount.Add(1)
		return false
	}
	return r.StatusCode == http.StatusOK
}

func handler(w http.ResponseWriter, r *http.Request) {
	hitCount.Add(1)
	state.RLock()
	data := struct {
		Yes bool
		URL string
	}{
		Yes: state.yes,
		URL: changeURL,
	}
	state.RUnlock()
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Print(err)
	}
}