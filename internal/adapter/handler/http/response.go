package http

import (
	"api/internal/core/domain"
	"encoding/json"
	"github.com/go-chi/render"
	"net/http"
)

type response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

func newResponse(success bool, message string, data any) response {
	return response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

func handleSuccess(w http.ResponseWriter, data any) {
	rsp := newResponse(true, "success", data)
	jsonRsp, _ := json.Marshal(rsp)

	_, _ = w.Write(jsonRsp)
}

type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText string `json:"status"`
	AppCode    int64  `json:"code,omitempty"`
	ErrorText  string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(_ http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func handleErrBadRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     "invalid request",
		ErrorText:      err.Error(),
	}
}

var errorStatusMap = map[error]int{
	domain.ErrInternal:        http.StatusInternalServerError,
	domain.ErrConflictingData: http.StatusConflict,
}

func handleError(text string, err error) render.Renderer {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: statusCode,
		StatusText:     text,
		ErrorText:      err.Error(),
	}
}
