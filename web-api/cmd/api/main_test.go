// web-api/cmd/api/main_test.go
package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestHandlers(t *testing.T) {
    tests := []struct {
        name           string
        handler        http.HandlerFunc
        expectedStatus int
        expectedBody   string
        shouldContain  bool
    }{
        {
            name:           "Home Handler",
            handler:        home,
            expectedStatus: http.StatusOK,
            expectedBody:   "Welcome to the Shapes API",
            shouldContain:  false,
        },
        {
            name:           "Health Handler",
            handler:        health,
            expectedStatus: http.StatusOK,
            expectedBody:   "Server is running",
            shouldContain:  false,
        },
        {
            name:           "About Handler",
            handler:        about,
            expectedStatus: http.StatusOK,
            expectedBody:   "John Doe", // CHANGE TO YOUR NAME
            shouldContain:  false,
        },
        {
            name:           "Time Handler",
            handler:        timeHandler,
            expectedStatus: http.StatusOK,
            expectedBody:   "",
            shouldContain:  false,
        },
        {
            name:           "Random Handler",
            handler:        random,
            expectedStatus: http.StatusOK,
            expectedBody:   "Random number between 1 and 100:",
            shouldContain:  true,
        },
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            req := httptest.NewRequest("GET", "/", nil)
            rr := httptest.NewRecorder()
            
            tc.handler(rr, req)
            
            if rr.Code != tc.expectedStatus {
                t.Errorf("%s: wrong status - got %v, want %v",
                    tc.name, rr.Code, tc.expectedStatus)
            }
            
            body := rr.Body.String()
            
            if tc.name == "Time Handler" {
                if body == "" {
                    t.Errorf("%s: expected time, got empty", tc.name)
                }
            } else if tc.shouldContain {
                if !strings.Contains(body, tc.expectedBody) {
                    t.Errorf("%s: body doesn't contain %q - got %q",
                        tc.name, tc.expectedBody, body)
                }
            } else {
                if body != tc.expectedBody {
                    t.Errorf("%s: wrong body - got %q, want %q",
                        tc.name, body, tc.expectedBody)
                }
            }
        })
    }
}