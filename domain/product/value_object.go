package product

import (
	"fmt"
	"strings"
)

type SKU string

func NewSKU(raw string) (SKU, error) {
	normalized := strings.TrimSpace(strings.ToUpper(raw))
	if len(normalized) < 4 {
		return SKU(""), ErrSKUMustBeAtLeast4Chars
	}

	return SKU(normalized), nil
}

type Description string

func NewDescription(raw string) (Description, error) {
	trimmed := strings.TrimSpace(raw)

	// Maximum length of 3000 characters
	if len(trimmed) > 3000 {
		return Description(""), ErrDescriptionTooLong
	}

	return Description(trimmed), nil
}

type Money struct {
	Cents int
}

func (m Money) IsZero() bool {
	return m.Cents == 0
}

func (m Money) String() string {
	dollars := float64(m.Cents) / 100.0
	return fmt.Sprintf("$%.2f", dollars)
}

func (m Money) Add(other Money) Money {
	return Money{Cents: m.Cents + other.Cents}
}

func (m Money) Subtract(other Money) Money {
	return Money{Cents: m.Cents - other.Cents}
}
