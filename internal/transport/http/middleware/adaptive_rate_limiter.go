package middleware

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AdaptiveRateLimiter struct {
	requests map[string][]time.Time
	mutex    sync.RWMutex
	limits   map[string]int
	window   time.Duration
}

func NewAdaptiveRateLimiter(window time.Duration) *AdaptiveRateLimiter {
	return &AdaptiveRateLimiter{
		requests: make(map[string][]time.Time),
		limits: map[string]int{
			"auth":      5,
			"api":       60,
			"websocket": 100,
		},
		window: window,
	}
}

func (arl *AdaptiveRateLimiter) RateLimit() fiber.Handler {
	return func(c *fiber.Ctx) error {
		clientIP := c.IP()
		path := c.Path()

		requestType := arl.getRequestType(path)
		limit := arl.limits[requestType]

		key := clientIP + ":" + requestType

		arl.mutex.Lock()
		defer arl.mutex.Unlock()

		now := time.Now()
		windowStart := now.Add(-arl.window)

		var validRequests []time.Time
		for _, reqTime := range arl.requests[key] {
			if reqTime.After(windowStart) {
				validRequests = append(validRequests, reqTime)
			}
		}

		if len(validRequests) >= limit {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error":       "rate_limit_exceeded",
				"message":     "Too many requests, please try again later",
				"retry_after": arl.window.Seconds(),
			})
		}

		validRequests = append(validRequests, now)
		arl.requests[key] = validRequests

		c.Set("X-RateLimit-Limit", fmt.Sprintf("%d", limit))
		c.Set("X-RateLimit-Remaining", fmt.Sprintf("%d", limit-len(validRequests)))
		c.Set("X-RateLimit-Reset", fmt.Sprintf("%d", now.Add(arl.window).Unix()))

		return c.Next()
	}
}

func (arl *AdaptiveRateLimiter) getRequestType(path string) string {
	if strings.HasPrefix(path, "/auth/") {
		return "auth"
	}
	if strings.HasPrefix(path, "/ws") {
		return "websocket"
	}
	return "api"
}

func (arl *AdaptiveRateLimiter) SetLimit(requestType string, limit int) {
	arl.mutex.Lock()
	defer arl.mutex.Unlock()
	arl.limits[requestType] = limit
}

func (arl *AdaptiveRateLimiter) Cleanup() {
	arl.mutex.Lock()
	defer arl.mutex.Unlock()

	now := time.Now()
	windowStart := now.Add(-arl.window)

	for key, requests := range arl.requests {
		var validRequests []time.Time
		for _, reqTime := range requests {
			if reqTime.After(windowStart) {
				validRequests = append(validRequests, reqTime)
			}
		}

		if len(validRequests) == 0 {
			delete(arl.requests, key)
		} else {
			arl.requests[key] = validRequests
		}
	}
}
