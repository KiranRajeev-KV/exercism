package twofer

// ShareWith returns two fer lines.
func ShareWith(name string) string {
	if name == "" {
        return "One for you, one for me."
    } else {
        return "One for " + name + ", one for me."
    }
}
