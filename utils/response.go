package utils

import (
	"encoding/json"
	"math"
	"net/http"
)

type response map[string]interface{}

// RespondwithJSON write json response format
func RespondwithJSON(w http.ResponseWriter, code int, payload interface{}, success bool, args ...string) {
	var resmsg = response{}

	if payload != nil {
		if len(args) == 0 {
			resmsg["data"] = payload
		} else {
			typ := args[0]
			resmsg[typ] = []interface{}{payload}
		}
	}
	resmsg["success"] = success

	res, _ := json.Marshal(resmsg)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(res)
}

// RespondwithPagination : with pagination
func RespondwithPagination(w http.ResponseWriter, payload interface{}, count int, limit int, offset int) {
	var resp = response{}
	// calculating total Page
	totalPage := math.Ceil(float64(count) / float64(limit))
	// next page
	nextPage, hasNext := func() (int, bool) {
		if int(totalPage)-offset > 0 {
			next := offset + 1
			return next, true
		}
		return 0, false
	}()
	// previous page
	previousPage, hasPre := func() (int, bool) {
		if int(totalPage)-offset >= 0 && offset-1 > 0 {
			pre := offset - 1
			return pre, true
		}
		return 0, false
	}()

	if payload != nil {
		if hasNext {
			resp["next"] = nextPage
		} else {
			resp["next"] = nil
		}
		if hasPre {
			resp["previous"] = previousPage
		} else {
			resp["previous"] = nil
		}
		resp["data"] = payload
		resp["page_size"] = limit
		resp["success"] = true
	}
	res, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(res)
}

// RespondWithError : return error message
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondwithJSON(w, code, map[string]string{"message": msg}, false, "errors")
}

// RespondWithValidationError return error message
func RespondWithValidationError(w http.ResponseWriter, code int, i interface{}) {
	RespondwithJSON(w, code, i, false, "errors")
}
