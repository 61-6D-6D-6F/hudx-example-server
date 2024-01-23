package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("Hudx example server is running.")

	responseHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<p class=\"response\">server response</p>")
	}

	response2Handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<p class=\"response\">server response 2</p>")
	}

	response3Handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<p class=\"response\">server response 3</p>")
	}

	noResponseHandler := func(w http.ResponseWriter, r *http.Request) {}

	addContactHandler := func(w http.ResponseWriter, r *http.Request) {
		fullName := r.PostFormValue("fullName")
		email := r.PostFormValue("email")
		fmt.Fprintf(w, "<p class=\"response\">%s - %s</p>", fullName, email)
	}

	addCustomerHandler := func(w http.ResponseWriter, r *http.Request) {
		firstName := r.PostFormValue("firstName")
		lastName := r.PostFormValue("lastName")
		phoneNumber := r.PostFormValue("phoneNumber")
		email := r.PostFormValue("email")
		fmt.Fprintf(w, "<div><h3>New customer:</h3><p class=\"response\">%s %s : %s - %s</p></div>",
			firstName, lastName, phoneNumber, email)
	}

	slowResponseHandler := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1200 * time.Millisecond)
		fmt.Fprintf(w, "<p class=\"response\">server response</p>")
	}

	slowResponse2Handler := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2400 * time.Millisecond)
		fmt.Fprintf(w, "<p class=\"response\">server response 2</p>")
	}

	notFoundHandler := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1200 * time.Millisecond)
		w.WriteHeader(404)
	}

	slowNotFoundHandler := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2400 * time.Millisecond)
		w.WriteHeader(404)
	}

	imageHandler := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		http.ServeFile(w, r, "example/image.png")
	}

	customerName := "John James"
	customerEmail := "johnjames@email.com"

	customerUpdateHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			customerName = r.PostFormValue("name")
			customerEmail = r.PostFormValue("email")
		}
		fmt.Fprintf(w,
			`<div id="target14">
                <label for="name"><b>Name:</b></label>
                <span>%s</span>
                <br>
                <label for="email"><b>Email:</b></label>
                <span>%s</span>
                <br>
                <br>
                <div hudx hdx="click GET /customer/1/edit #target14 outerHTML">
                    Edit
                </div>
            </div>`, customerName, customerEmail)
	}

	customerEditHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,
			`<div id="target14">
                <form id="myForm7">
                    <label for="name"><b>Name:</b></label>
                    <input type="text" name="name" value="%s"/>
                    <br>
                    <label for="email"><b>Email:</b></label>
                    <input type="text" name="email" value="%s"/>
                </form>
                <br>
                <div hudx hdx="click POST /customer/1 #target14 outerHTML #myForm7">
                    Save
                </div>
                <div hudx hdx="click GET /customer/1 #target14 outerHTML">
                    Cancel
                </div>
            </div>`, customerName, customerEmail)
	}

	modalHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,
			`<div class="panel">
                <p>Press the close button.</p>
                <div hudx
                    hdsx="click #target15 remove showModal"
                    hdsx-2="click #target15 add hideModal"
                    >Close
                </div>
            </div>`)
	}

	http.HandleFunc("/image.png", imageHandler)
	http.Handle("/", http.FileServer(http.Dir("./example")))

	http.HandleFunc("/response", responseHandler)
	http.HandleFunc("/response2", response2Handler)
	http.HandleFunc("/response3", response3Handler)
	http.HandleFunc("/no-response", noResponseHandler)
	http.HandleFunc("/add-contact", addContactHandler)
	http.HandleFunc("/add-customer", addCustomerHandler)
	http.HandleFunc("/slow-response", slowResponseHandler)
	http.HandleFunc("/slow-response2", slowResponse2Handler)
	http.HandleFunc("/not-found", notFoundHandler)
	http.HandleFunc("/slow-not-found", slowNotFoundHandler)
	http.HandleFunc("/customer/1", customerUpdateHandler)
	http.HandleFunc("/customer/1/edit", customerEditHandler)
	http.HandleFunc("/modal", modalHandler)
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("dist"))))

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
