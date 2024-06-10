package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ship struct {
	ID      string `json:"id"`
	Pilot   string `json:"pilot"`
	Model   string `json:"model"`
	Alias   string `json:"alias"`
	Bullets int16  `json:"bullets"`
	Landed  bool   `json:"landed"`
}

var ships = []ship{
	{ID: "1", Pilot: "Biggs Darklighter", Model: "T-65B", Bullets: 6, Alias: "Red Three", Landed: false},
	{ID: "2", Pilot: "Garven Dreis", Model: "T-75C", Bullets: 4, Alias: "Red Leader", Landed: false},
	{ID: "3", Pilot: "Wedge Antilles", Model: "T-65B", Bullets: 9, Alias: "Red Two", Landed: false},
	{ID: "4", Pilot: "Dutch Vander", Model: "T-99A", Bullets: 8, Alias: "Gold Leader", Landed: false},
	{ID: "5", Pilot: "Luke Skywalker", Model: "T-65B", Bullets: 8, Alias: "Red Five", Landed: false},
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})
	router.GET("/x-wings", getShips)
	router.POST("/x-wings", postShip)
	router.GET("/x-wings/:id", getShipByID)
	router.PUT("/x-wings/:id/land", landShip)
	router.PUT("/x-wings/:id/shoot", landShoot)
	router.PUT("/x-wings/reload", reloadShip)

	router.Run(":8080")
}

func landShip(c *gin.Context) {
	id := c.Param("id")

	for x, a := range ships {
		if a.ID == id {
			landingShip(x)
			c.IndentedJSON(http.StatusOK, gin.H{"shipLanded": a.Alias})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ship was shot down"})
}

func landingShip(x int) {
	ships[x].Landed = true
}

func landShoot(c *gin.Context) {
	id := c.Param("id")
	for x, a := range ships {
		if a.ID == id {
			if a.Bullets > 0 && !a.Landed {
				ships[x].Bullets = ships[x].Bullets - 1
				c.IndentedJSON(http.StatusOK, gin.H{"shipShoot": a.Alias, "bullets": a.Bullets})
			} else {
				if a.Landed {
					c.IndentedJSON(http.StatusOK, "x-wing landed, can not shoot!")
				} else {
					landingShip(x)
					c.IndentedJSON(http.StatusOK, "x-wing with no bullets, probably should land!")
				}
			}
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ship was shot down"})
}

func getShips(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ships)
}

func reloadShip(c *gin.Context) {
	var loadedShip ship
	if err := c.BindJSON(&loadedShip); err != nil {
		return
	}
	for x, a := range ships {
		if a.ID == loadedShip.ID {
			ships[x].Bullets = loadedShip.Bullets
			ships[x].Landed = false
			ships[x].Pilot = loadedShip.Pilot
			ships[x].Alias = loadedShip.Alias
			c.IndentedJSON(http.StatusAccepted, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Ship does not exist, probably it was destroyed"})
}

func postShip(c *gin.Context) {
	var newShip ship

	// Call BindJSON to bind the received JSON to
	// newShip.
	if err := c.BindJSON(&newShip); err != nil {
		return
	}

	// Add the new album to the slice.
	ships = append(ships, newShip)
	c.IndentedJSON(http.StatusCreated, newShip)
}

func getShipByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range ships {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ship was shot down"})
}
