package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type Service interface {
	Daysleft() int64
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {
	d := e.s.Daysleft()

	s := fmt.Sprintf("Days: %d", d)

	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}

	return nil
}
