package fractal

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lys091112/gopiers/loura/fractal/mandelbrot"
)

var fractals map[string]Fractal

// 注册分型图示例
func init() {
	fractals = make(map[string]Fractal)
	register(&mandelbrot.Mandelbrot{})
}

func NewRoute(r *gin.Engine) {

	subR := r.Group("/fractal")
	subR.GET("/:fractalName", fractalHandler)

	/******** mux **********/
	// subR := r.PathPrefix("/fractal").Subrouter()
	// subR.HandleFunc("/{key:[a-z0-9]+}",fractalHandler)
	// r.NewRoute()

}

func fractalHandler(c *gin.Context) {
	vars := c.Params
	v, _ := vars.Get("fractalName")
	if b, err := read(v); err == nil {
		c.Data(http.StatusOK, "image/png", b)
	} else {
		fmt.Print(err)
		c.JSON(http.StatusNotFound, v)
	}

	// w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, vars.Get("key"))
}

// func fractalHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, vars["key"])
// }

func register(f Fractal) {
	fractals[f.Name()] = f
}

func read(name string) ([]byte, error) {
	path := "/home/langle/xianyue/workspace/go/gopiers/loura/fractal/pictures"
	return os.ReadFile(path + "/" + name + ".png")
}

func draw(name string, param map[string]string) error {
	if f, ok := fractals[name]; ok {
		writer, err := os.OpenFile("./pictures/"+name+".png", os.O_CREATE|os.O_WRONLY, 0666)
		if nil != err {
			return err
		}
		defer writer.Close()
		f.DrawWithParam(writer, param)
	}
	return nil
}

type Fractal interface {
	DrawWithParam(writer io.Writer, param map[string]string) error
	Name() string
}
