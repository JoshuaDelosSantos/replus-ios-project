package auth

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testLogger *log.Logger

func init() {
	testLogger = log.New(os.Stdout, "[MIDDLEWARE_TEST] ", log.LstdFlags|log.Lshortfile)
}

func TestAuthMiddleware(t *testing.T) {
	// Use the mock validator for testing
	mockValidator := &MockValidator{}

	testLogger.Println("Starting middleware test")

	// Setup test handler
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user_id")
		if userID == nil {
			http.Error(w, "Missing user ID", http.StatusUnauthorized)
			return
		}
		assert.Equal(t, 1, userID) // Ensure that userID from context matches
		w.WriteHeader(http.StatusOK)
	})

	// Pass the mock validator into the AuthMiddleware
	middleware := AuthMiddleware(mockValidator, testHandler)

	// Test with valid token
	testLogger.Println("Testing with valid token")
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer validtoken")
	rr := httptest.NewRecorder()
	middleware.ServeHTTP(rr, req)
	testLogger.Printf("Valid token test result: %d", rr.Code)
	assert.Equal(t, http.StatusOK, rr.Code)

	// Test with missing token
	testLogger.Println("Testing with missing token")
	req = httptest.NewRequest(http.MethodGet, "/", nil)
	rr = httptest.NewRecorder()
	middleware.ServeHTTP(rr, req)
	testLogger.Printf("Missing token test result: %d", rr.Code)
	assert.Equal(t, http.StatusUnauthorized, rr.Code)

	// Test with invalid token
	testLogger.Println("Testing with invalid token")
	req = httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer invalidtoken")
	rr = httptest.NewRecorder()
	middleware.ServeHTTP(rr, req)
	testLogger.Printf("Invalid token test result: %d", rr.Code)
	assert.Equal(t, http.StatusUnauthorized, rr.Code)

	testLogger.Println("Middleware tests completed")
}