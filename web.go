package main

import (
    "os"
    "github.com/hoisie/web"
    "math/rand"
    "time"
    "fmt"
)

var sloths = []string{
    "http://cuteoverload.files.wordpress.com/2012/03/6852916110_3a06e758d9_z.jpg",
    "http://static.environmentalgraffiti.com/sites/default/files/images/Sloth2.img_assist_custom-600x450.jpg",
    "http://www.veganreader.com/images/pictureofsloth.gif",
    "http://www.jakobdezwart.com/wp-content/uploads/2012/11/jakobdezwart-bestof-sloths-9517.jpg",
    "http://media.smithsonianmag.com/images/Three-toed-sloth-Panama-631.jpg",
}

func hello(val string) string {
    return fmt.Sprintf(`
    <!DOCTYPE html><html><body>
    <h1>DocSloth</h1>
    <img src="%s" />
    </body></html>
`, sloths[rand.Intn(len(sloths))])
}

func main() {
    rand.Seed(time.Now().UnixNano())
    port := os.Getenv("PORT")
    if port == "" { port = "9999" }
    web.Get("/(.*)", hello)
    web.Run(":"+port)
}

