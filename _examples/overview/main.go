package main

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"

	"github.com/get-ion/sessions"
)

var (
	cookieNameForSessionID = "mycookiesessionnameid"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)

func secret(ctx context.Context) {

	// Check if user is authenticated
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(ion.StatusForbidden)
		return
	}

	// Print secret message
	ctx.WriteString("The cake is a lie!")
}

func login(ctx context.Context) {
	session := sess.Start(ctx)

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Set("authenticated", true)
}

func logout(ctx context.Context) {
	session := sess.Start(ctx)

	// Revoke users authentication
	session.Set("authenticated", false)
}

func main() {
	app := ion.New()

	app.Get("/secret", secret)
	app.Get("/login", login)
	app.Get("/logout", logout)

	app.Run(ion.Addr(":8080"))
}
