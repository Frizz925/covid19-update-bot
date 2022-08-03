package formatters

import (
	"fmt"
)

func IntToNumber(value int) string {
	head, tail := fmt.Sprintf("%d", value), ""
	for tmp := value; tmp >= 1000; tmp /= 1000 {
		idx := len(head) - 3
		tail = "," + head[idx:] + tail
		head = head[:idx]
	}
	return head + tail
}
