package main

import (
	"encoding/json"
	"neversitup-test/internal/core/implements/services/counter"
	"neversitup-test/internal/core/implements/services/odd"
	"neversitup-test/internal/core/implements/services/permutations"
	"neversitup-test/internal/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var port string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("cannot read environment variables")
	}
}

func counterRoute(r *fiber.App) {
	counterSrv := counter.NewCounterServiceImplement()
	counterHdr := handlers.NewCounterHandler(counterSrv)
	r.Get("/counter/smileys", counterHdr.CountSmileys)
}

func oddRoute(r *fiber.App) {
	oddSrv := odd.NewOddServiceImplement()
	oddHdr := handlers.NewOddHander(oddSrv)
	r.Get("/odd/find-int", oddHdr.FindIntInSlice)
}

func permutationRoute(r *fiber.App) {
	permutationsSrv := permutations.NewPermutationsServiceImplement()
	permuhtaionsHdr := handlers.NewPermutationsHandler(permutationsSrv)
	r.Get("/permutation", permuhtaionsHdr.Generate)
}

func main() {
	port := os.Getenv("APP_PORT")
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		BodyLimit:   0,
	})

	permutationRoute(app)
	oddRoute(app)
	counterRoute(app)

	app.Listen(":" + port)
}
