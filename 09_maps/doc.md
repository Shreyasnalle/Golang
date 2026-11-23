# 09 - Maps

A **map** stores **key → value** pairs, like a Python `dict` or a hashmap. Lookup is by key, not a linear search, so it's fast.

---

## 1. Declaring & Initialising

```go
var m map[string]vertex        // declaration only → m is nil
m = make(map[string]vertex)    // allocation (must happen before writing)
m["Bell Labs"] = vertex{40.12, -75.67}

var mapping = map[string]vertex{        // declare + initialise in one go
    "bitcoin": {58000.12, 120000.15},   // value type can be omitted inside the literal
}
```

---

## 2. Operations

| Operation | Syntax |
| :--- | :--- |
| insert / update | `m[key] = value` |
| read | `v := m[key]` (missing key → **zero value**) |
| delete | `delete(m, key)` |
| exists? | `v, ok := m[key]` |
| size | `len(m)` |

### The "comma ok" idiom
```go
if qty, ok := stock["apple"]; ok {
    // key exists (even if qty is 0)
}
```
You need `ok` because a missing key also returns `0`, and `0` could be a real stored value.

---

## 3. Iterating: Order Is Random

```go
for key, value := range ages { }   // random order on every run
```
For a fixed order, collect the keys, `sort.Strings(keys)`, then loop over the keys.

---

## 4. Nil Map

* `var m map[string]int` → `m == nil`
* **Reading** a nil map works (gives the zero value).
* **Writing** to a nil map → **panic: assignment to entry in nil map**. Always `make` it first.

---

## 5. Common Patterns

```go
counts[word]++                                 // word count: missing key starts at 0
skills["shreyas"] = append(skills["shreyas"], "go") // map of slices
```

---

## 6. Map vs Struct

| | Struct | Map |
| :--- | :--- | :--- |
| Keys/fields | fixed, known at **compile time** | added/removed at **runtime** |
| Access | dot: `c.host` | key: `m["host"]` |
| Value types | each field its own type | all values have **one** type |
| Loop over it? | no | yes, with `range` |
| Use for | fixed, structured data (a user, a JSON payload) | lookups, dynamic data, unknown keys |

---

## 7. Files in This Folder

| File | What it covers |
| :--- | :--- |
| `maps.go` | declare vs make, literal, insert/update/delete, comma ok, map vs struct notes |
| `maps_more.go` | short literals, `len`, random vs sorted iteration, comma ok in `if`, nil map, word count, map of slices, map vs struct in code |
| `main.go` | runs every example |

> Fix made: `mapping_in_declaration` printed `m` instead of `mapping`. This folder used to be `slices/map.go`.

## 8. How to Run

```bash
go run ./09_maps
```

## 9. Revision Checklist

- [ ] What happens when you write to a map declared only with `var`?
- [ ] What does reading a missing key return?
- [ ] Why do we need `v, ok := m[k]`?
- [ ] Is map iteration order guaranteed?
- [ ] When would you pick a struct over a map?
