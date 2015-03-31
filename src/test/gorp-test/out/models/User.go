type User struct {
  Id int64 `db:"user_id"`
  Created int64
  Deleted int64
  Username string
}