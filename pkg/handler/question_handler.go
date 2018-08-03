package handler

import (
	"net/http"
	"fmt"
)

type QuestionHandler struct {

}

func (h *QuestionHandler) GetQuestion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "lvyang")
}
