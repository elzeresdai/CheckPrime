package internal

import (
	"check_primes/pkg"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

// CheckPrimes main function, returning result in json
// returns res to user
func CheckPrimes(e echo.Context) error {
	resp, err := getIntegers(e)
	if err != nil {
		e.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		e.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(e.Response()).Encode(err.Error())
	}
	e.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	e.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(e.Response()).Encode(resp)

}

// func works with request data, checking is data valid or not
// and returning result
func getIntegers(e echo.Context) ([]string, error) {

	body, err := io.ReadAll(e.Request().Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(e.Request().Body)
	if err != nil {
		return nil, err
	}
	numbers := requestStruct{}
	if err = json.Unmarshal(body, &numbers); err != nil {
		return nil, err
	}

	var list []int

	for i, num := range numbers.Numbers {
		if i = isIntNumber(num); i != 0 {
			list = append(list, i)
			continue
		} else {
			return nil, fmt.Errorf("the given message is invalid in index: %+v", i+1)
		}

	}

	result := pkg.CheckPrimes(list)

	return result, nil
}

// func for checking int values
func isIntNumber(n any) int {
	switch n.(type) {
	case int:
		return n.(int)
	case float64:
		if n == float64(int(n.(float64))) {
			return int(n.(float64))
		} else {
			return 0
		}
	default:
		return 0
	}
}
