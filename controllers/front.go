package controllers

import "net/http"

/*RegisterControllers registers a new controler */
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
