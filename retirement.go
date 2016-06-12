package main

import (
	"fmt"
	"github.com/joeriggs/retirement/social_security"
)

func main() {
	fmt.Printf("1985 returned %d.\n", ss_max_earnings.MaxEarnings(1985))
	fmt.Printf("1995 returned %d.\n", ss_max_earnings.MaxEarnings(1995))
	fmt.Printf("2005 returned %d.\n", ss_max_earnings.MaxEarnings(2005))
	fmt.Printf("2015 returned %d.\n", ss_max_earnings.MaxEarnings(2015))
	fmt.Printf("2025 returned %d.\n", ss_max_earnings.MaxEarnings(2025))
}

