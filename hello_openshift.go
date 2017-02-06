package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, "Hello OpenShift!  These are my env variables:")
	env := os.Environ()

	for _, value := range env {
		name := strings.Split(value, "=")
		fmt.Fprintf(w, "%s : %s\n", name[0], name[1])
	}

	fmt.Fprintln(w, "These are my network addresses:\n")
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, addr := range addrs {
		fmt.Fprintf(w, "%s\n", addr.String())
	}
	
	fmt.Println("Servicing request.")
}

func listenAndServe(port string) {
	fmt.Printf("serving on %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	go listenAndServe(port)

	port = os.Getenv("SECOND_PORT")
	if len(port) == 0 {
		port = "8888"
	}
	go listenAndServe(port)

	select {}
}
