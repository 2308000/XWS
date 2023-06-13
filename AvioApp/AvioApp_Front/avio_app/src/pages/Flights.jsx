import React from "react";
import dayjs, { Dayjs } from "dayjs";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import { useState, useRef, useEffect } from "react";
import classes from "./Flights.module.css";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Modal from "@mui/material/Modal";

const flights = [
  {
    id: 1,
    date: new Date(),
    duration: 1,
    start: "Belgrade",
    destination: "Paris",
    price: 50,
    tickets: 10,
    remainingTickets: 2,
  },
  {
    id: 2,
    date: new Date(),
    duration: 2,
    start: "Belgrade",
    destination: "Paris",
    price: 250,
    tickets: 20,
    remainingTickets: 10,
  },
  {
    id: 3,
    date: new Date(),
    duration: 3,
    start: "Belgrade",
    destination: "Paris",
    price: 150,
    tickets: 15,
    remainingTickets: 5,
  },
];

const Flights = () => {
  const [value, setValue] = useState(dayjs(Date.now()));
  const [open, setOpen] = useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);
  const [selectedFlight, setSelectedFlight] = useState();
  const [numberOfTickets, setNumberOfTickets] = useState(1);
  const startRef = useRef();
  const destRef = useRef();
  const reqTicketRef = useRef();

  const style = {
    position: "absolute",
    top: "50%",
    left: "50%",
    transform: "translate(-50%, -50%)",
    bgcolor: "background.paper",
    boxShadow: 24,
    borderRadius: 3,
  };

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
        requiredTickets: reqTicketRef.current.value,
      }),
    })
      .then((response) => response.json())
      .then((actualData) => {
        console.log(actualData.items);
      });
  };

  const buyTicketHandler = () => {
    event.preventDefault();
    fetch(
      "http://localhost:5041/api/Ticket/" +
        selectedFlight.id +
        "/" +
        numberOfTickets,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: "Bearer ",
        },
        body: JSON.stringify({}),
      }
    )
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
            <input className={classes.inputNumber} ref={reqTicketRef}></input>
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
                <th>Price</th>

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

                  <td>
                    <button
                      className={classes.buyButton}
                      onClick={() => {
                        handleOpen();
                        setSelectedFlight(flight);
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
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={style}>
          <div>
            <div className={classes.modalTitle}>Buy ticket</div>
            <form className={classes.register}>
              <div className={classes.spanReserve}>
                <label>Date:</label>
                <span>{dayjs(selectedFlight?.date).format("DD.MM.YYYY")}</span>
              </div>
              <div className={classes.spanReserve}>
                <label>From:</label>
                <span>{selectedFlight?.start}</span>
              </div>
              <div className={classes.spanReserve}>
                <label>Destination:</label>
                <span>{selectedFlight?.destination}</span>
              </div>
              <div className={classes.spanReserve}>
                <label>Price per ticket:</label>
                <span>{selectedFlight?.price}</span>
              </div>
              <div className={classes.spanReserve}>
                <label>Number of tickets: </label>
                <input
                  className={classes.inputNumber}
                  value={numberOfTickets}
                  onChange={changeNumberHandler}
                ></input>
              </div>
              <div className={classes.spanReserve}>
                <label>Total price:</label>
                <span>{selectedFlight?.price * numberOfTickets}</span>
              </div>
              <button
                className={classes.reserveButton}
                onClick={buyTicketHandler}
              >
                Buy
              </button>
            </form>
          </div>
        </Box>
      </Modal>
    </div>
  );
};

export default Flights;
