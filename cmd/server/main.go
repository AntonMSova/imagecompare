package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	archParam = "arch"
	osParam = "os"
)

func main() {
	m := http.NewServeMux()
	fmt.Println("Starting server")
	m.Handle("/", http.FileServer(http.Dir("./public")))
	m.HandleFunc("/imagecompare", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("in /imagecompare")
		arch := request.URL.Query().Get(archParam)
		runtimeOs := request.URL.Query().Get(osParam)
		if arch == "" || runtimeOs == "" {
			fmt.Println("missing parameters")
			_, _ = writer.Write([]byte("must provide arch and os query parameters\n"))
			writer.WriteHeader(404)
			_ = request.Body.Close()
		}
		writer.Header().Set("Content-Disposition", "attachment; filename=compare")
		http.ServeFile(writer, request, fmt.Sprintf("/arch=%s+os=%s", arch, runtimeOs))
	})
	err := http.ListenAndServe(":8080", m)
	if err != nil {
		fmt.Printf("error starting meta server: %v", err)
		os.Exit(1)
	}
}