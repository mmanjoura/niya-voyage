import dynamic from "next/dynamic";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import ModalVideo from "react-modal-video";
import "photoswipe/dist/photoswipe.css";
import { Gallery, Item } from "react-photoswipe-gallery";
// import { hotelsData } from "../../../data/hotels";
import Seo from "../../../components/common/Seo";
import Header from "../../../components/header/header";
import Overview from "../../../components/hotel-single/Overview";
import PopularFacilities from "../../../components/hotel-single/PopularFacilities";
import PropertyHighlights from "../../../components/hotel-single/PropertyHighlights";
import RatingTag from "../../../components/hotel-single/RatingTag";
import StickyHeader from "../../../components/hotel-single/StickyHeader";
import TopBreadCrumb from "../../../components/hotel-single/TopBreadCrumb";
import SidebarRight from "../../../components/hotel-single/SidebarRight";
import AvailableRooms from "../../../components/hotel-single/AvailableRooms";
import ReviewProgress from "../../../components/hotel-single/guest-reviews/ReviewProgress";
import DetailsReview from "../../../components/hotel-single/guest-reviews/DetailsReview";
import ReplyForm from "../../../components/hotel-single/ReplyForm";
import ReplyFormReview from "../../../components/hotel-single/ReplyFormReview";
import Facilities from "../../../components/hotel-single/Facilities";
import Image from "next/image";
import Surroundings from "../../../components/hotel-single/Surroundings";
import HelpfulFacts from "../../../components/hotel-single/HelpfulFacts";
import Faq from "../../../components/faq/Faq";
import Hotels2 from "../../../components/hotels/Hotels";
import CallToActions from "../../../components/common/CallToActions";
import DefaultFooter from "../../../components/footer/default";
import Link from "next/link";

import axios from "axios";
import React from "react";

const baseURL = process.env.NEXT_PUBLIC_API_URL;

