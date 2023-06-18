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
          {authCtx.isLoggedIn && (
            <li className={classes.navListItem}>
              <NavLink to="/profile">Profile</NavLink>
            </li>
          )}
          {authCtx.role == "guest" && (
            <li className={classes.navListItem}>
              <NavLink to="/my-reservations">My reservations</NavLink>
            </li>
          )}
          {authCtx.role == "guest" && (
            <li className={classes.navListItem}>
              <NavLink to="/my-grades">My grades</NavLink>
            </li>
          )}
          {authCtx.isLoggedIn && (
            <li className={classes.navListItem}>
              <NavLink to="/accommodations">Accommodations</NavLink>
            </li>
          )}
          {authCtx.role == "host" && (
            <li className={classes.navListItem}>
              <NavLink to="/reservation-requests">Reservation requests</NavLink>
            </li>
          )}
          {authCtx.role == "host" && (
            <li className={classes.navListItem}>
              <NavLink to="/my-accommodations">My accommodations</NavLink>
            </li>
          )}
          {!authCtx.isLoggedIn && (
            <li>
              <NavLink to="/login">
                <button className={classes.loginButton}>Login</button>
              </NavLink>
            </li>
          )}
          {!authCtx.isLoggedIn && (
            <li>
              <NavLink to="/register">
                <button className={classes.registerButton}>Register</button>
              </NavLink>
            </li>
          )}
          {authCtx.isLoggedIn && (
            <li>
              <button
                className={classes.registerButton}
                onClick={logoutHandler}
              >
                Logout
              </button>
            </li>
          )}
        </ul>
      </div>
    </div>
  );
};

export default Navbar;
