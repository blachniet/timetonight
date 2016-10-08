package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"time"

	"github.com/blachniet/timetonight/toggl"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("Debug", false)
	viper.SetDefault("TemplatesGlobPattern", "./templates/*.tmpl")
	viper.SetDefault("HoursPerDay", 8)
	viper.SetDefault("TogglAPIToken", "")
	viper.SetDefault("Host", "")
	viper.SetDefault("Port", 3000)

	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/timetonight")
	viper.AddConfigPath("$HOME/.timetonight")
	viper.AddConfigPath(".")

	// Configure apply flag overrides
	flagConfigPath := flag.String("config", "", "Path to config file. By default, searches for config file named 'config.[toml|yaml|json]' at '/etc/timetonight/', '$HOME/.timetonight/' and './'")
	flagDebug := flag.Bool("debug", false, "Enables debugging. Overrides any setting in config file.")
	flag.Parse()
	if *flagConfigPath != "" {
		viper.SetConfigFile(*flagConfigPath)
	}
	if *flagDebug {
		viper.Set("Debug", *flagDebug)
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	tmplPattern := viper.GetString("TemplatesGlobPattern")
	tmpls, err := template.ParseGlob(tmplPattern)
	if err != nil {
		log.Fatalf("Failed to parse templates at %q: %+v", tmplPattern, err)
	}

	togglAPIToken := viper.GetString("TogglAPIToken")
	if togglAPIToken == "" {
		log.Fatal("TogglAPIToken not set")
	}
	timer, err := toggl.NewTimer(togglAPIToken)
	if err != nil {
		log.Printf("Unable to connect to Toggl with API token: %q\n", togglAPIToken)
		log.Fatalf("%+v", err)
	}

	hrsPerDay := viper.GetFloat64("HoursPerDay")
	timePerDay := time.Duration(hrsPerDay * float64(time.Hour))
	if timePerDay <= 0 {
		log.Fatal("Invalid 'HoursPerDay'. Must round to >0")
	}

	app := &app{
		Debug:            viper.GetBool("Debug"),
		Timer:            timer,
		Templ:            tmpls,
		TemplGlobPattern: viper.GetString("TemplatesGlobPattern"),
		TimePerDay:       timePerDay,
	}

	// Echo Setup
	e := echo.New()
	e.SetDebug(app.Debug)
	e.SetRenderer(app)

	// Echo Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())

	// Controllers
	homeController := &homeController{app}
	homeController.setup(e)

	// Echo Run
	e.Run(standard.New(fmt.Sprintf("%v:%v", viper.GetString("Host"), viper.GetInt("Port"))))
}
