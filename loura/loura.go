package loura

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/lys091112/gopiers/loura/fractal"
)

func Start() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Static("/css", "resources/css")
	r.Static("/js", "resources/js")
	r.LoadHTMLGlob("resources/template/*")

	// TODO 开启pprof分析
	pprof.Register(r)

	r.GET("/", indexHandler)

	// fractal route
	fractal.NewRoute(r)

	server := &http.Server{Addr: ":8080", Handler: r}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("start failed! error=%s\n", err)
		}
		log.Fatal("Start loura success!")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutdown server")

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("server shutDown:", err)
	}

	log.Print("service exit")

}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"result": "Hello World!",
	})
}
