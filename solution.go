package solution

import (
	emoji "github.com/kyokomi/emoji"
)

func GetMessage() string {
	return emoji.Sprint("Hello :world_map:!")
}
