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

type Extension string

var Extensions = struct {
	GOLANG Extension
	JAVA   Extension
	RUST   Extension
	PYTHON Extension
}{
	GOLANG: ".go",
	JAVA:   ".java",
	RUST:   ".rs",
	PYTHON: ".py",
}

type Command []string

var Commands = struct {
	GOLANG Command
	JAVA   Command
	RUST   Command
	PYTHON Command
}{
	GOLANG: Command{"go", "run", "/code/main.go"},
	JAVA:   Command{"sh", "-c", "javac /code/Main.java && java -cp /code Main"},
	RUST:   Command{"sh", "-c", "rustc -o /tmp/output /code/main.rs && /tmp/output"},
	PYTHON: Command{"python", "/code/main.py"},
}
