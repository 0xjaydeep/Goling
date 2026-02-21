# Hints

Use these when you are genuinely stuck — after at least 10 minutes of trying.
Each hint nudges you in the right direction without giving away the full answer.
Run `go run solution.go` only as a last resort.

---

## Exercise 01 — Calculator

<details>
<summary>Hint 1 — overall structure</summary>

Use a `switch` statement on the `op` string, not a chain of `if/else`.
Each `case` should do the arithmetic and `return` the result with `nil` error.

```go
switch op {
case "+":
    return a + b, nil
// ...
}
```

</details>

<details>
<summary>Hint 2 — division by zero</summary>

Inside the `"/"` case, check `b == 0` *before* dividing.
Return the zero value `0` and an error created with `fmt.Errorf`.

```go
case "/":
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
```

</details>

<details>
<summary>Hint 3 — unknown operator</summary>

After all valid cases, add a `default` that returns an error.
Use `%q` in the format string to quote the operator for clarity.

```go
default:
    return 0, fmt.Errorf("unknown operator %q", op)
```

</details>

---

## Exercise 02 — Phonebook

<details>
<summary>Hint 1 — addContact</summary>

A map is like a dictionary. Store the contact using its name as the key:

```go
book[c.Name] = c
```

</details>

<details>
<summary>Hint 2 — findContact (two-value map lookup)</summary>

Always use the two-value form when you need to distinguish "not found" from
a stored zero value:

```go
c, ok := book[name]
return c, ok
```

</details>

<details>
<summary>Hint 3 — deleteContact</summary>

Go has a built-in `delete` function. Deleting a key that doesn't exist is a no-op.

```go
delete(book, name)
```

</details>

<details>
<summary>Hint 4 — listContacts</summary>

Check `len(book) == 0` first. Then `range` over the map:

```go
for _, c := range book {
    fmt.Printf("  %-8s  %s\n", c.Name, c.Phone)
}
```

</details>

---

## Exercise 03 — Word Frequency

<details>
<summary>Hint 1 — countWords: splitting the text</summary>

`strings.Fields` splits on any whitespace (spaces, tabs, newlines) and handles
multiple consecutive spaces automatically. It is better than `strings.Split(text, " ")`.

```go
for _, word := range strings.Fields(text) { ... }
```

</details>

<details>
<summary>Hint 2 — countWords: normalising words</summary>

Apply transformations in this order — lowercase first, then strip punctuation:

```go
word = strings.ToLower(word)
word = strings.Trim(word, ".,!?;:\"'")
if word == "" {
    continue
}
freq[word]++
```

</details>

<details>
<summary>Hint 3 — topN: building the word slice</summary>

You cannot range over a map and sort it in place. Extract keys into a slice first:

```go
words := make([]string, 0, len(freq))
for w := range freq {
    words = append(words, w)
}
```

</details>

<details>
<summary>Hint 4 — topN: the sort comparator</summary>

`sort.Slice` takes a "less" function. Primary: higher count comes first.
Tie-break: alphabetical order (so results are deterministic):

```go
sort.Slice(words, func(i, j int) bool {
    if freq[words[i]] != freq[words[j]] {
        return freq[words[i]] > freq[words[j]]
    }
    return words[i] < words[j]
})
```

</details>

---

## Exercise 04 — Student Grades

<details>
<summary>Hint 1 — Score(): averaging grades</summary>

Guard against an empty slice or you will get a divide-by-zero:

```go
if len(s.Grades) == 0 {
    return 0
}
sum := 0.0
for _, g := range s.Grades {
    sum += g
}
return sum / float64(len(s.Grades))
```

</details>

<details>
<summary>Hint 2 — leaderboard: sorting</summary>

Make a copy of the slice before sorting so you do not mutate the caller's data.
Then use `sort.Slice` with a comparator that puts higher scores first:

```go
sorted := make([]Rankable, len(students))
copy(sorted, students)
sort.Slice(sorted, func(i, j int) bool {
    return sorted[i].Score() > sorted[j].Score()
})
```

</details>

<details>
<summary>Hint 3 — leaderboard: printing</summary>

Use `%-10s` to left-pad the name column for aligned output:

```go
fmt.Printf("  %d. %-10s — %.2f\n", i+1, s.GetName(), s.Score())
```

</details>

---

## Exercise 05 — Safe Stack

<details>
<summary>Hint 1 — the Stack struct field</summary>

A slice is the natural backing store for a stack:

```go
type Stack struct {
    items []int
}
```

</details>

<details>
<summary>Hint 2 — Push</summary>

`append` is all you need. The slice grows automatically:

```go
s.items = append(s.items, val)
```

</details>

<details>
<summary>Hint 3 — Pop: reading and shrinking the slice</summary>

The "top" of the stack is the last element. Shrink by re-slicing:

```go
top := s.items[len(s.items)-1]
s.items = s.items[:len(s.items)-1]
return top, nil
```

Always call `IsEmpty()` first and return `ErrEmptyStack` if true.

</details>

<details>
<summary>Hint 4 — Peek vs Pop</summary>

`Peek` reads the top element without modifying the slice.
`Pop` reads *and* removes it. The only difference is whether you re-slice.

</details>

---

## Exercise 06 — Mini Bank

<details>
<summary>Hint 1 — custom error types</summary>

Any type with an `Error() string` method satisfies the `error` interface.
Use pointer receivers so `errors.As` can match them:

```go
func (e *InsufficientFundsError) Error() string {
    return fmt.Sprintf("insufficient funds: have $%.2f, need $%.2f",
        e.Available, e.Requested)
}
```

</details>

<details>
<summary>Hint 2 — NewAccount: recording the first transaction</summary>

If `initialBalance > 0`, append a Transaction so it shows up in the statement:

```go
a.History = append(a.History, Transaction{
    Type:    "initial deposit",
    Amount:  initialBalance,
    Balance: initialBalance,
})
```

</details>

<details>
<summary>Hint 3 — Deposit and Withdraw: validation order</summary>

Always validate first, then modify state:
1. Check `amount <= 0` → `InvalidAmountError`
2. (Withdraw only) Check `amount > a.Balance` → `InsufficientFundsError`
3. Update `a.Balance`
4. Append to `a.History`

</details>

<details>
<summary>Hint 4 — PrintStatement: formatting</summary>

Use `%-20s` to left-align the transaction type column:

```go
fmt.Printf("  %-20s $%8.2f  → balance $%.2f\n",
    t.Type, t.Amount, t.Balance)
```

</details>
