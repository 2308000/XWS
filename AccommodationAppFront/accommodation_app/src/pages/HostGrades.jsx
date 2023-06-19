import React from "react";
import classes from "./Flights.module.css";
import utils from "./Utils.module.css";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Modal from "@mui/material/Modal";
import { useState, useRef, useEffect, useContext } from "react";
import dayjs, { Dayjs } from "dayjs";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import AuthContext from "../store/auth-context";
import { useNavigate } from "react-router-dom";

const HostGrades = () => {
  const [grades, setGrades] = useState();
  const authCtx = useContext(AuthContext);

  useEffect(() => {
    fetch("http://localhost:8000/profile/" + authCtx.id, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: authCtx.token,
      },
    })
      .then((response) => response.json())
      .then((actualData) => {
        console.log(actualData);
        setGrades(actualData.profile.hostGrades);
      });
  }, []);

  return (
    <div className={classes.body}>
      <div>
        <div className={classes.home}>
          <br></br>
          <h1>My grades</h1>
          <br></br>
          <br></br>
          <table className={classes.styledTable}>
            <thead>
              <tr>
                <th>Name</th>
                <th>Grade</th>
                <th>Date</th>
              </tr>
            </thead>
            <tbody>
              {grades?.map((app) => (
                <tr key={app.id}>
                  <td>{app.hostName}</td>
                  <td>{app.grade}</td>
                  <td>{dayjs(app.date).format("DD-MM-YYYY")}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
};

export default HostGrades;
