package main

import (
	"github.com/erzha/http/server"
	"github.com/erzha/http/plugin/view"
	pluginSession "github.com/erzha/http/plugin/session"

	"github.com/walu/unknown/action"
)

func setRouter() {
	server.Router("/_session/login", "login", func()server.ActionInterface{ return &action.LoginAction{} })
	server.Router("/_session/logout", "logout", func()server.ActionInterface{ return &action.LogoutAction{} })
	server.Router("/#name#", "index", func()server.ActionInterface{ return &action.IndexAction{} })
}

//every minutes.
func crontabGitSync() {
	//get the configs
	//sync every config: gitrepo <-> local dirs
}

func main() {
	//register plugins
	pluginSession.RegisterPlugin()
	view.RegisterPlugin()

	//set url router
	setRouter()

	//boot
	server.Boot()
}
