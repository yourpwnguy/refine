package common

import "github.com/yourpwnguy/gostyle"

var (
	// Some colored messages ( mainly prefix )
	G       = gostyle.New()
	Succfix = G.Inf()
	Errfix  = G.Err()
)
