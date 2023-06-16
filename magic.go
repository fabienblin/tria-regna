package main

import (

)

const (
	RED int = 1
	BLUE int = 2
	GREEN int = 3
)

const (
	WEAK float64 = 1.5
	EQUAL float64 = 1.0
	RESIST float64 = 0.5
)

/**
 * Get weakness coefficient on magic damage
 * 
 * @param active is the attacking magic, use RED | BLUE | GREEN
 * @param passive is the defending magic, use RED | BLUE | GREEN
 * @return coefficient to apply on magic damage 
 */
func weaknesses(active int, passive int) float64 {
	if active == RED {
		if passive == RED {
			return EQUAL
		} else if passive == BLUE {
			return RESIST
		} else if passive == GREEN {
			return WEAK
		}
	} else if active == BLUE {
		if passive == RED {
			return WEAK
		} else if passive == BLUE {
			return EQUAL
		} else if passive == GREEN {
			return RESIST
		}
	} else if active == GREEN {
		if passive == RED {
			return RESIST
		} else if passive == BLUE {
			return WEAK
		} else if passive == GREEN {
			return EQUAL
		}
	}
	return 0
}