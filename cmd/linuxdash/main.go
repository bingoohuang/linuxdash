package main

import (
	"flag"
	"fmt"
	"github.com/bingoohuang/linuxdash"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.FS(linuxdash.DashStatic)))
	http.HandleFunc("/server/", linuxdash.MakeDashServe(linuxdash.ExecuteShell))

	listen := flag.String("listen", ":8081", "Where the server listens for connections. [interface]:port")
	version := flag.Bool("v", false, "Print version and exit")
	flag.Parse()

	if *version {
		fmt.Println("v1.0.0 2021-07-21 09:57:34")
		os.Exit(0)
	}

	fmt.Println("Starting http server at:", *listen)
	if err := http.ListenAndServe(*listen, nil); err != nil {
		fmt.Println("Error starting http server:", err)
		os.Exit(1)
	}
}
