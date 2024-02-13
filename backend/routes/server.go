package routes

import(
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2"
)

func Server(app fiber.Router) {
	app.Get("/status", monitor.New(monitor.Config{
		Title: "Server Status",
		FontURL:"https://fonts.googleapis.com/css2?family=Roboto:wght@400;900&display=swap",
		ChartJsURL:"https://cdn.jsdelivr.net/npm/chart.js@2.9/dist/Chart.bundle.min.js",
	}))
}