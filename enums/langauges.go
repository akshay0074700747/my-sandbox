package enums

type Langage string

var Languages = struct {
	GOLANG Langage
	JAVA   Langage
	RUST   Langage
	PYTHON Langage
}{
	GOLANG: "GOLANG",
	JAVA:   "JAVA",
	RUST:   "RUST",
	PYTHON: "PYTHON",
}
