package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"sync"
)

type ServerConfig struct {
	Host           string
	Port           int
	WelcomeMessage string
}

var (
	serverConfig ServerConfig
	mutex        sync.RWMutex
)

func loadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	if err := viper.UnmarshalKey("server", &serverConfig); err != nil {
		log.Fatalf("Unable to decode server config into struct: %s", err)
	}
}

func main() {
	loadConfig()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed: ", e.Name)
		mutex.Lock()
		loadConfig()
		mutex.Unlock()
	})

	log.Printf("Welcome message: %s\n", serverConfig.WelcomeMessage)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mutex.RLock()
		welcomeMessage := serverConfig.WelcomeMessage
		mutex.RUnlock()
		fmt.Fprint(w, welcomeMessage)
	})

	address := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)
	fmt.Printf("Starting server at %s\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
