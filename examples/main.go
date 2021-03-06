package main

import (
	"awesomeProject/link-parser"
	"fmt"
	"strings"
)

var exampleHtml = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
<h1>Servus</h1>
<a href="https://www.youtube.com/watch?v=wMKpYxhI2KI&t=168s">Awesome <span> Human Being</span></a>
<a href="https://www.youtube.com/watch?v=OdtCuC5mLeQ">King of Minchinland</a>
<a href="https://www.youtube.com/watch?v=dQw4w9WgXcQ">Why are u even here?</a>
<a href="www.linkedin.at"><span>linkedin Link a lot of link</span></a>
<a href="www.facebook.com">Zuckerberg is a robot convince me otherwise</a>
<a href="www.spacex.com">We love memelord</a>

</body>
</html>
`
func main() {
	r := strings.NewReader(exampleHtml)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
