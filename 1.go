//我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

import (
   "database/sql"
   "fmt"

   "github.com/pkg/errors"
)



func GetSql() error {
   return errors.Wrap(sql.ErrNoRows, "GetSql failed")
}

func Call() error {
   return errors.WithMessage(GetSql(), "bar failed")
}

func main() {
   err := Call()
   if errors.Cause(err) == sql.ErrNoRows {
      fmt.Printf("data not found, %v\n", err)
      //fmt.Printf("%+v\n", err)
      return
   }
   if err != nil {
      // unknown error
   }
}