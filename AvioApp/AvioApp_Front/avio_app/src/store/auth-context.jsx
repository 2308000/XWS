import React from "react";
import { useState } from "react";

const AuthContext = React.createContext({
  token: "",
  role: "unauthorized",
  email: "",
  isLoggedIn: false,
  login: (role, email, token) => {},
  logout: () => {},
});

export const AuthContextProvider = (props) => {
  const tokenn = localStorage.getItem("token");
  const [token, setToken] = useState(tokenn);
  const [role, setRole] = useState(localStorage.getItem("role"));
  const [email, setEmail] = useState(localStorage.getItem("email"));
  const userIsLoggedIn = token != null ? true : false;

  const loginHandler = (role, email, token) => {
    setToken(token);
    setRole(role);
    setEmail(email);
    localStorage.setItem("token", token);
    localStorage.setItem("role", role);
    localStorage.setItem("email", email);
  };
  const logoutHandler = () => {
    localStorage.removeItem("token");
    localStorage.removeItem("email");
    localStorage.removeItem("role");
    setEmail(null);
    setRole(null);
    setToken(null);
  };

  const contextValue = {
    token: token,
    role: role,
    email: email,
    isLoggedIn: userIsLoggedIn,
    login: loginHandler,
    logout: logoutHandler,
  };

  return (
    <AuthContext.Provider value={contextValue}>
      {props.children}
    </AuthContext.Provider>
  );
};

export default AuthContext;
