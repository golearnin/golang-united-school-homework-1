package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	wd := emoji.Sprint("Hello :world_map:")
	return wd
}
