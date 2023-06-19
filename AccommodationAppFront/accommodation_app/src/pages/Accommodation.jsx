import React from "react";
import classes from "./Property.module.css";
import utils from "./Utils.module.css";
import dayjs, { Dayjs } from "dayjs";
import { useState, useRef, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useParams } from "react-router-dom";
import AuthContext from "../store/auth-context";
import { useContext } from "react";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Modal from "@mui/material/Modal";

const style = {
  position: "absolute",
  top: "50%",
  left: "50%",
  transform: "translate(-50%, -50%)",
  bgcolor: "background.paper",
  boxShadow: 24,
  borderRadius: 3,
};

const Accommodation = () => {
  const [accommodation, setAccommodation] = useState();
  const navigate = useNavigate();
  const [open, setOpen] = useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);
  const [HGrades, setHGrades] = useState();
  const [openGrades, setOpenGrades] = useState(false);
  const handleOpenGrades = () => setOpenGrades(true);
  const handleCloseGrades = () => setOpenGrades(false);
  let { id } = useParams();
  const authCtx = useContext(AuthContext);
  useEffect(() => {
    console.log(id);
    fetch("http://localhost:8000/accommodation/" + id, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: authCtx.token,
      },
    })
      .then((response) => response.json())
      .then((actualData) => {
        console.log(actualData);
        setAccommodation(actualData);
      });
  }, []);

  const getHostGrades = () => {
    console.log(accommodation);
    fetch(
      "http://localhost:8000/grade/graded/" +
        accommodation.accommodation.host.hostId,
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
        setHGrades(actualData.grades);
      });
  };

  const reserveHandler = () => {
    console.log(localStorage.getItem("startDate"));
    if (authCtx.role !== "guest") {
      alert("Log in as guest to reserve accommodation!");
      navigate("/login");
    } else {
      const start = dayjs(localStorage.getItem("startDate")).format(
        "YYYY-MM-DDTHH:mm:ss.SSSZ"
      );
      const end = dayjs(localStorage.getItem("endDate")).format(
        "YYYY-MM-DDTHH:mm:ss.SSSZ"
      );

      fetch("http://localhost:8000/reservation", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: authCtx.token,
        },
        body: JSON.stringify({
          accommodationId: id,
          beginning: start,
          ending: end,
          guests: localStorage.getItem("numberOfGuests"),
        }),
      })
        .then((res) => {
          if (res.ok) {
            return res.json();
          }
        })
        .then((data) => {
          if (data) console.log(data);
          navigate("/my-reservations");
        })
        .catch((error) => {
          alert(error);
        });
    }
  };

  const viewHostHandler = () => {};

  return (
    <div className={classes.body}>
      <br></br>
      <div className={classes.home}>
        <div className={classes.imgTitle}>
          <div className={classes.priceReserveContainer}>
            <div className={classes.reserve}>
              <h3>
                {localStorage.getItem("numberOfGuests")} guests, from{" "}
                {dayjs(localStorage.getItem("startDate")).format("DD.MM.YYYY")}{" "}
                to {dayjs(localStorage.getItem("endDate")).format("DD.MM.YYYY")}
              </h3>
            </div>
            <div className={classes.reserve}>
              <h3>Total price : EUR {localStorage.getItem("totalPrice")}</h3>
              <button className={utils.blueButton} onClick={reserveHandler}>
                Reserve
              </button>
            </div>
          </div>
          <div className={classes.topContainer}>
            <div className={classes.image}></div>
            <div className={classes.title}>
              <h1>{accommodation?.accommodation.name}</h1>
              <h3 className={classes.underlinedBlueText}>
                {accommodation?.accommodation.location.street},{" "}
                {accommodation?.accommodation.location.city},{" "}
                {accommodation?.accommodation.location.country}
              </h3>
              <br></br>
              <h5>
                Lorem ipsum dolor sit amet consectetur adipisicing elit.
                Molestias, nulla? Porro vitae voluptatum rem esse possimus sunt
                soluta minima animi, hic corporis tempora at nesciunt minus odio
                qui quidem quam quasi est voluptatem praesentium culpa officiis
                eum officia. Dolor quam ducimus ex quaerat adipisci quae sequi
                impedit quidem, optio sapiente!
              </h5>
              <br></br>
              <button
                className={utils.blueButton}
                onClick={() => {
                  getHostGrades();
                  setOpen(true);
                }}
              >
                Host details
              </button>
            </div>
          </div>
        </div>
        <div>
          <h2>Benefits</h2>
          <div class={classes.gridContainer}>
            <div class={classes.gridItem}>
              <span>Parking: </span>
              {accommodation?.accommodation.hasParking ? "Yes" : "No"}
            </div>
            <div class={classes.gridItem}>
              <span>Wifi: </span>
              {accommodation?.accommodation.hasWifi ? "Yes" : "No"}
            </div>
            <div class={classes.gridItem}>
              <span>Balcony: </span>
              {accommodation?.accommodation.hasBalcony ? "Yes" : "No"}
            </div>
            <div class={classes.gridItem}>
              <span>Washing machine: </span>
              {accommodation?.accommodation.hasWashingMachine ? "Yes" : "No"}
            </div>
            <div class={classes.gridItem}>
              <span>Kithcen facilities: </span>
              {accommodation?.accommodation.hasKitchen ? "Yes" : "No"}
            </div>
            <div class={classes.gridItem}>
              <span>Bathtub: </span>
              {accommodation?.accommodation.hasBathtub ? "Yes" : "No"}
            </div>
            <div class={classes.gridItem}>
              <span>Air Conditioning: </span>
              {accommodation?.accommodation.hasAirConditioning ? "Yes" : "No"}
            </div>
            <div class={classes.gridItem}>
              <span>TV: </span>
              <span>Yes</span>
            </div>
          </div>
        </div>
        <div>
          <div className={classes.allGradesContainer}>
            <h2>Last User Reviews</h2>
            <button
              className={utils.lightBlueButton}
              onClick={() => {
                setOpenGrades(true);
              }}
            >
              View All
            </button>
          </div>
          <br></br>
          <br></br>
          <div className={classes.gradesContainer}>
            {accommodation?.accommodation.grades.map((grade) => (
              <div className={classes.gradeContainer}>
                <h2>{grade.grade} </h2>
                <span>{grade.guestName} </span>

                <span>{dayjs(grade.date).format("dddd, DD.MM.YYYY")}</span>
              </div>
            ))}
          </div>
        </div>
      </div>
      <Modal
        open={openGrades}
        onClose={handleCloseGrades}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={style}>
          <div>
            <div className={classes.modalTitle}>Accomodation Reviews</div>
            <div className={classes.register}>
              <table className={classes.styledTable}>
                <thead>
                  <tr>
                    <th>User</th>
                    <th>Date</th>
                    <th>Grade</th>
                  </tr>
                </thead>
                <tbody>
                  {accommodation?.accommodation.grades.map((app, index) => (
                    <tr key={index}>
                      <td>{app.guestName}</td>
                      <td>{dayjs(app.date).format("DD-MM-YYYY")}</td>
                      <td>{app.grade}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
              <div>
                <button
                  className={utils.blueButton}
                  onClick={() => {
                    setOpenGrades(false);
                  }}
                >
                  Close
                </button>
              </div>
            </div>
          </div>
        </Box>
      </Modal>
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={style}>
          <div>
            <div className={classes.modalTitle}>Host Details</div>
            <div className={classes.register}>
              <h3>Host name: {accommodation?.accommodation.host.username}</h3>
              <h3>
                Phone number: {accommodation?.accommodation.host.phoneNumber}
              </h3>
              <h3>Grades</h3>
              <table className={classes.styledTable}>
                <thead>
                  <tr>
                    <th>User</th>
                    <th>Date</th>
                    <th>Grade</th>
                  </tr>
                </thead>
                <tbody>
                  {HGrades?.map((app, index) => (
                    <tr key={index}>
                      <td>{app.gradedName}</td>
                      <td>{dayjs(app.date).format("DD-MM-YYYY")}</td>
                      <td>{app.grade}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
              <div>
                <button
                  className={utils.blueButton}
                  onClick={() => {
                    setOpen(false);
                  }}
                >
                  Close
                </button>
              </div>
            </div>
          </div>
        </Box>
      </Modal>
    </div>
  );
};

export default Accommodation;
