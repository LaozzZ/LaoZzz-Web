package main

import (
	"encoding/json"
	"os"
	"testing"
	"time"
)

func TestRotationScheduleClipsDailyBoundaries(t *testing.T) {
	t.Parallel()

	data, err := os.ReadFile("data/rotation.json")
	if err != nil {
		t.Fatalf("read data/rotation.json: %v", err)
	}

	var configs []rotationConfig
	if err := json.Unmarshal(data, &configs); err != nil {
		t.Fatalf("parse data/rotation.json: %v", err)
	}
	if len(configs) == 0 {
		t.Fatal("expected at least one rotation config")
	}

	cfg := configs[0]
	wantAnchor, err := time.Parse(time.RFC3339, "2026-09-21T02:30:00+08:00")
	if err != nil {
		t.Fatal(err)
	}
	if !cfg.StandardTime.Equal(wantAnchor) {
		t.Fatalf("standardTime = %s, want %s", cfg.StandardTime, wantAnchor)
	}
	if cfg.LastTime != 270 {
		t.Fatalf("lastTime = %d, want 270 minutes", cfg.LastTime)
	}

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		t.Fatal(err)
	}
	date := time.Date(2026, time.October, 22, 0, 0, 0, 0, loc)
	items := buildDayRotation(cfg, date)

	want := []struct {
		mapID int
		start string
		end   string
	}{
		{mapID: 3, start: "00:00", end: "01:00"},
		{mapID: 2, start: "01:00", end: "05:30"},
		{mapID: 6, start: "05:30", end: "10:00"},
		{mapID: 3, start: "10:00", end: "14:30"},
		{mapID: 2, start: "14:30", end: "19:00"},
		{mapID: 6, start: "19:00", end: "23:30"},
		{mapID: 3, start: "23:30", end: "24:00"},
	}
	if len(items) != len(want) {
		t.Fatalf("rotation item count = %d, want %d", len(items), len(want))
	}

	for i, expected := range want {
		item := items[i]
		if item.MapID != expected.mapID || item.StartTime != expected.start || item.EndTime != expected.end {
			t.Errorf(
				"item %d = mapId %d %s-%s, want mapId %d %s-%s",
				i+1,
				item.MapID,
				item.StartTime,
				item.EndTime,
				expected.mapID,
				expected.start,
				expected.end,
			)
		}
	}
	if items[0].NameCN != "残月" || items[0].NameEN != "Broken Moon" {
		t.Errorf("first map = %s (%s), want 残月 (Broken Moon)", items[0].NameCN, items[0].NameEN)
	}

	previousDayItems := buildDayRotation(cfg, date.AddDate(0, 0, -1))
	previousDayLast := previousDayItems[len(previousDayItems)-1]
	if previousDayLast.MapID != 3 || previousDayLast.StartTime != "20:30" || previousDayLast.EndTime != "24:00" {
		t.Errorf(
			"previous day last interval = mapId %d %s-%s, want mapId 3 20:30-24:00",
			previousDayLast.MapID,
			previousDayLast.StartTime,
			previousDayLast.EndTime,
		)
	}
}
