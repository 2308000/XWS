import React from "react";
import dayjs, { Dayjs } from "dayjs";
import classes from "./Flights.module.css";
import { useNavigate } from "react-router-dom";
import { useEffect } from "react";
const flights = [
  {
    ticketId: 1,
    flightId: 3,
    date: "2023-06-10T20:28:31.278Z",
    start: "Belgrade",
    destination: "Paris",
    price: 200,
  },
  {
    ticketId: 2,
    flightId: 3,
    date: "2023-06-10T20:28:31.278Z",
    start: "Belgrade",
    destination: "Paris",
    price: 200,
  },
  {
    ticketId: 3,
    flightId: 3,
    date: "2023-06-10T20:28:31.278Z",
    start: "Belgrade",
    destination: "Paris",
    price: 200,
  },
];

const Reservations = () => {
  useEffect(() => {
    fetch("http://localhost:5041/api/Ticket", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer ",
      },
    })
      .then((response) => response.json())
      .then((actualData) => {
        console.log(actualData.items);
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
                <th>Ticket</th>
                <th>Flight</th>
                <th>Start</th>
                <th>Destination</th>
                <th>Price</th>
              </tr>
            </thead>
            <tbody>
              {flights.map((flight) => (
                <tr key={flight.id}>
                  <td>{dayjs(flight.duration).format("DD.MM.YYYY")}</td>
                  <td>{flight.ticketId}</td>
                  <td>{flight.flightId}</td>
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
