module.exports = [
  {
    id: 1,
    title: "Company",
    menuList: [
      { name: "About Us", routerPath: "/static/about" },    
      { name: "Become Expert", routerPath: "/static/expert"  },
      { name: "Blog", routerPath: "/static/blog" },
      // { name: "Gift Cards", routerPath: "/" },
    ],
  },
  {
    id: 2,
    title: "Support",
    menuList: [
      { name: "Contact", routerPath: "static/contact" },
      { name: "Terms and Conditions", routerPath: "/static/terms" },
      { name: "Help Center", routerPath: "/static/help-center" },
    ],
  },
  {
    id: 3,
    title: "Other Services",
    menuList: [
      { name: "Car hire", routerPath: "/car/car-list-v1" },
      { name: "Activity Finder", routerPath: "/home" },
      { name: "Tour List", routerPath: "/tour/tour-list-v1" },
      { name: "Flight finder", routerPath: "/flight/flight-list-v1" },
      { name: "Golf Courses", routerPath: "/golf/golf-list-v1" },
      { name: "Holiday Rental", routerPath: "/hotel/hotel-list-v1" },
    ],
  },
];
