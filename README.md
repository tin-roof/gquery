# GQuery
Eloquent esq query builder for Go

# Import
To start using GQuery, install Go and run `go get`:

```sh
$ go get -u github.com/tin-roof/gquery
```

This will retrieve the library.

# Usage
Building your first query

```go
package main

import "github.com/tin-roof/gquery"

func main() {
  var query = query.Init("users") // query the users table
  query.Select("name", "address").Where("id", "=", 1).Fetch() // for the users name and address where id is 1
}
```

#Switch to $1 style placeholders for use with Postgresql

```go
package main

import "github.com/tin-roof/gquery"

func main() {
  var query = query.Init("users", "postgres") // query the users table
  query.Select("name", "address").Where("id", "=", 1).Fetch() // for the users name and address where id is 1
}
