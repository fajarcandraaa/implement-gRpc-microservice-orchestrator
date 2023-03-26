package helpers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
)

// StatusAPIXXX are statuses for API Response.
const (
	StatusAPISuccess = "SUCCESS"
	StatusAPIError   = "ERROR"
	StatusAPIFailure = "FAILURE"
)

// APIFailureMessage is a default message for failure state.
const APIFailureMessage = "Internal Server Error"

// API represents respond body for HTTP API.
type API struct {
	statusCode int
	Code       int    `json:"code,omitempty"`
	Status     string `json:"status"`
	// Entity     string `json:"entity,omitempty"`
	State   string `json:"state,omitempty"`
	Message string `json:"message,omitempty"`
}

// APISuccess represents body for API on success.
type APISuccess struct {
	*API
	Meta interface{} `json:"meta,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

// PaginationParams represents body for payload for pagination.
type PaginationParams struct {
	Path        string
	Page        string
	TotalRows   int32
	TotalPages  int32
	PerPage     int32
	OrderBy     string
	SortBy      string
	CurrentPage int32
}

// GetPagination to set pagination
func GetPagination(params PaginationParams) (PaginationParams, error) {

	totalPages := int32(math.Ceil(float64(params.TotalRows) / float64(params.PerPage)))

	return PaginationParams{
		Path:        params.Path,
		Page:        params.Page,
		TotalRows:   params.TotalRows,
		TotalPages:  totalPages,
		PerPage:     params.PerPage,
		OrderBy:     params.OrderBy,
		SortBy:      params.SortBy,
		CurrentPage: params.CurrentPage,
	}, nil
}

// SetDefaultPginationParam to set parametes of pagination
func SetDefaultPginationParam(pageParam, perPageParam, orderByParam, sortByParam string) (*PaginationParams, error) {
	if !IsNumber(pageParam) {
		return nil, fmt.Errorf("Page value is string, numeric needed")
	}

	if !IsNumber(perPageParam) {
		return nil, fmt.Errorf("PerPage value is string, numeric needed")
	}
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		page = 1
	}
	perPage, err := strconv.Atoi(perPageParam)
	if err != nil {
		perPage = 10
	}
	orderBy := orderByParam
	if orderBy == "" {
		orderBy = "created_at"
	}
	sortBy := sortByParam
	if sortBy == "" {
		sortBy = "desc"
	}

	result := &PaginationParams{
		Page:    strconv.Itoa(page),
		PerPage: int32(perPage),
		OrderBy: orderBy,
		SortBy:  sortBy,
	}

	return result, nil
}

// APIError represents response body for API on error.
// e.q: Validation Error, Not Found Error, etc.
type APIError struct {
	*API
	Errors error `json:"errors,omitempty"`
}

// APIFailure represents body for API on failure. (e.g. Internal Server Error)
type APIFailure struct {
	*API
	causer error
}

// StatusCode returns status code.
func (a *API) StatusCode() int {
	return a.statusCode
}

// NewHTTPResponse creates a new HTTP Response.
func NewHTTPResponse(entity string) *API {
	return &API{
		Status: StatusAPISuccess,
	}
}

func formatState(status string) string {
	status = strings.Title(strings.ToLower(status))
	return status
}

// ============================= HANDLE SUCCESS RESPONSE ===================================
// Success setting response format for success state.
func (a *API) Success(data interface{}, code int, message string) *APISuccess {
	a.statusCode = code
	a.Status = StatusAPISuccess
	a.Message = message
	a.State = formatState(a.Status)
	return &APISuccess{
		API:  a,
		Data: data,
	}
}

// SuccessResponseJSON setting response for success condition
func SuccessResponseJSON(w http.ResponseWriter, statusCode int, data *APISuccess) error {
	// If there is nothing to marshal then set status code and return.
	if statusCode == http.StatusNoContent {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}

	// Encode the data to JSON.
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	// Set the content type and headers once we know marshaling has succeeded.
	w.Header().Set("Content-Type", "application/json")

	// Write the status code to the response.
	w.WriteHeader(statusCode)

	// Send the result back to the client.
	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}

// Success returns response format for success state.
func (a *API) SuccessJSON(w http.ResponseWriter, data interface{}, code int, message string) {
	a.statusCode = code
	a.Status = StatusAPISuccess
	a.Message = message
	a.State = formatState(a.Status)
	response := &APISuccess{
		API:  a,
		Data: data,
	}

	SuccessResponseJSON(w, a.statusCode, response)
}

// SuccessWithMeta returns response format for success state but with metadata.
func (a *API) SuccessWithMeta(w http.ResponseWriter, data interface{}, meta PaginationParams, code int, message string) {
	res := a.Success(data, code, message)
	res.Meta = meta

	SuccessResponseJSON(w, res.statusCode, res)
}

// SuccessWithoutData returns response format for success state without data.
func (a *API) SuccessWithoutData(w http.ResponseWriter, code int, message string) {
	a.statusCode = code
	a.Status = StatusAPISuccess
	a.Message = message
	a.State = formatState(a.Status)
	response := &APISuccess{
		API: a,
	}
	SuccessResponseJSON(w, a.statusCode, response)
}

// ============================= ======================= ===================================

// ============================== HANDLE ERROR RESPONSE ====================================
// Error returns response format for error state.
func (a *API) Error(code int, message string) *APIError {
	a.statusCode = code
	a.Status = StatusAPIError
	a.Message = message
	a.State = formatState(a.Status)
	return &APIError{
		API: a,
	}
}

// ErrorResponseJSON setting response for error condition
func ErrorResponseJSON(w http.ResponseWriter, statusCode int, data *APIError) error {
	// If there is nothing to marshal then set status code and return.
	if statusCode == http.StatusNoContent {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}

	// Encode the data to JSON.
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	// Set the content type and headers once we know marshaling has succeeded.
	w.Header().Set("Content-Type", "application/json")

	// Write the status code to the response.
	w.WriteHeader(statusCode)

	// Send the result back to the client.
	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}

// Error returns response format for error state.
func (a *API) ErrorJSON(w http.ResponseWriter, code int, message string) {
	a.statusCode = code
	a.Status = StatusAPIError
	a.Message = message
	a.State = formatState(a.Status)
	response := &APIError{
		API: a,
	}

	ErrorResponseJSON(w, a.statusCode, response)
}

// FieldErrors returns response format error.
func (a *API) FieldErrors(w http.ResponseWriter, err error, code int, message string) {
	fe := a.Error(code, message)
	fe.Errors = err

	ErrorResponseJSON(w, fe.statusCode, fe)
}

// ErrorWithStatusCode returns response format error.
func (a *API) ErrorWithStatusCode(w http.ResponseWriter, code int, message string) {
	a.statusCode = code
	a.Code = code
	a.Status = StatusAPIError
	a.Message = strings.Title(message)
	a.State = formatState(a.Status)
	response := &APIError{
		API: a,
	}

	ErrorResponseJSON(w, a.statusCode, response)
}

// ============================= ======================= ===================================

// ============================= HANDLE FAILURE RESPONSE ===================================
// Failure returns response format for failure state.
func (a *API) Failure(err error, code int) *APIFailure {
	a.statusCode = code
	a.Status = StatusAPIFailure
	a.Message = APIFailureMessage
	a.State = formatState(a.Status)
	return &APIFailure{
		API:    a,
		causer: nil,
	}
}

// FailureResponseJSON setting response for failure condition
func FailureResponseJSON(w http.ResponseWriter, statusCode int, data *APIFailure) error {
	// If there is nothing to marshal then set status code and return.
	if statusCode == http.StatusNoContent {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}

	// Encode the data to JSON.
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	// Set the content type and headers once we know marshaling has succeeded.
	w.Header().Set("Content-Type", "application/json")

	// Write the status code to the response.
	w.WriteHeader(statusCode)

	// Send the result back to the client.
	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}

// Failure returns response format for failure state.
func (a *API) FailureJSON(w http.ResponseWriter, err error, code int) {
	a.statusCode = code
	a.Status = StatusAPIFailure
	a.Message = APIFailureMessage
	a.State = formatState(a.Status)
	response := &APIFailure{
		API:    a,
		causer: nil,
	}

	FailureResponseJSON(w, a.statusCode, response)
}

// ============================= ======================= ===================================

// Error implements error interface.
func (f *APIFailure) Error() string {
	b, err := json.Marshal(f) // {"", ""}
	if err != nil {
		return err.Error()
	}
	return string(b)
}

// Causer returns error causer.
// The Causer error is needed for logging.
func (f *APIFailure) Causer() error {
	return f.causer
}

// Error implement error interface.
func (e *APIError) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return err.Error()
	}

	return string(b)
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
