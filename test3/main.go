package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/beef/summary", beefSummaryHandler)

	go func() {
		if err := e.Start(":8000"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down server: %v\n", err)
	}
}

func beefSummaryHandler(c echo.Context) error {
	text, err := fetchBaconIpsum()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch data"})
	}

	summary := processText(text)
	return c.JSON(http.StatusOK, map[string]map[string]int{"beef": summary})
}

func fetchBaconIpsum() (string, error) {
	resp, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func processText(text string) map[string]int {
	meats := []string{"t-bone", "fatback", "pastrami", "pork", "meatloaf", "jowl", "bresaola", "enim"}

	re := regexp.MustCompile(`[^\w\s-]`)
	text = re.ReplaceAllString(text, "")
	text = strings.ToLower(text)

	words := strings.Fields(text)

	counts := make(map[string]int)
	for _, word := range words {
		for _, meat := range meats {
			if word == meat {
				counts[meat]++
			}
		}
	}

	return counts
}
