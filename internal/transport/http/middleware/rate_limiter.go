package middleware

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

type RateLimiter struct {
	requests map[string][]time.Time
	mutex    sync.RWMutex
	limit    int
	window   time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}
}

func (rl *RateLimiter) RateLimit() fiber.Handler {
	return func(c *fiber.Ctx) error {
		clientIP := c.IP()

		rl.mutex.Lock()
		defer rl.mutex.Unlock()

		now := time.Now()
		windowStart := now.Add(-rl.window)

		var validRequests []time.Time
		for _, reqTime := range rl.requests[clientIP] {
			if reqTime.After(windowStart) {
				validRequests = append(validRequests, reqTime)
			}
		}

		if len(validRequests) >= rl.limit {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error":   "rate_limit_exceeded",
				"message": "Too many requests, please try again later",
			})
		}

		validRequests = append(validRequests, now)
		rl.requests[clientIP] = validRequests

		return c.Next()
	}
}

func (rl *RateLimiter) Cleanup() {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	now := time.Now()
	windowStart := now.Add(-rl.window)

	for clientIP, requests := range rl.requests {
		var validRequests []time.Time
		for _, reqTime := range requests {
			if reqTime.After(windowStart) {
				validRequests = append(validRequests, reqTime)
			}
		}

		if len(validRequests) == 0 {
			delete(rl.requests, clientIP)
		} else {
			rl.requests[clientIP] = validRequests
		}
	}
}
