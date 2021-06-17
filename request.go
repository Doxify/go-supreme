package gosupreme

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Supreme) makeRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	// TODO: Add more headers here...

	res, err := s.c.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("http request was unsuccessful, recieved status: %d", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(&v)
	if err != nil {
		return err
	}

	return nil
}
