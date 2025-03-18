package controller

import (
	"FlexcityTest/domain"
	"errors"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	errorLog = log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.LUTC|log.Lshortfile)
	validate = validator.New(validator.WithRequiredStructEnabled())
)

func init() {
	if err := validate.RegisterValidation("notBeforeNow", notBeforeNow); err != nil {
		log.Fatalf("Failed to register validator notBeforeNow: %s", err)
	}
}

func handleError(writer http.ResponseWriter, request *http.Request, err error) {
	if err == nil {
		return
	}

	var errResp domain.ErrorResponse
	if ok := errors.As(err, &errResp); !ok {
		errResp = domain.ErrorResponse{
			NativeError: err,
			Type:        domain.ErrInternal,
		}
	}

	if errResp.NativeError != nil {
		errorLog.Printf("%s: %s", errResp, errResp.NativeError)
	} else {
		errorLog.Println(errResp)
	}

	writer.WriteHeader(errResp.StatusCode())
	render.JSON(writer, request, errResp)
}

func validatePayload(payload any) error {
	if payload == nil {
		return domain.ErrorResponse{
			Type: domain.ErrInvalidPayload,
		}
	}

	if err := validate.Struct(payload); err != nil {
		return domain.ErrorResponse{
			NativeError: err,
			Type:        domain.ErrInvalidPayload,
		}
	}

	return nil
}

func notBeforeNow(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if !ok {
		return false
	}
	return !date.Before(time.Now())
}
