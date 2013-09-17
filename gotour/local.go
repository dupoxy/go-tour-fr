// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !appengine

package main

import (
	"flag"
	"fmt"
	"go/build"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"code.google.com/p/go.talks/pkg/socket"

	// Imports so that go build/install automatically installs them.
	_ "code.google.com/p/go-tour-french/pic"
	_ "code.google.com/p/go-tour-french/tree"
	_ "code.google.com/p/go-tour-french/wc"
)

const (
	basePkg    = "code.google.com/p/go-tour-french/"
	socketPath = "/socket"
)

var (
	httpListen  = flag.String("http", "127.0.0.1:3999", "host:port à écouter")
	openBrowser = flag.Bool("openbrowser", true, "Ouvrir le navigateur automatiquement")
)

var (
	// GOPATH containing the tour packages
	gopath = os.Getenv("GOPATH")

	httpAddr string
)

// isRoot reports whether path is the root directory of the tour tree.
// To be the root, it must have content and template subdirectories.
func isRoot(path string) bool {
	_, err := os.Stat(filepath.Join(path, "content", "tour.article"))
	if err == nil {
		_, err = os.Stat(filepath.Join(path, "template", "tour.tmpl"))
	}
	return err == nil
}

func findRoot() (string, error) {
	ctx := build.Default
	p, err := ctx.Import(basePkg, "", build.FindOnly)
	if err == nil && isRoot(p.Dir) {
		return p.Dir, nil
	}
	tourRoot := filepath.Join(runtime.GOROOT(), "misc", "tour")
	ctx.GOPATH = tourRoot
	p, err = ctx.Import(basePkg, "", build.FindOnly)
	if err == nil && isRoot(tourRoot) {
		gopath = tourRoot
		return tourRoot, nil
	}
	return "", fmt.Errorf("n'a pas pu trouver le contenu de go-tour-french; vérifier $GOROOT et $GOPATH")
}

func main() {
	flag.Parse()

	// find and serve the go tour files
	root, err := findRoot()
	if err != nil {
		log.Fatalf("n'a pas pu trouver les fichiers du tour: %v", err)
	}

	log.Println("Sers le contenu de", root)

	host, port, err := net.SplitHostPort(*httpListen)
	if err != nil {
		log.Fatal(err)
	}
	if host == "" {
		host = "localhost"
	}
	if host != "127.0.0.1" && host != "localhost" {
		log.Print(localhostWarning)
	}
	httpAddr = host + ":" + port

	if err := initTour(root); err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir(root))
	http.Handle("/favicon.ico", fs)
	http.Handle("/static/", fs)
	http.Handle("/talks/", fs)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			if err := renderTour(w); err != nil {
				log.Println(err)
			}
			return
		}
		http.Error(w, "not found", 404)
	})

	http.Handle(socketPath, socket.Handler)

	err = serveScripts(filepath.Join(root, "js"), "SocketTransport")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		url := "http://" + httpAddr
		if waitServer(url) && *openBrowser && startBrowser(url) {
			log.Printf("Une fenêtre de navigateur devrait s'ouvrir. Si non, s'il vous plaît visitez %s", url)
		} else {
			log.Printf("S'il vous plaît ouvrir votre navigateur Web et visitez %s", url)
		}
	}()
	log.Fatal(http.ListenAndServe(httpAddr, nil))
}

const localhostWarning = `
ATTENTION! ATTENTION! ATTENTION!

Je semble être à l'écoute sur une adresse qui n'est pas localhost.
Toute personne ayant accès à cette adresse et au port aura accès
à cette machine comme l'utilisateur exécutant gotour.

Si vous ne comprenez pas ce message, appuyez sur Control-C pour mettre fin à ce processus.

ATTENTION! ATTENTION! ATTENTION!
`

type response struct {
	Output string `json:"output"`
	Errors string `json:"compile_errors"`
}

func init() {
	socket.Environ = environ
}

// environ returns the original execution environment with GOPATH
// replaced (or added) with the value of the global var gopath.
func environ() (env []string) {
	for _, v := range os.Environ() {
		if !strings.HasPrefix(v, "GOPATH=") {
			env = append(env, v)
		}
	}
	env = append(env, "GOPATH="+gopath)
	return
}

// waitServer waits some time for the http Server to start
// serving url. The return value reports whether it starts.
func waitServer(url string) bool {
	tries := 20
	for tries > 0 {
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
			return true
		}
		time.Sleep(100 * time.Millisecond)
		tries--
	}
	return false
}

// startBrowser tries to open the URL in a browser, and returns
// whether it succeed.
func startBrowser(url string) bool {
	// try to start the browser
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}

// prepContent for the local tour simply returns the content as-is.
func prepContent(r io.Reader) io.Reader { return r }

// socketAddr returns the WebSocket handler address.
func socketAddr() string { return "ws://" + httpAddr + socketPath }
