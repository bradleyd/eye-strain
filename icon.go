package main

import (
	"encoding/base64"
	"fmt"
)

const icon = `iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAAXNSR0IArs4c6QAAAERlWElmTU0AKgAAAAgAAYdpAAQAAAABAAAAGgAAAAAAA6ABAAMAAAABAAEAAKACAAQAAAABAAAAEKADAAQAAAABAAAAEAAAAAA0VXHyAAABWklEQVQ4Ee1QzUoCURQ+d+6dOyNjmsLoDBNlKRS6cOMjRC8QBG16AHftc9c2kHqG9u3qHQqCSClKk2oCR6McTRybube5E0R7tx44H+ec7/zc+wHMbWYFkNjQKJVo0fMyMPQWbYKJhfAIKO+hdnsgeJ5OJ0BL6zbzFyykBqDSj6YSOKVGY0pEg8nkzWsqn9YVX3sYuph9jSdbqvbct/KXnHF0JPPKxdhd7g8GioExO1zJj9c43Q1Hz8U8cD1nVBOpDmDMRSq8QGR+Z+bO7kPfUGN/dcFV48kWz6xmwxgkAdDrOFmQbnAQRKmAMlX4OsEHOSzViihaHHE4REMit+A8OaIQfSEUgjVV7WQIrHzlTUwLE3dHS9ThpdWkYdOebtZiEtq3fT9ZUdS37Zh2jNx38SqIRBSBMG4U9JHMs3GGPpH9+Ppb/UVuFZamiKXoN+qibiu6/p+fxzMo8AMvrHmZC6n4XwAAAABJRU5ErkJggg==`

func decodeIcon() []byte {
	decoded, err := base64.StdEncoding.DecodeString(icon)
	if err != nil {
		fmt.Println("decode error:", err)
		return []byte("")
	}
	return decoded
}
