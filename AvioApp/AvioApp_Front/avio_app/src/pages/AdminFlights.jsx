import React from "react";
import dayjs, { Dayjs } from "dayjs";
import classes from "./Flights.module.css";
import { useNavigate } from "react-router-dom";
import { useEffect, useContext, useState } from "react";
import AuthContext from "../store/auth-context";

const AdminFlights = () => {
  const navigate = useNavigate();
  const [flights, setFlights] = useState();
  const authCtx = useContext(AuthContext);
  const [refresh, setRefresh] = useState(false);

  useEffect(() => {
    fetch("https://localhost:5000/api/Flight", {
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
  }, [refresh]);

  const deleteFlightHandler = (id) => {
    console.log(id);
    fetch("https://localhost:5000/api/Flight/" + id, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + authCtx.token,
      },
    })
      .then((response) => response)
      .then((actualData) => {
        alert("Success!");
        setRefresh(true);
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
              {flights?.map((flight) => (
                <tr key={flight.id}>
                  <td>{dayjs(flight.date).format("DD.MM.YYYY")}</td>
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
