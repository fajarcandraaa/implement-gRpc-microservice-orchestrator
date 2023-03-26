package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/helpers"
	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/internal/entity"
	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/internal/entity/bookentity"
	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/internal/src/book"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type BookHandler struct {
	service book.Service
}

func NewBookHandler(service book.Service) *BookHandler {
	return &BookHandler{
		service: service,
	}
}

func (bu *BookHandler) InsertNewBook(w http.ResponseWriter, r *http.Request) {
	responder := helpers.NewHTTPResponse("registerNewUser")
	ctx := r.Context()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, fmt.Sprint(err))
		return
	}

	var payload bookentity.BookRequest
	err = json.Unmarshal(body, &payload)
	if err != nil {
		responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, fmt.Sprint(err))
		return
	}

	err = bu.service.AddNewBook(ctx, &payload)
	if err != nil {
		causer := errors.Cause(err)
		switch causer {
		case entity.ErrBookAlreadyExist:
			responder.FieldErrors(w, err, http.StatusNotAcceptable, err.Error())
			return
		default:
			responder.FieldErrors(w, err, http.StatusInternalServerError, fmt.Sprint(err))
			return
		}
	}
	responder.SuccessWithoutData(w, http.StatusCreated, "Succes to add new book")
	return
}

func (bu *BookHandler) FindBookById(w http.ResponseWriter, r *http.Request) {
	var (
		bookID    = mux.Vars(r)["id"]
		responder = helpers.NewHTTPResponse("registerNewUser")
		ctx       = r.Context()
	)

	findBook, err := bu.service.FindBook(ctx, bookID)
	if err != nil {
		causer := errors.Cause(err)
		switch causer {
		case entity.ErrBookNotExist:
			responder.FieldErrors(w, err, http.StatusNotAcceptable, err.Error())
			return
		default:
			responder.FieldErrors(w, err, http.StatusInternalServerError, fmt.Sprint(err))
			return
		}
	}

	responder.SuccessJSON(w, findBook, http.StatusOK, "Book found")
	return
}
