import Link from "next/link";

import {
  homeItems,
  blogItems,
  pageItems,
  dashboardItems,
} from "../../data/mainMenuData";
import CategoriesMegaMenu from "./CategoriesMegaMenu";
import {
  isActiveParent,
  isActiveLink,
  isActiveParentChaild,
} from "../../utils/linkActiveChecker";
import { useRouter } from "next/router";

const MainMenu = ({ style = "" }) => {
  const router = useRouter();

  return (
    <nav className="menu js-navList">
      <ul className={`menu__nav ${style} -is-active`}>
      <li className="menu-item-has-children -has-mega-menu">
          <a href="#">
            <span className="mr-10">Activites</span>
            <i className="icon icon-chevron-sm-down" />
          </a>
          <div className="mega">
            <CategoriesMegaMenu />
          </div>
        </li>
        {/* End categories menu items */}
        <li
          className={`${
            isActiveParentChaild(homeItems, router.asPath) ? "current" : ""
          } menu-item-has-children`}
        >
          <a href="#">
            <span className="mr-10">Holidays</span>
            <i className="icon icon-chevron-sm-down" />
          </a>
          <ul className="subnav">
            {homeItems.map((menu, i) => (
              <li
                key={i}
                className={
                  isActiveLink(menu.routePath, router.asPath) ? "current" : ""
                }
              >
                <Link href={menu.routePath}>{menu.name}</Link>
              </li>
            ))}
          </ul>
        </li>
        {/* End home page menu */}

        <li className={router.pathname === "tour/tour-list-v1" ? "current" : ""}>
          <Link href="/tour/tour-list-v1">Tours</Link>
        </li>
        {/* End Tours single menu */}
        <li className={router.pathname === "/car/car-list-v1" ? "current" : ""}>
          <Link href="/car/car-list-v1">Cars</Link>
        </li>
  
        {/* End Cars single menu */}
        <li className={router.pathname === "/flight/flight-list-v1" ? "current" : ""}>
          <Link href="/flight/flight-list-v1">Flights</Link>
        </li>
        {/* End Flights single menu */}


        <li className={router.pathname === "/contact" ? "current" : ""}>
          <Link href="/contact">Contact</Link>
        </li>
      </ul>
    </nav>
  );
};

export default MainMenu;
