package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, "<h1>Hello OpenShift!</h1><br/>\n")
	fmt.Fprintln(w, "<html><div style='border:1px solid black; align:center; width:80%'>\n")

	env := os.Environ()

	for _, value := range env {
		name := strings.Split(value, "=")
		fmt.Fprintf(w, "%s : %s <br />", name[0], name[1])
	}

	fmt.Fprintln(w, "</div>\n<p>&nbsp;</p>\n<div style='border:1px solid black; align:center; width:80%'>\n")
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, addr := range addrs {
		fmt.Fprintf(w, "%s<br />\n", addr.String())
	}
	fmt.Fprintln(w, "</div>\n")

	fmt.Fprintln(w, "</div>\n<p>&nbsp;</p>\n<div style='border:1px solid black; align:center; width:80%'>\n")
	ifs, err := net.Interfaces()
	if err != nil {
		return
	}

	for _, ifc := range ifs {
		fmt.Fprintf(w, "%s -> %s<br />\n", ifc.Name, ifc.HardwareAddr.String())
	}
	fmt.Fprintln(w, "</div>\n")

	fmt.Fprintln(w, "</html>\n")
	
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
