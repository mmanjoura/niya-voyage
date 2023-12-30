// const PropertyHighlights2 = () => {
//   const highlightsContent = [
//     {
//       id: 1,
//       icon: "icon-city",
//       text: `In London City Centre`,
//     },
//     {
//       id: 2,
//       icon: "icon-airplane",
//       text: `Airport transfer`,
//     },
//     {
//       id: 3,
//       icon: "icon-bell-ring",
//       text: `Front desk [24-hour]`,
//     },
//     {
//       id: 4,
//       icon: "icon-tv",
//       text: `Premium TV channels`,
//     },
//   ];
import { useEffect, useState } from "react";

const PropertyHighlights = ({ hotel }) => {
  const [highlights, setHighlights] = useState([]);

  useEffect(() => {
    setHighlights(hotel);
  }, []);

  

  if (!hotel) return null;
console.log(hotel);


  return (
    <div className="row y-gap-20 pt-30">
      {hotel?.hotel_info?.hotel_facility.map((item, id) => (
        <div className="col-lg-3 col-6" key={id}>
          <div className="text-center">
            <i className={`${item?.class_name} text-24 text-blue-1`} />
            <div className="text-15 lh-1 mt-10">{item?.facility_name}</div>
          </div>
        </div>
      ))}
    </div>
  );
};

export default PropertyHighlights;
