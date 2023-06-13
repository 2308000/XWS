import React from "react";
import dayjs, { Dayjs } from "dayjs";
import classes from "./Flights.module.css";
import { useNavigate } from "react-router-dom";
import { useEffect } from "react";

const flights = [
  {
    id: 1,
    date: new Date(),
    duration: 1,
    start: "bg",
    destination: "ns",
    price: 50,
    tickets: 10,
    remainingTickets: 2,
  },
  {
    id: 2,
    date: new Date(),
    duration: 2,
    start: "par",
    destination: "lnd",
    price: 250,
    tickets: 20,
    remainingTickets: 10,
  },
  {
    id: 3,
    date: new Date(),
    duration: 3,
    start: "tok",
    destination: "lsb",
    price: 150,
    tickets: 15,
    remainingTickets: 5,
  },
];

const AdminFlights = () => {
  const navigate = useNavigate();

  useEffect(() => {
    fetch("http://localhost:5041/api/Flight", {
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

  const deleteFlightHandler = (id) => {
    console.log(id);
    fetch("http://localhost:5041/api/Flight/" + id, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer ",
      },
    })
      .then((response) => response.json())
      .then((actualData) => {
        console.log(actualData.items);
      });
  };

  const addFlightHandler = () => {
    navigate("/new-flight");
  };

  return (
    <div>
      <div>
        <div className={classes.buttonContainerRight}>
          <button
            className={classes.addFlightButton}
            onClick={addFlightHandler}
          >
            Add
          </button>
        </div>
        <div className={classes.tableContainer}>
          <table className={classes.styledTable}>
            <thead>
              <tr>
                <th>Date</th>
                <th>Duration</th>
                <th>Start</th>
                <th>Destination</th>
                <th>Price</th>
                <th>Tickets</th>
                <th>Remaining Tickets</th>
                <th></th>
              </tr>
            </thead>
            <tbody>
              {flights.map((flight) => (
                <tr key={flight.id}>
                  <td>{dayjs(flight.duration).format("DD.MM.YYYY")}</td>
                  <td>{flight.duration}</td>
                  <td>{flight.start}</td>
                  <td>{flight.destination}</td>
                  <td>{flight.price}</td>
                  <td>{flight.tickets}</td>
                  <td>{flight.remainingTickets}</td>
                  <td>
                    <button
                      className={classes.deleteFlightButton}
                      onClick={() => deleteFlightHandler(flight.id)}
                    >
                      Delete
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
};

export default AdminFlights;
