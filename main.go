package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

var (
	server = &http.Server{Addr: ":8080"}
)

type Img struct {
	Path string `json:"path"`
}
type NsfwResult struct {
	Score string
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
	cmd := exec.Command("python", "/workspace/classify_nsfw.py", "--model_def", "/workspace/nsfw_model/deploy.prototxt", "--pretrained_model", "/workspace/nsfw_model/resnet_50_1by2_nsfw.caffemodel", fmt.Sprint(img.Path))
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("error in Output():", fmt.Sprint(err)))
	}
	score := strings.Trim(after(string(output), "NSFW_SCORE="), " ")
	result := &NsfwResult{Score: strings.TrimSuffix(score, "\n")}
	json, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func after(value string, a string) string {
	// Get substring after a string.
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}
