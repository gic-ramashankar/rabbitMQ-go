package router

import (
	"demo/controller"
	"net/http"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("/rabbitMQ/send_to_queue", http.HandlerFunc(controller.SendToQueue))

	return router
}
