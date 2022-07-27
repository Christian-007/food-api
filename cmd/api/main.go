package main

import (
	"food/internal/menu"
	"food/internal/restaurant"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	router := chi.NewRouter()
	
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&restaurant.Restaurant{}, &menu.Menu{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	// restaurant
	restaurantRepository := restaurant.NewRepositoryGorm(db)
	restaurantHandler := restaurant.NewHttpHandler(restaurantRepository)
	router.Mount("/restaurants", restaurantHandler.Routes())

	// menu
	menuRepository := menu.NewRepositoryGorm(db)
	menuHandler := menu.NewHttpHandler(menuRepository)
	router.Mount("/menus", menuHandler.Routes())
	router.Mount("/restaurants/{restaurantId}/menus", menuHandler.SubRoutes())

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Failed to start the server: %v", err)
	}
}
