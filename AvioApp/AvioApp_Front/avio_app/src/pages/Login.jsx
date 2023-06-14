import React from "react";
import classes from "./Login.module.css";
import { Link } from "react-router-dom";
import { useRef, useContext } from "react";
import AuthContext from "../store/auth-context";
import { useNavigate } from "react-router-dom";

const Login = () => {
  const emailRef = useRef();
  const permaRef = useRef();
  const pwRef = useRef();
  const authCtx = useContext(AuthContext);
  const navigate = useNavigate();
  const loginHandler = () => {
    event.preventDefault();

    fetch("https://localhost:5000/api/User/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        email: emailRef.current.value,
        password: pwRef.current.value,
        isLoginPermanent: permaRef.current.checked,
      }),
    })
      .then((res) => {
        if (res.ok) {
          console.log("c");
          return res.json();
        }
      })
      .then((data) => {
        console.log(data);
        const parsedJWT = parseJwt(data.token);
        authCtx.login(parsedJWT.role, parsedJWT.email, data.token);
        setTimeout(() => {
          navigateLogin(parsedJWT.role);
        }, 100);
      })
      .catch((error) => {
        alert("Wrong credentials!");
      });
  };

  function parseJwt(token) {
    if (!token) {
      return;
    }
    const base64Url = token.split(".")[1];
    const base64 = base64Url.replace("-", "+").replace("_", "/");
    return JSON.parse(window.atob(base64));
  }

  const navigateLogin = (role) => {
    if (role == "ADMIN") {
      navigate("/admin-flights", { replace: true });
    } else {
      navigate("/flights", { replace: true });
    }
  };

  return (
    <div className={classes.centerDiv}>
      <div>
        <form className={classes.login}>
          <div className={classes.span}>
            <label>Email</label>
            <input className={classes.input} ref={emailRef}></input>
          </div>
          <div className={classes.span}>
            <label>Password</label>
            <input className={classes.input} ref={pwRef}></input>
          </div>
          <div className={classes.keepLog}>
            <input type="checkbox" ref={permaRef}></input>
            <span>Keep me logged in </span>
          </div>
          <div className={classes.buttonContainerCenter}>
            <button className={classes.loginButton} onClick={loginHandler}>
              Login
            </button>
          </div>
          <span className={classes.registerSpan}>
            <Link to={"/register"}>Not a member? Register here!</Link>
          </span>
        </form>
      </div>
    </div>
  );
};

export default Login;
