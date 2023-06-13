import React from "react";
import classes from "./Login.module.css";
import { Link } from "react-router-dom";
import { useRef } from "react";

const Login = () => {
  const emailRef = useRef();
  const permaRef = useRef();
  const pwRef = useRef();

  const loginHandler = () => {
    event.preventDefault();
    fetch("http://localhost:5041/api/User/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer ",
      },
      body: JSON.stringify({
        email: emailRef.current.value,
        password: pwRef.current.value,
        isLoginPermanent: permaRef.current.value,
      }),
    })
      .then((response) => response.json())
      .then((actualData) => {
        console.log(actualData.items);
      });
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
