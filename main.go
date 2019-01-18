package main

import (
	"mapserver/initialrenderer"
	"github.com/sirupsen/logrus"
	"mapserver/layerconfig"
	"mapserver/app"
	"mapserver/params"
	"fmt"
)



func main() {
	logrus.SetLevel(logrus.InfoLevel)

	//Parse command line

  p := params.Parse()

  if p.Help {
    params.PrintHelp()
    return
  }

  if p.Version {
    fmt.Print("Mapserver version: ")
    fmt.Println(app.Version)
    return
  }

	//setup app context

	ctx, err := app.Setup(p)

	if err != nil {
		//error case
		panic(err)
	}

	//run initial rendering

	initialrenderer.Render(ctx.Tilerenderer, layerconfig.DefaultLayers)

}
