package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/helpers"
	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/internal/entity"
	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/internal/entity/userentity"
	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/internal/src/user"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type UserHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// RegisterNewUser is handler function to Handle user registration
func (uh *UserHandler) RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	responder := helpers.NewHTTPResponse("registerNewUser")
	ctx := r.Context()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, fmt.Sprint(err))
		return
	}

	var payload userentity.UserRequest
	err = json.Unmarshal(body, &payload)
	if err != nil {
		responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, fmt.Sprint(err))
		return
	}

	err = uh.service.InsertNewUser(ctx, &payload) //uh.service.InsertNewUser(&payload)
	if err != nil {
		causer := errors.Cause(err)
		switch causer {
		case entity.ErrUserAlreadyExist:
			responder.FieldErrors(w, err, http.StatusNotAcceptable, err.Error())
			return
		default:
			responder.FieldErrors(w, err, http.StatusInternalServerError, fmt.Sprint(err))
			return
		}
	}
	responder.SuccessWithoutData(w, http.StatusCreated, "Succes to register new user")
	return
}

// FindUserByUserID is handler function to Handle find user
func (ud *UserHandler) FindUserByUserID(w http.ResponseWriter, r *http.Request) {
	var (
		userID    = mux.Vars(r)["id"]
		responder = helpers.NewHTTPResponse("registerNewUser")
		ctx       = r.Context()
	)

	findUser, err := ud.service.FindUser(ctx, userID)
	if err != nil {
		causer := errors.Cause(err)
		switch causer {
		case entity.ErrUserNotExist:
			responder.ErrorJSON(w, http.StatusNotFound, "user not found")
			return
		default:
			responder.FailureJSON(w, err, http.StatusInternalServerError)
			return
		}
	}

	responder.SuccessJSON(w, findUser, http.StatusOK, "User found")
	return
}
