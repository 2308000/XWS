import React from "react";
import classes from "./Navbar.module.css";
import { NavLink } from "react-router-dom";
const Navbar = () => {
  return (
    <div>
      <div className={classes.navbar}>
        <ul className={classes.list}>
          <li className={classes.navListItem}>
            <NavLink to="/flights">Flights</NavLink>
          </li>
          <li className={classes.navListItem}>
            <NavLink to="/reservations">Reservations</NavLink>
          </li>

          <li>
            <NavLink to="/login">
              <button className={classes.loginButton}>Login</button>
            </NavLink>
          </li>
          <li>
            <NavLink to="/register">
              <button className={classes.registerButton}>Register</button>
            </NavLink>
          </li>
        </ul>
      </div>
    </div>
  );
};

export default Navbar;
