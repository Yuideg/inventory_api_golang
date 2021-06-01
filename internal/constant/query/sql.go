package query

const (
	// Role Insert,Update ,Retrieve and delete operation
	    RoleInsert= `INSERT INTO role (id,name) values($1, $2);`
		RoleSelectOne=`SELECT * FROM role WHERE id = $1;`
        RoleSelectAll=`SELECT * FROM role;`
        RoleUpdate=`UPDATE role SET id=$1, id=$2 WHERE id=$3;`
        RoleDelete=`DELETE FROM role WHERE id=$1;`
	// User Insert,Update ,Retrieve and delete operation
		UserInsert = `INSERT INTO user_tb(username,password,first_name,last_name,gender,email,country,city,state,zipcode,street,latitude,longitude,phone,role_id) VALUES($1, $2, $3,$4,$5,$6,$7, $8, $9,$10,$11,$12,$13,$14,$15);`
		UserSelectOne=`SELECT * FROM user_tb WHERE id = $1;`
		UserSelectAll=`SELECT * FROM user_tb;`
		UserUpdate=`UPDATE user_tb SET id=$1, username=$2,password=$3,first_name=$4,last_name=$5,gender=$6,email=$7,country=$8,city=$9,state=$10,zipcode=$11,street=$12,latitude=$13,longitude=$14,phone=$15,role_id=$16 WHERE id=$17;`
		UserDelete=`DELETE FROM user_tb WHERE id=$1;`

	// Order Insert,Update ,Retrieve and delete operation
		OrderInsert=`INSERT INTO orders(quantity,unit,confirm_by,user_id,product_id,status,expired_on) VALUES($1, $2, $3,$4,$5,$6,$7);`
		OrderSelectOneBy_ID=`SELECT * FROM orders WHERE  id = $1;`
	    OrderSelectBy_User_ID=`SELECT * FROM orders WHERE  user_id = $1;`
		OrderSelectAll=`SELECT * FROM orders;`
		OrderUpdate=`UPDATE orders SET id=$1, quantity=$2,unit=$3,confirm_by=$4,customer_id=$5,product_id=$6,status=$7,expired_on=$8 WHERE id=$9;`
		OrderDelete=`DELETE FROM orders WHERE id=$1;`
	// Product Insert,Update ,Retrieve and delete operation
		ProductInsert=`INSERT INTO product(barcode, product_name,stock_id,brand,image,cost,price,unit,tax_rate,discount,mfg_date,expired_on) VALUES($1, $2, $3,$4,$5,$6,$7, $8, $9,$10,$11,$12,$13);`
		ProductSelectOne=`SELECT * FROM product WHERE id = $1;`
		ProductSelectAll=`SELECT * FROM product;`
		ProductUpdate=`UPDATE product SET id=$1,barcode=$2, product_name=$3,stock_id=$4,brand=$5,image=$6,cost=$7,price=$8,unit=$9,tax_rate=$10,mfg_date=$11,expired_on=$12 WHERE barcode=$13;`
		ProductDelete=`DELETE FROM product WHERE barcode=$1;`
	// Stock Insert,Update ,Retrieve and delete operation
		StockInsert=`INSERT INTO stock(products_common_code,products_common_name, number_of_products,supplier_id,country,city,state,zipcode,street,latitude,longitude) VALUES($1, $2, $3,$4,$5,$6,$7,$8,$9,$10,$11);`
		StockSelectOne=`SELECT * FROM stock WHERE id = $1;`
		StockSelectAll=`SELECT * FROM stock;`
		StockUpdate=`UPDATE stock SET id=$1 ,products_code=$2,products_name=$3, number_of_products=$4,supplier_id=$5,country=$6,city=$7,state=$8,zipcode=$9,street=$10,latitude=$11,longitude=$12 WHERE id=$13;`
		StockDelete=`DELETE FROM stock WHERE id=$1;`
	// Supplier Insert,Update ,Retrieve and delete operation
		SupplierInsert=`INSERT INTO supplier(supplier_name,country,city,state,zipcode,street,latitude,longitude,fax,po_box,email,website,phone) VALUES($1, $2, $3,$4,$5,$6,$7, $8, $9,$10,$11,$12,$13);`
		SupplierSelectOne=`SELECT * FROM supplier WHERE id = $1;`
		SupplierSelectAll=`SELECT * FROM supplier;`
		SupplierUpdate=`UPDATE supplier SET id=$1 ,supplier_name=$2,country=$3,city=$4,state=$5,zipcode=$6,street=$7,latitude=$8,longitude=$9,fax=$10,po_box=$11,email=$12,website=$13,phone=$4 WHERE id=$15;`
		SupplierDelete=`DELETE FROM supplier WHERE id=$1;`
)
