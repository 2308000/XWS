import React from "react";
import dayjs, { Dayjs } from "dayjs";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import { useState, useRef, useEffect, useContext } from "react";
import classes from "./Flights.module.css";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Modal from "@mui/material/Modal";
import AuthContext from "../store/auth-context";
import { useNavigate } from "react-router-dom";

const Flights = () => {
  const today = new Date();
  const tomorrow = new Date(today);
  tomorrow.setDate(tomorrow.getDate() + 1);
  const [value, setValue] = useState(dayjs(tomorrow));
  const [open, setOpen] = useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);
  const [selectedFlight, setSelectedFlight] = useState();
  const [numberOfTickets, setNumberOfTickets] = useState(1);
  const startRef = useRef();
  const destRef = useRef();
  const reqTicketRef = useRef();
  const [flights, setFlights] = useState();
  const authCtx = useContext(AuthContext);
  const navigate = useNavigate();
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

  const buyTicketHandler = () => {
    event.preventDefault();
    fetch(
      "https://localhost:5000/api/Ticket/" +
        selectedFlight.id +
        "/" +
        numberOfTickets,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: "Bearer " + authCtx.token,
        },
        body: JSON.stringify({}),
      }
    )
      .then((response) => response)
      .then((actualData) => {
        alert("Succes!");
        navigate("/reservations");
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
          ) : (
            <div>No flights </div>
          )}
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
                <span>{numberOfTickets}</span>
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
