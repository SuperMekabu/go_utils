# Go generics utils

## Slices

### Utils for slice

#### Filter
- Filter

#### Map

- Map

#### Includes

- Includes

#### Remove

- RemoveFirst
- RemoveAll

#### Every

- Every

#### Some

- Some

---

## Maps

### Utils for map

#### Filter

- Filter

#### Map

- Map

#### Has

- HasKey
- HasValue

#### Remove

- Remove

#### Every

- Every

#### Some

- Some

---

## ID

Depends on

- [`github.com/google/uuid`](https://github.com/google/uuid) to generate uuid
- [`github.com/oklog/ulid`](https://github.com/oklog/ulid) to generate ulid

### Utils for generate an id

#### NewUUID

generate UUID v4 from unix nano.

#### NewUUIDFromObj

generate UUID v5 from byte slice.  
same UUID generates when same byte slice is given.

#### NewULID

generate ULID from given entropy.  
usually make entropy to using `NewEntropy`.

#### NewEntropy

make entropy for ULID generator from unix nano.