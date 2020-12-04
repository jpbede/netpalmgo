package transport

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// CheckForHTTPError does some generic response checks and raises a error when it detects some
func CheckForHTTPError(response *http.Response) error {
	if response.StatusCode >= 300 {
		read, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}

		var errArr map[string]string
		if err := json.Unmarshal(read, &errArr); err == nil {
			return errors.New(errArr["detail"])
		}
		return errors.New(string(read))
	}

	return nil
}
