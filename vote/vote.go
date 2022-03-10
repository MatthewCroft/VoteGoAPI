package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// VoteCard represents data about votes
type votecard struct {
	ID      string         `json:"id"`
	OPTIONS []string       `json:"options"`
	VOTES   map[string]int `json:"votes"`
}

// hold votecards in memory for now
var votecards = []votecard{}

func setupRouter() *gin.Engine {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.New()

	router.GET("/votecard/:id", getVoteCardById)
	router.PUT("/votecard/:id", updateVoteCount)
	router.POST("/votecard", createVoteCard)

	return router
}

func main() {
	router := setupRouter()

	router.Run("localhost:8080")
}

// returning vote card
func getVoteCardById(c *gin.Context) {
	id := c.Param("id")

	for _, votecard := range votecards {
		if votecard.ID == id {
			c.IndentedJSON(http.StatusOK, votecard)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "vote card not found"})
}

// create vote card
func createVoteCard(c *gin.Context) {
	var newVoteCard votecard

	if err := c.BindJSON(&newVoteCard); err != nil {
		return
	}

	votecards = append(votecards, newVoteCard)
	c.IndentedJSON(http.StatusCreated, newVoteCard)
}

// update votes
func updateVoteCount(c *gin.Context) {
	option := c.Query("option")
	id := c.Param("id")
	var updateVoteCard = votecard{}

	for _, votecard := range votecards {
		if votecard.ID == id {
			updateVoteCard = votecard
		}
	}

	if updateVoteCard.ID == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "vote card not found"})
		return
	}

	for _, op := range updateVoteCard.OPTIONS {
		if option == op {
			updateVoteCard.VOTES[option] = updateVoteCard.VOTES[option] + 1
			c.IndentedJSON(http.StatusOK, updateVoteCard)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not a valid option"})
}
