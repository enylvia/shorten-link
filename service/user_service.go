package service

import (
	"fmt"
	"os"
	"time"

	"github.com/enylvia/shorten-link/app"
	"github.com/enylvia/shorten-link/model/web"
	"github.com/enylvia/shorten-link/repository"
	"github.com/gofrs/uuid"
)

type ShortenService interface {
	Shorten(data web.URLRequest, ip string) (web.URLResponse, error)
	Resolve(url string) (web.URLResponse, error)
}

// Create User Service and Inject User Repository and Redis DB
type ShortenServiceImplement struct {
	shortenRepository repository.ShortenRepository
}

// implementation of UserService
func NewShortenService(shortenRepository repository.ShortenRepository) ShortenService {
	return &ShortenServiceImplement{shortenRepository}
}

func (s *ShortenServiceImplement) Shorten(data web.URLRequest, ip string) (web.URLResponse, error) {
	// do something here
	// check if custom short is not empty or empty
	if data.CustomShort == "" {
		data.CustomShort = uuid.Must(uuid.NewV4()).String()
	}

	// check if custom short is already exist
	val, _ := s.shortenRepository.CheckShorten(data.CustomShort, 0)

	fmt.Println(val)
	if val != "null" {
		return web.URLResponse{
			URL: "Custom short is already exist",
		}, nil
	}

	// check expiry time
	if data.ExpiryDate == 0 {
		data.ExpiryDate = 24
	}

	// set limiter for ip
	limit := s.shortenRepository.CheckLimiter(ip, 1)

	if limit <= 0 {
		s.shortenRepository.SetLimiter(ip, 1)
	}
	exec, err := s.shortenRepository.ShortenUrl(data.CustomShort, data.URL, data.ExpiryDate*time.Hour, 0)
	if err != nil {
		return web.URLResponse{
			URL: "Failed to shorten url",
		}, err
	}
	// decrement quota
	s.shortenRepository.DecrementQuota(ip, 1)

	limit = s.shortenRepository.CheckLimiter(ip, 1)

	var response web.URLResponse
	response.URL = os.Getenv("APP_URL") + exec
	response.CustomShort = data.CustomShort
	response.ExpiryDate = data.ExpiryDate * time.Hour
	response.XRateRemaining = limit
	response.XrateLimitRest = 30 * time.Minute

	return response, nil
}

func (s *ShortenServiceImplement) Resolve(url string) (web.URLResponse, error) {
	// do something here

	response := web.URLResponse{}
	data, err := s.shortenRepository.ResolveUrl(url, 0)
	if err != nil {
		return web.URLResponse{
			URL: "Failed to resolve url",
		}, err
	}

	rInr := app.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(app.Ctx, "counter")

	response.URL = data
	return response, nil
}
