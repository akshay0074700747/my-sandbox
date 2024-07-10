package enums

type Container string

var Containers = struct {
	GOLANG Container
	JAVA   Container
	RUST   Container
	PYTHON Container
}{
	GOLANG: "golang:alpine",
	JAVA:   "openjdk:alpine",
	RUST:   "rust:alpine",
	PYTHON: "python:alpine",
}

type Bind string

var Binds = struct {
	GOLANG Bind
	JAVA   Bind
	RUST   Bind
	PYTHON Bind
}{
	GOLANG: "%s:/code/main.go",
	JAVA:   "%s:/code/main.rs",
	RUST:   "%s:/code/Main.java",
	PYTHON: "%s:/code/main.py",
}