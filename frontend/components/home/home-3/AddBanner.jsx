import Image from "next/image";
import Link from "next/link";

import axios from "axios";
import React from "react";

const baseURL = process.env.NEXT_PUBLIC_API_URL;


export default function AddBanner() {
  const [addContent, setAddContent] = React.useState(null);

  React.useEffect(() => {
    axios.get(baseURL+'/addBanners').then((response) => {
      setAddContent(response.data);
    });
  }, []);
console.log("Home, AddBanner: ", addContent?.data)
  if (!addContent) return null;

  return (
    <>
      {addContent.data.map((item) => (
        <div
          className="col-lg-4 col-sm-6"
          data-aos="fade"
          data-aos-delay={item.delayAnimation}
          key={item.id}
        >
          <div className="ctaCard -type-1 rounded-4 ">
            <div className="ctaCard__image ratio ratio-41:45">
              <Image
                width={410}
                height={455}
                className="js-lazy img-ratio"
                src={item.img}
                alt="image"
              />
            </div>
            <div className="ctaCard__content py-50 px-50 lg:py-30 lg:px-30">
              {item.meta ? (
                <>
                  <div className="text-15 fw-500 text-white mb-10">
                    Enjoy Summer Deals
                  </div>
                </>
              ) : (
                ""
              )}

              <h4 className="text-30 lg:text-24 text-white">{item.title}</h4>
              <div className="d-inline-block mt-30">
                <Link
                  href={item.routerPath}
                  className="button px-48 py-15 -blue-1 -min-180 bg-white text-dark-1"
                >
                  Experiences
                </Link>
              </div>
            </div>
          </div>
        </div>
      ))}
    </>
  );
};


