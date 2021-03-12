package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		for key, value := range r.Form {
			fmt.Printf("%s = %s\n\n", key, value)
		}

		fmt.Printf("event => %s\n", r.FormValue("ev"))
		fmt.Printf("event_type => %s\n", r.FormValue("et"))
		fmt.Printf("app_id => %s\n", r.FormValue("id"))
		fmt.Printf("user_id => %s\n", r.FormValue("uid"))
		fmt.Printf("message_id => %s\n", r.FormValue("mid"))
		fmt.Printf("page_title => %s\n", r.FormValue("t"))
		fmt.Printf("page_url => %s\n", r.FormValue("p"))
		fmt.Printf("browser_language => %s\n", r.FormValue("l"))
		fmt.Printf("screen_size => %s\n", r.FormValue("sc"))

		fmt.Printf("attributes => %s\n", r.FormValue("atrk1"))
		fmt.Printf("value => %s\n", r.FormValue("atrv1"))
		fmt.Printf("type => %s\n", r.FormValue("atrt1"))
		fmt.Printf("attributes => %s\n", r.FormValue("atrk2"))
		fmt.Printf("value => %s\n", r.FormValue("atrv2"))
		fmt.Printf("type => %s\n", r.FormValue("atrt2"))

		fmt.Printf("traits => %s\n", r.FormValue("uatrk1"))
		fmt.Printf("value => %s\n", r.FormValue("uatrv1"))
		fmt.Printf("type => %s\n", r.FormValue("uatrt1"))
		fmt.Printf("traits => %s\n", r.FormValue("uatrk2"))
		fmt.Printf("value => %s\n", r.FormValue("uatrv2"))
		fmt.Printf("type => %s\n", r.FormValue("uatrt2"))
		fmt.Printf("traits => %s\n", r.FormValue("uatrk3"))
		fmt.Printf("value => %s\n", r.FormValue("uatrv3"))
		fmt.Printf("type => %s\n", r.FormValue("uatrv3"))

		fmt.Fprintln(w, "HTTP Request, Response - Succeed")
	})

	http.ListenAndServe(":8080", mux)
}
