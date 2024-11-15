package configs

import "github.com/swaggo/swag"

func SwaggerConfigure(infos *swag.Spec) {
	infos.Title = "E-commerce API Orders"
	infos.Description = "API Rest for order management"
	infos.Host = "localhost:8092"
	infos.BasePath = "/"
	infos.Version = "1.0"
	infos.Schemes = []string{"http", "https"}
}
