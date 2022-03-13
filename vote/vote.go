package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles"
	// swagger embed files
)

// VoteCard represents data about votes
type VoteCard struct {
	ID      string         `json:"id"`
	OPTIONS []string       `json:"options"`
	VOTES   map[string]int `json:"votes"`
}

type CreateVoteCardRequest struct {
	ID      string   `json:"id"`
	OPTIONS []string `json:"options"`
}

type HttpErrorMessage struct {
	MESSAGE string `json:"message"`
}

// hold votecards in memory for now
var votecards = []VoteCard{}

func setupRouter() *gin.Engine {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.New()

	router.GET("/votecard/:id", getVoteCardById)
	router.PUT("/votecard/:id", updateVoteCount)
	router.POST("/votecard", createVoteCard)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

// @title           Survey Voting API
// @version         1.0
// @description     This is a Survey Voting API
// @contact.name   Matthew Croft
// @contact.url    https://www.linkedin.com/in/matthew-croft-44a5a5b3/
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
func main() {
	router := setupRouter()

	router.Run("localhost:8080")
}

// GetVoteCard godoc
// @Summary		Get VoteCard
// @Description Returns a VoteCard
// @Produce		json
// @Param		id		path	int		true	"VoteCard ID"
// @Success		200	{object}	VoteCard
// @Failure		404 {object} 	HttpErrorMessage "VoteCard not found"
// @Router       /votecard/{id} [get]
func getVoteCardById(c *gin.Context) {
	id := c.Param("id")

	for _, votecard := range votecards {
		if votecard.ID == id {
			c.IndentedJSON(http.StatusOK, votecard)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, HttpErrorMessage{MESSAGE: "vote card not found"})
}

// CreateVoteCard godoc
// @Summary		Create VoteCard
// @Description Creates a VoteCard that can be used in a survey
// @Accept		json
// @Produce		json
// @Param		createVoteCardRequest	body	CreateVoteCardRequest	true	"Create VoteCard request body"
// @Success		200 {object}	VoteCard
// @Failure		400 {object}	HttpErrorMessage	"Incorrect request body"
// @Router       /votecard [post]
func createVoteCard(c *gin.Context) {
	var newVoteCardRequest CreateVoteCardRequest
	optionMap := make(map[string]int)

	if err := c.BindJSON(&newVoteCardRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, HttpErrorMessage{MESSAGE: "incorrect request body, should be VoteCard body"})
		return
	}

	for _, option := range newVoteCardRequest.OPTIONS {
		optionMap[option] = 0
	}

	var newVoteCard = VoteCard{ID: newVoteCardRequest.ID, OPTIONS: newVoteCardRequest.OPTIONS, VOTES: optionMap}

	votecards = append(votecards, newVoteCard)
	c.IndentedJSON(http.StatusCreated, newVoteCard)
}

// UpdateVoteCount godoc
// @Summary		Update count on a VoteCard
// @Description Updates count for a certain option in the VoteCard
// @Accept		json
// @Produce		json
// @Param		id		path	int		true	"VoteCard ID"
// @Param		option	query	string	true	"Option to update vote for"
// @Success		200 {object}	VoteCard
// @Failure		404 {object}	HttpErrorMessage	"VoteCard not found"
// @Failure		400	{object}	HttpErrorMessage	"Not a valid option"
// @Router       /votecard/{id} [put]
func updateVoteCount(c *gin.Context) {
	option := c.Query("option")
	id := c.Param("id")
	var updateVoteCard = VoteCard{}

	for _, votecard := range votecards {
		if votecard.ID == id {
			updateVoteCard = votecard
		}
	}

	if updateVoteCard.ID == "" {
		c.IndentedJSON(http.StatusNotFound, HttpErrorMessage{MESSAGE: "vote card not found"})
		return
	}

	for _, op := range updateVoteCard.OPTIONS {
		if option == op {
			updateVoteCard.VOTES[option] = updateVoteCard.VOTES[option] + 1
			c.IndentedJSON(http.StatusOK, updateVoteCard)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, HttpErrorMessage{MESSAGE: "not a valid option"})
}
