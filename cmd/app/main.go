package main

import (
	"log"
	"os"
	"vk_server/configs"
	"vk_server/internal/repository/ad"
)

func main() {
	storage, err := configs.NewStorage()
	if err != nil {
		log.Fatal("Failed to create storage")
		os.Exit(1)
	}
	s := ad.NewAddStorage(storage)
	s.SaveAd("he", "asd", 12, "2e094c8c-bbf5-4e0e-8316-569128bd282d")

}
