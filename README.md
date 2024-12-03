# ncat 🐾

## A Modern `cat` Implementation in Go 🚀

A blazing-fast, Go-powered reimplementation of the UNIX `cat` command-line tool.

<p align="center">
  <img src="https://raw.githubusercontent.com/AdiPat/ncat/refs/heads/main/assets/ncat_logo_art.png" alt="ncat logo" />
</p>

`ncat` brings the versatility of the traditional `cat` command to the modern era. Written in Go, it’s fast, efficient, and packed with extra features to level up your file-handling game.

---

## 📜 What Is `cat`?

> The `cat` command reads files sequentially, writing their contents to standard output.  
> It's commonly used to display, concatenate, and process files in UNIX-like systems.

---

## 🚀 Features

- **Go-Powered**: Fast, efficient, and robust for all your file-handling needs.
- **Flexible Input/Output**: Works seamlessly with files, standard input, and pipes.
- **Modern Enhancements**: New options like line numbering, reversing, and more.
- **Lightweight & Portable**: Compile once, run anywhere. 🖥️

---

## ⚡️ Quick Start

### 1️⃣ Build 🛠️

Compile the code using Go:

```bash
go build -o ncat
```

### 2️⃣ Run 🚀

Start using `ncat` just like the traditional `cat`:

```bash
./ncat file1.txt file2.txt
```

### 3️⃣ Combine Files 📝

Concatenate multiple files into one:

```bash
./ncat file1.txt file2.txt > combined.txt
```

### 4️⃣ Add Line Numbers 📖

Show line numbers alongside content:

```bash
./ncat -n file.txt
```

### 5️⃣ Run Tests 🧪

Ensure everything works perfectly by running the tests:

```bash
go test ./...
```

This will execute all test cases and confirm that your `ncat` implementation is reliable.

---

## 🌟 Why ncat?

`ncat` is more than just a reimplementation—it’s an upgrade:

- **🚀 Optimized Performance**: Harness the power of Go for lightning-fast operations.
- **🔧 Enhanced Features**: New flags like `--number`, `--reverse`, and more for better control.
- **📂 Large File Handling**: Efficiently process multi-GB files without breaking a sweat.

---

## 🛠️ Additional Features

- **Line Numbering (`-n`)**: Display line numbers alongside the content.
- **Reverse (`-r`)**: Output file content in reverse order.
- **Color Highlights (`--color`)**: Highlight specific patterns in the content.
- **Interactive Mode (`--interactive`)**: Browse through the file interactively.

---

## 🌟 Inspiration

This project was inspired by [John Crickett's Coding Challenges](https://codingchallenges.fyi/), a fantastic resource for developers to sharpen their skills by reimplementing classic tools like `cat`. The idea of modernizing `cat` while staying true to its core utility aligns with the philosophy of learning by building practical projects. Special thanks to John Crickett for creating such a thought-provoking repository of challenges!

---

`ncat` is your new favorite file companion. 🐾 Happy hacking! 🎉
