package main

import (
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v9"

	dummy_handler "github.com/0Delta/echo_srv/internal/Handler/dummy"
	usecase_handler "github.com/0Delta/echo_srv/internal/Usecase/Handler"
)

type Env struct {
	Port string `envconfig:"PORT"`
}

// validator
type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

// main
func main() {
	var goenv Env
	envconfig.Process("", &goenv)

	e := echo.New()
	e.Validator = &Validator{validator: validator.New()}

	e.GET("/", GetHandler(dummy_handler.DummyHandler{})) // TODO: change subhandler
	// e.POST("/", Handler(c))

	e.Logger.Fatal(e.Start(":" + goenv.Port))
}

func GetHandler(h usecase_handler.Handler) func(echo.Context) error {
	return func(c echo.Context) error {
		var data struct {
			test string
		}
		if err := c.Bind(&data); err != nil {
			return c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		}
		if err := c.Validate(&data); err != nil {
			return c.String(http.StatusBadRequest, "Validate is failed: "+err.Error())
		}
		resp, err := h.Run(data)
		if err != nil {
			return c.String(http.StatusBadRequest, "Logic failed: "+err.Error())
		}
		return c.String(http.StatusOK, resp)
	}
}
