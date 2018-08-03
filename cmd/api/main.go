package main

import (
	"flag"
	"net/http"

	"github.com/go-chi/chi"
	"fmt"
	"context"
	"time"
	"os"
	"os/signal"
	"syscall"
)

var (
	banner = `
		 ┏┓　 ┏┓
		┏┛┻━━━┛┻┓
		┃　　　  ┃ 　
		┃　  ━　 ┃
		┃ ┳┛　┗┳ ┃
		┃　　　　 ┃
		┃　  ┻　　┃
		┃　　　　　┃
		┗━┓　　　┏━┛
		　┃　　　┃　　　　　　　　　
		　┃　　　┃
		　┃　　　┗━┓
		　┃　　　　┣┓
		　┃　　　　┏┛
		　┗┓┓┏━┳┓┏┛
		　 ┃┫┫ ┃┫┫
　　		   ┗┻┛ ┗┻┛
	`
)

func main() {
	var port string
	flag.StringVar(&port, "port", ":9090", "Port api server will listen on")
	flag.Parse()

	questionHandler := NewQuestionHandler()

	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Route("/question", func(r chi.Router) {
			r.Get("/", questionHandler.GetQuestion)
		})
	})

	run(port, r)
}

func run(port string, route chi.Router) {
	hs := http.Server{
		Addr: port,
		Handler: route,
	}

	go func() {
		fmt.Println(banner)
		fmt.Println("Mimiron is watching you", hs.Addr)
		if err := hs.ListenAndServe(); err != nil {
			fmt.Println("Mimiron exiting with error: ", err)
		}
	} ()

	// Graceful shutdown
	kill := make(chan os.Signal)
	signal.Notify(kill, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	sig := <-kill
	fmt.Printf("Received stop signal from system: %s\n", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	hs.Shutdown(ctx)
	time.Sleep(100 * time.Millisecond)
}
