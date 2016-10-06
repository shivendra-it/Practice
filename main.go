package main
import (
	"fmt"
	"github.com/goLanguageDetection"
)
func main() {
  var detect goLanguageDetection.Detect = *goLanguageDetection.New()
  
  result, validity := detect.Text("i am a good boy i do a lot of work")

  fmt.Println(result) // English
  fmt.Println(validity) // Ratio of words found in the returned language : 0.92
}

