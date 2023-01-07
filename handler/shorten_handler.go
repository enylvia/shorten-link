package handler

import (
	"fmt"

	"github.com/enylvia/shorten-link/model/web"
	"github.com/enylvia/shorten-link/service"
	"github.com/enylvia/shorten-link/utils"
	"github.com/labstack/echo/v4"
)

type ShortenHandler interface {
	ShortenURL(echo.Context) error
	ResolveURL(echo.Context) error
}

// implement ShortenHandler interface here
type ShortenHandlerImplement struct {
	shortenService service.ShortenService
}

// create new instance of ShortenHandler
func NewShortenHandler(shortenService service.ShortenService) ShortenHandler {
	return &ShortenHandlerImplement{shortenService}
}

func (h *ShortenHandlerImplement) ShortenURL(c echo.Context) error {
	//  Get Request From Body
	data := web.URLRequest{}
	err := c.Bind(&data)
	if err != nil {
		formatter := utils.ValidateRequest(data)
		return c.JSON(403, formatter)
	}
	// check the IP and set rate limiter
	exec, err := h.shortenService.Shorten(data, c.RealIP())
	fmt.Println(c.RealIP())
	if err != nil {
		formatter := utils.BadRequestResponse(err.Error(), "Failed to shorten url")
		return c.JSON(403, formatter)
	}
	formatter := utils.SuccessResponse(exec)
	return c.JSON(200, formatter)

}

func (h *ShortenHandlerImplement) ResolveURL(c echo.Context) error {
	url := c.Param("url")
	resolve, err := h.shortenService.Resolve(url)
	if err != nil {
		formatter := utils.BadRequestResponse(err.Error(), "Failed to resolve url")
		return c.JSON(403, formatter)
	}
	return c.Redirect(301, resolve.URL)
}
