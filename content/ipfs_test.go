package content

import "testing"

func TestIPFS(t *testing.T) {
	client := NewIPFSClient("10.235.64.47:5001")
	hash, err := client.AddContent("hello world")
	if err != nil {
		t.Fatalf("failed to add content: %v", err)
	}
	if hash != "Qmf412jQZiuVUtdgnB36FXFX7xg5V6KEbSJ4dpQuhkLyfD" {
		t.Fatalf("expect hash Qmf412jQZiuVUtdgnB36FXFX7xg5V6KEbSJ4dpQuhkLyfD but got %s", hash)
	}

	content, err := client.GetContent(hash)
	if err != nil {
		t.Fatalf("failed to get content for hash %s: %v", hash, err)
	}

	if content != "hello world" {
		t.Fatalf("expect hello world but got %s", content)
	}
}
