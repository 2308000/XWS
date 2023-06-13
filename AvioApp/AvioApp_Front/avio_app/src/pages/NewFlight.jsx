import React, { useState, useRef } from "react";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DateTimePicker } from "@mui/x-date-pickers/DateTimePicker";
import dayjs from "dayjs";
import classes from "./Flights.module.css";

const NewFlight = () => {
  const [value, setValue] = useState(dayjs(Date.now()));
  const startRef = useRef();
  const destRef = useRef();
  const durationRef = useRef();
  const priceRef = useRef();
  const ticketsRef = useRef();

  const addFlightHandler = () => {
    event.preventDefault();
    fetch("http://localhost:5041/api/Flight", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer ",
      },
      body: JSON.stringify({
        date: value,
        duration: durationRef.current.value,
        start: startRef.current.value,
        destination: destRef.current.value,
        price: priceRef.current.value,
        tickets: ticketsRef.current.value,
      }),
    })
      .then((response) => response.json())
      .then((actualData) => {
        console.log(actualData.items);
      });
  };

  return (
    <div className={classes.centerDiv}>
      <div>
        <form className={classes.login}>
          <div className={classes.span}>
            <label>Date</label>
            <LocalizationProvider dateAdapter={AdapterDayjs}>
              <DateTimePicker
                value={value}
                ampm={false}
                onChange={(newValue) => {
                  setValue(newValue);
                }}
              />
            </LocalizationProvider>
          </div>
          <div className={classes.span}>
            <label>Duration</label>
            <input className={classes.newFlightInput} ref={durationRef}></input>
          </div>
          <div className={classes.span}>
            <label>Start</label>
            <input className={classes.newFlightInput} ref={startRef}></input>
          </div>
          <div className={classes.span}>
            <label>Destination</label>
            <input className={classes.newFlightInput} ref={destRef}></input>
          </div>
          <div className={classes.span}>
            <label>Ticket number</label>
            <input className={classes.newFlightInput} ref={ticketsRef}></input>
          </div>
          <div className={classes.span}>
            <label>Price</label>
            <input className={classes.newFlightInput} ref={priceRef}></input>
          </div>
          <div className={classes.buttonContainerCenter}>
            <button
              className={classes.addFlightButton}
              onClick={addFlightHandler}
            >
              Add
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default NewFlight;
