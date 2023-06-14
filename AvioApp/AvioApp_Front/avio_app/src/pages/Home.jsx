import React, { useEffect } from "react";
import dayjs, { Dayjs } from "dayjs";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import { useState, useRef } from "react";
import classes from "./Flights.module.css";
import { useNavigate } from "react-router-dom";

const Home = () => {
  const today = new Date();
  const tomorrow = new Date(today);
  tomorrow.setDate(tomorrow.getDate() + 1);
  const [value, setValue] = useState(dayjs(tomorrow));
  const [numberOfTickets, setNumberOfTickets] = useState(1);
  const startRef = useRef();
  const destRef = useRef();
  const navigate = useNavigate();
  const [flights, setFlights] = useState();

  const changeNumberHandler = () => {
    setNumberOfTickets(event.target.value);
  };

  useEffect(() => {
    fetch("https://localhost:5000/api/Flight/search", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        date: tomorrow,
        start: "",
        destination: "",
        requiredTickets: 1,
      }),
    })
      .then((response) => response.json())
      .then((actualData) => {
        console.log(actualData);
        setFlights(actualData);
      });
  }, []);

  const searchHandler = () => {
    if (numberOfTickets <= 0) {
      alert("You must enter at least 1 for required tickets");
      return;
    }
    fetch("https://localhost:5000/api/Flight/search", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        date: value,
        start: startRef.current.value,
        destination: destRef.current.value,
        requiredTickets: numberOfTickets,
      }),
    })
      .then((response) => response.json())
      .then((actualData) => {
        console.log(actualData);
        setFlights(actualData);
      });
  };

  return (
    <div className={classes.home}>
      <div>
        <div className={classes.filters}>
          <div className={classes.filter}>
            <label>Date</label>
            <LocalizationProvider dateAdapter={AdapterDayjs}>
              <DatePicker
                value={value}
                onChange={(newValue) => {
                  setValue(newValue);
                }}
                className={classes.DatePicker}
              />
            </LocalizationProvider>
          </div>
          <div className={classes.filter}>
            <label>From</label>
            <input className={classes.input} ref={startRef}></input>
          </div>
          <div className={classes.filter}>
            <label>To</label>
            <input className={classes.input} ref={destRef}></input>
          </div>
          <div className={classes.filter}>
            <label>Required tickets</label>
            <input
              className={classes.inputNumber}
              value={numberOfTickets}
              onChange={changeNumberHandler}
            ></input>
          </div>
          <button className={classes.searchButton} onClick={searchHandler}>
            Search
          </button>
        </div>

        <div className={classes.tableContainer}>
          {flights?.length > 0 ? (
            <table className={classes.styledTable}>
              <thead>
                <tr>
                  <th>Date</th>
                  <th>Duration</th>
                  <th>Start</th>
                  <th>Destination</th>
                  <th>Price</th>
                  <th>Total Price</th>
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
                    <td>
                      {numberOfTickets > 0
                        ? flight.price * numberOfTickets
                        : flight.price}
                    </td>
                    <td>
                      <button
                        className={classes.buyButton}
                        onClick={() => {
                          navigate("/login");
                        }}
                      >
                        Buy
                      </button>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          ) : (
            <div>No flights </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default Home;
