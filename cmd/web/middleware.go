package main

import (
	"github.com/justinas/nosurf"
	"net/http"

	)

//func WriteToConsole(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		log.Println("from middleware")
//		next.ServeHTTP(w,r)
//	})
//}

func SessionLoad(next http.Handler) http.Handler{
	return session.LoadAndSave(next)
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: false,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}