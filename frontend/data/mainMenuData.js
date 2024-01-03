export const homeItems = [
  {
    name: "Hotels",
    routePath: "/hotel/hotel-list-v1",
  },
  {
    name: "Riads",
    routePath: "/hotel/hotel-list-v1",
  },
  {
    name: "Holiday Rentals",
    routePath: "/hotel/hotel-list-v1",
  },
];
export const blogItems = [
  {
    name: "Blog List V1",
    routePath: "/blog/blog-list-v1",
  },
  {
    name: "Blog List V2",
    routePath: "/blog/blog-list-v2",
  },
  {
    name: "Blog Single",
    routePath: "/blog/blog-details/1",
  },
];
export const pageItems = [
  {
    name: "404",
    routePath: "/404",
  },
  {
    name: "About",
    routePath: "/static/about",
  },
  {
    name: "Become Expert",
    routePath: "/static/become-expert",
  },
  {
    name: "Help Center",
    routePath: "/static/help-center",
  },
  {
    name: "Login",
    routePath: "/static/login",
  },
  {
    name: "Register",
    routePath: "/static/signup",
  },
  {
    name: "Terms",
    routePath: "/static/terms",
  },
  {
    name: "Invoice",
    routePath: "/static/invoice",
  },
  {
    name: "UI Elements",
    routePath: "/static/ui-elements",
  },
];
export const dashboardItems = [
  {
    name: "Dashboard",
    routePath: "/dashboard/db-dashboard",
  },
  {
    name: "Booking History",
    routePath: "/dashboard/db-booking",
  },
  {
    name: "Wishlist",
    routePath: "/dashboard/db-wishlist",
  },
  {
    name: "Settings",
    routePath: "/dashboard/db-settings",
  },
  {
    name: "Vendor Dashboard",
    routePath: "/vendor-dashboard/dashboard",
  },
  {
    name: "Vendor Add Hotel",
    routePath: "/vendor-dashboard/add-hotel",
  },
  {
    name: "Vendor Booking",
    routePath: "/vendor-dashboard/booking",
  },
  {
    name: "Vendor Hotels",
    routePath: "/vendor-dashboard/hotels",
  },
  {
    name: "Vendor Recovery",
    routePath: "/vendor-dashboard/recovery",
  },
  {
    name: "Logout",
    routePath: "/static/login",
  },
];

export const categorieMegaMenuItems = [
  {
    id: 1,
    menuCol: [
      {
        id: 1,
        megaBanner: "/img/backgrounds/new/2.png",
        title: "Things to do on your hotel",
        btnText: "See Hotel",
        btnRoute: "/hotel/hotel-list-v1",
        menuItems: [
          {
            id: 1,
            title: "",
            menuList: [
              {
                name: "Golfing",
                routePath: "/golf/golf-list-v1",
              }
            ],
          }
        ],
      },
    ],
  },
  {
    id: 2,
    menuCol: [
      {
        id: 1,
        megaBanner: "/img/backgrounds/new/4.png",
        title: "Things while hiking",
        btnText: "See hiking",
        btnRoute: "/golf/golf-list-v1",
        menuItems: [
          {
            id: 1,
            title: "",
            menuList: [
              {
                name: "5-Day High Atlas",
                routePath: "/golf/golf-list-v1",
              },
              {
                name: "4-Day Central Atlas Discovery",
                routePath: "/golf/golf-list-v1",
              }
            ],
          },
      
        ],
      },
    ],
  },
  {
    id: 3,
    menuCol: [
      {
        id: 1,
        megaBanner: "/img/backgrounds/new/5.png",
        title: "Things to do on your activity",
        btnText: "See Activity",
        btnRoute: "/activity/activity-list-v1",
        menuItems: [
          {
            id: 1,
            title: "",
            menuList: [
              {
                name: "Taghzout",
                routePath: "/activity/activity-list-v1",
              },
              {
                name: "Agadir",
                routePath: "/activity/activity-list-v2",
              },
            ],
          }
        ],
      },
    ],
  },
  {
    id: 4,
    menuCol: [
      {
        id: 1,
        megaBanner: "/img/backgrounds/new/5.png",
        title: "Things while kiting",
        btnText: "See Kitings",
        btnRoute: "/golf/golf-list-v1",
        menuItems: [
          {
            id: 1,
            title: "Kiting List",
            menuList: [
              {
                name: "5 days Essaouira, Marrakesh-Safi",
                routePath: "/golf/golf-list-v1",
              },
              {
                name: "8 days Levels in Dakhla",
                routePath: "/golf/golf-list-v1",
              },
            ],
          }
        ],
      },
    ],
  },

];

