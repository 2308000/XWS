import React from "react";
import classes from "./Navbar.module.css";
import { NavLink } from "react-router-dom";
import AuthContext from "../store/auth-context";
import { useContext } from "react";
import { useNavigate } from "react-router-dom";

const Navbar = () => {
  const authCtx = useContext(AuthContext);
  const navigate = useNavigate();

  const logoutHandler = () => {
    authCtx.logout();
    navigate("/", { replace: true });
  };
  return (
    <div>
      <div className={classes.navbar}>
        <ul className={classes.list}>
          <li className={classes.navListItem}>
            <NavLink to="/profile">Profile</NavLink>
          </li>
          <li className={classes.navListItem}>
            <NavLink to="/my-reservations">My reservations</NavLink>
          </li>
          <li className={classes.navListItem}>
            <NavLink to="/accommodations">Accommodations</NavLink>
          </li>
          <li className={classes.navListItem}>
            <NavLink to="/reservation-requests">Reservation requests</NavLink>
          </li>
          <li className={classes.navListItem}>
            <NavLink to="/my-accommodations">My accommodations</NavLink>
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

          <li>
            <button className={classes.registerButton} onClick={logoutHandler}>
              Logout
            </button>
          </li>
        </ul>
      </div>
    </div>
  );
};

export default Navbar;
