import Image from "next/image";
import { useState } from "react";
import ModalVideo from "react-modal-video";

const Block2 = () => {
  const [isOpen, setOpen] = useState(false);

  const expertContent = [
    {
      id: 1,
      icon: "/img/featureIcons/1/1.svg",
      title: "Best Price Guarantee",
      text: `Our Best Price Guarantee ensures you get the most competitive prices, offering peace of mind that you're receiving the best deal available`,
    },
    {
      id: 2,
      icon: "/img/featureIcons/1/2.svg",
      title: "Easy & Quick Booking",
      text: `Experience hassle-free and swift booking with our user-friendly platform, making your reservation process effortless and efficient.`,
    },
    {
      id: 3,
      icon: "/img/featureIcons/1/3.svg",
      title: "Customer Care 24/7",
      text: `Count on our round-the-clock customer care, ensuring assistance and support whenever you need it, day or night, for a seamless and reliable experience`,
    },
  ];

  return (
    <>
      <ModalVideo
        channel="youtube"
        autoplay
        isOpen={isOpen}
        videoId="qr8u0-iwkXI"
        onClose={() => setOpen(false)}
      />

      <div className="section-bg__item -right -image col-5 md:mb-60 sm:mb-40 d-flex z-2">
        <Image
          width={450}
          height={350}
          src="/img/backgrounds/new/10.png"
          alt="image"
        />
        <div className="absolute col-12 h-full flex-center z-1">
          <div
            onClick={() => setOpen(true)}
            className="d-flex items-center js-gallery"
            role="button"
          >
            <span className="button -outline-white text-white size-50 rounded-full flex-center">
              <i className="icon-play text-16" />
            </span>
            <span className="fw-500 text-white ml-15">Watch Video</span>
          </div>
        </div>
      </div>
      {/* End right video popup icon with image */}

      <div className="container">
        <div className="row">
          <div className="col-xl-4 col-md-7">
            <h2 className="text-30 fw-600">Why be a Local Expert</h2>
            <p className="mt-5">
              These popular destinations have a lot to offer
            </p>
            <div className="row y-gap-30 pt-60 md:pt-40">
              {expertContent.map((item) => (
                <div className="col-12" key={item.id}>
                  <div className="d-flex pr-30">
                    <Image
                      width={50}
                      height={50}
                      className="size-50"
                      src={item.icon}
                      alt="image"
                    />
                    <div className="ml-15">
                      <h4 className="text-18 fw-500">{item.title}</h4>
                      <p className="text-15 mt-10">{item.text}</p>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>
      </div>
      {/* End left local expert content */}
    </>
  );
};

export default Block2;
