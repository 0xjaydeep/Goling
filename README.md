# Golings — Go Exercises for Beginners (Rustlings-style)

> **Learn Go by fixing small, focused programs — one exercise at a time.**
> The hands-on Go practice project inspired by [Rustlings](https://github.com/rust-lang/rustlings).

[![Go Version](https://img.shields.io/badge/go-1.21+-00ADD8?logo=go&logoColor=white)](https://go.dev/dl/)
[![Exercises](https://img.shields.io/badge/exercises-6-brightgreen)]()
[![Topics](https://img.shields.io/badge/topics-14-blue)]()
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

---

**Golings** is a collection of small, self-contained Go exercises that teach you the
Go programming language through *doing* rather than reading. If you have ever searched
for *golang exercises*, *go programming practice problems*, *rustlings for Go*,
or *learn Go by doing* — this is the project for you.

Each exercise gives you:
- A clear **problem statement** in the file comments
- **Function signatures** already written for you
- `TODO` comments showing exactly what to implement
- A `solution.go` you can check after a genuine attempt

---

## Table of Contents

- [Who Is This For?](#who-is-this-for)
- [Prerequisites](#prerequisites)
- [Quick Start](#quick-start)
- [How Golings Works](#how-golings-works)
- [Topic Reference (01–14)](#topic-reference-0114)
- [Exercises (Hands-On Coding)](#exercises-hands-on-coding)
- [Check Your Progress](#check-your-progress)
- [Hints](#hints)
- [Recommended Learning Path](#recommended-learning-path)
- [FAQ](#faq)
- [What to Learn Next](#what-to-learn-next)

---

## Who Is This For?

Golings is designed for:

- **Developers switching to Go** from Python, JavaScript, Java, Rust, or any other language
- **Absolute beginners** who learn best by writing code rather than reading documentation
- **Rustlings users** who want the same small-exercise experience but for Go
- Anyone who has Googled *go lang exercises*, *golang beginner practice*, or *how to learn golang fast*

You do **not** need prior Go experience. You need basic programming knowledge (variables, loops, functions in any language) and Go installed.

---

## Prerequisites

| Requirement | Version | Link |
|-------------|---------|------|
| Go | 1.21 or newer | [go.dev/dl](https://go.dev/dl/) |
| A text editor | any | VS Code + [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.go) recommended |

Verify your installation:

```bash
go version
# go version go1.21.x ...
```

No external libraries. No `go get`. Every exercise uses only the Go standard library.

---

## Quick Start

```bash
# 1. Clone the repository
git clone https://github.com/yourname/golings.git
cd golings

# 2. See your current progress
bash verify.sh

# 3. Start at the beginning — read and run the first topic
cd 01-hello-world
go run main.go

# 4. When you feel ready, tackle the first exercise
cd ../exercises/01-calculator
go run main.go        # run your work-in-progress
go run solution.go    # reveal the full answer (only after trying!)
```

---

## How Golings Works

Golings has two layers that work together:

### Layer 1 — Topics (read + run + experiment)

Folders `01-hello-world` through `14-packages` each contain a complete, fully runnable
`main.go` that demonstrates one Go concept from scratch. Every example prints output so
you can verify it works immediately.

**The topic loop:**
1. Open `main.go` and read the code top-to-bottom
2. Run it: `go run main.go`
3. Modify values — break things intentionally — read the compiler errors
4. Restore it and move on

### Layer 2 — Exercises (implement + verify)

Folders under `exercises/` give you a problem to solve. The starter `main.go` has:
- A problem description in comments at the top
- Function signatures with correct parameter and return types already written
- `// TODO` comments pointing to exactly what logic to add
- A `main()` that calls every function so you can test as you go

```
exercises/
└── 01-calculator/
    ├── main.go       ← you edit this
    └── solution.go   ← //go:build ignore — run only with: go run solution.go
```

**The exercise loop:**
1. Read the problem statement at the top of `main.go`
2. Implement each `TODO` one at a time
3. Run: `go run main.go` — fix compiler errors and wrong output
4. Stuck for 10+ minutes? Check `HINTS.md` for a nudge
5. Still stuck? Run `go run solution.go` to see the full working solution

> **The Golings rule:** spend at least 10 minutes on each TODO before looking at hints,
> and at least 20 minutes before running the solution. The struggle is the learning.

---

## Topic Reference (01–14)

Each topic is a standalone Go module. `cd` into the folder and run `go run main.go`.

| # | Folder | What you will learn |
|---|--------|---------------------|
| 01 | [01-hello-world](01-hello-world/) | `package main`, `import`, `fmt.Println`, `fmt.Printf`, format verbs (`%s %d %f %v %T`), zero values |
| 02 | [02-variables](02-variables/) | `var` with type, `var` inferred, `:=` short declaration, multiple assignment, swap, explicit type conversion |
| 03 | [03-constants](03-constants/) | `const`, typed vs untyped constants, `iota` for enums (weekdays), `iota` with bit-shift (KB/MB/GB) |
| 04 | [04-control-flow](04-control-flow/) | `if/else`, `if` with init statement, `switch` on value, `switch` without condition, `fallthrough` |
| 05 | [05-loops](05-loops/) | traditional `for`, while-style `for`, infinite `for` with `break`, `range` over slice / map / string, `continue` |
| 06 | [06-functions](06-functions/) | shorthand params, multiple return values, named returns + naked return, variadic `...`, spread operator, function variables, higher-order functions, closures |
| 07 | [07-arrays-slices](07-arrays-slices/) | array as value type, slice literal, `make(len, cap)`, `append`, slice expressions, shared backing-array gotcha, `copy()`, nil slice |
| 08 | [08-maps](08-maps/) | map literal, `make`, zero value for missing key, two-value form `val, ok`, `delete`, `range`, nested maps |
| 09 | [09-structs-methods](09-structs-methods/) | struct definition, named/zero-value init, pointer to struct, embedded struct (composition), anonymous struct, value vs pointer receiver, `String()` / `fmt.Stringer` |
| 10 | [10-pointers](10-pointers/) | `&` address-of, `*` dereference, pass-by-value vs pass-by-pointer, `new()`, nil pointer |
| 11 | [11-interfaces](11-interfaces/) | implicit satisfaction, polymorphism, type assertion with `ok`, type switch, empty interface `any` |
| 12 | [12-error-handling](12-error-handling/) | `errors.New`, sentinel errors, custom error type, `fmt.Errorf` with `%w`, `errors.Is`, `errors.As` |
| 13 | [13-defer-panic-recover](13-defer-panic-recover/) | `defer` LIFO order, `defer` for cleanup, `panic`, `recover` in deferred function |
| 14 | [14-packages](14-packages/) | package declaration, exported vs unexported identifiers, local package import |

---

## Exercises (Hands-On Coding)

Six progressively harder exercises. Each builds on concepts from the topics above.

| # | Exercise | Practises | Estimated difficulty |
|---|----------|-----------|----------------------|
| [01](exercises/01-calculator/) | **Calculator** | `switch`, error returns | ★☆☆ beginner |
| [02](exercises/02-phonebook/) | **Phonebook** | `map`, `struct`, CRUD operations | ★☆☆ beginner |
| [03](exercises/03-word-frequency/) | **Word Frequency** | `map`, `sort.Slice`, string processing | ★★☆ intermediate |
| [04](exercises/04-student-grades/) | **Student Grades** | `interface`, sorting, polymorphism | ★★☆ intermediate |
| [05](exercises/05-safe-stack/) | **Safe Stack** | pointer receivers, sentinel errors | ★★☆ intermediate |
| [06](exercises/06-mini-bank/) | **Mini Bank** | custom error types, method design, history | ★★★ challenging |

### Running an exercise

```bash
cd exercises/01-calculator
go run main.go        # test your implementation
go run solution.go    # see the complete solution
```

### Why does `go run solution.go` work separately?

Every `solution.go` has `//go:build ignore` at line 1. This tells the Go toolchain
to skip the file during a normal build. It only compiles when you name it explicitly.

---

## Check Your Progress

```bash
bash verify.sh
```

The script builds and runs every exercise, then reports:

```
╔══════════════════════════════════════╗
║        Golings Progress Check        ║
╚══════════════════════════════════════╝

  ✓ PASS   01-calculator
  ○ TODO   02-phonebook
  ○ TODO   03-word-frequency
  ○ TODO   04-student-grades
  ○ TODO   05-safe-stack
  ○ TODO   06-mini-bank

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  Done: 1   Todo: 5   Fail: 0
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
```

- **PASS** — compiles and runs without errors, no unfinished stubs
- **TODO** — still has `not implemented` stubs in the file
- **FAIL** — code does not compile or panics at runtime

---

## Hints

Stuck on an exercise? Before looking at the solution, read the relevant hint in
[HINTS.md](HINTS.md). Hints give you a directional nudge without spoiling the answer.

---

## Recommended Learning Path

Follow topics in order, then do the matching exercise before moving on:

```
01 Hello World      ──▶  run it, read the output
02 Variables        ──┐
03 Constants        ──┘  ▶ Exercise 01: Calculator

04 Control Flow     ──┐
05 Loops            ──┘  ▶ Exercise 02: Phonebook

06 Functions        ──▶  Exercise 03: Word Frequency

07 Arrays & Slices  ──┐
08 Maps             ──┘  ▶ (revisit Exercise 02 with maps fresh in mind)

09 Structs          ──┐
10 Pointers         ──┘  ▶ Exercise 05: Safe Stack

11 Interfaces       ──▶  Exercise 04: Student Grades

12 Error Handling   ──┐
13 defer/panic      ──┘  ▶ Exercise 06: Mini Bank

14 Packages         ──▶  read, then write your own package
```

Total time estimate: 6–12 hours depending on your background.

---

## FAQ

**Q: Is this like Rustlings but for Go?**

Yes. The philosophy is identical — small files, TODO stubs, solution files, no framework.
The workflow (read → run → break → fix → exercise → solution) is directly inspired by
Rustlings. If you enjoyed Rustlings for Rust, Golings gives you the same experience for Go.

---

**Q: Do I need to install any dependencies?**

No. `go run main.go` is all you ever need. Every exercise uses only the Go standard
library. There is no `go get`, no vendor directory, no Docker.

---

**Q: I come from Python / JavaScript / Java — is Go hard to pick up?**

Go is deliberately simple. It has no classes (only structs + methods), no inheritance,
no generics complexity at the beginner level, and one way to do most things. Most
developers are writing useful Go within a day.

---

**Q: Why are the topic files in separate modules (each with their own go.mod)?**

Isolation. You can `cd` into any folder and run it independently without anything else
needing to be set up. Each folder is its own mini project, just like a real Go module.

---

**Q: Can I use this to prepare for a Go interview?**

The topics cover all the basics interviewers test: data structures (slices, maps, structs),
interfaces, error handling, pointers, and closures. After completing Golings, work through
common Go interview questions on concurrency and goroutines as a natural next step.

---

**Q: Something is wrong or unclear — how do I report it?**

Open an issue on GitHub. Bug reports, unclear exercise descriptions, and better hint
suggestions are all welcome.

---

## What to Learn Next

After finishing all 14 topics and 6 exercises, you have a solid Go foundation. Level up with:

| Topic | Why it matters |
|-------|----------------|
| Goroutines & Channels | Go's unique concurrency model — the main reason people choose Go |
| `go test` & table-driven tests | How every real Go project tests its code |
| `net/http` | Build web servers in ~10 lines |
| `encoding/json` | Marshalling/unmarshalling — used in virtually every Go service |
| Generics (Go 1.18+) | Write reusable functions without sacrificing type safety |
| `context.Context` | Cancel long-running operations — essential for real applications |

---

## Contributing

Pull requests are welcome. To add an exercise:
1. Create a new folder under `exercises/` following the existing naming pattern
2. Add `main.go` (starter with TODOs) and `solution.go` (with `//go:build ignore`)
3. Add a `go.mod` with `module go-basics-practice/exercises/<name>` and `go 1.21`
4. Add a hint entry to `HINTS.md`
5. Open a PR with a description of what Go concept the exercise targets

---

## License

MIT — free to use, fork, and share.
