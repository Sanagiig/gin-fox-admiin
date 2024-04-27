package common

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func PaginateAndLike(page, pageSize int, columns []string, vals []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Scopes(Paginate(page, pageSize), AndLike(columns, vals))
	}
}

func OrLike(columns []string, vals []string) func(db *gorm.DB) *gorm.DB {
	if len(columns) != len(vals) {
		return nil
	}
	return func(db *gorm.DB) *gorm.DB {
		var likeVals []any
		colSB := strings.Builder{}

		for i := 0; i < len(columns); i++ {
			if vals[i] != "" {
				if colSB.Len() > 0 {
					colSB.WriteString(" OR ")
				}
				colSB.WriteString(columns[i])
				colSB.WriteString(" LIKE ?")
				likeVals = append(likeVals, "%"+vals[i]+"%")
			}
		}
		return db.Or(colSB.String(), likeVals...)
	}
}

func AndLike(columns []string, vals []string) func(db *gorm.DB) *gorm.DB {
	if len(columns) != len(vals) {
		return nil
	}
	return func(db *gorm.DB) *gorm.DB {
		var likeVals []any
		colSB := strings.Builder{}

		for i := 0; i < len(columns); i++ {
			if vals[i] != "" {
				if colSB.Len() > 0 {
					colSB.WriteString(" And ")
				}
				colSB.WriteString(columns[i])
				colSB.WriteString(" LIKE ?")
				likeVals = append(likeVals, "%"+vals[i]+"%")
			}
		}
		fmt.Println(colSB.String(), likeVals)
		return db.Where(colSB.String(), likeVals...)
	}
}
