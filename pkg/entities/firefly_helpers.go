package entities

import "github.com/randaalex/finance_bot/pkg/firefly"

func NewNullableInt32(val *int32) *firefly.NullableInt32 {
	result := firefly.NewNullableInt32(val)
	if val == nil || *val == 0 {
		result.Unset()
	}

	return result
}

func NewNullableString(val *string) *firefly.NullableString {
	result := firefly.NewNullableString(val)
	if val == nil {
		result.Unset()
	}

	return result
}
