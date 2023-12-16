------------Disable/Enable Foreign Key Checks-----------------
PRAGMA foreign_keys = OFF;
PRAGMA foreign_keys = ON;

-------------- Destinations Table--------------------
CREATE TABLE Destinations (
    id INTEGER PRIMARY KEY,
    class TEXT,
    title TEXT,
    name TEXT,
    location TEXT,
    travellers TEXT,
    hover TEXT,
    img TEXT,
    city TEXT,
    properties TEXT,
    region TEXT,
    Animation TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO Destinations (class, img, name, properties, animation) VALUES
('col-xl-3 col-md-4 col-sm-6', '/img/destinations/2/1.png', 'los angeles', '1714', '200'),
('col-xl-6 col-md-4 col-sm-6', '/img/destinations/2/2.png', 'london', '1714', '400'),
('col-xl-3 col-md-4 col-sm-6', '/img/destinations/2/3.png', 'reykjavik', '1714', '600'),
('col-xl-6 col-md-4 col-sm-6', '/img/destinations/2/4.png', 'paris', '1714', '200'),
('col-xl-3 col-md-4 col-sm-6', '/img/destinations/2/5.png', 'amsterdam', '1714', '400'),
('col-xl-3 col-md-4 col-sm-6', '/img/destinations/2/6.png', 'istanbul', '1714', '600');

-- Hotels Table

CREATE TABLE hotels (
    id INTEGER PRIMARY KEY,
    category_id INTEGER,
    tag TEXT,
    img TEXT,
    title TEXT,
    location TEXT,
    ratings TEXT,
    Reviews TEXT,
    price TEXT,
    Animation TEXT,
    city TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP,
     FOREIGN KEY (category_id) REFERENCES Categories(id)
);

INSERT INTO hotels (id, category_id, tag, img, title, location, ratings, Reviews, price, Animation, city)
VALUES
(1, 3, 'Breakfast Included', '/img/hotels/1.png', 'The Montcalm At Brewery London City', 'Westminster Borough, London', 4.7, 3014, 72.0, 100, 'new_york'),
(2, 3,'', '/img/hotels/2.png', 'Staycity Aparthotels Deptford Bridge Station', 'Ciutat Vella, Barcelona', 4.8, 2345, 85.0, 200, 'london'),
(3, 3,'best seller', '/img/hotels/3.png', 'The Westin New York at Times Square West', 'Manhattan, New York', 4.7, 3014, 68.0, 300, 'new_york'),
(4, 3,'top rated', '/img/hotels/4.png', 'DoubleTree by Hilton Hotel New York Times Square West', 'Vaticano Prati, Rome', 4.5, 5633, 89.0, 400, 'new_york'),
(5, 3,'Breakfast Included', '/img/hotels/5.png', 'The Montcalm At Brewery London City', 'Westminster Borough, London', 4.8, 3014, 99.0, 500, 'london'),
(6, 3,'-25% today', '/img/hotels/6.png', 'The Montcalm At Brewery London City', 'Westminster Borough, London', 4.7, 3467, 79.0, 600, 'new_york'),
(7, 3,'best seller', '/img/hotels/7.png', 'Staycity Aparthotels Deptford Bridge Station', 'Ciutat Vella, Barcelona', 4.8, 3014, 88.0, 700, 'new_york'),
(8, 3,'top rated', '/img/hotels/8.png', 'The Westin New York at Times Square West', 'Manhattan, New York', 4.9, 7654, 68.0, 800, 'london'),
(9, 3,'Breakfast Included', '/img/hotels/9.png', 'The Montcalm At Brewery London City', 'Westminster Borough, London', 4.7, 3245, 72.0, 900, 'new_york'),
(10, 3,'', '/img/hotels/10.png', 'Staycity Aparthotels Deptford Bridge Station', 'Ciutat Vella, Barcelona', 4.8, 3432, 85.0, 1000, 'paris'),
(11, 3,'best seller', '/img/hotels/11.png', 'The Westin New York at Times Square West', 'Manhattan, New York', 4.7, 3014, 68.0, 1100, 'new_york'),
(12, 3,'top rated', '/img/hotels/12.png', 'DoubleTree by Hilton Hotel New York Times Square West', 'Vaticano Prati, Rome', 4.5, 2343, 89.0, 1200, 'paris'),
(13, 3,'Breakfast Included', '/img/hotels/13.png', 'The Montcalm At Brewery London City', 'Westminster Borough, London', 4.8, 3014, 99.0, 500, 'new_york'),
(14, 3,'-25% today', '/img/hotels/14.png', 'The Montcalm At Brewery London City', 'Westminster Borough, London', 4.7, 3467, 79.0, 600, 'istanbul'),
(15, 3,'best seller', '/img/hotels/15.png', 'Staycity Aparthotels Deptford Bridge Station', 'Ciutat Vella, Barcelona', 4.8, 3014, 88.0, 700, 'istanbul'),
(16, 3,'top rated', '/img/hotels/16.png', 'The Westin New York at Times Square West', 'Manhattan, New York', 4.9, 7654, 68.0, 800, 'new_york');

------------------Hotel Images Table ------------------------------
CREATE TABLE hotel_images (
    id INTEGER PRIMARY KEY,
    hotel_id INTEGER,
    slideImg TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (hotel_id) REFERENCES hotels(id)
);

INSERT INTO hotel_images (hotel_id, slideImg) VALUES 
(1, '/img/hotels/1.png'),
(2, '/img/hotels/2.png'),
(2, '/img/hotels/1.png'),
(2, '/img/hotels/3.png'),
(3, '/img/hotels/3.png'),
(4, '/img/hotels/4.png'),
(5, '/img/hotels/5.png'),
(6, '/img/hotels/6.png'),
(7, '/img/hotels/7.png'),
(8, '/img/hotels/8.png'),
(9, '/img/hotels/9.png'),
(10, '/img/hotels/10.png'),
(10, '/img/hotels/11.png'),
(10, '/img/hotels/12.png'),
(11, '/img/hotels/11.png'),
(12, '/img/hotels/12.png'),
(13, '/img/hotels/13.png'),
(14, '/img/hotels/14.png'),
(15, '/img/hotels/15.png'),
(16, '/img/hotels/16.png');


---------------- Tours Table -----------------------------
CREATE TABLE Tours (
    ID INTEGER PRIMARY KEY,
    tag TEXT,
    title TEXT,
    price REAL,
    location TEXT,
    duration TEXT,
    reviews TEXT,
    price TEXT, 
    tourType TEXT,
    animation TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO Tours (ID, tag, title, price, location, duration, reviews, price, tourType, animation)
VALUES
(1, 'LIKELY TO SELL OUT*', 'Stonehenge, Windsor Castle and Bath with Pub Lunch in Lacock', 72.0, 'Westminster Borough, London', '16', '3014', '72', 'Full-day Tours', '100'),
(2, '', 'Westminster Walking Tour & Westminster Abbey Entry', 65.0, 'Ciutat Vella, Barcelona', '14', '2045', '65', 'Attractions & Museums', '200'),
(3, 'best seller', 'High-Speed Thames River RIB Cruise in London', 87.0, 'Manhattan, New York', '18', '2163', '87', 'Private and Luxury', '300'),
(4, 'top rated', 'Edinburgh Darkside Walking Tour: Mysteries, Murder and Legends', 99.0, 'Vaticano Prati, Rome', '20', '1458', '99', 'Bus Tours', '400'),
(5, 'LIKELY TO SELL OUT*', 'Stonehenge, Windsor Castle and Bath with Pub Lunch in Lacock', 72.0, 'Westminster Borough, London', '16', '3014', '72', 'Full-day Tours', '100'),
(6, '', 'Westminster Walking Tour & Westminster Abbey Entry', 65.0, 'Ciutat Vella, Barcelona', '14', '2045', '65', 'Attractions & Museums', '200'),
(7, 'best seller', 'High-Speed Thames River RIB Cruise in London', 87.0, 'Manhattan, New York', '18', '2163', '87', 'Private and Luxury', '300'),
(8, 'top rated', 'Edinburgh Darkside Walking Tour: Mysteries, Murder and Legends', 99.0, 'Vaticano Prati, Rome', '20', '1458', '99', 'Bus Tours', '400'),
(9, 'best seller', 'High-Speed Thames River RIB Cruise in London', 87.0, 'Manhattan, New York', '18', '2163', '87', 'Private and Luxury', '500');


--------------Activities--------------------------------
---------------- Tours Table -----------------------------
CREATE TABLE Activities (
    ID INTEGER PRIMARY KEY,
    tag TEXT,
    title TEXT,
    price REAL,
    location TEXT,
    duration TEXT,
    reviews TEXT,
    ratings TEXT,
    animation TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO Activities (ID, tag, title, price, location, duration, reviews, ratings, animation)
VALUES
(1, 'LIKELY TO SELL OUT*', 'Stonehenge, Windsor Castle and Bath with Pub Lunch in Lacock', 67.0, 'Westminster Borough, London', '16', '94', '4.82', '100'),
(2, '', 'Westminster Walking Tour & Westminster Abbey Entry', 99.0, 'Ciutat Vella, Barcelona', '14', '2045', '4.82', '200'),
(3, 'best seller', 'High-Speed Thames River RIB Cruise in London', 88.0, 'Manhattan, New York', '18', '2163', '4.82', '300'),
(4, 'top rated', 'Edinburgh Darkside Walking Tour: Mysteries, Murder and Legends', 55.0, 'Vaticano Prati, Rome', '20', '1458', '4.82', '400'),
(5, 'LIKELY TO SELL OUT*', 'Stonehenge, Windsor Castle and Bath with Pub Lunch in Lacock', 67.0, 'Westminster Borough, London', '16', '94', '4.82', '100'),
(6, 'LIKELY TO SELL OUT*', 'Stonehenge, Windsor Castle and Bath with Pub Lunch in Lacock', 67.0, 'Westminster Borough, London', '16', '94', '4.82', '200'),
(7, '', 'Westminster Walking Tour & Westminster Abbey Entry', 99.0, 'Ciutat Vella, Barcelona', '14', '2045', '4.82', '300'),
(8, 'best seller', 'High-Speed Thames River RIB Cruise in London', 88.0, 'Manhattan, New York', '18', '2163', '4.82', '400'),
(9, 'top rated', 'Edinburgh Darkside Walking Tour: Mysteries, Murder and Legends', 55.0, 'Vaticano Prati, Rome', '20', '1458', '4.82', '500');

----------------Rentals Table--------------------------------------------------------------------------------------------

CREATE TABLE Rentals (
    ID INTEGER PRIMARY KEY,
    tag TEXT,
    title TEXT,
    price REAL,
    location TEXT,
    duration TEXT,
    reviews TEXT,
    ratings TEXT,
    animation TEXT,
    guest TEXT,
    bedroom TEXT,
    bed Text,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO Rentals (ID, tag, title, price, location, ratings, numberOfReviews, guest, bedroom, bed, delayAnimation)
VALUES
(1, '', 'Luxury New Apartment With Private Garden', 72.0, 'Westminster Borough, London', '4.8', '3014', '2', '1', '1', '100'),
(2, '', 'Premium One Bedroom Luxury Living in the Heart of Mayfair', 72.0, 'Ciutat Vella, Barcelona', '4.8', '3014', '4', '2', '1', '200'),
(3, 'best seller', 'Style, Charm & Comfort in Camberwell', 72.0, 'Manhattan, New York', '4.5', '3014', '2', '1', '1', '300'),
(4, 'top rated', 'Marylebone - Oxford Street 1 bed apt with WiFi', 72.0, 'Vaticano Prati, Rome', '4.8', '3014', '3', '2', '1', '400'),
(5, '', 'Luxury New Apartment With Private Garden', 72.0, 'Westminster Borough, London', '4.9', '3014', '2', '1', '1', '100'),
(6, '', 'Premium One Bedroom Luxury Living in the Heart of Mayfair', 72.0, 'Ciutat Vella, Barcelona', '4.8', '3014', '4', '2', '1', '200'),
(7, 'best seller', 'Style, Charm & Comfort in Camberwell', 72.0, 'Manhattan, New York', '4.65', '3014', '2', '1', '1', '300'),
(8, 'top rated', 'Marylebone - Oxford Street 1 bed apt with WiFi', 72.0, 'Vaticano Prati, Rome', '4.8', '3014', '3', '2', '1', '400'),
(9, '', 'Luxury New Apartment With Private Garden', 72.0, 'Westminster Borough, London', '4.7', '3014', '2', '1', '1', '100');

-------------------------Cars Table----------------------------------------------
CREATE TABLE Cars (
    ID INTEGER PRIMARY KEY,
    tag TEXT,
    title TEXT,
    price REAL,
    location TEXT,
    reviews TEXT,
    ratings TEXT,
    animation TEXT,
    seat TEXT,
    type Text,
    luggage TEXT,
    transmission TEXT,
    speed TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO Cars (ID, tag, title, price, location, type, ratings, reviews, seat, luggage, transmission, speed, animation)
VALUES
(1, '', 'Mercedes-Benz E-Class', 72.0, 'Heathrow Airport', 'LUXURY', '4.8', '3014', '4', '1', 'Automatic', 'Unlimited', '100'),
(2, '', 'Jaguar F-Pace', 99.0, 'Heathrow Airport', 'SUV', '4.6', '2345', '3', '1', 'Manual', 'Unlimited', '200'),
(3, 'best seller', 'Volvo XC90', 87.0, 'Heathrow Airport', 'SUV', '4.9', '4321', '5', '1', 'Automatic', 'Unlimited', '300'),
(4, '', 'BMW 5 Series', 78.0, 'Heathrow Airport', 'LUXURY', '5.0', '2432', '3', '1', 'Automatic', 'Unlimited', '400'),
(5, '', 'Mercedes-Benz E-Class', 72.0, 'Heathrow Airport', 'LUXURY', '4.8', '3014', '4', '1', 'Automatic', 'Unlimited', '100'),
(6, '', 'Jaguar F-Pace', 99.0, 'Heathrow Airport', 'SUV', '4.6', '2345', '3', '1', 'Manual', 'Unlimited', '200'),
(7, '', 'Mercedes-Benz E-Class', 72.0, 'Heathrow Airport', 'LUXURY', '4.8', '3014', '4', '1', 'Automatic', 'Unlimited', '100'),
(8, '', 'Jaguar F-Pace', 99.0, 'Heathrow Airport', 'SUV', '4.6', '2345', '3', '1', 'Manual', 'Unlimited', '200'),
(9, 'best seller', 'Volvo XC90', 87.0, 'Heathrow Airport', 'SUV', '4.9', '4321', '5', '1', 'Automatic', 'Unlimited', '300'),
(10, '', 'BMW 5 Series', 78.0, 'Heathrow Airport', 'LUXURY', '5.0', '2432', '3', '1', 'Automatic', 'Unlimited', '400'),
(11, '', 'Mercedes-Benz E-Class', 72.0, 'Heathrow Airport', 'LUXURY', '4.8', '3014', '4', '1', 'Automatic', 'Unlimited', '100'),
(12, '', 'Jaguar F-Pace', 99.0, 'Heathrow Airport', 'SUV', '4.6', '2345', '3', '1', 'Manual', 'Unlimited', '200'),
(13, 'best seller', 'Volvo XC90', 87.0, 'Heathrow Airport', 'SUV', '4.9', '4321', '5', '1', 'Automatic', 'Unlimited', '300'),
(14, '', 'BMW 5 Series', 78.0, 'Heathrow Airport', 'LUXURY', '5.0', '2432', '3', '1', 'Automatic', 'Unlimited', '400'),
(15, '', 'Mercedes-Benz E-Class', 72.0, 'Heathrow Airport', 'LUXURY', '4.8', '3014', '4', '1', 'Automatic', 'Unlimited', '100');


------------------







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
(1, 4, Hotel Equatorial Melaka, /img/avatars/testimonials/1.png, Annette Black, UX / UI Designer, The place is in a great location in Gumbet. The area is safe and beautiful. The apartment was comfortable and the host was kind and responsive to our requests. Really a nice place., 100),
(2, 5,  Hotel Equatorial Melaka, /img/avatars/testimonials/2.png, Annette Black, UX / UI Designer, The place is in a great location in Gumbet. The area is safe and beautiful. The apartment was comfortable and the host was kind and responsive to our requests. Really a nice place., 200),
(3, 4,  Hotel Equatorial Melaka, /img/avatars/testimonials/3.png, Annette Black, UX / UI Designer, The place is in a great location in Gumbet. The area is safe and beautiful. The apartment was comfortable and the host was kind and responsive to our requests. Really a nice place., 300),
(4, 5, Hotel Equatorial Melaka, /img/avatars/testimonials/1.png, Annette Black, UX / UI Designer, The place is in a great location in Gumbet. The area is safe and beautiful. The apartment was comfortable and the host was kind and responsive to our requests. Really a nice place., 400);


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
(1, /img/blog/1.png, 10 European ski destinations you should visit this winter, Jan 06, 2023, 100, For decades, travelers have been escaping to the Catskills — a mountainous region in upstate New York — whenever they’ve needed a reset., art, adventure_travel, food_drink),
(2, /img/blog/2.png, Booking travel during Corona: good advice in an uncertain time, April 06, 2022, 200, For decades, travelers have been escaping to the Catskills — a mountainous region in upstate New York — whenever they’ve needed a reset., beaches, beaches, family_holidays);
