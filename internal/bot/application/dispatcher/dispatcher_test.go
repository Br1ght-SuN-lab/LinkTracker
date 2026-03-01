package dispatcher

import (
	"testing"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application/handlers"
)


func TestStart(t *testing.T) {
	d := New()
	
	reply, ok := d.Dispatch("/start")
	if !ok {
		t.Fatal("expected ok=true for /start command")
	}

	expected := handlers.StartHandler()
	if reply != expected {
		t.Fatalf("expected %s, got %s", expected, reply)
	}
}


func TestHelp(t *testing.T) {
	d := New()
	
	reply, ok := d.Dispatch("/help")
	if !ok {
		t.Fatal("expected ok=true for /help command")
	}

	expected := handlers.HelpHandler()
	if reply != expected {
		t.Fatalf("expected %s, got %s", expected, reply)
	}
}


func TestUnknow(t *testing.T) {
	d := New()
	
	reply, ok := d.Dispatch("/unknown")
	if !ok {
		t.Fatal("expected ok=true for /unknown command")
	}

	expected := handlers.UnknownHandler()
	if reply != expected {
		t.Fatalf("expected %s, got %s", expected, reply)
	}
}