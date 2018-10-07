package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

var (
	docPathOpt = flag.String("d", "/html", "document root path")
	portOpt    = flag.Int("p", 8080, "port to listen")
)

func exitWithUsage() {
	flag.Usage()
	os.Exit(0)
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Printf("  %s [OPTIONS] [\"run\"]\n", os.Args[0])
		flag.PrintDefaults()
	}
}

type HtmlServer struct {
	Server *http.Server
	Handler http.Handler
	Port    string
	DocPath string
}

func initHtmlServer (port int, path string) *HtmlServer {
	portStr := strconv.Itoa(port)
	addr := net.JoinHostPort("", portStr)
	fileHandler := http.FileServer(http.Dir(path))
	fileServer := &http.Server{Addr: addr, Handler: fileHandler}
	return &HtmlServer{fileServer, fileHandler, portStr, path}
}

func main() {
	flag.Parse()

	if len(flag.Args()) != 1 {
		exitWithUsage()
	}
	if flag.Args()[0] != "run" {
		exitWithUsage()
	}
	server := initHtmlServer(*portOpt, *docPathOpt)

	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	var ips []net.IP
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Fatal(err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			ips = append(ips, ip)
		}
	}

	fmt.Printf("htmlserver start, ips: %s, port: %s, doc root: %s\n", ips, server.Port, server.DocPath)
	log.Fatal(server.Server.ListenAndServe())
}
