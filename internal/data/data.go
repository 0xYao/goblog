package data

import "fmt"

var (
	RandomImageUrl = "https://placeimg.com/256/256/any"
)

func NewRandomImageUrl(width, height int) string {
	return fmt.Sprintf("https://placeimg.com/%d/%d/any", width, height)
}
