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

const MyReservations = () => {
  const handleCancel = () => {};

  return (
    <div className={classes.body}>
      <div className={classes.home}>
        <br></br>
        <h1>Accommodation name</h1>

        <div className={classes.buttonContainerRight}></div>
        <table className={classes.styledTable}>
          <thead>
            <tr>
              <th>Name</th>
              <th>City</th>
              <th>Start</th>
              <th>End</th>
              <th>Price</th>
              <th>price</th>
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
                    className={utils.redButton}
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

export default MyReservations;
