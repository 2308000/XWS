import React from "react";
import classes from "./Accommodation.module.css";
import utils from "./Utils.module.css";
import { useNavigate } from "react-router-dom";

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

const MyAccommodations = () => {
  const navigate = useNavigate();

  const selectPropertyHandler = (id) => {
    navigate("/my-accommodations/" + id);
  };

  const addAccommodationHandler = () => {
    navigate("/new-accommodation/");
  };

  return (
    <div className={classes.body}>
      <br></br>
      <div className={classes.home}>
        <div className={classes.tableContainer}>
          <h1>My accomodations</h1>
          <div className={utils.buttonContainerRight}>
            <button
              className={utils.greenButton}
              onClick={addAccommodationHandler}
            >
              Add
            </button>
          </div>
          {properties?.map((property) => (
            <div className={classes.propertyContainer}>
              <div className={classes.imgTitle}>
                <div className={classes.image}></div>
                <div>
                  <h2>{property.Name}</h2>
                </div>
              </div>
              <div className={classes.property}>
                <div className={classes.propertyCont}>
                  <div className={classes.checkButtonContainer}>
                    <button
                      className={utils.greenButton}
                      onClick={() => selectPropertyHandler(property.Id)}
                    >
                      Manage
                    </button>
                  </div>
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default MyAccommodations;
