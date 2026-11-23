# 06 - Structs

A **struct** groups named fields, possibly of **different types**, into one value. Go has no classes; structs plus methods (see `11_methods`) fill that role.

---

## 1. Array vs Struct

| | Array | Struct |
| :--- | :--- | :--- |
| Element types | all the **same** type | each field can be a **different** type |
| Access | by index `a[0]` | by name `v.x` |
| Closest Python idea | fixed-size list | dataclass / class with only fields |

---

## 2. Declaring & Using

```go
type vertex struct {
    x int
    y int
}

v := vertex{1, 2}  // positional literal
v.x = 10           // dot to read/write a field
```

### Literal forms (`struct.go`)
```go
v1 = vertex{1, 2}    // {1 2}
v2 = vertex{x: 1}    // {1 0}, missing field gets its zero value
v3 = vertex{}        // {0 0}
p  = &vertex{1, 2}   // &{1 2}, pointer to a new struct
```

---

## 3. Pointer to Struct

```go
v := vertex{1, 5}
p := &v
p.y = 75      // shortcut for (*p).y = 75, Go dereferences automatically
fmt.Println(v) // {1 75}
```

---

## 4. Nested & Anonymous Structs

```go
type student struct {
    name    string
    address address   // nested struct
}
s.address.city        // chain the dots

point := struct{ x, y int }{3, 4}   // anonymous: no type name, used once
```

---

## 5. Comparing

`a == b` is `true` when **every field** is equal (works when all fields are comparable).

---

## 6. Exported Fields & JSON

* A field/type starting with a **capital letter** is exported and visible to other packages.
* `encoding/json` is another package, so it only sees **exported** fields.
* **Struct tags** give the JSON key name.

```go
type Payload struct {
    UserID int    `json:"user_id"`
    Email  string `json:"email"`
    secret string               // unexported → ignored by json
}

bytes, err := json.Marshal(p)            // struct → JSON
err = json.Unmarshal(body, &incoming)    // JSON → struct (needs a pointer!)
```

This is how Go APIs receive and send JSON payloads: decode into a struct, work with the struct, encode back.

---

## 7. Files in This Folder

| File | What it covers |
| :--- | :--- |
| `struct.go` | declaring a struct, field access, pointer to struct, literal forms, export rule |
| `struct_more.go` | nested struct, anonymous struct, `==` comparison, JSON with struct tags |
| `main.go` | runs every example |

## 8. How to Run

```bash
go run ./06_structs
```

## 9. Revision Checklist

- [ ] What value does a missing field get in `vertex{x: 1}`?
- [ ] Why can we write `p.y` instead of `(*p).y`?
- [ ] Why does json ignore the `secret` field?
- [ ] What is a struct tag for?
- [ ] Why does `json.Unmarshal` need `&incoming` and not `incoming`?
