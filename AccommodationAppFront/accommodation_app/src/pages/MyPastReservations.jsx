import React from "react";
import classes from "./Flights.module.css";
import utils from "./Utils.module.css";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Modal from "@mui/material/Modal";
import { useState, useRef, useEffect, useContext } from "react";
import dayjs, { Dayjs } from "dayjs";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import AuthContext from "../store/auth-context";
import { useNavigate } from "react-router-dom";
const style = {
  position: "absolute",
  top: "50%",
  left: "50%",
  transform: "translate(-50%, -50%)",
  bgcolor: "background.paper",
  boxShadow: 24,
  borderRadius: 3,
};
const MyPastReservations = () => {
  const authCtx = useContext(AuthContext);
  const [open, setOpen] = useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);
  const [refresh, setRefresh] = useState(false);
  const [reservations, setReservations] = useState();
  const navigate = useNavigate();
  const [selectedAccommodation, setSelectedAccommodation] = useState();
  const gradeRef = useRef();

  useEffect(() => {
    fetch("http://localhost:8000/reservation/my/past", {
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
      });
  }, [refresh]);

  const gradeHandler = () => {
    fetch("http://localhost:8000/grade", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: authCtx.token,
      },
      body: JSON.stringify({
        guestId: authCtx.id,
        gradedId: selectedAccommodation.accommodation.id,
        value: gradeRef.current.value,
        isHostGrade: false,
      }),
    })
      .then((response) => response.json())
      .then((actualData) => {
        console.log(actualData);
      });
  };

  return (
    <div className={classes.body}>
      <div className={classes.home}>
        <br></br>
        <h1>My reservations</h1>

        <div className={classes.buttonContainerRight}></div>
        <table className={classes.styledTable}>
          <thead>
            <tr>
              <th>Name</th>
              <th>Start</th>
              <th>End</th>
              <th>Number of guests</th>
              <th>Status</th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            {reservations?.map((app) => (
              <tr key={app.id}>
                <td>{app.accommodation.name}</td>
                <td>{dayjs(app.beginning).format("DD-MM-YYYY")}</td>
                <td>{dayjs(app.ending).format("DD-MM-YYYY")}</td>
                <td>{app.guests}</td>
                <td>{app.reservationStatus == 1 ? "Approved" : "Pending"}</td>
                <td>
                  <button
                    className={utils.greenButton}
                    onClick={() => {
                      setSelectedAccommodation(app);
                      setOpen(true);
                    }}
                  >
                    Grade
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={style}>
          <div>
            <div className={classes.modalTitle}>Grade acoommodation</div>
            <div className={classes.register}>
              <div className={classes.grading}>
                <div>
                  <label>Grade: </label>
                  <input className={utils.input} ref={gradeRef}></input>
                </div>
              </div>
              <button className={classes.reserveButton} onClick={gradeHandler}>
                Confirm
              </button>
            </div>
          </div>
        </Box>
      </Modal>
    </div>
  );
};

export default MyPastReservations;
