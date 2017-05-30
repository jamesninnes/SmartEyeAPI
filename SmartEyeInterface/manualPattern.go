package SmartEyeInterface

import "time"

type manualPattern struct {
	name      string    `json:"name"`
	pattern	  string    `json:"pattern"`
	created   time.Time `json:"created"`
}

func NewManualPattern(name string, pattern string) manualPattern {
	var currManualPattern manualPattern

	currManualPattern.name = name
	currManualPattern.pattern = pattern
	currManualPattern.created = time.Now()

	return currManualPattern
}

