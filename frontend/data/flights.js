module.exports = [
  {
    id: 1,
    price: "934",
    deals: "16",
    animation: "100",
    selectId: "collapse_1",
    flightList: [
      {
        id: 1,
        avatar: "/img/flightIcons/1.png",
        arrivalAirport: "SAW",
        departureAirport: "STN",
        departureTime: "14:00",
        arrivalTime: "22:00",
        duration: "3h 05m- Nonstop",
      },
      {
        id: 2,
        avatar: "/img/flightIcons/2.png",
        arrivalAirport: "SAW",
        departureAirport: "STN",
        departureTime: "14:00",
        arrivalTime: "22:00",
        duration: "5h 05m- Nonstop",
      },
    ],
  },
  {
    id: 2,
    price: "690",
    deals: "12",
    selectId: "collapse_2",
    animation: "200",
    flightList: [
      {
        id: 1,
        avatar: "/img/flightIcons/1.png",
        arrivalAirport: "SAW",
        departureAirport: "STN",
        departureTime: "14:00",
        arrivalTime: "22:00",
        duration: "4h 05m- Nonstop",
      },
      {
        id: 2,
        avatar: "/img/flightIcons/2.png",
        arrivalAirport: "SAW",
        departureAirport: "STN",
        departureTime: "14:00",
        arrivalTime: "22:00",
        duration: "6h 05m- Nonstop",
      },
    ],
  },
  {
    id: 3,
    price: "999",
    deals: "17",
    animation: "300",
    selectId: "collapse_3",
    flightList: [
      {
        id: 1,
        avatar: "/img/flightIcons/1.png",
        arrivalAirport: "SAW",
        departureAirport: "STN",
        departureTime: "14:00",
        arrivalTime: "22:00",
        duration: "4h 05m- Nonstop",
      },
      {
        id: 2,
        avatar: "/img/flightIcons/2.png",
        arrivalAirport: "SAW",
        departureAirport: "STN",
        departureTime: "14:00",
        arrivalTime: "22:00",
        duration: "7h 05m- Nonstop",
      },
    ],
  },
  {
    id: 4,
    animation: "400",
    price: "859",
    deals: "15",
    selectId: "collapse_4",
    flightList: [
      {
        id: 1,
        avatar: "/img/flightIcons/1.png",
        arrivalAirport: "SAW",
        departureAirport: "STN",
        departureTime: "14:00",
        arrivalTime: "22:00",
        duration: "3h 05m- Nonstop",
      },
      {
        id: 2,
        avatar: "/img/flightIcons/2.png",
        arrivalAirport: "SAW",
        departureAirport: "STN",
        departureTime: "14:00",
        arrivalTime: "22:00",
        duration: "9h 05m- Nonstop",
      },
    ],
  },
  {
    id: 5,
    price: "934",
    deals: "16",
    selectId: "collapse_5",
    animation: "500",
    flightList: [
      {
        id: 1,
        avatar: "/img/flightIcons/1.png",
        arrivalAirport: "SAW",
        departureAirport: "STN",
        departureTime: "14:00",
        arrivalTime: "22:00",
        duration: "3h 05m- Nonstop",
      },
      {
        id: 2,
        avatar: "/img/flightIcons/2.png",
        arrivalAirport: "SAW",
        departureAirport: "STN",
        departureTime: "14:00",
        arrivalTime: "22:00",
        duration: "5h 05m- Nonstop",
      },
    ],
  },
  {
    id: 6,
    price: "690",
    deals: "12",
    selectId: "collapse_6",
    animation: "600",
    flightList: [
      {
        id: 1,
        avatar: "/img/flightIcons/1.png",
        arrivalAirport: "SAW",
        departureAirport: "STN",
        departureTime: "14:00",
        arrivalTime: "22:00",
        duration: "4h 05m- Nonstop",
      },
      {
        id: 2,
        avatar: "/img/flightIcons/2.png",
        arrivalAirport: "SAW",
        departureAirport: "STN",
        departureTime: "14:00",
        arrivalTime: "22:00",
        duration: "6h 05m- Nonstop",
      },
    ],
  },
];

CREATE TABLE Flights (
  ID INTEGER PRIMARY KEY,
  price REAL,
  deals TEXT,
  animation TEXT,
  SelectId TEXT,
  Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
  Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
);  



CREATE TABLE FlightList (
  ID INTEGER PRIMARY KEY,
  avatar TEXT,
  arrivalAirport TEXT,
  departureAirport TEXT,
  departureTime TEXT,
  arrivalTime TEXT,
  duration TEXT,
  Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
  Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
);

Can you make the insert findPlaceSlice.