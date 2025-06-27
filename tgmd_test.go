package tgmd

import (
	"log/slog"
	"os"
	"slices"
	"strings"
	"testing"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

func TestRenderE2E(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TEST_BOT_TOKEN"))
	receiverID, _ := strconv.ParseInt(os.Getenv("RECEIVER_PEER_ID"), 10, 64)
	if err != nil {
		t.Errorf("failed to create bot API with %s", err)
		t.Fatal()
	}

	input, _ := os.ReadFile("_testdata/input1.md")
	got := Telegramify(string(input))
	msg := tgbotapi.NewMessage(receiverID, got)
	msg.ParseMode = "MarkdownV2"
	_, err = bot.Send(msg)
	if err != nil {
		t.Errorf("failed to send telegram messge with %s\nMessage: %s", err, got)
	}
}
