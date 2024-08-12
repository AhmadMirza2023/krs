package spc

import (
	"net/http"

	"github.com/AhmadMirza2023/krs/types"
	"github.com/AhmadMirza2023/krs/utils"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.SPCStore
}

func NewHandler(store types.SPCStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/createSPC", h.handleCreateSPC).Methods("POST")
}

func (h *Handler) handleCreateSPC(w http.ResponseWriter, r *http.Request) {
	var payload types.SPCPayload

	// get JSON payload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "failed", err.Error(), nil)
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := "invalid payload " + err.(validator.ValidationErrors).Error()
		utils.FormatResponse(w, http.StatusBadRequest, "failed", errors, nil)
		return
	}

	// check spc if already exists
	_, err := h.store.GetSPCByUserId(payload.UserId)
	if err == nil {
		utils.FormatResponse(w, http.StatusBadRequest, "failed", err.Error(), nil)
		return
	}

	// create spc
	newSPC, err := h.store.CreateSPC(types.SPC{
		UserId:   payload.UserId,
		CourseId: payload.CourseId,
	})
	if err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "failed", err.Error(), nil)
		return
	}

	// create response
	utils.FormatResponse(w, http.StatusCreated, "success", "SPC created", newSPC)
}
