package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type rotationItem struct {
	ID        int    `json:"id"`
	NameCN    string `json:"NameCN"`
	NameEN    string `json:"NameEN"`
	MapID     int    `json:"mapId"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type rotationConfig struct {
	Maps         []int     `json:"maps"`
	StandardTime time.Time `json:"standardTime"`
	LastTime     int       `json:"lastTime"`
}

type mapInfo struct {
	NameCN string
	NameEN string
	Image  string
}

var (
	maps = map[int]mapInfo{
		1: {NameCN: "诸王峡谷", NameEN: "Kings Canyon", Image: "kings-canyon.avif"},
		2: {NameCN: "世界尽头", NameEN: "World's Edge", Image: "worlds-edge.avif"},
		3: {NameCN: "残月", NameEN: "Broken Moon", Image: "broken-moon.avif"},
		4: {NameCN: "风暴点", NameEN: "Storm Point", Image: "storm-point.avif"},
		5: {NameCN: "电气区", NameEN: "E-District", Image: "e-district.avif"},
		6: {NameCN: "奥林匹斯", NameEN: "Olympus", Image: "Olympus.avif"},
	}
	rotationConfigs []rotationConfig
	timezone        = loadLocation(getenv("TIMEZONE", "Asia/Shanghai"))
)

func main() {
	loadRotation()

	r := gin.Default()
	r.Use(corsMiddleware())

	api := r.Group("/api")
	api.GET("/health", handleHealth)
	api.GET("/rotations", handleRotations)
	api.GET("/image/:mapId", handleImage)

	port := getenv("PORT", "8080")
	log.Printf("laozzz-backend listening on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}

func loadRotation() {
	data, err := os.ReadFile("data/rotation.json")
	if err != nil {
		log.Fatalf("read data/rotation.json: %v", err)
	}

	var configs []rotationConfig
	if err := json.Unmarshal(data, &configs); err != nil {
		log.Fatalf("parse data/rotation.json: %v", err)
	}

	for _, cfg := range configs {
		if cfg.LastTime <= 0 || len(cfg.Maps) == 0 {
			log.Printf("warning: skip invalid rotation config (maps=%v, lastTime=%d)", cfg.Maps, cfg.LastTime)
			continue
		}
		rotationConfigs = append(rotationConfigs, cfg)
	}
	if len(rotationConfigs) == 0 {
		log.Fatal("no valid rotation config in data/rotation.json")
	}
}

func handleHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func handleRotations(c *gin.Context) {
	date, err := parseDateQuery(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, buildDayRotation(rotationConfigs[0], date))
}

func handleImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("mapId"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "unknown map id"})
		return
	}
	meta, ok := maps[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "unknown map id"})
		return
	}
	dir := getenv("MAP_IMAGE_DIR", "static/images")
	c.Header("Content-Type", "image/avif")
	c.File(filepath.Join(dir, meta.Image))
}

func parseDateQuery(c *gin.Context) (time.Time, error) {
	yearStr := c.Query("year")
	monthStr := c.Query("month")
	dayStr := c.Query("day")

	if yearStr == "" && monthStr == "" && dayStr == "" {
		now := time.Now().In(timezone)
		return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, timezone), nil
	}

	year, errYear := strconv.Atoi(yearStr)
	month, errMonth := strconv.Atoi(monthStr)
	day, errDay := strconv.Atoi(dayStr)
	if errYear != nil || errMonth != nil || errDay != nil {
		return time.Time{}, errors.New("year, month and day must all be integers")
	}

	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, timezone)
	if date.Year() != year || int(date.Month()) != month || date.Day() != day {
		return time.Time{}, errors.New("invalid date")
	}
	return date, nil
}

func buildDayRotation(cfg rotationConfig, date time.Time) []rotationItem {
	dayStart := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, timezone)
	dayEnd := time.Date(date.Year(), date.Month(), date.Day()+1, 0, 0, 0, 0, timezone)
	period := time.Duration(cfg.LastTime) * time.Minute

	elapsed := dayStart.Sub(cfg.StandardTime)
	periodIndex := floorDiv(int64(elapsed), int64(period))
	cursor := cfg.StandardTime.Add(time.Duration(periodIndex) * period)

	items := make([]rotationItem, 0, 8)
	for cursor.Before(dayEnd) {
		segmentEnd := cursor.Add(period)
		if segmentEnd.After(dayStart) {
			start := cursor
			if start.Before(dayStart) {
				start = dayStart
			}
			end := segmentEnd
			if end.After(dayEnd) {
				end = dayEnd
			}

			mapIndex := periodIndex % int64(len(cfg.Maps))
			if mapIndex < 0 {
				mapIndex += int64(len(cfg.Maps))
			}
			mapID := cfg.Maps[mapIndex]

			item := rotationItem{
				ID:        len(items) + 1,
				MapID:     mapID,
				StartTime: formatClock(start, dayEnd),
				EndTime:   formatClock(end, dayEnd),
			}
			if meta, ok := maps[mapID]; ok {
				item.NameCN = meta.NameCN
				item.NameEN = meta.NameEN
			}
			items = append(items, item)
		}
		cursor = segmentEnd
		periodIndex++
	}
	return items
}

func formatClock(t, dayEnd time.Time) string {
	t = t.In(timezone)
	if t.Equal(dayEnd) {
		return "24:00"
	}
	return t.Format("15:04")
}

func floorDiv(a, b int64) int64 {
	q := a / b
	if a%b != 0 && (a < 0) != (b < 0) {
		q--
	}
	return q
}

func loadLocation(name string) *time.Location {
	loc, err := time.LoadLocation(name)
	if err != nil {
		log.Printf("warning: cannot load timezone %q (%v), using server local time", name, err)
		return time.Local
	}
	return loc
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if origin := c.GetHeader("Origin"); origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Header("Access-Control-Max-Age", "86400")
		}
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
