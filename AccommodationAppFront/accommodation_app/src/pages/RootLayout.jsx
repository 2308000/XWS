import React from "react";

import { Outlet } from "react-router-dom";
import Navbar from "../components/Navbar";

const RootLayout = () => {
  return (
    <>
      <div>
        <Navbar></Navbar>
        <main>
          <Outlet></Outlet>
        </main>
      </div>
    </>
  );
};

export default RootLayout;
