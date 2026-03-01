package dispatcher 

import (
	"testing"
)


func TestStart(t *testing.T) {
	d := New()
	d.Register("start", "desc", func() string { return "START_MSG" })

	reply, ok := d.Dispatch("start")
	if !ok {
		t.Fatal("expected ok=true for start command")
	}

	if reply != "START_MSG" {
		t.Fatalf("expected %q, got %q", "START_MSG", reply)
	}
}

func TestHelp(t *testing.T) {
	d := New()

	d.Register("help", "desc", func() string { return "HELP_MSG" })
	reply, ok := d.Dispatch("help")
	if !ok {
		t.Fatal("expected ok=true for /help command")
	}

	if reply != "HELP_MSG" {
		t.Fatalf("expected %s, got %s", "HELP_MSG", reply)
	}
}


func TestUnknow(t *testing.T) {
	d := New()
	
	reply, ok := d.Dispatch("unknown")
	if !ok {
		t.Fatal("expected ok=true for /unknown command")
	}

	if reply != "Неизвестная команда. Воспользуйтесь командой /help" {
		t.Fatalf("expected %s, got %s", "Неизвестная команда. Воспользуйтесь командой /help", reply)
	}
}