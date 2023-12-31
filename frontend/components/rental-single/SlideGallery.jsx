import Image from "next/image";
import { Gallery, Item } from "react-photoswipe-gallery";

const SlideGallery = (slides) => {
  console.log("slides: ", slides);

  return (
    <>
      <Gallery>
        <div className="galleryGrid -type-1 relative">
        {slides?.data?.gallery_img.map((slide, index) => {
                return (
                  <div className="galleryGrid__item" key={index}>
                  <Item original={slide} thumbnail={slide} width={1006} height={765}  key={index}>
                  {({ ref, open }) => (
                <Image
                  width={600}
                  height={500}
                  src={slide}
                  ref={ref}
                  onClick={open}
                  alt="image"
                  role="button"
                  className="rounded-4"
                />
              )}
                  </Item>
                  </div>
                );
              })}

          {/* End .galleryGrid__item */}

          <div className="galleryGrid__item relative">
            <Image
              width={450}
              height={375}
              src="/img/rentals/single/new/1.png"
              alt="image"
              className="rounded-4"
            />
            <div className="absolute h-full col-12 z-2 px-20 py-20 d-flex justify-end items-end bottom-0 end-0">
              <Item
                original="/img/rentals/single/new/1.png"
                thumbnail="/img/rentals/single/new/1.png"
                width={450}
                height={375}
              >
                {({ ref, open }) => (
                  <div
                    className="button -blue-1 px-24 py-15 bg-white text-dark-1 js-gallery"
                    ref={ref}
                    onClick={open}
                    role="button"
                  >
                    See All Photos
                  </div>
                )}
              </Item>
            </div>
          </div>
          {/* End .galleryGrid__item */}
        </div>
      </Gallery>
    </>
  );
};

export default SlideGallery;
