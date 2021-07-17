package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/justinas/nosurf"
)


func WriteToConsol(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000"),
		r.RemoteAddr,
		 "Hit the page",  
		 r.Host, 
		 r.URL, 
		 r.Method, 
		 r.Response,)
		next.ServeHTTP(w, r)
	})
}


func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
			HttpOnly: true,
			Path: "/",
			Secure: app.InProduction,
			SameSite: http.SameSiteLaxMode,
		})
		return csrfHandler
}

func SessionLoad(next http.Handler) http.Handler{
	return session.LoadAndSave(next)
}