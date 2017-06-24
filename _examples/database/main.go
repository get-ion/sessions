package main

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"

	"github.com/get-ion/sessions"
	"github.com/get-ion/sessions/sessiondb/redis"
	"github.com/get-ion/sessions/sessiondb/redis/service"
)

func main() {
	// replace with your running redis' server settings:
	db := redis.New(service.Config{Network: service.DefaultRedisNetwork,
		Addr:          service.DefaultRedisAddr,
		Password:      "",
		Database:      "",
		MaxIdle:       0,
		MaxActive:     0,
		IdleTimeout:   service.DefaultRedisIdleTimeout,
		Prefix:        "",
		MaxAgeSeconds: service.DefaultRedisMaxAgeSeconds}) // optionally configure the bridge between your redis server

	sess := sessions.New(sessions.Config{Cookie: "sessionscookieid"})

	//
	// IMPORTANT:
	//
	sess.UseDatabase(db)

	// the rest of the code stays the same.
	app := ion.New()

	app.Get("/", func(ctx context.Context) {
		ctx.Writef("You should navigate to the /set, /get, /delete, /clear,/destroy instead")
	})
	app.Get("/set", func(ctx context.Context) {

		//set session values
		sess.Start(ctx).Set("name", "ion")

		//test if setted here
		ctx.Writef("All ok session setted to: %s", sess.Start(ctx).GetString("name"))
	})

	app.Get("/get", func(ctx context.Context) {
		// get a specific key, as string, if no found returns just an empty string
		name := sess.Start(ctx).GetString("name")

		ctx.Writef("The name on the /set was: %s", name)
	})

	app.Get("/delete", func(ctx context.Context) {
		// delete a specific key
		sess.Start(ctx).Delete("name")
	})

	app.Get("/clear", func(ctx context.Context) {
		// removes all entries
		sess.Start(ctx).Clear()
	})

	app.Get("/destroy", func(ctx context.Context) {
		//destroy, removes the entire session data and cookie
		sess.Destroy(ctx)
	})

	app.Run(ion.Addr(":8080"))
}
