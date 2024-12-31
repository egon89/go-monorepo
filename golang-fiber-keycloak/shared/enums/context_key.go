package enums

// This custom type is often used to define strongly-typed keys for contexts,
// which helps avoid accidental key collisions in Go's context.Context.
type contextKey int

// iota is a special identifier used within constant declarations. It starts at 0 and increments by 1 for each constant in the block.
const (
	ContextKeyRequestId contextKey = iota // 0
	ContextKeyClaims    contextKey = iota // 1
)
