import { Tab, Tabs, TabList, TabPanel } from "react-tabs";

const TermsConent = () => {
  return (
    <Tabs>
      <div className="row y-gap-30">
        <div className="col-lg-3">
          <div className="px-30 py-30 rounded-4 border-light">
            <TabList className="tabs__controls row y-gap-10 js-tabs-controls">
              <Tab className="col-12 tabs__button js-tabs-button">
                General Terms of Use
              </Tab>
              <Tab className="col-12 tabs__button js-tabs-button">
                Privacy policy
              </Tab>
              <Tab className="col-12 tabs__button js-tabs-button">
                Cookie Policy
              </Tab>
            </TabList>
          </div>
        </div>
        {/* End .col-lg-3 */}

        <div className="col-lg-9">
          <TabPanel>
            <div className="tabs__content js-tabs-content" data-aos="fade">
              <h1 className="text-30 fw-600 mb-15">General Terms of Use</h1>
              <h2 className="text-16 fw-500">1. Your Agreement</h2>
              <p className="text-15 text-dark-1 mt-5">
              Niya Voyage Provide Hosting Service agrees to provide hosting and distribution services for 
              the tourists activities and  packages offered by the Merchant Provider Client on the terms and conditions stated below:
                <br />
                <br />
                <h2 className="text-16 fw-500 mt-35">1. Hosting Services</h2>
                The Hosting Service will provide a platform for the Service Provider Client to display and market their packages.

                <h2 className="text-16 fw-500 mt-35">2. Distribution</h2>
                The Hosting Service will distribute the tour packages through various channels, 
                including but not limited to its website, affiliated partners, and promotional materials.

                <h2 className="text-16 fw-500 mt-35">3. Booking Management</h2>
                The Hosting Service will facilitate booking management systems, allowing customers to reserve and pay for the packages.

                <h2 className="text-16 fw-500 mt-35">4. Merchant Package Information</h2>
                The merchant Provider Client will provide accurate and up-to-date information regarding the packages, including pricing, 
                availability, itineraries, and terms and conditions.

                <h2 className="text-16 fw-500 mt-35">5. Customer Service</h2>
                The Merchant Provider Client will be responsible for addressing customer inquiries, handling bookings, 
                and ensuring customer satisfaction related to the provided offered packages.

                <h2 className="text-16 fw-500 mt-35">6. Compliance</h2>
                The Merchant Provider Client will comply with all laws, regulations, and industry standards related to the provision and advertisement of provided services.


                It was popularised in the 1960s with the release of Letraset
                sheets containing Lorem Ipsum passages, and more recently with
                desktop publishing software like Aldus PageMaker including
                versions of Lorem Ipsum.
              </p>
            
            </div>
          </TabPanel>
          {/* End  General Terms of Use */}

          <TabPanel>
          








            <div className="tabs__content js-tabs-content" data-aos="fade">
              <h1 className="text-30 fw-600 mb-15"> Privacy policy</h1>
              <h2 className="text-16 fw-500">1. Your Agreement</h2>
              <p className="text-15 text-dark-1 mt-5">
              This Privacy Policy outlines how Niya Voyage collects, uses, maintains, and protects information 
              gathered from users ("Users") of our website [www.niyavoyage.com] ("Site").
           
              </p>
              <h2 className="text-16 fw-500 mt-35">
                2. Information Collection
              </h2>
              <p className="text-15 text-dark-1 mt-5">
              We may collect personal identification information from Users in various ways, including, but not limited to, 
              when Users visit our Site, register on the Site, place an order, subscribe to the newsletter, respond to a survey, 
              fill out a form, or interact with other activities, services, features, or resources we make available on our Site. 
              Users may be asked for, as appropriate, name, email address, mailing address, phone number, credit card information, or other details
       
              </p>
              <h2 className="text-16 fw-500 mt-35">
                3. Information Usage
              </h2>
              <p className="text-15 text-dark-1 mt-5">
              Any information we collect from Users may be used in one or more of the following ways:
              <br />
                <br />

              To personalize user experience: We may use information to understand how our Users as a group use the services and resources provided on our Site.
              To improve our Site: We continually strive to improve our website offerings based on the information and feedback we receive from Users.
              To process transactions: We may use information Users provide about themselves when placing an order only to provide service to that order. 
              We do not share this information with outside parties except to the extent necessary to provide the service.
              To send periodic emails: We may use the email address to send information and updates pertaining to their order. 
              It may also be used to respond to their inquiries, questions, and/or other requests. If Users decide to opt-in to our mailing list, 
              they will receive emails that may include company news, updates, related product or service information, etc.
               </p>
            </div>
          </TabPanel>
          {/* End  Privacy policy */}

          <TabPanel>
            <div className="tabs__content js-tabs-content" data-aos="fade">
              <h1 className="text-30 fw-600 mb-15"> Cookie Policy</h1>
              <h2 className="text-16 fw-500">1. Your Agreement</h2>
              <p className="text-15 text-dark-1 mt-5">
              This Cookie Policy explains how [Your Company Name] ("we", "us", or "our") uses cookies and similar 
              technologies when you visit our website [www.niyavoyage.com] ("Site"). It explains what these technologies are and why we use them, 
              as well as your rights to control our use of them.
    
              </p>
              <h2 className="text-16 fw-500 mt-35">
                2. What Are Cookies
              </h2>
              <p className="text-15 text-dark-1 mt-5">
              Cookies are small text files that are stored on your computer or mobile device when you visit a website. 
              They are widely used to make websites work or to work more efficiently, 
              as well as to provide reporting information and assist with personalized advertising.
  
              </p>
              <h2 className="text-16 fw-500 mt-35">
                3. How We Use Cookies
              </h2>
              <p className="text-15 text-dark-1 mt-5">
              We use cookies for the following purposes:

              To ensure the proper functioning of our website.
              To analyze and improve the performance of our website.
              To provide personalized advertising.
         
              </p>
            </div>
          </TabPanel>
          {/* End  Cookie Policy */}

       
          {/* End  Best Price Guarantee */}
        </div>
        {/* End col-lg-9 */}
      </div>
    </Tabs>
  );
};

export default TermsConent;
