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
          {authCtx.isLoggedIn && authCtx.role == "USER" && (
            <li className={classes.navListItem}>
              <NavLink to="/flights">Flights</NavLink>
            </li>
          )}

          {authCtx.isLoggedIn && authCtx.role == "USER" && (
            <li className={classes.navListItem}>
              <NavLink to="/reservations">Reservations</NavLink>
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
