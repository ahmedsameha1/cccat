# cccat

**cccat** is a custom implementation of the Unix `cat` command,  written in **Standard ML (SML)**. This implementation does not address all the `cat` command options.

This project was created as a solution to the ["Build Your Own cat Tool"](https://codingchallenges.fyi/challenges/challenge-cat/) coding challenge. After doing the [Programming Languages, Part A](https://www.coursera.org/learn/programming-languages) course and solving multiple problems using the Standard ML programming language, I thought that it would be good if I did a project using the Standard ML programming language. So I did that challenge.

## 🚀 Features

- Read and output the contents of a single file.
- Concatenate and display the contents of multiple files.
- Read from standard input (stdin) when no files are provided or when `-` is used.
- (Optional) Line numbering flags (depending on implementation status):
    - `-n`: Number all output lines.
    - `-b`: Number nonempty output lines.

## 📋 Prerequisites

To build and run this project, you will need:

* **[Moscow ML](https://mosml.org/)**: An implementation of Standard ML used to compile the source code.
* **[Go](https://go.dev/)**: Required to run the end-to-end test suite.

## 🛠️ Installation & Building

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/ahmedsameha1/cccat.git](https://github.com/ahmedsameha1/cccat.git)
    cd cccat
    ```

2.  **Compile the source code:**
    Use the Moscow ML compiler (`mosmlc`) to build the executable.
    ```bash
    mosmlc -o cccat cccat.sml
    ```
    *This command creates an executable named `cccat` (or `cccat.exe` on Windows).*

## 📖 Usage

After compiling, you can use the tool just like the standard `cat` command.

### Read a single file
```bash
./cccat hello.txt
```

### Concatenate multiple files
```bash
./cccat hello.txt quotes.txt
```

### Read from Standard Input
```
./cccat
./cccat -
echo "Hello World" | ./cccat
```

## 🧪 Testing

This project uses Go for end-to-end integration testing.

### To run the tests:
```
go test ./...
```

## 📂 Project Structure

```text
.
├── endtoendtests/      # Directory containing Go end-to-end tests
├── cccat.sml           # Main application source code (Standard ML)
├── go.mod              # Go module definition for test dependencies
├── go.sum              # Go checksums for dependencies
├── *.txt               # Sample text files used for testing.
└── README.md
```

## ⚠️ Known Issues

* **Line Numbering:** The line numbering flags (`-n`, `-b`) currently behave incorrectly when using this command with multiple files.