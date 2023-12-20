import { useRouter } from "next/navigation";
import { useState } from "react";
const baseURL = process.env.NEXT_PUBLIC_API_URL;

const LocationInfo = () => {

  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);
  const [show, setShow] = useState(false);
  const [Address_Line_1, setAddress_Line_1] = useState("");
  const [Address_Line_2, setAddress_Line_2] = useState("");
  const [City, setCity] = useState("");
  const [State, setState] = useState("");
  const [Country, setCountry] = useState("");
  const [ZIP_Code, setZIP_Code] = useState("");

const handleAddLocation = async (e) => {
  e.preventDefault();
  setIsLoading(true);
  const location = {
    Address_Line_1,
    Address_Line_2,
    City,
    State,
    Country,
    ZIP_Code
  };
  const res = await fetch(`${baseURL}/locations`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(location),
  });
 if(res.status === 201){
   setIsLoading(false);
   router.push("/hotel-list");
   
}
}

  return (
    <form className="col-xl-9" onSubmit={handleAddLocation}>
      <div className="row x-gap-20 y-gap-20">
        <div className="col-12">
          <div className="form-input ">
            <input type="text" required onChange={(e) => setAddress_Line_1(e.target.value)} value={Address_Line_1}/>
            <label className="lh-1 text-16 text-light-1">Address Line 1</label>
          </div>
        </div>
        {/* End col-12 */}

        <div className="col-12">
          <div className="form-input ">
            <input type="text" required onChange={(e) => setAddress_Line_2(e.target.value)} value={Address_Line_2}/>
            <label className="lh-1 text-16 text-light-1">Address Line 2</label>
          </div>
        </div>
        {/* End col-12 */}

        <div className="col-md-6">
          <div className="form-input ">
            <input type="text" required onChange={(e) => setCity(e.target.value)} value={City}/>
            <label className="lh-1 text-16 text-light-1">City</label>
          </div>
        </div>
        {/* End col-6 */}

        <div className="col-md-6">
          <div className="form-input ">
            <input type="text" required onChange={(e) => setState(e.target.value)} value={State}/>
            <label className="lh-1 text-16 text-light-1">State</label>
          </div>
        </div>
        {/* End col-6 */}

        <div className="col-md-6">
          <div className="form-input ">
            <input type="text" required onChange={(e) => setCountry(e.target.value)} value={Country}/>
            <label className="lh-1 text-16 text-light-1">Select Country</label>
          </div>
        </div>
        {/* End col-6 */}

        <div className="col-md-6">
          <div className="form-input ">
            <input type="text" required onChange={(e) => setZIP_Code(e.target.value)} value={ZIP_Code}/>
            <label className="lh-1 text-16 text-light-1">ZIP Code</label>
          </div>
        </div>
        {/* End col-6 */}

        <div className="col-12">
          <div className="d-inline-block">
            <button
            disabled={isLoading}
              type="submit"
              className="button h-50 px-24 -dark-1 bg-blue-1 text-white"
              onClick={() => setShow(true)}
            >
              Save Changes <div className="icon-arrow-top-right ml-15" />
              {isLoading && <div className="spinner-border spinner-border-sm" role="status">Adding Hotel...</div>}
              {!isLoading && <div className="icon-arrow-top-right ml-15" > Add Hotel</div>}
            </button>
          </div>
        </div>
      </div>
    </form>
  );
};

export default LocationInfo;
