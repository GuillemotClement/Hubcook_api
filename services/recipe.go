package services

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/GuillemotClement/Hubcook_api/models"
	"github.com/gin-gonic/gin"
)

func GetRecipes(c *gin.Context) {

	var recipes []models.Recipe

	query := `
						SELECT recipe.id, recipe.title, 
										recipe.describ, 
										recipe.time_prep, 
										recipe.image, 
										category.title, 
										users.username
						FROM recipe
							LEFT JOIN category ON recipe.category_id  = category.id  
							LEFT JOIN users on user_id = users.id
						`

	rows, err := db.Query(query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Erreur": "Erreur de recuperation des donnees"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var recipe models.Recipe
		if err := rows.Scan(&recipe.Id, &recipe.Title, &recipe.Describ, &recipe.TimePrep, &recipe.Image, &recipe.Author, &recipe.Category); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Erreur": "Erreur de recuperation"})
		}
		recipes = append(recipes, recipe)
	}

	if len(recipes) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"Erreur": "Pas de donnees"})
		return
	}

	fmt.Print(recipes)

	c.JSON(http.StatusOK, gin.H{"message": "Succes", "data": recipes})

}

func GetDetailRecipe(c *gin.Context) {
	type Recipe struct {
		Id       uint   `json:"id"`
		Title    string `json:"title"`
		Describ  string `json:"describ"`
		TimePrep uint   `json:"timePrep"`
		Image    string `json:"image"`
		Author   string `json:"author"`
		Category string `json:"category"`
	}

	var recipe Recipe

	id := c.Param("id")

	query := `
		SELECT recipe.id, 
					recipe.title, 
					recipe.describ, 
					recipe.time_prep, 
					recipe.image, 
					category.title, 
					users.username
		FROM recipe
			LEFT JOIN category ON recipe.category_id  = category.id  
			LEFT JOIN users on user_id = users.id
		WHERE recipe.id = $1
	`

	err := db.QueryRow(query, id).Scan(&recipe.Id, &recipe.Title, &recipe.Describ, &recipe.TimePrep, &recipe.Image, &recipe.Category, &recipe.Author)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"Erreur": "Recette non trouver"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"Erreur": "Erreur serveur"})
		}
	}

	c.JSON(http.StatusOK, recipe)
}
