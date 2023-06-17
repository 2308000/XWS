import React from "react";
import { useState } from "react";

const AuthContext = React.createContext({
  token: "",
  id: "",
  role: "unauthorized",
  username: "",
  isLoggedIn: false,
  login: (id, role, username, token) => {},
  logout: () => {},
});

export const AuthContextProvider = (props) => {
  const tokenn = localStorage.getItem("token");
  const [token, setToken] = useState(tokenn);
  const [id, setId] = useState(localStorage.getItem("id"));
  const [role, setRole] = useState(localStorage.getItem("role"));
  const [username, setusername] = useState(localStorage.getItem("username"));
  const userIsLoggedIn = token != null ? true : false;

  const loginHandler = (id, role, username, token) => {
    setToken(token);
    setId(id);
    setRole(role);
    setusername(username);
    localStorage.setItem("id", id);
    localStorage.setItem("token", token);
    localStorage.setItem("role", role);
    localStorage.setItem("username", username);
  };
  const logoutHandler = () => {
    localStorage.removeItem("id");
    localStorage.removeItem("token");
    localStorage.removeItem("username");
    localStorage.removeItem("role");
    setusername(null);
    setRole(null);
    setToken(null);
  };

  const contextValue = {
    token: token,
    id: id,
    role: role,
    username: username,
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
