import { useEffect, useState } from "react";

const PopularFacilities = ({ hotel }) => {
  const [facilities, setFacilities] = useState([]);

  useEffect(() => {
    setFacilities(hotel);
  }, []);

  

  if (!hotel) return null;
console.log(hotel);
  return (
    <>
      {hotel?.hotel_info?.hotel_facility.map((item) => (

        <div className="col-md-5" key={item.id}>
          <div className="d-flex x-gap-15 y-gap-15 items-center">
            <i className={`${item?.class_name} text-24 text-blue-1`}></i>
            <div className="text-15">{item?.facility_name}</div>
          </div>
        </div>

      ))}
    </>
  );
};

export default PopularFacilities;
