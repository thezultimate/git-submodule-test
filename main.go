package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const directory = "./alpine-git-lfs"

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The path you entered: %s \n", r.URL.Path[1:])

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Fprintf(w, "Could not read submodule directory: %v", err)
	}

	fmt.Fprintf(w, "Contents of %s submodule directory: \n", directory)
	for _, file := range files {
		fmt.Fprintln(w, file.Name())
	}
}
