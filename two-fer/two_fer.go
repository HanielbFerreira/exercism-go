package twofer

//Change name's person, if name is empty, you is replaced
func ShareWith(name string) string {

	if name == "" {
		name = "you"
	}

	return "One for " + name + ", one for me."
}
