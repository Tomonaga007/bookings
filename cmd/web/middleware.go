package main

import "net/http"

//func WriteToConsole(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		log.Println("from middleware")
//		next.ServeHTTP(w,r)
//	})
//}

func SessionLoad(next http.Handler) http.Handler{
	return session.LoadAndSave(next)
}