package internal

import (
	"check_primes/pkg"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CheckPrimes(e echo.Context) error {
	numbers := requestStruct{}
	err := json.NewDecoder(e.Request().Body).Decode(&numbers)
	defer e.Request().Body.Close()
	if err != nil {
		e.Response().WriteHeader(http.StatusBadRequest)
		e.Response().Write([]byte("invalid request number"))
	}
	res := pkg.CheckPrimes(numbers.Numbers)
	fmt.Println(res)
	//res,err = json.Marshal(res)
	e.Response().WriteHeader(http.StatusOK)
	//e.Response().Write([]byte())
	return nil
}
