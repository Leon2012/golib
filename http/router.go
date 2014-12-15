package router

import (
	"log"
	"net/http"
	"strings"
)

type Router struct {
	handlers map[string]*parser
	NotFound Handle
	PanicHandler
}
