package rectanglesolver

import (
	"encoding/json"
	"log"
	"time"

	"github.com/3ep-one/rectangle/rediswraper"
)

//Rectangle contain dimenstions and creation time
type Rectangle struct {
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Time   string `json:"time"`
}

//Jsoninput contain json input
type Jsoninput struct {
	Main  Rectangle   `json:"main"`
	Input []Rectangle `json:"input"`
}

//Haveoverlap check for overlaping rectangles
func Haveoverlap(jsoninput Jsoninput) {
	ans := []string{}
	mainRectangle := jsoninput.Main
	rectangleList := jsoninput.Input
	for _, rectangle := range rectangleList {
		doOverlap := true
		rectangle.Time = time.Now().Format("2006-01-02 15:04:05")
		if mainRectangle.X > rectangle.X+rectangle.Width || rectangle.X > mainRectangle.X+mainRectangle.Width {
			doOverlap = false
		}

		if mainRectangle.Y > rectangle.Y+rectangle.Height || rectangle.Y > mainRectangle.Y+mainRectangle.Height {
			doOverlap = false
		}

		if doOverlap {
			b, err := json.Marshal(&rectangle)
			if err != nil {
				log.Println(err)
			}
			ans = append(ans, string(b))
		}
	}
	redisClient := rediswraper.Makeredisclient()
	rediswraper.Setkeyvalue(redisClient, ans)
}
