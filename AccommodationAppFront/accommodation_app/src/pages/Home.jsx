import React from "react";
import classes from "./Accommodation.module.css";
import utils from "./Utils.module.css";
import dayjs, { Dayjs } from "dayjs";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import { useState, useRef, useEffect, useContext } from "react";

const properties = [
  {
    Id: 1,
    Host: 2,
    Name: "penthaus neki",
    HasWifi: true,
    HasFreeParking: true,
    HasWashingMachine: true,
    MinNumberOfGuests: 2,
    MaxNumberOfGuests: 7,
    Availability: [
      {
        Price: 20,
        IsPricePerGuest: true,
      },
      {
        Price: 60,
        IsPricePerGuest: true,
      },
      {
        IsPricePerGuest: true,
      },
      {
        Price: 40,
        IsPricePerGuest: true,
      },
    ],
    IsReservationAcceptenceManual: true,
  },
  {
    Id: 1,
    Host: 2,
    Name: "penthaus neki",
    HasWifi: true,
    HasFreeParking: true,
    HasWashingMachine: true,
    MinNumberOfGuests: 2,
    MaxNumberOfGuests: 7,
    Availability: [
      {
        Price: 20,
        IsPricePerGuest: true,
      },
      {
        Price: 60,
        IsPricePerGuest: true,
      },
      {
        Price: 55,
        IsPricePerGuest: true,
      },
      {
        Price: 40,
        IsPricePerGuest: true,
      },
    ],
    IsReservationAcceptenceManual: true,
  },
];

const Home = () => {
  const today = new Date();
  const tomorrow = new Date(today);
  tomorrow.setDate(tomorrow.getDate() + 1);
  const [value, setValue] = useState(dayjs(tomorrow));
  const startRef = useRef();
  const destRef = useRef();
  const reqTicketRef = useRef();

  return (
    <div className={classes.body}>
      <br></br>
      <div className={classes.home}>
        <div className={classes.search}>
          <div className={classes.searchItem}>
            <label>Location</label>
            <input className={classes.inputLoc} ref={startRef}></input>
          </div>
          <div className={classes.searchItem}>
            <label>From</label>
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
          <div className={classes.searchItem}>
            <label>To</label>
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
          <div className={classes.searchItem}>
            <label>Number of guests</label>
            <input className={classes.inputNumber}></input>
          </div>
          <button className={classes.searchButton}>Search</button>
        </div>
        <div className={classes.accomodationsContainer}>
          <div className={classes.filters}>
            <h4>Filter by:</h4>
            <h5>Price range: </h5>
            <div className={classes.filterPrice}>
              <label>From</label>
              <input className={classes.inputPrice}></input>
            </div>
            <div className={classes.filterPrice}>
              <label>To</label>
              <input className={classes.inputPrice}></input>
            </div>
            <h5>Benefits: </h5>
            <div className={classes.filter}>
              <input type="checkbox"></input>
              <label>Free Parking</label>
            </div>
            <div className={classes.filter}>
              <input type="checkbox"></input>
              <label>WiFi</label>
            </div>
            <div className={classes.filter}>
              <input type="checkbox"></input>
              <label>Air Conditioning</label>
            </div>
            <div className={classes.filter}>
              <input type="checkbox"></input>
              <label>Kitchen facilities</label>
            </div>
            <div className={classes.filter}>
              <input type="checkbox"></input>
              <label>Washing Machine</label>
            </div>
            <div className={classes.filter}>
              <input type="checkbox"></input>
              <label>Balcony</label>
            </div>
            <div className={classes.filter}>
              <input type="checkbox"></input>
              <label>Bathtub</label>
            </div>
            <h5>Host: </h5>
            <div className={classes.filter}>
              <input type="checkbox"></input>
              <label>Outstanding host</label>
            </div>
            <button className={classes.searchButton}>Search</button>
          </div>
          <div className={classes.tableContainer}>
            <h2>Properties in Beograd</h2>
            {properties?.map((property) => (
              <div className={classes.propertyContainer}>
                <div className={classes.imgTitle}>
                  <div className={classes.image}></div>
                  <div>
                    <h2>{property.Name}</h2>
                    <h4>{property.Name}</h4>
                    <h4>{property.Name}</h4>
                    <h4>{property.Name}</h4>
                  </div>
                </div>
                <div className={classes.property}>
                  <div className={classes.propertyCont}>
                    <div className={classes.checkButtonContainer}>
                      <div className={classes.pricesContainer}>
                        <h4>{property.Name}</h4>
                        <h4>{property.Name}</h4>
                        <h4>{property.Name}</h4>
                      </div>
                      <button className={utils.greenButton}>Check</button>
                    </div>
                  </div>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
};

export default Home;
