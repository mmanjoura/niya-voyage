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

-- Hotels Table
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

---------------- The new Schema----------------------------------

CREATE TABLE Categories (
    ID                 INTEGER PRIMARY KEY,
    Name               TEXT    NOT NULL,
    Description        TEXT,
    Parent_Category_ID INTEGER,
    FOREIGN KEY (
        Parent_Category_ID
    )
    REFERENCES Categories (ID) 
);

-- Hotels Table

CREATE TABLE hotels (
    id INTEGER PRIMARY KEY,
    category_id INTEGER,
    tag TEXT,
    img TEXT,
    title TEXT,
    location TEXT,
    ratings REAL,
    numberOfReviews INTEGER,
    price REAL,
    delayAnimation INTEGER,
    city TEXT,
     FOREIGN KEY (category_id) REFERENCES Categories(id)
);

CREATE TABLE hotel_images (
    id INTEGER PRIMARY KEY,
    hotel_id INTEGER,
    slideImg TEXT,
    FOREIGN KEY (hotel_id) REFERENCES hotels(id)
);

INSERT INTO hotels (id, category_id, tag, img, title, location, ratings, numberOfReviews, price, delayAnimation, city)
VALUES 
(4, 3, 'top rated', '/img/hotels/4.png', 'DoubleTree by Hilton Hotel New York Times Square West', 'Vaticano Prati, Rome', 4.5, 5633, 89, 400, 'new_york'),
(5, 3, 'Breakfast Included', '/img/hotels/5.png', 'The Montcalm At Brewery London City', 'Westminster Borough, London', 4.8, 3014, 99, 500, 'london');

INSERT INTO hotel_images (hotel_id, slideImg)
VALUES 
(4, '["/img/hotels/4.png"]'),
(5, '["/img/hotels/5.png"]');

select * from Hotels h inner join hotel_images hi on h.id = hi.hotel_id;

-------------- Destinations Table--------------------
CREATE TABLE Destinations (
    id INTEGER PRIMARY KEY,
    colClass TEXT,
    title TEXT,
    location TEXT,
    travellers TEXT,
    hoverText TEXT,
    img TEXT,
    city TEXT,
    properties TEXT,
    region TEXT,
    delayAnimation INTEGER
);

INSERT INTO Destinations (id, colClass, title, location, travellers, hoverText, img, city, properties, region, delayAnimation)
VALUES 
(1, "col-xl-3 col-md-4 col-sm-6", "United Kingdom", "London, UK", "147,681", "14 Hotel - 22 Cars - 18 Tours - 95 Activity", "/img/destinations/1/1.png", "Hawai", "12,683", "north_america", 0),
(2, "col-xl-3 col-md-4 col-sm-6", "Italy", "Italy", "147,681", "14 Hotel - 22 Cars - 18 Tours - 95 Activity", "/img/destinations/1/2.png", "Istanbul", "12,683", "europe", 100),
(3, "col-xl-3 col-md-4 col-sm-6", "France", "France", "147,681", "14 Hotel - 22 Cars - 18 Tours - 95 Activity", "/img/destinations/1/3.png", "San Diego", "12,683", "north_america", 200);

------------------Testimonials Table -------------------------------------------
CREATE TABLE Testimonials (
    id INTEGER PRIMARY KEY,
    hotel_id INTEGER,
    meta TEXT,
    avatar TEXT,
    name TEXT,
    designation TEXT,
    text TEXT,
    delayAnimation INTEGER,
     FOREIGN KEY (hotel_id) REFERENCES hotels(id)
);

INSERT INTO Testimonials (id, hotel_id, meta, avatar, name, designation, text, delayAnimation)
VALUES 
(1, 4, "Hotel Equatorial Melaka", "/img/avatars/testimonials/1.png", "Annette Black", "UX / UI Designer", "The place is in a great location in Gumbet. The area is safe and beautiful. The apartment was comfortable and the host was kind and responsive to our requests. Really a nice place.", 100),
(2, 5,  "Hotel Equatorial Melaka", "/img/avatars/testimonials/2.png", "Annette Black", "UX / UI Designer", "The place is in a great location in Gumbet. The area is safe and beautiful. The apartment was comfortable and the host was kind and responsive to our requests. Really a nice place.", 200),
(3, 4,  "Hotel Equatorial Melaka", "/img/avatars/testimonials/3.png", "Annette Black", "UX / UI Designer", "The place is in a great location in Gumbet. The area is safe and beautiful. The apartment was comfortable and the host was kind and responsive to our requests. Really a nice place.", 300),
(4, 5, "Hotel Equatorial Melaka", "/img/avatars/testimonials/1.png", "Annette Black", "UX / UI Designer", "The place is in a great location in Gumbet. The area is safe and beautiful. The apartment was comfortable and the host was kind and responsive to our requests. Really a nice place.", 400);


---------------------Blogs table --------------------------
CREATE TABLE Blogs (
    id INTEGER PRIMARY KEY,
    img TEXT,
    title TEXT,
    date TEXT,
    delayAnimation INTEGER,
    details TEXT,
    tag TEXT,
    tags TEXT
);


INSERT INTO Blogs (id, img, title, date, delayAnimation, details, tag, tags)
VALUES 
(1, "/img/blog/1.png", "10 European ski destinations you should visit this winter", "Jan 06, 2023", 100, "For decades, travelers have been escaping to the Catskills — a mountainous region in upstate New York — whenever they’ve needed a reset.", "art", "adventure_travel, food_drink"),
(2, "/img/blog/2.png", "Booking travel during Corona: good advice in an uncertain time", "April 06, 2022", 200, "For decades, travelers have been escaping to the Catskills — a mountainous region in upstate New York — whenever they’ve needed a reset.", "beaches", "beaches, family_holidays");
