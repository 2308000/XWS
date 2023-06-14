import React from "react";
import dayjs, { Dayjs } from "dayjs";
import classes from "./Flights.module.css";
import { useNavigate } from "react-router-dom";
import { useEffect, useState, useContext } from "react";
import AuthContext from "../store/auth-context";

const Reservations = () => {
  const [flights, setFlights] = useState();
  const authCtx = useContext(AuthContext);

  useEffect(() => {
    fetch("https://localhost:5000/api/Ticket", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + authCtx.token,
      },
    })
      .then((response) => response.json())
      .then((actualData) => {
        console.log(actualData);
        setFlights(actualData);
      });
  }, []);

  return (
    <div>
      <div>
        <br></br>
        <div className={classes.tableContainer}>
          <table className={classes.styledTable}>
            <thead>
              <tr>
                <th>Date</th>

                <th>Start</th>
                <th>Destination</th>
                <th>Price</th>
              </tr>
            </thead>
            <tbody>
              {flights?.map((flight) => (
                <tr key={flight.id}>
                  <td>{dayjs(flight.date).format("DD.MM.YYYY")}</td>

                  <td>{flight.start}</td>
                  <td>{flight.destination}</td>
                  <td>{flight.price}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
};

export default Reservations;
