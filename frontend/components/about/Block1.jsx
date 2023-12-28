import Image from "next/image";

const Block1 = () => {
  return (
    <>
      <div className="col-lg-5">
        <h2 className="text-30 fw-600">About niyavoyage.com</h2>
        <p className="mt-5">These popular destinations have a lot to offer</p>
        <p className="text-dark-1 mt-60 lg:mt-40 md:mt-20">
         
        Morocco is a mesmerizing tapestry of beauty, blending vibrant colors, 
        intricate architecture, and diverse landscapes. From the bustling 
        markets of Marrakech to the serene Sahara dunes, it enchants with the kaleidoscope of its culture, 
        the majestic Atlas Mountains, and the ancient medinas that echo with history.
          <br />
          <br />
          The fusion of Berber, Arab, and European influences creates an alluring mosaic, 
          where every corner invites you to explore and indulge in its rich heritage and natural wonders.
        </p>
      </div>
      {/* End .col */}

      <div className="col-lg-6">
        <Image
          width={400}
          height={400}
          src="/img/pages/about/2.png"
          alt="image"
          className="rounded-4 w-100"
        />
      </div>
      {/* End .col */}
    </>
  );
};

export default Block1;
