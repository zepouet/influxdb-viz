package models

import (
	"fmt"
	"testing"
	"encoding/json"
)

func TestUnmarshalBubble(t *testing.T) {
	bubble = &Bubble{}
	res1B, _ := json.Marshal(bubble)
	fmt.Println(string(res1B))
}

