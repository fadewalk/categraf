package filter


import (
	"testing"
)

func TestCompile_EmptyInput_ReturnsNil(t *testing.T) {
	f, err := Compile([]string{})
	if err != nil {
		t.Errorf("Compile returned an error: %v", err)
	}
	if f != nil {
		t.Errorf("Compile should return nil for empty input")
	}
}

func TestCompile_NoWildcards_ReturnsSimpleFilter(t *testing.T) {
	f, err := Compile([]string{"cpu", "mem"})
	if err != nil {
		t.Errorf("Compile returned an error: %v", err)
	}
	if !f.Match("cpu") {
		t.Errorf("Expected 'cpu' to match")
	}
	if !f.Match("mem") {
		t.Errorf("Expected 'mem' to match")
	}
	if f.Match("network") {
		t.Errorf("Expected 'network' not to match")
	}
}

func TestCompile_SingleWildcard_ReturnsGlobFilter(t *testing.T) {
	f, err := Compile([]string{"net*"})
	if err != nil {
		t.Errorf("Compile returned an error: %v", err)
	}
	if !f.Match("network") {
		t.Errorf("Expected 'network' to match")
	}
	if f.Match("memory") {
		t.Errorf("Expected 'memory' not to match")
	}
}

func TestCompile_MultipleWildcards_ReturnsGlobFilter(t *testing.T) {
	f, err := Compile([]string{"net*", "cpu*"})
	if err != nil {
		t.Errorf("Compile returned an error: %v", err)
	}
	if !f.Match("network") {
		t.Errorf("Expected 'network' to match")
	}
	if !f.Match("cpu0") {
		t.Errorf("Expected 'cpu0' to match")
	}
	if f.Match("memory") {
		t.Errorf("Expected 'memory' not to match")
	}
}

func TestCompile_MixedFilters_ReturnsGlobFilter(t *testing.T) {
	f, err := Compile([]string{"cpu", "mem", "net*"})
	if err != nil {
		t.Errorf("Compile returned an error: %v", err)
	}
	if !f.Match("cpu") {
		t.Errorf("Expected 'cpu' to match")
	}
	if !f.Match("network") {
		t.Errorf("Expected 'network' to match")
	}
	if f.Match("memory") {
		t.Errorf("Expected 'memory' not to match")
	}
}