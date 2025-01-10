package pkg

import "github.com/ochom/gutils/sqlx"

func Seed() error {
	err := sqlx.Conn().Raw("DELETE FROM custom_columns").Error
	if err != nil {
		return err
	}

	err = sqlx.Conn().Raw("DELETE FROM orders").Error
	if err != nil {
		return err
	}

	err = sqlx.Conn().Raw("DELETE FROM users").Error
	if err != nil {
		return err
	}

	if err := createUsers(); err != nil {
		return err
	}

	if err := createOrders(); err != nil {
		return err
	}

	return nil
}

func createUsers() error {

	customColumns := []CustomColumns{
		{
			TableName:  "users",
			ColumnName: "Age",
			ColumnType: "number",
			IsNullable: false,
		},
		{
			TableName:  "users",
			ColumnName: "Sex",
			ColumnType: "text",
			IsNullable: false,
		},
		{
			TableName:  "users",
			ColumnName: "AccountBalance",
			ColumnType: "number",
			IsNullable: true,
		},
	}

	if err := sqlx.Conn().CreateInBatches(&customColumns, 100).Error; err != nil {
		return err
	}

	users := []User{
		{
			Name: "John Doe",
			Meta: map[string]any{
				"Age":            30,
				"Sex":            "male",
				"AccountBalance": 1000,
			},
		},
		{
			Name: "Jane Doe",
			Meta: map[string]any{
				"Age": 25,
				"Sex": "female",
			},
		},
	}

	for _, user := range users {
		if err := sqlx.Conn().Create(&user).Error; err != nil {
			return err
		}
	}

	return nil
}

func createOrders() error {
	customColumns := []CustomColumns{
		{
			TableName:  "orders",
			ColumnName: "OrderDate",
			ColumnType: "date",
			IsNullable: false,
		},
		{
			TableName:  "orders",
			ColumnName: "OrderStatus",
			ColumnType: "text",
			IsNullable: false,
		},
		{
			TableName:  "orders",
			ColumnName: "OrderAmount",
			ColumnType: "number",
			IsNullable: false,
		},
	}

	if err := sqlx.Conn().CreateInBatches(&customColumns, 100).Error; err != nil {
		return err
	}

	orders := []Order{
		{
			UserId: 1,
			Meta: map[string]any{
				"order_date":   "2022-01-01",
				"order_status": "pending",
				"order_amount": 100,
			},
		},
		{
			UserId: 1,
			Meta: map[string]any{
				"order_date":   "2022-02-01",
				"order_status": "shipped",
				"order_amount": 200,
			},
		},
		{
			UserId: 2,
			Meta: map[string]any{
				"order_date":   "2022-03-01",
				"order_status": "delivered",
				"order_amount": 300,
			},
		},
	}

	for _, order := range orders {
		if err := sqlx.Conn().Create(&order).Error; err != nil {
			return err
		}
	}

	return nil
}
