package middlewares

import (
	"sync"
)

var limiters = sync.Map{}
