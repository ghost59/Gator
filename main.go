package main


import (
	_ "github.com/lib/pq"
	"gator/internal/config"
		
		"log"
		"os"
		"database/sql"
		"gator/internal/database"
		)

type state struct {
	db *database.Queries
	cfg *config.Config
}

func main() {
	

read, err := config.Read()

if err != nil {
	log.Fatal(err)
}

dbURL := read.Db_url

db, errs := sql.Open("postgres", dbURL)

if errs != nil {
	log.Fatal(errs)
}
dbQueries := database.New(db)

prog := &state{cfg: &read,db:dbQueries}

cmds := Commands{
	com: make(map[string]func(*state,command)error),
}
if len(os.Args) <2 {
	log.Fatal("not enought args")
}
comma := os.Args[1]
split := os.Args[2:]
cmd := command{name:comma, args:split}
cmds.register("login", handlerLogin)
cmds.register("register",handlerRegister)
cmds.register("reset",handleReset)
cmds.register("users",handlerListUsers)
cmds.register("agg",Agg)
cmds.register("addfeed", middlewareLoggedIn(addFeed))
cmds.register("feeds",Feeds)
cmds.register("follow", middlewareLoggedIn(follow))
cmds.register("following",middlewareLoggedIn(following))
cmds.register("unfollow",middlewareLoggedIn(unfollow))
cmds.register("browse",middlewareLoggedIn(Browse))
err = cmds.run(prog, cmd)
if err != nil {
	log.Fatal(err)
}
}

