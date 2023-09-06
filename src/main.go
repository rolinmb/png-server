package main

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"math"
    "math/rand"
	"image"
	"image/color"
    "image/png"
	"time"
	"fmt"
	"log"
	"os"
	"github.com/rs/cors"
)

const (
	PORT = "8080"
)

type clientData struct {
	Filename string  `json:"filename"`
	Width    int     `json:"width"`
	Height   int     `json:"height"`
	Scale    float64 `json:"scale"`
}

func savePng(fname string, newPng *image.RGBA) {
    out,err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
    err = png.Encode(out, newPng)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("\nSuccessfully created/rewritten "+fname+"\n")
}


func calcColor(x, y int, scale float64) (uint8, uint8, uint8) {
    floatX := float64(x)
    floatY := float64(y)
	r := uint8((math.Sin((floatX*floatY*floatY)*scale) + 1) * 0.5 * 255)
	g := uint8((math.Sin((floatX*floatX*floatY)*scale) + 1) * 0.5 * 255)
	b := uint8((math.Sin((math.Exp(floatX+floatY))*scale) + 1) * 0.5 * 255)
	return r, g, b
}

func trippyPng(fname string, width, height int, scale float64) {
    newPng := image.NewRGBA(image.Rect(0, 0, width, height))
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            r, g, b := calcColor(x, y, scale)
            newPng.Set(x,y,color.RGBA{r, g, b, 255})
        }
    }
    savePng(fname, newPng)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodPost:
			fmt.Fprintf(w, "You made a POST request.")
			fmt.Printf("* Client made a POST request:\n")
			rand.Seed(time.Now().UnixNano())
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatal(err)
			}
			// log.Println(string(body)) // preview POSTed JSON request body
			var data clientData
			err = json.Unmarshal(body, &data)
			if err != nil {
				log.Fatal(err)
			}
			if data.Filename == "" {
				data.Filename = "new.png"
				fmt.Println("\nNo .png output filename entered; defaulting to 'new.png'\n")
			}
			fmt.Println("\t-> Client POSTed Filename: ", data.Filename)
			fmt.Println("\t-> Client POSTed .PNG Width: ", data.Width)
			fmt.Println("\t-> Client POSTed .PNG Height: ", data.Height)
			fmt.Println("\t-> Client POSTed Scale: ", data.Scale)
			trippyPng(data.Filename, data.Width, data.Height, data.Scale)
			// send newly created .png as server response to client
			http.ServeFile(w, r, data.Filename)
			fmt.Println("\nSuccessfully sent generated .png "+data.Filename+" to the client\n")
			// finally delete the newly created .png from server local storage (ideally POSTed data fields would be in a db)
			err = os.Remove(data.Filename)
			if err != nil {
				log.Println("Error deleting the generated .PNG "+data.Filename, err)
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	c := cors.Default()
	postHandler := http.HandlerFunc(handlePost)
	http.Handle("/generate", c.Handler(postHandler))
	fmt.Printf("[HTTP Server started on port %s]\n", PORT)
	http.ListenAndServe(":"+PORT, nil)
}