const HotelSingleV1Dynamic = () => {

  const [isOpen, setOpen] = useState(false);
  const router = useRouter();
  const [hotel, setHotel] = useState({});
  const id = router.query.id;

  useEffect(() => {
    if (!id) {<h1>Loading...</h1>;}
    else 
    {
      
      axios.get(`${baseURL}/hotels/${id}`).then((response) => {
        setHotel(response.data);
      });
     
    }

    return () => {};
  }, [id]);
    console.log("hotel/hotel-single-v1/1", hotel.data)

if (!hotel) return null;
  return (
    <>
      <ModalVideo
        channel="youtube"
        autoplay
        isOpen={isOpen}
        videoId="oqNZOOWF8qM"
        onClose={() => setOpen(false)}
      />

      <Seo pageTitle="Hotel Single v1" />
      {/* End Page Title */}

      <div className="header-margin"></div>
      {/* header top margin */}

      <Header />
      {/* End Header 1 */}

      <TopBreadCrumb />
      {/* End top breadcrumb */}

      <StickyHeader hotel={hotel.data} />
      {/* sticky single header for hotel single */}

      <section className="pt-40">
        <div className="container">
          <div className="row y-gap-20 justify-between items-end">
            <div className="col-auto">
              <div className="row x-gap-20  items-center">
                <div className="col-auto">
                  <h1 className="text-30 sm:text-25 fw-600">{hotel?.data?.title}</h1>
                </div>
                {/* End .col */}
                <div className="col-auto">
                  <i className="icon-star text-10 text-yellow-1" />
                  <i className="icon-star text-10 text-yellow-1" />
                  <i className="icon-star text-10 text-yellow-1" />
                  <i className="icon-star text-10 text-yellow-1" />
                  <i className="icon-star text-10 text-yellow-1" />
                </div>
              </div>
              {/* End .row */}

              <div className="row x-gap-20 y-gap-20 items-center">
                <div className="col-auto">
                  <div className="d-flex items-center text-15 text-light-1">
                    <i className="icon-location-2 text-16 mr-5" />
                    {hotel?.data?.title}
                  </div>
                </div>
                <div className="col-auto">
                  <button
                    data-x-click="mapFilter"
                    className="text-blue-1 text-15 underline"
                  >
                    Show on map
                  </button>
                </div>
              </div>
              {/* End .row */}
            </div>
            {/* End .col */}

            <div className="col-auto">
              <div className="row x-gap-15 y-gap-15 items-center">
                <div className="col-auto">
                  <div className="text-14">
                    From{" "}
                    <span className="text-22 text-dark-1 fw-500">
                      US${hotel?.data?.price}
                    </span>
                  </div>
                </div>
                <div className="col-auto">
                  <Link
                    href="/hotel/booking-page"
                    className="button h-50 px-24 -dark-1 bg-blue-1 text-white"
                  >
                    Select Room <div className="icon-arrow-top-right ml-15" />
                  </Link>
                </div>
              </div>
            </div>
            {/* End .col */}
          </div>
          {/* End .row */}
      
          <Gallery>
         
            <div className="galleryGrid -type-1 pt-30">
            { hotel?.data?.gallery_img?.slice(0, 8).map((item, id) => (
              <>
              <div className="galleryGrid__item relative d-flex" key={id}>
                <Item
                  original={item}
                  thumbnail={item}
                  width={660}
                  height={660}
                >
                  {({ ref, open }) => (
                    <img
                      src={item}
                      ref={ref}
                      onClick={open}
                      alt="image"
                      role="button"
                      className="rounded-4"
                    />
                  )}
                </Item>
                <div className="absolute px-20 py-20 col-12 d-flex justify-end">
                  <button className="button -blue-1 size-40 rounded-full flex-center bg-white text-dark-1">
                    <i className="icon-heart text-16" />
                  </button>
                </div>
              </div>
     
             
              <div className="galleryGrid__item">
                <Item
                  original={item}
                  thumbnail={item}
                  width={450}
                  height={375}
                >
                  {({ ref, open }) => (
                    <img
                      ref={ref}
                      onClick={open}
                      src={item}
                      alt="image"
                      className="rounded-4"
                      role="button"
                    />
                  )}
                </Item>
              </div>
              </>
              ))}
      
            </div>
           
          </Gallery>
         
        </div>
        {/* End .container */}
      </section>
      {/* End gallery grid wrapper */}

      <section className="pt-30">
        <div className="container">
          <div className="row y-gap-30">
            <div className="col-xl-8">
              <div className="row y-gap-40">
                <div className="col-12">
                  <h3 className="text-22 fw-500">Property highlights</h3>
                  <PropertyHighlights  />
                </div>
                {/* End .col-12 Property highlights */}

                <div id="overview" className="col-12">
                  <Overview />
                </div>
                {/* End .col-12  Overview */}

                <div className="col-12">
                  <h3 className="text-22 fw-500 pt-40 border-top-light">
                    Most Popular Facilities
                  </h3>
                  <div className="row y-gap-10 pt-20">
                    <PopularFacilities  />
                  </div>
                </div>
                {/* End .col-12 Most Popular Facilities */}

                <div className="col-12">
                  <RatingTag />
                </div>
                {/* End .col-12 This property is in high demand! */}
              </div>
              {/* End .row */}
            </div>
            {/* End .col-xl-8 */}

            <div className="col-xl-4">
              <SidebarRight hotel={hotel} />
            </div>
            {/* End .col-xl-4 */}
          </div>
          {/* End .row */}
        </div>
        {/* End container */}
      </section>
      {/* End single page content */}

      <section id="rooms" className="pt-30">
        <div className="container">
          <div className="row pb-20">
            <div className="col-auto">
              <h3 className="text-22 fw-500">Available Rooms</h3>
            </div>
          </div>
          {/* End .row */}
          <AvailableRooms hotel={hotel} />
        </div>
        {/* End .container */}
      </section>
      {/* End Available Rooms */}

      <section className="pt-40" id="reviews">
        <div className="container">
          <div className="row">
            <div className="col-12">
              <h3 className="text-22 fw-500">Guest reviews</h3>
            </div>
          </div>
          {/* End .row */}

          <ReviewProgress />
          {/* End review with progress */}

          <div className="pt-40">
            <DetailsReview />
            {/* End review with details */}
          </div>

          <div className="row pt-30">
            <div className="col-auto">
              <a href="#" className="button -md -outline-blue-1 text-blue-1">
                Show all 116 reviews{" "}
                <div className="icon-arrow-top-right ml-15"></div>
              </a>
            </div>
          </div>
          {/* End .row */}
        </div>
        {/* End .container */}
        {/* End container */}
      </section>
      {/* End Review section */}

      <section className="pt-40">
        <div className="container">
          <div className="row">
            <div className="col-xl-8 col-lg-10">
              <div className="row">
                <div className="col-auto">
                  <h3 className="text-22 fw-500">Leave a Reply</h3>
                  <p className="text-15 text-dark-1 mt-5">
                    Your email address will not be published.
                  </p>
                </div>
              </div>
              {/* End .row */}

              <ReplyFormReview />
              {/* End ReplyFormReview */}

              <ReplyForm />
            </div>
          </div>
        </div>
      </section>
      {/* End Reply Comment box section */}

      <section className="mt-40" id="facilities">
        <div className="container">
          <div className="row x-gap-40 y-gap-40">
            <div className="col-12">
              <h3 className="text-22 fw-500">Facilities of this Hotel</h3>
              <div className="row x-gap-40 y-gap-40 pt-20">
                <Facilities />
              </div>
              {/* End .row */}
            </div>
            {/* End .col-12 */}
          </div>
          {/* End .row */}
        </div>
        {/* End .container */}
      </section>
      {/* End facilites section */}

      <section className="pt-40">
        <div className="container">
          <div className="row">
            <div className="col-12">
              <div className="px-24 py-20 rounded-4 bg-light-2">
                <div className="row x-gap-20 y-gap-20 items-center">
                  <div className="col-auto">
                    <div className="flex-center size-60 rounded-full bg-white">
                      <Image
                        width={30}
                        height={30}
                        src="/img/icons/health.svg"
                        alt="icon"
                      />
                    </div>
                  </div>
                  <div className="col-auto">
                    <h4 className="text-18 lh-15 fw-500">
                      Extra health &amp; safety measures
                    </h4>
                    <div className="text-15 lh-15">
                      This property has taken extra health and hygiene measures
                      to ensure that your safety is their priority
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
      {/* End health &  safety measures section */}

      <section className="pt-40">
        <div className="container">
          <div className="row">
            <div className="col-12">
              <h3 className="text-22 fw-500">Hotel surroundings</h3>
            </div>
          </div>
          {/* End .row */}

          <div className="row x-gap-50 y-gap-30 pt-20">
            <Surroundings />
          </div>
          {/* End .row */}
        </div>
        {/* End .container */}
      </section>
      {/* End hotel surroundings */}

      <section className="pt-40">
        <div className="container">
          <div className="pt-40 border-top-light">
            <div className="row">
              <div className="col-12">
                <h3 className="text-22 fw-500">Some helpful facts</h3>
              </div>
            </div>
            {/* End .row */}

            <div className="row x-gap-50 y-gap-30 pt-20">
              <HelpfulFacts />
            </div>
            {/* End .row */}
          </div>
          {/* End .pt-40 */}
        </div>
        {/* End .container */}
      </section>
      {/* End helpful facts surroundings */}

      <section id="faq" className="pt-40 layout-pb-md">
        <div className="container">
          <div className="pt-40 border-top-light">
            <div className="row y-gap-20">
              <div className="col-lg-4">
                <h2 className="text-22 fw-500">
                  FAQs about
                  <br /> The Crown Hotel
                </h2>
              </div>
              {/* End .row */}

              <div className="col-lg-8">
                <div className="accordion -simple row y-gap-20 js-accordion">
                  <Faq />
                </div>
              </div>
              {/* End .col */}
            </div>
            {/* End .row */}
          </div>
          {/* End .pt-40 */}
        </div>
        {/* End .container */}
      </section>
      {/* End Faq about sections */}

      <section className="layout-pt-md layout-pb-lg">
        <div className="container">
          <div className="row justify-center text-center">
            <div className="col-auto">
              <div className="sectionTitle -md">
                <h2 className="sectionTitle__title">
                  Popular properties similar to The Crown Hotel
                </h2>
                <p className=" sectionTitle__text mt-5 sm:mt-0">
                  Interdum et malesuada fames ac ante ipsum
                </p>
              </div>
              {/* End sectionTitle */}
            </div>
            {/* End .col */}
          </div>
          {/* End .row */}

          <div className="pt-40 sm:pt-20 item_gap-x30">
            <Hotels2 />
          </div>
          {/* End slide hotel */}
        </div>
        {/* End .container */}
      </section>
      {/* End similar hotel */}

      <CallToActions />
      {/* End Call To Actions Section */}

      <DefaultFooter />
    </>
  );
};

export default dynamic(() => Promise.resolve(HotelSingleV1Dynamic), {
  ssr: false,
});
