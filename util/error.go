package util

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"net/http"
)

// CheckForHTTPError does some generic response checks and raises a error when it detects some
func CheckForHTTPError(response *resty.Response) error {
	// check if authentication error
	if response.StatusCode() == http.StatusForbidden {
		var errArr map[string]string
		json.Unmarshal(response.Body(), &errArr)
		return errors.New(errArr["detail"])
	}
	return nil
}
