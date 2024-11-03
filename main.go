package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ProfileResponse struct {
	Name       string   `json:"Name"`
	Education  string   `json:"Education"`
	Field      string   `json:"Field"`
	LookingFor string   `json:"Looking_for"`
	Skills     []string `json:"Skills"`
	Experience string   `json:"Experience"`
	OpenSource string   `json:"Open_source"`
	Project    Project  `json:"Project"`
	Contact    Contact  `json:"Contact"`
	Note       string   `json:"Note"`
}

type Project struct {
	Name        string   `json:"Name"`
	Description string   `json:"Description"`
	Backend     string   `json:"Backend"`
	Frontend    string   `json:"Frontend"`
	Features    []string `json:"Features"`
}

type Contact struct {
	Phone    string `json:"Phone"`
	Resume   string `json:"Resume"`
	LinkedIn string `json:"LinkedIn"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.Default()
	config := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}

	router.Use(cors.New(config))

	router.GET("/profile", func(c *gin.Context) {
		response := ProfileResponse{
			Name:       "Hemanth Kumar Reddy",
			Education:  "junior at IIT Bhilai",
			Field:      "Computer Science and Engineering",
			LookingFor: "Backend Engineering Intern at Fam",
			Skills: []string{
				"React.js",
				"Node.js",
				"GoLang",
				"MongoDB",
				"React Native",
			},
			Experience: "Developed dynamic web applications during internship at Litadz Media.",
			OpenSource: "Contributed to C2SI and Avanti Fellows, tackling easy to medium-level challenges; currently working on medium-hard level challenges with a mentor at Avanti Fellows.",
			Project: Project{
				Name:        "Techque",
				Description: "A queue management system designed for restaurants.",
				Backend:     "Built in Go",
				Frontend:    "Utilizes React and Tailwind CSS",
				Features: []string{
					"virtual waitlisting",
					"real-time wait time tracking",
				},
			},
			Contact: Contact{
				Phone:    "+91 9398481212",
				Resume:   "https://drive.google.com/file/d/1obnsLpnYesYonIfrvktA7mP6Z_kpPGnI/view?usp=sharing",
				LinkedIn: "https://www.linkedin.com/in/hemanth-kumar-reddy-89668b252/",
			},
			Note: "This JSON structure is written in Golang using the Gin framework.",
		}

		c.JSON(200, response)
	})

	err := router.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
