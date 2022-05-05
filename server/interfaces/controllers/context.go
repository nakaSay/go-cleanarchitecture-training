package controllers

type Context interface {
    QueryParam(string) string
	FormValue(string) string
}