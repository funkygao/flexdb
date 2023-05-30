package dto

// H is a universal data container.
type H map[string]interface{}

// S returns string value of a key. if no present, returns empty string.
func (h H) S(k string) string {
	if v, present := h[k]; present {
		return v.(string)
	}

	return ""
}

// B returns boolean value of a key.
func (h H) B(k string) bool {
	if v, present := h[k]; present {
		return v.(bool)
	}

	return false
}

// V returns value of a key and whether the key exists.
func (h H) V(k string) (interface{}, bool) {
	v, present := h[k]
	return v, present
}
