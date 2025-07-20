package timePkg

import (
	"fmt"
	"time"
)

var (
	misures = [4]string{"ns","μs", "ms", "s"}
)
func Since(start time.Time) string {
	end := time.Since(start).Nanoseconds()
	index :=0 
	for end<1000 && index < 4 {
		end /= 1000
		index++
	} 
	return fmt.Sprintf("%d %s", end, misures[index])
}
