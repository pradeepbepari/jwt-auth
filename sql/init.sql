USE users;

CREATE TABLE IF NOT EXISTS USERS (
    UUID VARCHAR(255),
    FirstName VARCHAR(255),
    LastName VARCHAR(255),
    Password VARCHAR(255),
    Email VARCHAR(255),
    Phone VARCHAR(255),
    Role VARCHAR(255),
    User_id VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME,
    PRIMARY KEY (UUID)
);

CREATE TABLE IF NOT EXISTS PRODUCTS (
    Product_UUID VARCHAR(255),
    Product_Name VARCHAR(255),   
    Product_Price DECIMAL(18,2),
    Product_Quantity INT,
    created_at DATETIME,
    updated_at DATETIME,
    PRIMARY KEY (Product_UUID)
);

CREATE TABLE IF NOT EXISTS CartItems (
    Cart_ID VARCHAR(255),
    User_Id VARCHAR(255),   
    Product_ID VARCHAR(225),
    Product_Price DECIMAL(18,2),
    Quantity INT,
    created_at DATETIME,
    updated_at DATETIME,
    PRIMARY KEY (Product_ID)
    -- FOREIGN KEY (User_Id) REFERENCES USERS(UUID)
);

CREATE TABLE IF NOT EXISTS Orders (
    Order_ID VARCHAR(255),
    User_Id VARCHAR(255),   
    Product_ID VARCHAR(225),
    Quantity INT,
    Address VARCHAR(225),   
    Total_Price DECIMAL(18,2),
    created_at DATETIME,
    PRIMARY KEY (Order_ID)
   
);
