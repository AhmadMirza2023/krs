package course

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AhmadMirza2023/krs/types"
	"github.com/AhmadMirza2023/krs/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.CourseStore
}

func NewHandler(store types.CourseStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/courses", h.handleGetCourses).Methods("GET")
	router.HandleFunc("/course/{id}", h.handleGetCourseById).Methods("GET")
}

func (h *Handler) handleGetCourses(w http.ResponseWriter, r *http.Request) {
	// get all the courses in db
	courses, err := h.store.GetCourses()
	if err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "failed", err.Error(), nil)
		return
	}

	// create response
	utils.FormatResponse(w, http.StatusOK, "success", "successfully retrieved courses data", courses)
}

func (h *Handler) handleGetCourseById(w http.ResponseWriter, r *http.Request) {
	// get id from request
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		er := fmt.Sprintf("failed to convert string to int: %v", err)
		utils.FormatResponse(w, http.StatusBadRequest, "failed", er, nil)
	}

	// get course by id
	course, err := h.store.GetCourseById(id)
	if err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "failed", err.Error(), nil)
		return
	}

	// create response
	utils.FormatResponse(w, http.StatusOK, "success", "success get course", course)
}
