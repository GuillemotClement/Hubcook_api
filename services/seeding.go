package services

import (
	"fmt"
	"os"

	"math/rand"

	"github.com/go-faker/faker/v4"
	"golang.org/x/crypto/bcrypt"
)

func GenerateFileSeedingDB() {
	// creation du fichier
	f, err := os.Create("./database/02_seedDB.sql")

	if err != nil {
		fmt.Println(err)
		return
	}

	// requete a ecrire
	dataToWrite := "BEGIN;\n\n"
	queryRole := GenerateRole()  // creation de la requete pour generer des roles
	queryUser := GenerateUser(5) // creation de la requete pour generer les users
	queryCategory := GenerateCategory()
	queryRecipe := GenerateRecipe(30)

	dataToWrite += fmt.Sprintf(queryRole)
	dataToWrite += fmt.Sprintf(queryUser)
	dataToWrite += fmt.Sprintf(queryCategory)
	dataToWrite += fmt.Sprintf(queryRecipe)

	dataToWrite += fmt.Sprintf("COMMIT;")

	l, err := f.WriteString(dataToWrite)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	fmt.Println(l, "Ecriture ok")

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GenerateUser(numberUserToGenerate uint) (userQuery string) {
	// creer un struct qui represente un utilisateur
	type UserSeed struct {
		Username string
		Email    string
		Password string
		Image    string
		RoleId   uint
	}

	// mdp basique pour le dev
	mdp := "123456"

	// generation du mot de passe hasher
	hashed, err := bcrypt.GenerateFromPassword([]byte(mdp), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Erreur pour le hash du mot de passe")
		return
	}

	userList := []UserSeed{
		{"gizmo", "gizmo@mail.com", string(hashed), "https://randomuser.me/api/portraits/men/9.jpg", 1},
	}

	for i := 0; i < int(numberUserToGenerate); i++ {
		// generation d'un nombre aleatoire de 1 a 10 pour choper un image de profil
		randomValue := rand.Intn(10-0) + 1
		randomRole := rand.Intn(2-0) + 1

		// generation de l'url pour l'image
		urlImage := fmt.Sprintf("https://randomuser.me/api/portraits/men/%d.jpg", randomValue)

		username := faker.FirstName()
		email := fmt.Sprintf("%v@mail.com", username)

		// creation de l'user
		userData := UserSeed{
			Username: username,
			Email:    email,
			Password: string(hashed),
			Image:    urlImage,
			RoleId:   uint(randomRole),
		}

		userList = append(userList, userData)

	}

	query := fmt.Sprintf("INSERT INTO users (username, email, password, image, role_id) VALUES \n")

	for i, user := range userList {
		if i != 0 {
			query += fmt.Sprintf(",\n")
		}
		query += fmt.Sprintf("('%v', '%v', '%v', '%v', '%v')", user.Username, user.Email, user.Password, user.Image, user.RoleId)
	}

	query += fmt.Sprintf(";\n\n")

	return query
}

func GenerateRole() (roleQuery string) {
	type RoleSeed struct {
		title string
	}

	roleList := []RoleSeed{
		{"admin"},
		{"user"},
		{"moderator"},
		{"writter"},
	}

	query := fmt.Sprintf("INSERT INTO role (title) VALUES \n")

	for i, role := range roleList {
		if i != 0 {
			query += fmt.Sprintf(",\n")
		}
		query += fmt.Sprintf("('%v')", role.title)
	}

	query += fmt.Sprintf(";\n\n")

	return query
}

func GenerateCategory() (categoryQuery string) {
	type CategorySeed struct {
		title string
	}

	categoryList := []CategorySeed{
		{"salade"},
		{"fast food"},
		{"vegan"},
		{"viande"},
		{"desert"},
		{"entree"},
		{"plat"},
		{"gouter"},
	}

	query := fmt.Sprintf("INSERT INTO category (title) VALUES \n")

	for i, category := range categoryList {
		if i != 0 {
			query += fmt.Sprintf(",\n")
		}
		query += fmt.Sprintf("('%v')", category.title)
	}

	query += fmt.Sprintf(";\n\n")

	return query
}

func GenerateRecipe(numberToGenerate uint) (queryRecipe string) {

	type RecipeSeed struct {
		Title      string
		Describ    string
		TimePrep   uint
		Image      string
		UserId     uint
		CategoryId uint
	}

	recipeList := []RecipeSeed{}

	for i := 0; i < int(numberToGenerate); i++ {
		randomValue := rand.Intn(1000-0) + 1 // generation du nombre aleatoire pour image
		randomTimePrep := rand.Intn(200-4) + 1
		randomUserId := rand.Intn(3-0) + 1
		randomCategory := rand.Intn(8-1) + 1

		urlPicture := fmt.Sprintf("https://fastly.picsum.photos/id/%v/400/400.jpg", randomValue)

		recipeData := RecipeSeed{
			Title:      fmt.Sprintf(faker.Word() + faker.Word()),
			Describ:    fmt.Sprintf(faker.Paragraph()),
			TimePrep:   uint(randomTimePrep),
			Image:      urlPicture,
			UserId:     uint(randomUserId),
			CategoryId: uint(randomCategory),
		}

		recipeList = append(recipeList, recipeData)
	}

	query := "INSERT INTO recipe (title, describ, time_prep, image, user_id, category_id) VALUES\n"

	for i, recipe := range recipeList {
		if i != 0 {
			query += fmt.Sprintf(",\n")
		}
		query += fmt.Sprintf("('%v', '%v', '%v', '%v', '%v', '%v')", recipe.Title, recipe.Describ, recipe.TimePrep, recipe.Image, recipe.UserId, recipe.CategoryId)
	}

	query += fmt.Sprintf(";\n\n")

	return query
}
