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

const MyGrades = () => {
  const authCtx = useContext(AuthContext);
  const [open, setOpen] = useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);
  const [refresh, setRefresh] = useState(false);
  const [reservations, setReservations] = useState();
  const [HGrades, setHGrades] = useState();
  const navigate = useNavigate();
  const [selectedAccommodation, setSelectedAccommodation] = useState();
  const [selectedHost, setSelectedHost] = useState();
  const gradeRef = useRef();
  const gradeHostRef = useRef();
  const [toggleACC, setToggleACC] = useState(false);

  useEffect(() => {
    fetch(
      "http://localhost:8000/grade/guest/" + authCtx.id + "/accommodations",
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: authCtx.token,
        },
      }
    )
      .then((response) => response.json())
      .then((actualData) => {
        console.log(actualData);
        setReservations(actualData.grades);
      });
  }, [refresh]);

  const getHostGrades = () => {
    fetch("http://localhost:8000/grade/guest/" + authCtx.id + "/hosts", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: authCtx.token,
      },
    })
      .then((response) => response.json())
      .then((actualData) => {
        console.log(actualData);
        setHGrades(actualData.grades);
      });
  };

  const deleteGradeHandler = (id, type) => {
    fetch("http://localhost:8000/grade/" + id, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
        Authorization: authCtx.token,
      },
    })
      .then((response) => response.json())
      .then((actualData) => {
        console.log(actualData);
        if (type === "acc") {
          setRefresh(true);
        } else {
          getHostGrades();
        }
      });
  };

  const updateGradeHandler = (type) => {
    fetch(
      "http://localhost:8000/grade/" +
        (type === "acc" ? selectedAccommodation.id : selectedHost.id),
      {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          Authorization: authCtx.token,
        },
        body: JSON.stringify({
          id: type === "acc" ? selectedAccommodation.id : selectedHost.id,
          value:
            type === "acc"
              ? gradeRef.current.value
              : gradeHostRef.current.value,
        }),
      }
    )
      .then((response) => response.json())
      .then((actualData) => {
        console.log(actualData);
        if (type === "acc") {
          setRefresh(true);
        } else {
          getHostGrades();
        }
      });
  };

  return (
    <div className={classes.body}>
      {!toggleACC && (
        <div>
          <div className={classes.home}>
            <br></br>
            <h1>Accommodation grades</h1>
            <div className={classes.buttonContainerRight}>
              <button
                className={utils.lightBlueButton}
                onClick={() => {
                  setToggleACC(true);
                  getHostGrades();
                }}
              >
                Host grades
              </button>
            </div>
            <table className={classes.styledTable}>
              <thead>
                <tr>
                  <th>Name</th>
                  <th>Grade</th>
                  <th>Date</th>
                  <th></th>
                  <th></th>
                </tr>
              </thead>
              <tbody>
                {reservations?.map((app) => (
                  <tr key={app.id}>
                    <td>{app.gradedName}</td>
                    <td>{app.grade}</td>
                    <td>{dayjs(app.date).format("DD-MM-YYYY")}</td>
                    <td>
                      <button
                        className={utils.redButton}
                        onClick={() => {
                          deleteGradeHandler(app.id, "acc");
                        }}
                      >
                        Delete
                      </button>
                    </td>
                    <td>
                      <button
                        className={utils.blueButton}
                        onClick={() => {
                          setSelectedAccommodation(app);
                          setOpen(true);
                        }}
                      >
                        Change
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
                  <button
                    className={classes.reserveButton}
                    onClick={() => updateGradeHandler("acc")}
                  >
                    Confirm
                  </button>
                </div>
              </div>
            </Box>
          </Modal>
        </div>
      )}
      {toggleACC && (
        <div>
          <div className={classes.home}>
            <br></br>
            <h1>Host grades</h1>
            <div className={classes.buttonContainerRight}>
              <button
                className={utils.lightBlueButton}
                onClick={() => {
                  setToggleACC(false);
                }}
              >
                Accommodation grades
              </button>
            </div>
            <table className={classes.styledTable}>
              <thead>
                <tr>
                  <th>Name</th>
                  <th>Grade</th>
                  <th>Date</th>
                  <th></th>
                  <th></th>
                </tr>
              </thead>
              <tbody>
                {HGrades?.map((app) => (
                  <tr key={app.id}>
                    <td>{app.gradedName}</td>
                    <td>{app.grade}</td>
                    <td>{dayjs(app.date).format("DD-MM-YYYY")}</td>
                    <td>
                      <button
                        className={utils.redButton}
                        onClick={() => {
                          deleteGradeHandler(app.id, "host");
                        }}
                      >
                        Delete
                      </button>
                    </td>
                    <td>
                      <button
                        className={utils.blueButton}
                        onClick={() => {
                          setSelectedHost(app);
                          setOpen(true);
                        }}
                      >
                        Change
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
                <div className={classes.modalTitle}>Grade host</div>
                <div className={classes.register}>
                  <div className={classes.grading}>
                    <div>
                      <label>Grade: </label>
                      <input className={utils.input} ref={gradeHostRef}></input>
                    </div>
                  </div>
                  <button
                    className={classes.reserveButton}
                    onClick={() => updateGradeHandler("host")}
                  >
                    Confirm
                  </button>
                </div>
              </div>
            </Box>
          </Modal>
        </div>
      )}
    </div>
  );
};

export default MyGrades;
