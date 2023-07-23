package server

// "members/config"

// "github.com/gofiber/fiber/v2"
// "go.uber.org/fx"

// var New = fx.Annotate(
// 	create,
// )

// type (
// 	AppResult struct {
// 		*fiber.App
// 		PID  int
// 		Addr string
// 	}
// )

// func create(li fx.Lifecycle, prov config.ConfigProvider) *AppResult {
// 	// app := fiber.New()
// 	// cfg := prov.GetConfig()
// 	// addr := prov.HostPort(cfg.Members.Grpc)
// 	// app.Get("health", func(c *fiber.Ctx) error { return c.SendString("healthy ✨") })
// 	// pid := os.Getpid()
// 	// ar := &AppResult{
// 	// 	app,
// 	// 	pid,
// 	// 	addr,
// 	// }
// 	// go func() {
// 	// 	if err := ar.Listen(ar.Addr); err != nil {
// 	// 		panic(err)
// 	// 	}
// 	// }()
// 	// return ar
// }
