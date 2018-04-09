# GQuery
Eloquent esq query builder for Go

# Import
To start using GQuery, install Go and run `go get`:

```sh
$ go get -u github.com/warpinc/gquery
```

This will retrieve the library.

# Usage
Building your first query

```go
package main

import "github.com/warpinc/gquery"

func main() {
  var query = query.Init("user") // query the user table
  query.Select("name", "address").Where("id", "=", 1).Fetch() // for the users name and address where id is 1
}
```
