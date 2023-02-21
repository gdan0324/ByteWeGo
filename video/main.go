package main

import (
	videoservice "github.com/gdan0324/ByteWeGo/video/kitex_gen/videoservice/videoservice"
	"log"
)

func main() {
	svr := videoservice.NewServer(new(VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
