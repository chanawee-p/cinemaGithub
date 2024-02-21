package main

import (
	"fmt"

	"github.com/chanawee-p/cinemaGithub/movie"
	"github.com/chanawee-p/cinemaGithub/ticket"
)

func init() {
	fmt.Println("init: main")
}

func main(){
	movieName := movie.FindMovieName("tt4154796")
	fmt.Println(movieName)
	movie.RevieMovie(movieName, 8.4)
	ticket.BuyTicket(movieName)
}