# Micro Kernel
## Class Diagram
```mermaid
classDiagram
    class Kernel {
        -middlewares: []Middleware
        -handler: http.HandlerFunc
        +Use(m: Middleware)
        +ServeHTTP(w: http.ResponseWriter, r: *http.Request)
    }

    class Middleware {
        <<interface>>
        +func(http.HandlerFunc) http.HandlerFunc
    }

    class AuthenticationMiddleware {
        +func(next: http.HandlerFunc) http.HandlerFunc
    }

    class AuthorizationMiddleware {
        +func(next: http.HandlerFunc) http.HandlerFunc
    }

    class CORSMiddleware {
        +func(next: http.HandlerFunc) http.HandlerFunc
    }

    class LoggingMiddleware {
        +func(next: http.HandlerFunc) http.HandlerFunc
    }

    class RateLimitMiddleware {
        +func(rps: int) Middleware
    }

    class CircuitBreakerMiddleware {
        +func(threshold: int, timeout: time.Duration) Middleware
    }

    Kernel --> "0..*" Middleware : uses
    AuthenticationMiddleware ..|> Middleware
    AuthorizationMiddleware ..|> Middleware
    CORSMiddleware ..|> Middleware
    LoggingMiddleware ..|> Middleware
    RateLimitMiddleware ..|> Middleware
    CircuitBreakerMiddleware ..|> Middleware

```
## Flow Diagram

```mermaid
graph TD
    A[HTTP Request] --> B[Kernel]
    B --> C{Circuit Breaker}
    C -->|Open| D[Return 503 Service Unavailable]
    C -->|Closed| E[Rate Limiter]
    E -->|Limit Exceeded| F[Return 429 Too Many Requests]
    E -->|Within Limit| G[Authentication]
    G -->|Invalid Token| H[Return 401 Unauthorized]
    G -->|Valid Token| I[Authorization]
    I -->|Not Admin| J[Return 403 Forbidden]
    I -->|Is Admin| K[CORS]
    K -->|OPTIONS Request| L[Return 200 OK with CORS headers]
    K -->|Non-OPTIONS Request| M[Logging Start]
    M --> N[Main Handler]
    N --> O[Logging End]
    O --> P[HTTP Response]
    
    subgraph "Kernel"
    B
    end
    
    subgraph "Middleware Stack"
    C
    E
    G
    I
    K
    M
    O
    end
    
    subgraph "Main Handler"
    N
    end
```
## Code
```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// Kernel represents the core of the microkernel architecture
type Kernel struct {
	middlewares []Middleware
	handler     http.HandlerFunc
}

// Middleware is a higher-order function that wraps an http.HandlerFunc
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Use adds a middleware to the kernel
func (k *Kernel) Use(m Middleware) {
	k.middlewares = append(k.middlewares, m)
}

// ServeHTTP implements the http.Handler interface
func (k *Kernel) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler := k.handler
	for i := len(k.middlewares) - 1; i >= 0; i-- {
		handler = k.middlewares[i](handler)
	}
	handler(w, r)
}

// Authentication middleware
func AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "valid-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

// Authorization middleware
func AuthorizationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role := r.Header.Get("Role")
		if role != "admin" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}

// CORS middleware
func CORSMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

// Logging middleware
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	}
}

// RateLimit middleware
func RateLimitMiddleware(rps int) Middleware {
	limiter := rate.NewLimiter(rate.Limit(rps), rps)
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if !limiter.Allow() {
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}
			next(w, r)
		}
	}
}

// CircuitBreaker middleware
func CircuitBreakerMiddleware(threshold int, timeout time.Duration) Middleware {
	var (
		failures  int
		lastRetry time.Time
		mu        sync.Mutex
	)

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			mu.Lock()
			if failures >= threshold && time.Since(lastRetry) < timeout {
				mu.Unlock()
				http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
				return
			}
			mu.Unlock()

			ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
			defer cancel()

			done := make(chan bool)
			go func() {
				next(w, r.WithContext(ctx))
				done <- true
			}()

			select {
			case <-done:
				mu.Lock()
				failures = 0
				mu.Unlock()
			case <-ctx.Done():
				mu.Lock()
				failures++
				if failures == threshold {
					lastRetry = time.Now()
				}
				mu.Unlock()
				http.Error(w, "Request Timeout", http.StatusRequestTimeout)
			}
		}
	}
}

func main() {
	kernel := &Kernel{
		handler: func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello from the microkernel!")
		},
	}

	kernel.Use(LoggingMiddleware)
	kernel.Use(CORSMiddleware)
	kernel.Use(AuthenticationMiddleware)
	kernel.Use(AuthorizationMiddleware)
	kernel.Use(RateLimitMiddleware(10))
	kernel.Use(CircuitBreakerMiddleware(5, 30*time.Second))

	http.ListenAndServe(":8080", kernel)
}

```