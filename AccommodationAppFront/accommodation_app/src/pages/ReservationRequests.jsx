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

const ReservationRequests = () => {
  const authCtx = useContext(AuthContext);
  const [reservations, setReservations] = useState();
  const [refresh, setRefresh] = useState(false);

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
        setReservations(actualData.reservations);
      })
      .catch((error) => {
        alert(error);
      });
  }, [refresh]);

  const handleConfirm = (id) => {
    console.log(id);
    fetch("http://localhost:8000/reservation/approve/" + id, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: authCtx.token,
      },
    })
      .then((response) => response.json())
      .then((actualData) => {
        console.log(actualData);
        setRefresh(true);
      })
      .catch((error) => {
        alert(error);
      });
  };

  const handleCancel = (id) => {
    fetch("http://localhost:8000/reservation/reject/" + id, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: authCtx.token,
      },
      body: {},
    })
      .then((response) => response.json())
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
        <h1>Reservation requests</h1>

        <div className={classes.buttonContainerRight}></div>
        {reservations?.length > 0 ? (
          <table className={classes.styledTable}>
            <thead>
              <tr>
                <th>Name</th>
                <th>Start</th>
                <th>End</th>
                <th>Guests</th>
                <th>Guest</th>
                <th>Number of cancellations</th>
                <th></th>
                <th></th>
              </tr>
            </thead>
            <tbody>
              {reservations?.map((app) => (
                <tr key={app.accommodation.id}>
                  <td>{app.accommodation.name}</td>
                  <td>{dayjs(app.beginning).format("DD-MM-YYYY")}</td>
                  <td>{dayjs(app.ending).format("DD-MM-YYYY")}</td>
                  <td>{app.guests}</td>
                  <td>{app.user.username}</td>
                  <td>{app.user.cancellationCounter}</td>

                  <td>
                    <button
                      className={utils.greenTableButton}
                      onClick={() => {
                        handleConfirm(app.id);
                      }}
                    >
                      Accept
                    </button>
                  </td>
                  <td>
                    <button
                      className={utils.redTableButton}
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
        ) : (
          <div>
            <h2>No reservation requests</h2>
          </div>
        )}
      </div>
    </div>
  );
};

export default ReservationRequests;
