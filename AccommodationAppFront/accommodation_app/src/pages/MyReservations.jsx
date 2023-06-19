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

const MyReservations = () => {
  const authCtx = useContext(AuthContext);
  const [refresh, setRefresh] = useState(false);
  const [reservations, setReservations] = useState();
  const navigate = useNavigate();
  useEffect(() => {
    fetch("http://localhost:8000/reservation/my/future", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: authCtx.token,
      },
    })
      .then((response) => response.json())
      .then((actualData) => {
        console.log(actualData);
        setReservations(actualData.reservations);
      });
  }, [refresh]);

  const handleGrade = (app) => {
    navigate();
  };

  const handleCancel = (id) => {
    fetch("http://localhost:8000/reservation/cancel/" + id, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: authCtx.token,
      },
    })
      .then((res) => {
        if (res.ok) {
          return res.json();
        } else if (res.status == 500) {
          throw new Error("Can not cancel, less than 1 day left");
        } else if (res.status == 404) {
          throw new Error("Account with this email does not exist");
        }
      })
      .then((actualData) => {
        console.log(actualData);
        setRefresh(true);
      })
      .catch((error) => {
        alert(error);
      });
  };

  return (
    <div className={classes.body}>
      <div className={classes.home}>
        <br></br>
        <h1>My reservations</h1>

        <div className={classes.buttonContainerRight}>
          <button
            className={utils.lightBlueButton}
            onClick={() => {
              navigate("/my-past-reservations");
            }}
          >
            Past reservations
          </button>
        </div>
        <table className={classes.styledTable}>
          <thead>
            <tr>
              <th>Name</th>
              <th>Start</th>
              <th>End</th>
              <th>Number of guests</th>
              <th>Status</th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            {reservations?.map((app) => (
              <tr key={app.id}>
                <td>{app.accommodation.name}</td>
                <td>{dayjs(app.beginning).format("DD-MM-YYYY")}</td>
                <td>{dayjs(app.ending).format("DD-MM-YYYY")}</td>
                <td>{app.guests}</td>
                <td>{app.reservationStatus == 1 ? "Approved" : "Pending"}</td>

                <td>
                  <button
                    className={utils.redButton}
                    onClick={() => {
                      handleCancel(app.id);
                    }}
                  >
                    Cancel
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default MyReservations;
