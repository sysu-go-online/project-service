package controller

import "net/http"

// ErrorHandler is error handler for http
type ErrorHandler func(w http.ResponseWriter, r *http.Request) error
