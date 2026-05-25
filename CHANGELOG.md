# Changelog

This project does not publish formal releases. Entries below are keyed by date
and commit hash so that consumers can locate the exact point of change in the
git history.

---

## 2026-05-25 · commit `90188b8`

### Breaking changes

#### `antihax/optional` replaced with a built-in generic type

The dependency `github.com/antihax/optional` (unmaintained since 2019) has been
removed. Optional parameters in the Inventory API are now expressed with the
generic type `optional.Optional[T]` defined in
`github.com/vmware/go-vmware-nsxt/common/optional`.

**Minimum Go version is now 1.18** (required for generics).

##### Field types in `*Opts` structs

Every `*Opts` struct that previously used the antihax concrete types now uses
the generic form:

| Before | After |
|---|---|
| `optional.String` | `optional.Optional[string]` |
| `optional.Int64` | `optional.Optional[int64]` |
| `optional.Bool` | `optional.Optional[bool]` |

##### Constructors

| Before | After |
|---|---|
| `optional.NewString("foo")` | `optional.New("foo")` |
| `optional.NewInt64(42)` | `optional.New(int64(42))` |
| `optional.NewBool(true)` | `optional.New(true)` |

##### Value retrieval method

| Before | After |
|---|---|
| `.Value()` | `.Get()` |

Note: `Get()` on an unset `Optional` returns the zero value of `T` without
panicking. Always call `IsSet()` before calling `Get()`.

##### Example migration

```go
// Before
opts := &nsxt.ListContainerClustersOpts{
    Cursor:      optional.NewString("abc"),
    PageSize:    optional.NewInt64(25),
    SortAscending: optional.NewBool(true),
}

// After
opts := &nsxt.ListContainerClustersOpts{
    Cursor:        optional.New("abc"),
    PageSize:      optional.New(int64(25)),
    SortAscending: optional.New(true),
}
```

### Dependency upgrades

- `golang.org/x/oauth2`: `v0.0.0-20190604053449` → `v0.36.0`
- `golang.org/x/net` (indirect): removed (no longer required)
