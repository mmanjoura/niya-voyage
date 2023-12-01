import dynamic from "next/dynamic";
import PopularDestinations from "@/components/destinations/PopularDestinations";
import ActivityFooter from "@/components/footer/activity-footer";
import ActivityHeader from "@/components/header/activity-header";
import ActivityHero from "@/components/hero/activity-hero";
import BlockGuide from "@/components/activity/BlockGuide";
import AddBanner from "@/components/activity/AddBanner";
import TourCategories from "@/components/activity/TourCategories";
import Activity from "@/components/activity/Activity";
import Blog from "@/components/activity/Blog";
import AppBanner from "@/components/activity/AppBanner";
import Testimonials from "@/components/activity/Testimonials";
import Activity2 from "@/components/activity/Activity2";

export const metadata = {
  title: "Home-6 || GoTrip - Travel & Tour React NextJS Template",
  description: "GoTrip - Travel & Tour React NextJS Template",
};

const activity = () => {
  return (
    <>
      {/* End Page Title */}

      <ActivityHeader />
      {/* End Header 6 */}

      <ActivityHero />
      {/* End Hero 6 */}

      <section className="layout-pt-md layout-pb-md bg-light-2">
        <div className="container">
          <div className="row y-gap-30">
            <BlockGuide />
          </div>
          {/* End .row */}
        </div>
        {/* End .container */}
      </section>
      {/* End Block Guide */}

      <section className="layout-pt-md layout-pb-lg">
        <div className="container">
          <div className="row justify-center text-center">
            <div className="col-auto">
              <div className="sectionTitle -md">
                <h2 className="sectionTitle__title">Special Offers</h2>
                <p className=" sectionTitle__text mt-5 sm:mt-0">
                  These popular destinations have a lot to offer
                </p>
              </div>
            </div>
          </div>
          {/* End .row */}
          <div className="row y-gap-20 pt-40">
            <AddBanner />
          </div>
          {/* End .row */}
        </div>
        {/* End container */}
      </section>
      {/* End Special Offer Section */}

      <section className="layout-pt-md layout-pb-md">
        <div className="container">
          <div className="row justify-center text-center is-in-view">
            <div className="col-auto">
              <div className="sectionTitle -md">
                <h2 className="sectionTitle__title">Trending Activity</h2>
                <p className=" sectionTitle__text mt-5 sm:mt-0">
                  Interdum et malesuada fames ac ante ipsum
                </p>
              </div>
            </div>
            {/* End .col */}
          </div>
          {/* End .row */}

          <div className="row y-gap-30 pt-40 sm:pt-20 item_gap-x30">
            <Activity />
          </div>
          {/* End .row */}
        </div>
        {/* End .container */}
      </section>
      {/* Trending Activity Sections */}

      <section className="layout-pt-md layout-pb-md">
        <div className="container">
          <div className="row justify-center text-center">
            <div className="col-auto">
              <div className="sectionTitle -md">
                <h2 className="sectionTitle__title">
                  Adventure &amp; Activity
                </h2>
                <p className=" sectionTitle__text mt-5 sm:mt-0">
                  Interdum et malesuada fames ac ante ipsum
                </p>
              </div>
            </div>
          </div>
          {/* End .row */}

          <div className="row y-gap-30 pt-40 sm:pt-20 item_gap-x30">
            <TourCategories />
          </div>
          {/* End .row */}
        </div>
        {/* End .container */}
      </section>
      {/* Adventure and activity */}

      <section className="layout-pt-lg layout-pb-md" data-aos="fade-up">
        <div className="container">
          <div className="row y-gap-20 justify-between items-end">
            <div className="col-auto">
              <div className="sectionTitle -md">
                <h2 className="sectionTitle__title">Popular Destinations</h2>
                <p className=" sectionTitle__text mt-5 sm:mt-0">
                  These popular destinations have a lot to offer
                </p>
              </div>
            </div>
            {/* End col-auto */}

            <div className="col-auto md:d-none">
              <a
                href="#"
                className="button -md -blue-1 bg-blue-1-05 text-blue-1"
              >
                View All Destinations
                <div className="icon-arrow-top-right ml-15" />
              </a>
            </div>
            {/* End col-auto */}
          </div>
          {/* End .row */}

          <div className="relative pt-40 sm:pt-20">
            <PopularDestinations />
          </div>
        </div>
        {/* End .container */}
      </section>
      {/* End Popular Destinations */}

      <section className="section-bg layout-pt-lg layout-pb-lg bg-light-2">
        <div className="container">
          <div className="row justify-center text-center">
            <div className="col-auto">
              <div className="sectionTitle -md">
                <h2 className="sectionTitle__title">Testimonials</h2>
                <p className=" sectionTitle__text mt-5 sm:mt-0">
                  Interdum et malesuada fames ac ante ipsum
                </p>
              </div>
            </div>
          </div>
          {/* End .row */}

          <div className="row justify-center pt-50 md:pt-30">
            <div className="col-xl-7 col-lg-10">
              <div className="overflow-hidden">
                <Testimonials />
              </div>
            </div>
            {/* End .col */}
          </div>
          {/* End .row */}
        </div>
        {/* End .container */}
      </section>
      {/* End Testimonials Section */}

      <section className="layout-pt-lg layout-pb-md">
        <div className="container">
          <div className="row justify-center text-center is-in-view">
            <div className="col-auto">
              <div className="sectionTitle -md">
                <h2 className="sectionTitle__title">Recommended Activity</h2>
                <p className=" sectionTitle__text mt-5 sm:mt-0">
                  Interdum et malesuada fames ac ante ipsum
                </p>
              </div>
            </div>
            {/* End .col */}
          </div>
          {/* End .row */}

          <div className="row y-gap-30 pt-40 sm:pt-20 item_gap-x30">
            <Activity2 />
          </div>
          {/* End .row */}
        </div>
        {/* End .container */}
      </section>
      {/* Trending Activity Sections */}

      <AppBanner />
      {/* End DownloadAppBanner section */}

      <section className="layout-pt-md layout-pb-lg">
        <div className="container">
          <div className="row justify-center text-center">
            <div className="col-auto">
              <div className="sectionTitle -md">
                <h2 className="sectionTitle__title">
                  Get inspiration for your next trip
                </h2>
                <p className=" sectionTitle__text mt-5 sm:mt-0">
                  Interdum et malesuada fames
                </p>
              </div>
            </div>
          </div>
          {/* End .row */}

          <div className="row y-gap-30 pt-40">
            <Blog />
          </div>
          {/* End .row */}
        </div>
        {/* End .container */}
      </section>
      {/* End blog sections */}

      <ActivityFooter />
      {/* End Footer Section */}
    </>
  );
};

export default dynamic(() => Promise.resolve(activity), { ssr: false });
