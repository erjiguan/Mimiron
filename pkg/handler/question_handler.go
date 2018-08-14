package handler

import (
	"fmt"
	"net/http"
)

type QuestionHandler struct {
}

func (h *QuestionHandler) GetQuestion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "lvyang")
}
