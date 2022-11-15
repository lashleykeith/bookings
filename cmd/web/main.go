package main

import (
	"github.com/lashleykeith/bookings/pkg/config"
	"github.com/lashleykeith/bookings/pkg/handlers"
	"github.com/lashleykeith/bookings/pkg/render"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main function
func main() {
	// change this to true when in production
	app.InProduction = false

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}





// go mod init github.com/lashleykeith/go-course
// go run ./cmd/web/.

// https://github.com/bmizerany/pat
// go get github.com/bmizerany/pat

// https://github.com/go-chi/chi
// go get -u github.com/go-chi/chi/v5

// https://github.com/alexedwards/scs
// go get github.com/alexedwards/scs/v2


// Git Commands https://www.datacamp.com/tutorial/git-push-pull
// add images
// https://www.freecodecamp.org/news/error-src-refspec-master-does-not-match-any-how-to-fix-in-git/