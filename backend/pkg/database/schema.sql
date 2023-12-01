-- users Table
CREATE TABLE Users (
    ID INTEGER PRIMARY KEY,
    UserName TEXT UNIQUE,
    Password TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Customers Table
CREATE TABLE Customers (
    ID INTEGER PRIMARY KEY,
    First_Name TEXT,
    Last_Name TEXT,
    Email TEXT,
    Phone TEXT,
    Address TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Products Table
CREATE TABLE Products (
    ID INTEGER PRIMARY KEY,
    Product_Name TEXT,
    Description TEXT,
    Price REAL,
    Category TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Orders Table
CREATE TABLE Orders (
    ID INTEGER PRIMARY KEY,
    Customer_ID INTEGER,
    Order_Date TEXT,
    Total_Amount REAL,
    Status TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (Customer_ID) REFERENCES Customers(ID)
);

-- OrderDetails Table
CREATE TABLE Order_Details (
    ID INTEGER PRIMARY KEY,
    Order_ID INTEGER,
    Product_ID INTEGER,
    Quantity INTEGER,
    Subtotal REAL,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (Order_ID) REFERENCES Orders(ID),
    FOREIGN KEY (Product_ID) REFERENCES Products(ID)
);

-- TourPackages Table
CREATE TABLE Tour_Packages (
    ID INTEGER PRIMARY KEY,
    Package_Name TEXT,
    Description TEXT,
    Price REAL,
    Itinerary TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- TravelBookings Table
CREATE TABLE Travel_Bookings (
    ID INTEGER PRIMARY KEY,
    Customer_ID INTEGER,
    Package_ID INTEGER,
    Travel_Date TEXT,
    Status TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (Customer_ID) REFERENCES Customers(ID),
    FOREIGN KEY (Package_ID) REFERENCES Tour_Packages(ID)
);

-- Payments Table
CREATE TABLE Payments (
    ID INTEGER PRIMARY KEY,
    Order_ID INTEGER,
    Amount REAL,
    Payment_Date TEXT,
    Payment_Method TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (Order_ID) REFERENCES Orders(ID)
);

-- Reviews and Ratings Table
CREATE TABLE Review_Ratings (
    ID INTEGER PRIMARY KEY,
    Product_ID INTEGER,
    Customer_ID INTEGER,
    Rating INTEGER,
    Review_Text TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (Product_ID) REFERENCES Products(ID),
    FOREIGN KEY (Customer_ID) REFERENCES Customers(ID)
);
