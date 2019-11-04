// package main

// import (
// 	"log"
// 	"net/http"
// 	"net/http/httputil"
// )

// func main() {
// 	director := func(req *http.Request) {
// 		req.URL.Scheme = "http"
// 		req.URL.Host = ":9001"
// 	}
// 	// modifier := func(res *http.Response) error {
// 	// 	return nil
// 	// }
// 	reverseProxy := &httputil.ReverseProxy{
// 		Director: director,
// 		// ModifyResponse: modifier,
// 	}

// 	server := http.Server{
// 		Addr:    ":9000",
// 		Handler: reverseProxy,
// 	}
// 	if err := server.ListenAndServe(); err != nil {
// 		log.Fatal(err.Error())
// 	}

// }
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

type contextKey string

const tokenContextKey contextKey = "requestTime"

func setToken(parents context.Context, t time.Time) context.Context {
	return context.WithValue(parents, tokenContextKey, t)
}
func getToken(ctx context.Context) (time.Time, error) {
	v := ctx.Value(tokenContextKey)
	token, ok := v.(time.Time)
	if !ok {
		return time.Now(), fmt.Errorf("token not found")
	}
	return token, nil
}

func main() {
	director := func(request *http.Request) {
		fmt.Println("director")
		// ctx := setToken(request.Context(), time.Now())
		// request = request.WithContext(ctx)
		request.URL.Scheme = "http"
		request.URL.Host = ":8080"
		fmt.Println(request.URL.Scheme)
	}
	modifier := func(res *http.Response) error {
		fmt.Println("modifier")
		startTime, err := getToken(res.Request.Context())
		if err != nil {
			fmt.Println("err:", err)
		}
		fmt.Println(startTime)
		return nil
	}

	rp := &httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: modifier,
	}

	if err := http.ListenAndServe(":9000", rp); err != nil {
		log.Fatal(err.Error())
	}
}
