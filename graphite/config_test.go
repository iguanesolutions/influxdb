package graphite_test

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/influxdb/influxdb/httpd"
)

func TestConfig_Parse(t *testing.T) {
	// Parse configuration.
	var c httpd.Config
	if _, err := toml.Decode(`
bind-address = ":8080"
auth-enabled = true
log-enabled = true
write-tracing = true
pprof-enabled = true
`, &c); err != nil {
		t.Fatal(err)
	}

	// Validate configuration.
	if c.BindAddress != ":8080" {
		t.Fatalf("unexpected bind address: %s", c.BindAddress)
	} else if c.AuthEnabled != true {
		t.Fatalf("unexpected auth enabled: %v", c.AuthEnabled)
	} else if c.LogEnabled != true {
		t.Fatalf("unexpected log enabled: %v", c.LogEnabled)
	} else if c.WriteTracing != true {
		t.Fatalf("unexpected write tracing: %v", c.WriteTracing)
	} else if c.PprofEnabled != true {
		t.Fatalf("unexpected pprof enabled: %v", c.PprofEnabled)
	}
}
