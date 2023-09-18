package data

import (
	"strings"
)

// AttributeSet represents a set of OpenStreetMap attributes.
type AttributeSet struct {
	Bits uint64
}

// NewAttributeSet constructs an AttributeSet corresponding to the given bits value.
func NewAttributeSet(bits uint64) AttributeSet {
	return AttributeSet{Bits: bits}
}

// Of creates a AttributeSet containing the specified attributes.
func Of(attributes ...Attribute) AttributeSet {
	var bits uint64
	for _, a := range attributes {
		mask := uint64(1) << uint64(a.Ordinal)
		bits |= mask
	}
	return NewAttributeSet(bits)
}

// Contains checks if the AttributeSet contains the given attribute.
func (aset AttributeSet) Contains(attribute Attribute) bool {
	mask := uint64(1) << uint64(attribute.Ordinal)
	return (aset.Bits & mask) == mask
}

// Intersects checks if the intersection of two AttributeSets is not empty.
func (aset AttributeSet) Intersects(that AttributeSet) bool {
	return (aset.Bits & that.Bits) != 0
}

// String returns a string representation of the AttributeSet.
func (aset AttributeSet) String() string {
	var attrs []string
	for _, attr := range AllAttributes {
		if aset.Contains(attr) {
			attrs = append(attrs, attr.KeyValue)
		}
	}
	return "{" + strings.Join(attrs, ",") + "}"
}
