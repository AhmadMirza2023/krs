package user

import (
	"fmt"
	"net/http"

	"github.com/AhmadMirza2023/krs/config"
	"github.com/AhmadMirza2023/krs/service/auth"
	"github.com/AhmadMirza2023/krs/types"
	"github.com/AhmadMirza2023/krs/utils"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload

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

	// check email if is already exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		er := fmt.Sprintf("user with email %s is already exists", payload.Email)
		utils.FormatResponse(w, http.StatusBadRequest, "failed", er, nil)
		return
	}

	// hash password user
	HashPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.FormatResponse(w, http.StatusInternalServerError, "failed", err.Error(), nil)
		return
	}

	// create user
	newUser, err := h.store.CreateUser(types.User{
		Email:    payload.Email,
		Password: HashPassword,
		Name:     payload.Name,
		Nim:      payload.Nim,
		Semester: payload.Semester,
		Major:    payload.Major,
		Faculty:  payload.Faculty,
	})
	if err != nil {
		utils.FormatResponse(w, http.StatusInternalServerError, "failed", err.Error(), nil)
		return
	}
	response := types.FullResponse{
		Meta: types.MetaData{
			Code:    http.StatusCreated,
			Status:  "success",
			Message: "User Registered",
		},
		Data: newUser,
	}
	utils.WriteJson(w, http.StatusCreated, response)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginUserPayload

	// get JSON payload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "failed", "invalid payload "+err.Error(), nil)
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := "invalid payload " + err.(validator.ValidationErrors).Error()
		utils.FormatResponse(w, http.StatusBadRequest, "failed", errors, nil)
	}

	// check credentials
	user, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "failed", "not found, invalid email or password", nil)
		return
	}

	// check user password
	if !auth.ComparePassword(user.Password, []byte(payload.Password)) {
		utils.FormatResponse(w, http.StatusBadRequest, "failed", "not found, invalid email or password", nil)
		return
	}

	// generate JWT
	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, user.Id)
	if err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "failed", err.Error(), nil)
		return
	}

	// create response
	loginData := map[string]interface{}{
		"access_token": token,
		"token_type":   "JWT",
		"user":         user,
	}
	utils.FormatResponse(w, http.StatusOK, "success", "User Authenticated", loginData)
}
