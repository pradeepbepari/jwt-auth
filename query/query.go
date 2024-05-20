package query

const (
	GetEmail = `SELECT UUID, FirstName, LastName, Password, Email, Phone,  Role, User_id, created_at, updated_at 
	FROM USERS 
	WHERE Email = ?;
	`
	GetPhone = `SELECT UUID, FirstName, LastName, Password, Email, Phone,  Role, User_id, created_at, updated_at 
	FROM USERS 
	WHERE Phone = ?;
	`
	CreateUser = `Insert INTO USERS( UUID, FirstName, LastName, Password, Email, Phone,  Role, User_id, created_at, updated_at )
	values(?,?,?,?,?,?,?,?,?,?);
	`
	Getalluser = `SELECT UUID, FirstName, LastName, Password, Email, Phone,  Role, User_id, created_at, updated_at 
	FROM USERS 
	`
	GetUser = `SELECT UUID, FirstName, LastName, Password, Email, Phone,  Role, User_id, created_at, updated_at 
	FROM USERS 
	WHERE UUID = ?;
	`
	Deletealluser = `DELETE FROM USERS ;`
)

const (
	CreateProduct = `Insert INTO PRODUCTS(Product_UUID,Product_Name,Product_Price,Product_Quantity,created_at, updated_at)values(?,?,?,?,?,?);`

	GetAllProducts = `SELECT * FROM PRODUCTS;`

	GetProductByID = `SELECT * FROM PRODUCTS WHERE Product_UUID=?;`

	DeleteByid = `DELETE FROM PRODUCTS WHERE Product_UUID=?;`

	UpdateProduct = `UPDATE PRODUCTS SET Product_Quantity=?,updated_at=? WHERE Product_UUID=?;`
)
const (
	AddCart = `Insert INTO CartItems( Cart_ID, User_Id, Product_ID, Quantity, created_at,updated_at )
	values(?,?,?,?,?,?);`

	CheckProductInCart = `SELECT * FROM CartItems WHERE Product_ID=?`

	UpdateCart = `UPDATE CartItems SET  Quantity=?, updated_at=? WHERE Product_ID=? AND Cart_ID=?;`

	Orderitem = `INSERT INTO Orders(Order_ID,User_Id,Product_ID,Quantity,Address,Total_Price,created_at)
	values(?,?,?,?,?,?,?)`
)
