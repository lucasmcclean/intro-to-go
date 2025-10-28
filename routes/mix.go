package routes

import (
	"encoding/json"
	"net/http"
)

type Mixer interface {
	Mix() string
}

type LeftFirst struct {
	Left  string
	Right string
}

type RightFirst struct {
	Left  string
	Right string
}

type Weave struct {
	Left  string
	Right string
}

func (lf *LeftFirst) Mix() string {
	return lf.Left + lf.Right
}

func (rf *RightFirst) Mix() string {
	return rf.Right + rf.Left
}

func (w *Weave) Mix() string {
	res := ""
	for i := 0; i < max(len(w.Left), len(w.Right)); i++ {
		if i < len(w.Left) {
			res += string(w.Left[i])
		}
		if i < len(w.Right) {
			res += string(w.Right[i])
		}
	}
	return res
}

type MixRequest struct {
	Left  string `json:"left"`
	Right string `json:"right"`
	How   string `json:"how"`
}

func MixHandler(w http.ResponseWriter, r *http.Request) {
	var req MixRequest
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	var mixer Mixer
	switch req.How {
	case "left-first":
		mixer = &LeftFirst{Left: req.Left, Right: req.Right}
	case "right-first":
		mixer = &RightFirst{Left: req.Left, Right: req.Right}
	case "weave":
		mixer = &Weave{Left: req.Left, Right: req.Right}
	default:
		http.Error(w, "unknown mix type", http.StatusBadRequest)
		return
	}

	result := mixer.Mix()

	response := map[string]string{"result": result}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "error encoding JSON", http.StatusInternalServerError)
	}
}