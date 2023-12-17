import Image from "next/image";
import Link from "next/link";
//import flightsData from "../../data/flights";



import axios from "axios";
import React from "react";

const baseURL = process.env.NEXT_PUBLIC_API_URL;
var  flightList = [
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
]

export default function TopFlights() {
  const [flights, setFlights] = React.useState(null);

  React.useEffect(() => {
    axios.get(baseURL+'/flights').then((response) => {
      setFlights(response.data);
    });
  }, []);

  if (!flights) return null;

  return (
    <>
      {flights.data.slice(0, 4).map((item) => (
        <div
          className="col-12"
          key={item?.id}
          data-aos="fade"
          data-aos-delay={item?.animation}
        >
          <div className="px-20 py-20 rounded-4 border-light">
            <div className="row y-gap-30 justify-between xl:justify-">
              {flightList?.map((flight) => (
                <div className="col-xl-4 col-lg-6" key={flight.id}>
                  <div className="row y-gap-10 items-center">
                    <div className="col-sm-auto">
                      <Image
                        width={40}
                        height={40}
                        className="size-40"
                        src={flight?.avatar}
                        alt="image"
                      />
                    </div>
                    <div className="col">
                      <div className="row x-gap-20 items-end">
                        <div className="col-auto">
                          <div className="lh-15 fw-500">
                            {flight.departureTime}
                          </div>
                          <div className="text-15 lh-15 text-light-1">
                            {flight.arrivalAirport}
                          </div>
                        </div>
                        <div className="col text-center">
                          <div className="flightLine">
                            <div />
                            <div />
                          </div>
                          <div className="text-15 lh-15 text-light-1 mt-10">
                            {flight.duration}
                          </div>
                        </div>
                        <div className="col-auto">
                          <div className="lh-15 fw-500">
                            {flight.arrivalTime}
                          </div>
                          <div className="text-15 lh-15 text-light-1">
                            {flight.departureAirport}
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              ))}

              <div className="col-auto">
                <div className="d-flex items-center">
                  <div className="text-right mr-24">
                    <div className="lh-15 fw-500">US${item?.price}</div>
                    <div className="text-15 lh-15 text-light-1">
                      {item?.deals} deals
                    </div>
                  </div>
                  <Link
                    href="/flight/flight-list-v1"
                    className="button -outline-blue-1 px-30 h-50 text-blue-1"
                  >
                    View Deal <div className="icon-arrow-top-right ml-15" />
                  </Link>
                </div>
              </div>
              {/* End .col */}
            </div>
            {/* End .row */}
          </div>
          {/* End px-20 */}
        </div>
      ))}
    </>
  );
};
