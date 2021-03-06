package modules

import (
	"testing"
	//"os"
)

func TestParser(t *testing.T) {
	s := `
	@db.set "name"="john" "age" += 1 "cars" -= 2 in /tmp/users`

	p := NewParser()
	l, err := p.Parse(s)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(l.ToList())
}