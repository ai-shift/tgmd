package tgmd

import (
	"log/slog"
	"os"
	"slices"
	"strings"
	"testing"
)

func init() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
}

func TestRenderTelegram(t *testing.T) {
	input, _ := os.ReadFile("_testdata/input1.md")
	expected, _ := os.ReadFile("_testdata/expected1.tmd")
	got := Telegramify(string(input))
	gotLines := slices.Collect(strings.Lines(got))
	expectedLines := slices.Collect(strings.Lines(string(expected)))
	for i := 0; i < len(gotLines); i++ {
		if gotLines[i] != expectedLines[i] {
			t.Errorf("line %d differs\ngot     : '%s'\nexpected: '%s'\nFull:\n%s", i+1, gotLines[i], expectedLines[i], got)
			t.Fatal()
		}
	}
}