export const categorieMobileItems = [
  {
    id: 1,
    title: "Hotel",
    menuItems: [
      {
        id: 1,
        title: "Hotel List",
        menuList: [
          {
            name: "Hotel List v1",
            routePath: "/hotel/hotel-list-v1",
          },
          {
            name: "Hotel List v2",
            routePath: "/hotel/hotel-list-v2",
          },
          {
            name: "Hotel List v3",
            routePath: "/hotel/hotel-list-v3",
          },
          {
            name: "Hotel List v4",
            routePath: "/hotel/hotel-list-v4",
          },
          {
            name: "Hotel List v5",
            routePath: "/hotel/hotel-list-v5",
          },
        ],
      },
      {
        id: 2,
        title: "Hotel Single",
        menuList: [
          {
            name: "Hotel Single v1",
            routePath: "/hotel/hotel-single-v1/5",
          },
          {
            name: "Hotel Single v2",
            routePath: "/hotel/hotel-single-v2/5",
          },
        ],
      },
      {
        id: 3,
        title: "Hotel Booking",
        menuList: [
          {
            name: "Booking Page",
            routePath: "/hotel/booking-page",
          },
        ],
      },
    ],
  },
  {
    id: 2,
    title: "Tour",
    menuItems: [
      {
        id: 1,
        title: "Tour List",
        menuList: [
          {
            name: "Tour List v1",
            routePath: "/tour/tour-list-v1",
          },
          {
            name: "Tour List v2",
            routePath: "/tour/tour-list-v2",
          },
        ],
      },
      {
        id: 2,
        title: "Tour Pages",
        menuList: [
          {
            name: "Tour Map",
            routePath: "/tour/tour-list-v3",
          },
          {
            name: "Tour Single",
            routePath: "/tour/tour-single/5",
          },
        ],
      },
    ],
  },
  {
    id: 3,
    title: "Activity",
    menuItems: [
      {
        id: 1,
        title: "Activity List",
        menuList: [
          {
            name: "Activity List v1",
            routePath: "/activity/activity-list-v1",
          },
          {
            name: "Activity List v2",
            routePath: "/activity/activity-list-v2",
          },
        ],
      },
      {
        id: 2,
        title: "Activity Pages",
        menuList: [
          {
            name: "Activity Map",
            routePath: "/activity/activity-list-v3",
          },
          {
            name: "Activity Single",
            routePath: "/activity/activity-single/3",
          },
        ],
      },
    ],
  },
  {
    id: 4,
    title: "Hotel Rentals",
    menuItems: [
      {
        id: 1,
        title: "Rental List",
        menuList: [
          {
            name: "Rental List v1",
            routePath: "/rental/rental-list-v1",
          },
          {
            name: "Rental List v2",
            routePath: "/rental/rental-list-v2",
          },
        ],
      },
      {
        id: 2,
        title: "Rental Pages",
        menuList: [
          {
            name: "Rental Map",
            routePath: "/rental/rental-list-v3",
          },
          {
            name: "Rental Single",
            routePath: "/rental/rental-single/3",
          },
        ],
      },
    ],
  },
  {
    id: 5,
    title: "Car",
    menuItems: [
      {
        id: 1,
        title: "Car List",
        menuList: [
          {
            name: "Car List v1",
            routePath: "/car/car-list-v1",
          },
          {
            name: "Car List v2",
            routePath: "/car/car-list-v2",
          },
        ],
      },
      {
        id: 2,
        title: "Car Pages",
        menuList: [
          {
            name: "Car Map",
            routePath: "/car/car-list-v3",
          },
          {
            name: "Car Single",
            routePath: "/car/car-single/1",
          },
        ],
      },
    ],
  },
  {
    id: 6,
    title: "Cruise",
    menuItems: [
      {
        id: 1,
        title: "Cruise List",
        menuList: [
          {
            name: "Cruise List v1",
            routePath: "/cruise/cruise-list-v1",
          },
          {
            name: "Cruise List v2",
            routePath: "/cruise/cruise-list-v2",
          },
        ],
      },
      {
        id: 2,
        title: "Cruise Pages",
        menuList: [
          {
            name: "Cruise Map",
            routePath: "/cruise/cruise-list-v3",
          },
          {
            name: "Cruise Single",
            routePath: "/cruise/cruise-single/3",
          },
        ],
      },
    ],
  },
  {
    id: 7,
    title: "Flights",
    menuItems: [
      {
        id: 1,
        title: "Flight List",
        menuList: [
          {
            name: "Flight List v1",
            routePath: "/flight/flight-list-v1",
          },
        ],
      },
    ],
  },
];
