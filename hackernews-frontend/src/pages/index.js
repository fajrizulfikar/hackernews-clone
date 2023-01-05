import * as React from "react";

import Hackernews from "../gifs/hackernews.gif";

const menus = ["new", "past", "comments", "ask", "show", "jobs", "submit"];

const IndexPage = () => {
  return (
    <div>
      <div
        className="navbar h-14 min-h-0"
        style={{ backgroundColor: "#ff6600" }}
      >
        <div className="navbar-start lg:w-0">
          <div className="dropdown">
            <label tabIndex={0} className="btn btn-ghost lg:hidden">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                className="h-5 w-5"
                fill="none"
                viewBox="0 0 24 24"
                stroke="#fff"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth="2"
                  d="M4 6h16M4 12h8m-8 6h16"
                />
              </svg>
            </label>
            <ul
              tabIndex={0}
              className="menu menu-compact dropdown-content mt-3 p-2 shadow bg-base-100 rounded-box w-52"
            >
              <li>
                <a>Item 1</a>
              </li>
              <li>
                <a>Item 2</a>
              </li>
            </ul>
          </div>
        </div>
        <div>
          <a className="btn btn-ghost normal-case border border-white">
            <img src={Hackernews} alt="hackernews logo" />
          </a>
        </div>
        <div className="navbar hidden lg:flex">
          <ul className="menu menu-horizontal px-1">
            <li>
              <span className="font-bold">Hacker News</span>
            </li>
            {menus.map((menu, index) => (
              <li className="flex flex-row items-center">
                <a>{menu}</a>
                {index !== menus.length - 1 && "|"}
              </li>
            ))}
          </ul>
        </div>
        <div className="navbar-end">
          <a className="btn btn-link text-white">Login</a>
        </div>
      </div>
    </div>
  );
};

export default IndexPage;
