package gquery

// @TODO: work on building this out as we need it and get it connected to the db all as 1 nice package.

import (
  "encoding/json"
  "strconv"
)

type Query struct {
  DbType int // which type of db
  Table string // table we are querying
  QueryString string // full query to run
  SelectString string // string containing the columsn to select
  SetString string // string of values we are setting
  SetParams []interface{} // values that will fill the ? in the select query
  ColumnString string // colunns we are inserting to
  ValueString string // values we are inserting
  InsertParams []interface{} // values that will fill the ? in the insert query
  WhereString string // where clause of the query
  WhereParams []interface{} // values that will fill the ? in the Where clause
  OrderString string // order by clause of the query
  GroupString string // group claus of the query
  LimitNumber int // limit of results requested
  ReturningString string // returning query
  Params []interface{} // params used for the query
  PGPC int // number of parameters to count for POSTGRES $1, $2 type placeholders
}

func Init(args ...string) *Query {
  q := new(Query)

  // set the db type
  if len(args) == 2 {
    switch args[1] {
      case "pg":
        fallthrough
      case "postgres":
        q.DbType = 1
        q.PGPC = 1
      default: // mysql
        q.DbType = 0
    }
  } else {
    q.DbType = 0
  }

  q.Table = args[0]

  return q
}

// build the select query
func (q *Query) Fetch() {
  q.build(1)
}

// build the insert query
func (q *Query) Put() {
  q.build(2)
}

// build the update query
func (q *Query) Amend() {
  q.build(3)
}

// build the delete query
func (q *Query) Trash() {
  q.build(4)
}

func(q *Query) Run() {
}

// takes a full query string
func (q *Query) String(query string) *Query {
  q.QueryString = query
  return q
}

// select string
func (q *Query) Select(args ...interface{}) *Query {
  for index, element := range args {
    // add the comma for anything after the first element
    if index != 0 {
      q.SelectString += ", "
    }

    // add the item to the string
    q.SelectString += element.(string)
  }

  return q
}

// add where clause to the query
func (q *Query) Where(args ...interface{}) *Query {
  for index, element := range args {
    if index == 0 || index == 1 {
      q.WhereString += element.(string) + " "
    } else {

      switch q.DbType {
        case 1:
          q.WhereString += "$" + strconv.Itoa(q.PGPC)
          q.PGPC++
        default:
          q.WhereString += "?"
      }

      q.WhereParams = append(q.WhereParams, element)
    }
  }

  return q
}

// add AND where clause to the query
func (q *Query) Andwhere(args ...interface{}) *Query {
  q.WhereString += " AND "
  for index, element := range args {
    if index == 0 || index == 1 {
      q.WhereString += element.(string) + " "
    } else {

      switch q.DbType {
        case 1:
          q.WhereString += "$" + strconv.Itoa(q.PGPC)
          q.PGPC++
        default:
          q.WhereString += "?"
      }

      q.WhereParams = append(q.WhereParams, element)
    }
  }

  return q
}

// add OR where clause to the query
func (q *Query) Orwhere(args ...interface{}) *Query {
  q.WhereString += " OR "

  for index, element := range args {
    if index == 0 || index == 1 {
      q.WhereString += element.(string) + " "
    } else {

      switch q.DbType {
        case 1:
          q.WhereString += "$" + strconv.Itoa(q.PGPC)
          q.PGPC++
        default:
          q.WhereString += "?"
      }

      q.WhereParams = append(q.WhereParams, element)
    }
  }

  return q
}

// build the group string for the query
func (q *Query) Groupby(args ...interface{}) *Query {
  if q.GroupString != "" {
    q.GroupString += ", "
  }

  for index, element := range args {
    if index != 0 {
      q.GroupString += ", "
    }
    q.GroupString += element.(string)
  }

  return q
}

// build the order string for the query
func (q *Query) Orderby(args ...interface{}) *Query {
  if q.OrderString != "" {
    q.OrderString  += ", "
  }

  for index, element := range args {
    if index == 1 {
      q.OrderString += " "
    }

    q.OrderString += element.(string)
  }

  return q
}

// build the limit string for the query
func (q *Query) Limit(limit int) *Query {
  q.LimitNumber = limit

  return q
}

// build the set string for update or inserts
func (q *Query) Set(args ...interface{}) *Query {
  if q.SetString != "" {
    q.SetString += ", "
  }
  for index, element := range args {
    if index == 0 {
      q.SetString += element.(string)
    } else {

      switch q.DbType {
        case 1:
          q.SetString += " = $" + strconv.Itoa(q.PGPC)
          q.PGPC++
        default:
          q.SetString += " = ?"
      }

      q.SetParams = append(q.SetParams, element)
    }
  }
  return q
}

// build the set string for update or inserts
func (q *Query) Insert(args ...interface{}) *Query {
  if q.SetString != "" {
    q.SetString += ", "
  }
  for index, element := range args {
    if index == 0 {
      if q.ColumnString != "" {
        q.ColumnString += ", "
      }
      q.ColumnString += element.(string)
    } else {
      if q.ValueString != "" {
        q.ValueString += ", "
      }

      switch q.DbType {
        case 1:
          q.ValueString += "$" + strconv.Itoa(q.PGPC)
          q.PGPC++
        default:
          q.ValueString += "?"
      }

      q.InsertParams = append(q.InsertParams, element)
    }
  }
  return q
}

// build the returning string
func (q *Query) Returning(args ...interface{}) *Query {
  for index, element := range args {
    // add the comma for anything after the first element
    if index != 0 {
      q.ReturningString += ", "
    }

    // add the item to the string
    q.ReturningString += element.(string)
  }

  return q
}

// build the full query string
func (q *Query) build(qType int) {
  var query string

  switch qType {
    case 4:
      query = "DELETE FROM " + q.Table
    case 3:
      query = "UPDATE " + q.Table + " SET " + q.SetString
      q.Params = append(q.Params, q.SetParams...)
    case 2:
      query = "INSERT INTO " + q.Table + " (" + q.ColumnString + ") VALUES (" + q.ValueString + ")"
      q.Params = append(q.Params, q.InsertParams...)
    case 1:
      fallthrough
    default:
      if q.SelectString != "" {
        query = "SELECT " +  q.SelectString + " FROM " + q.Table
      } else {
        query = "SELECT * FROM " + q.Table
      }
  }

  // add where clause
  if q.WhereString != "" {
    query += " WHERE " + q.WhereString
    q.Params = append(q.Params, q.WhereParams...)
  }

  // add group by clause
  if q.GroupString != "" {
    query += " GROUP BY " + q.GroupString
  }

  // add order by clause
  if q.OrderString != "" {
    query += " ORDER BY " + q.OrderString
  }

  // add limit clause
  if q.LimitNumber != 0 {
    query += " LIMIT " + strconv.Itoa(q.LimitNumber)
  }

  // add order by clause
  if q.ReturningString != "" {
    query += " RETURNING " + q.ReturningString
  }

  query += ";"

  q.QueryString = query
}

// run the query
func query() {

}

// return the whole query object for review
func (q *Query) View() []byte {
	data, err := json.Marshal(q)

	if err != nil {
    return []byte("nothing here")
	}

  return data;
}
