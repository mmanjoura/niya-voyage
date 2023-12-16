import Link from "next/link";
import axios from "axios";
import React from "react";

const baseURL = process.env.NEXT_PUBLIC_API_URL;

export default function TopDestinations() {
  const [destinations, setDestinations] = React.useState(null);

  React.useEffect(() => {
    axios.get(baseURL+'/destinations').then((response) => {
      setDestinations(response.data);
    });
  }, []);

  if (!destinations) return null;

  return (
    <>
      {destinations.data.map((item) => (
        
        <div
          className={item.class}
          key={item.id}
          data-aos="fade"
          data-aos-delay={item.animation}
        >
          <Link
            href="/tour-list-v3"
            className="citiesCard -type-3 d-block h-full rounded-4 "
          >
            <div className="citiesCard__image ratio ratio-1:1">
              <img className="col-12 js-lazy" src={item.img} alt="image" />
            </div>
            <div className="citiesCard__content px-30 py-30">
              <h4 className="text-26 fw-600 text-white text-capitalize">
                {item.title}
              </h4>
              <div className="text-15 text-white">
                {item.properties} properties
              </div>
            </div>
          </Link>
        </div>
      ))}
    </>
  );
}