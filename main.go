package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/handlers"
	"gopkg.in/natefinch/lumberjack.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	. "github.com/newtechfellas/CurryPoint/currypoint"
	. "github.com/newtechfellas/CurryPoint/entity"
)

var configFile = flag.String("config", "config.json", "Config File")
var helpFlag = flag.Bool("help", false, "Print help text")

func main() {
	flag.Parse()
	if *helpFlag {
		fmt.Fprintf(os.Stderr, "Usage of %s\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
	config := loadConfig(*configFile)

	logger := setupLogging(config.Log.File)

	mysqlDbmap := PrepareDBMap(config.MySql.ConnectionString)
	defer mysqlDbmap.Db.Close()

	mysqlDbmap.TraceOn("[gorp]", &Foo{})

	router := NewRouter()

	//Logs the http requests status
	logHandler := handlers.LoggingHandler(logger, router)
	http.ListenAndServe(":8080", RecoveryHandler{Handler: logHandler})
}

type Foo struct{}

func (logger Foo) Printf(format string, v ...interface{}) {
	log.Printf(format, v)
}

func setupLogging(logfile string) *lumberjack.Logger {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	logger := &lumberjack.Logger{
		Filename:   logfile,
		MaxSize:    20, // megabytes
		MaxBackups: 20,
		MaxAge:     30, //days
	}
	log.SetOutput(logger)
	return logger
}

func loadConfig(filenameWithAbsPath string) *Config {
	file, err := ioutil.ReadFile(filenameWithAbsPath)
	if err != nil {
		panic("ERROR: Failed to load config file. Error is " + err.Error())
	}
	config := &GlobalConfig
	if err = json.Unmarshal(file, config); err != nil {
		panic("ERROR: Failed to load config.json. Error is " + err.Error())
	}
	fmt.Printf("Config file read is: %s\n", config)
	return config
}