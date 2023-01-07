package repository

import (
	"os"
	"strconv"
	"time"

	"github.com/enylvia/shorten-link/app"
	"github.com/go-redis/redis/v9"
)

// Create Contract for User Repository
type ShortenRepository interface {
	ShortenUrl(id string, url string, expiry time.Duration, dbNo int) (string, error)
	ResolveUrl(url string, dbNo int) (string, error)
	CheckLimiter(ip string, dbNo int) int
	DecrementQuota(ip string, dbNo int)
	CheckShorten(id string, dbNo int) (string, error)
	SetLimiter(ip string, dbNo int)
}

// Create User Repository and Inject GORM DB
type ShortenRepositoryImplement struct {
}

// Create New Implementation of User Repository
func NewShortenRepository() ShortenRepository {
	return &ShortenRepositoryImplement{}
}

func (r *ShortenRepositoryImplement) CheckLimiter(ip string, dbNo int) int {
	con := app.CreateClient(dbNo)
	defer con.Close()

	value, err := con.Get(app.Ctx, ip).Result()
	if err == redis.Nil {
		return 0
	} else if err != nil {
		return 0
	}
	quota, _ := strconv.Atoi(value)

	return quota
}

func (r *ShortenRepositoryImplement) SetLimiter(ip string, dbNo int) {
	con := app.CreateClient(dbNo)
	defer con.Close()

	con.Set(app.Ctx, ip, os.Getenv("API_QUOTA"), 30*time.Minute).Err()
}

func (r *ShortenRepositoryImplement) DecrementQuota(ip string, dbNo int) {
	con := app.CreateClient(dbNo)
	defer con.Close()

	con.Decr(app.Ctx, ip).Result()
}

func (r *ShortenRepositoryImplement) CheckShorten(id string, dbNo int) (string, error) {
	con := app.CreateClient(dbNo)
	defer con.Close()

	value, err := con.Get(app.Ctx, id).Result()
	if err == redis.Nil {
		return "null", err
	} else if err != nil {
		return "error to make a connnection", err
	}
	return value, nil

}
func (r *ShortenRepositoryImplement) ShortenUrl(id string, url string, expiry time.Duration, dbNo int) (string, error) {
	// do something here
	con := app.CreateClient(dbNo)
	defer con.Close()

	err := con.Set(app.Ctx, id, url, expiry).Err()
	if err != nil {
		return "error to make a connnection", err
	}
	return id, nil
}

func (r *ShortenRepositoryImplement) ResolveUrl(url string, dbNo int) (string, error) {
	// do something here
	con := app.CreateClient(dbNo)
	defer con.Close()

	value, err := con.Get(app.Ctx, url).Result()
	if err == redis.Nil {
		return "Shorten not found", err
	} else if err != nil {
		return "error to make a connnection", err
	}
	return value, nil
}
