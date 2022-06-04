package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type player struct {
	ID		string `json:"id"`
	Name	string `json:"name"`
	Team	string `json:"team"`
}

type score struct {
	Match	string	`json:"match"`
	Runs	int		`json:"runs"`
	Wickets	int 	`json:"wickets"`
}

var players = []player{
	{"1","Virat Kohli","RCB"},
	{"2","ABD","RCB"},
}

type scores struct {
	Id		string	`json:"id"`
	Scores	[]score	`json:"scores"`
}

var playerScores = []scores{
	{
		"1", []score{{"1",20,0}, {"2",12,1}},
	},
	{
		"2", []score{{"1",78,9}},
	},
}

func getPlayers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, players)
}

func postPlayers(c *gin.Context) {
	var newPlayer player
	if err := c.BindJSON(&newPlayer); err != nil {
		return
	}
	var newPlayerScore scores
	fmt.Println(newPlayer)
	fmt.Println(newPlayer.ID)
	newPlayerScore.Id=newPlayer.ID
	newPlayerScore.Scores=[]score{}
	players = append(players, newPlayer)
	playerScores = append(playerScores, newPlayerScore)
	c.IndentedJSON(http.StatusCreated, newPlayer)
}

func getPlayerById(c *gin.Context) {
	id := c.Param("id")

	for _, p := range players {
		if p.ID == id {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Player with this Id doesn't exist"})
}

func getScores(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, playerScores)
}

func postScores(c *gin.Context) {
	id := c.Param("id")
	var newScore score
	if err:=c.BindJSON(&newScore); err != nil {
		return
	}
	for _, p := range players {
		if p.ID == id {
			for i, playerScore := range playerScores {
				 if playerScore.Id == id {
					 newScores := append(playerScore.Scores, newScore)
					 playerScore.Scores=newScores
					 playerScores[i] = playerScore
					 c.IndentedJSON(http.StatusCreated, playerScores)
					 return
				 }
			}
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"Player with this Id exist, But unable to add scores"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Player with this Id doesn't exist"})
}

func main() {
	router := gin.Default()
	router.GET("/players", getPlayers)
	router.POST("/players", postPlayers)
	router.GET("/players/:id", getPlayerById)
	router.GET("/scores", getScores)
	router.POST("/scores/:id", postScores)
	router.Run("localhost:8080")
}