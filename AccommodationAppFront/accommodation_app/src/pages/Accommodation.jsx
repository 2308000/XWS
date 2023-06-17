import React from "react";
import classes from "./Property.module.css";
import utils from "./Utils.module.css";
import dayjs, { Dayjs } from "dayjs";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import { useState, useRef, useEffect } from "react";
import Property from "../components/Property";
import { useNavigate } from "react-router-dom";
import { useParams } from "react-router-dom";
import AuthContext from "../store/auth-context";
import { useContext } from "react";
const Accommodation = () => {
  const [accommodation, setAccommodation] = useState();

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

  return (
    <div className={classes.body}>
      <br></br>
      <div className={classes.home}>
        <div className={classes.container}>
          <div className={classes.imgTitle}>
            <div className={classes.image}></div>
            <div>
              <h1>{accommodation?.accommodation.name}</h1>
              <h3>
                {accommodation?.accommodation.location.street},{" "}
                {accommodation?.accommodation.location.city},{" "}
                {accommodation?.accommodation.location.country}
              </h3>
              <h5>
                Lorem ipsum dolor sit amet consectetur adipisicing elit.
                Molestias, nulla? Porro vitae voluptatum rem esse possimus sunt
                soluta minima animi, hic corporis tempora at nesciunt minus odio
                qui quidem quam quasi est voluptatem praesentium culpa officiis
                eum officia. Dolor quam ducimus ex quaerat adipisci quae sequi
                impedit quidem, optio sapiente!
              </h5>
              <h3>
                From
                {dayjs(localStorage.getItem("startDate")).format("DD-MM-YYYY")}
                to
                {dayjs(localStorage.getItem("endDate")).format("DD-MM-YYYY")}
              </h3>
              <h3>
                Number of guests : {localStorage.getItem("numberOfGuests")}
              </h3>
              <h3>Price per night : {localStorage.getItem("pricePerNight")}</h3>
              <h3>Total price : {localStorage.getItem("totalPrice")}</h3>
            </div>
          </div>
          <div className={classes.priceDate}>
            <button className={utils.greenButton}>Reserve</button>
          </div>
        </div>
        <br></br>
        <h2>Benefits</h2>
        <div>
          <div>
            <span>Wifi: </span>
            <span>{accommodation?.accommodation.hasWifi ? "Yes" : "No"}</span>
          </div>
          <div>
            <span>Parking: </span>
            <span>
              {accommodation?.accommodation.hasParking ? "Yes" : "No"}
            </span>
          </div>
          <div>
            <span>Balcony: </span>
            <span>
              {accommodation?.accommodation.hasBalcony ? "Yes" : "No"}
            </span>
          </div>
          <div>
            <span>Washing machine: </span>
            <span>
              {accommodation?.accommodation.hasWashingMachine ? "Yes" : "No"}
            </span>
          </div>
          <div>
            <span>Kithcen facilities: </span>
            {accommodation?.accommodation.hasKitchen ? "Yes" : "No"}
          </div>
          <div>
            <span>Bathtub: </span>
            {accommodation?.accommodation.hasBathtub ? "Yes" : "No"}
          </div>
          <div>
            <span>Air: </span>
            {accommodation?.accommodation.hasBathtub ? "Yes" : "No"}
          </div>
        </div>
        <h2>Reviews</h2>
        <h2>
          Average grade :{" "}
          {Math.round(
            accommodation?.accommodation.averageAccommodationGrade,
            2
          )}
        </h2>
        <div>
          <div>
            <span>User </span>
            <span>Grade </span>
            <span>Date</span>
          </div>
          {accommodation?.accommodation.grades.map((grade) => (
            <div>
              <span>{grade.guestName} </span>
              <span>{grade.grade} </span>
              <span>{dayjs(grade.date).format("DD-MM-YYYY")}</span>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default Accommodation;
