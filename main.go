package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

var (
	server = &http.Server{Addr: ":8080"}
)

type Img struct {
	Path string `json:"path"`
}

func main() {
	http.HandleFunc("/", imgHandler)
	log.Fatal(server.ListenAndServe())
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Use post", 405)
		return
	}
	img := Img{}

	err := json.NewDecoder(r.Body).Decode(&img)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("Error decoding body: %v", fmt.Sprint(err)))
	}
	imgUrl := r.URL.Path[1:]
	if len(imgUrl) > 2 {
		cmd := exec.Command("python", "/workspace/classify_nsfw.py", "--model_def", "/workspace/nsfw_model/deploy.prototxt", "--pretrained_model", "/workspace/nsfw_model/resnet_50_1by2_nsfw.caffemodel", fmt.Sprint(img.Path))
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintf(w, fmt.Sprintf("error in Output():", fmt.Sprint(err)))
		}
		fmt.Fprintf(w, fmt.Sprintf("Result: %v", string(output)))
	}
}
