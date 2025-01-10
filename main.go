package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ochom/gutils/sqlx"
	"github.com/ochom/projectx/pkg"
	"gorm.io/gorm"
)

func init() {
	dbConfig := sqlx.Config{
		Driver: sqlx.Sqlite,
	}

	if err := sqlx.New(&dbConfig); err != nil {
		panic(err)
	}

	if err := sqlx.Conn().AutoMigrate(&pkg.User{}, &pkg.Order{}, &pkg.CustomColumns{}); err != nil {
		panic(err)
	}

	// create test data
	if err := pkg.Seed(); err != nil {
		panic(err)
	}
}

func main() {
	app := fiber.New()
	app.Get("/users", getUsers)

	app.Get("/orders", getOrders)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}

func getUsers(c *fiber.Ctx) error {
	filters := pkg.GetFilters(c.Query("filters"))

	// get all custom user columns
	customColumns := sqlx.FindAll[pkg.CustomColumns](func(d *gorm.DB) *gorm.DB {
		return d.Where("table_name = 'users'")
	})

	inCustomColumns := func(column string) *pkg.CustomColumns {
		for _, customColumn := range customColumns {
			if customColumn.ColumnName == column {
				return customColumn
			}
		}

		return nil
	}

	tx := sqlx.Conn()
	for _, filter := range filters {
		column, operator, value := filter["column"].(string), filter["operator"].(string), filter["value"]
		if c := inCustomColumns(column); c != nil {
			condition := fmt.Sprintf("json_extract(meta, '$.%s') %s ?", column, operator)
			// condition := fmt.Sprintf("(meta->>%s)::%s %s ?", column, c.ColumnType, operator)
			tx = tx.Where(condition, value)
		} else {
			condition := fmt.Sprintf("%s %s ?", column, operator)
			tx = tx.Where(condition, value)
		}
	}

	var users []pkg.User = make([]pkg.User, 0)
	if err := tx.Find(&users).Error; err != nil {
		return err
	}

	fields := []map[string]any{}
	for _, c := range customColumns {
		fields = append(fields, map[string]any{
			"key":   fmt.Sprintf("meta.%s", c.ColumnName),
			"label": c.ColumnName,
			"type":  c.ColumnType,
			"show":  false,
		})
	}

	return c.JSON(map[string]any{
		"Fields": fields,
		"Data":   users,
	})
}

func getOrders(c *fiber.Ctx) error {
	return c.JSON(sqlx.FindAll[pkg.Order]())
}
