import React from "react";
import classes from "./Flights.module.css";
import utils from "./Utils.module.css";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Modal from "@mui/material/Modal";
import { useState, useRef, useEffect } from "react";
import dayjs, { Dayjs } from "dayjs";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import AuthContext from "../store/auth-context";
import { useContext } from "react";

const reservations = [
  {
    Id: 1,
    Host: 2,
    Name: "penthaus neki",
    HasWifi: true,
    HasFreeParking: true,
    HasWashingMachine: true,
    MinNumberOfGuests: 2,
    MaxNumberOfGuests: 7,
    Availability: [
      {
        Price: 20,
        IsPricePerGuest: true,
      },
      {
        Price: 60,
        IsPricePerGuest: true,
      },
      {
        IsPricePerGuest: true,
      },
      {
        Price: 40,
        IsPricePerGuest: true,
      },
    ],
    IsReservationAcceptenceManual: true,
  },
  {
    Id: 1,
    Host: 2,
    Name: "penthaus neki",
    HasWifi: true,
    HasFreeParking: true,
    HasWashingMachine: true,
    MinNumberOfGuests: 2,
    MaxNumberOfGuests: 7,
    Availability: [
      {
        Price: 20,
        IsPricePerGuest: true,
      },
      {
        Price: 60,
        IsPricePerGuest: true,
      },
      {
        Price: 55,
        IsPricePerGuest: true,
      },
      {
        Price: 40,
        IsPricePerGuest: true,
      },
    ],
    IsReservationAcceptenceManual: true,
  },
];

const ReservationRequests = () => {
  const authCtx = useContext(AuthContext);
  const [properties, setProperties] = useState();

  useEffect(() => {
    console.log(authCtx.token);

    fetch("http://localhost:8000/reservation/host", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: authCtx.token,
      },
    })
      .then((response) => response.json())
      .then((actualData) => {
        console.log(actualData);
        setProperties(actualData.reservations);
      })
      .catch((error) => {
        alert(error);
      });
  }, []);

  const handleCancel = () => {};

  return (
    <div className={classes.body}>
      <div className={classes.home}>
        <br></br>
        <h1>Reservation requests</h1>

        <div className={classes.buttonContainerRight}></div>
        <table className={classes.styledTable}>
          <thead>
            <tr>
              <th>Name</th>
              <th>Start</th>
              <th>End</th>
              <th>Guests</th>
              <th>Number of cancellations</th>
              <th></th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            {reservations?.map((app) => (
              <tr key={app.id}>
                <td>{app.Name}</td>
                <td>{app.City}</td>
                <td>{app.Start}</td>
                <td>{app.End}</td>
                <td>{app.Price}</td>
                <td>
                  <button
                    className={utils.greenTableButton}
                    onClick={() => {
                      handleCancel();
                    }}
                  >
                    Accept
                  </button>
                </td>
                <td>
                  <button
                    className={utils.redTableButton}
                    onClick={() => {
                      handleCancel();
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

export default ReservationRequests;
