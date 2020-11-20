package transport

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// CheckForHTTPError does some generic response checks and raises a error when it detects some
func CheckForHTTPError(response *http.Response) error {
	// check if authentication error
	if response.StatusCode == http.StatusForbidden {
		var errArr map[string]string
		read, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		json.Unmarshal(read, &errArr)
		return errors.New(errArr["detail"])
	}
	// invalid request
	if response.StatusCode == http.StatusUnprocessableEntity {
		read, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		return errors.New(string(read))
	}
	return nil
}
