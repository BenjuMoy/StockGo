package repository

import "database/sql"

const schema = `
	CREATE TABLE IF NOT EXISTS products (
    	id INTEGER PRIMARY KEY,

     name TEXT NOT NULL
         COLLATE NOCASE
         UNIQUE
         CHECK (length(trim(name)) > 0),

    quantity INTEGER NOT NULL
    	CHECK(quantity >= 0),

    unit_cost INTEGER NOT NULL
    	CHECK(unit_cost >= 0)
	);

	CREATE TABLE IF NOT EXISTS sales (
    	id INTEGER PRIMARY KEY,

    	product_id INTEGER NOT NULL,

     	product_name TEXT NOT NULL,

     	unit_cost INTEGER NOT NULL
      		CHECK(unit_cost >= 0),

    	unit_price INTEGER NOT NULL
    		CHECK(unit_price >= 0),

      	created_at TIMESTAMP NOT NULL
       		DEFAULT (datetime('now')),

    	FOREIGN KEY (product_id)
        	REFERENCES products(id)
        	ON DELETE RESTRICT
	);`

const indexes = `
	CREATE INDEX IF NOT EXISTS idx_sales_stock
		ON sales(product_id);

	CREATE INDEX IF NOT EXISTS idx_sales_created_at
		ON sales(created_at);`

func initSchema(db *sql.DB) error {
	if _, err := db.Exec(schema); err != nil {
		return err
	}

	if _, err := db.Exec(indexes); err != nil {
		return err
	}

	return nil
}
