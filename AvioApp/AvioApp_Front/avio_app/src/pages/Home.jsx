import React, { useEffect } from "react";
import dayjs, { Dayjs } from "dayjs";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import { useState, useRef } from "react";
import classes from "./Flights.module.css";
import { useNavigate } from "react-router-dom";

const flights = [
  {
    id: 1,
    date: new Date(),
    duration: 1,
    start: "Paris",
    destination: "London",
    price: 50,
    tickets: 10,
    remainingTickets: 2,
  },
  {
    id: 2,
    date: new Date(),
    duration: 2,
    start: "Paris",
    destination: "London",
    price: 250,
    tickets: 20,
    remainingTickets: 10,
  },
  {
    id: 3,
    date: new Date(),
    duration: 3,
    start: "Paris",
    destination: "London",
    price: 150,
    tickets: 15,
    remainingTickets: 5,
  },
];

const Home = () => {
  const [value, setValue] = useState(dayjs(Date.now()));
  const [numberOfTickets, setNumberOfTickets] = useState(1);
  const startRef = useRef();
  const destRef = useRef();
  const navigate = useNavigate();

  const changeNumberHandler = () => {
    setNumberOfTickets(event.target.value);
  };

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

  const searchHandler = () => {
    fetch("http://localhost:5041/api/Flight/search", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer ",
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
        console.log(actualData.items);
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
          <table className={classes.styledTable}>
            <thead>
              <tr>
                <th>Date</th>
                <th>Duration</th>
                <th>Start</th>
                <th>Destination</th>
                <th>Remaining tickets</th>
                <th>Price</th>
                <th>Total Price</th>
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
                  <td>{flight.remainingTickets}</td>
                  <td>{flight.price}</td>
                  <td>{flight.price * numberOfTickets}</td>
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
        </div>
      </div>
    </div>
  );
};

export default Home;
