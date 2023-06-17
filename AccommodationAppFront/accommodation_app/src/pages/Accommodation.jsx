import React from "react";
import classes from "./Property.module.css";
import utils from "./Utils.module.css";
import dayjs, { Dayjs } from "dayjs";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import { useState, useRef, useEffect, useContext } from "react";
import Property from "../components/Property";
import { useNavigate } from "react-router-dom";
const Accommodation = () => {
  return (
    <div className={classes.body}>
      <br></br>
      <div className={classes.home}>
        <div className={classes.container}>
          <div className={classes.imgTitle}>
            <div className={classes.image}></div>
            <div>
              <h1>Name</h1>
              <h3>Street, City, Country</h3>
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
            <span> Yes</span>
          </div>
          <div>
            <span>Parking: </span>
            <span> Yes</span>
          </div>
          <div>
            <span>Balcony: </span>
            <span> Yes</span>
          </div>
          <div>
            <span>Washing machine: </span>
            <span> Yes</span>
          </div>
          <div>
            <span>Kithcen facilities: </span>
            <span> Yes</span>
          </div>
          <div>
            <span>Bathtub: </span>
            <span> Yes</span>
          </div>
        </div>
        <h2>Reviews</h2>
        <div>
          <div>
            <span>ime </span>
            <span>ocena</span>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Accommodation;
