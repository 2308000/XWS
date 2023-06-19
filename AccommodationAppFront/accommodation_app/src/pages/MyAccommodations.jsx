import React from "react";
import classes from "./Accommodation.module.css";
import utils from "./Utils.module.css";
import { useNavigate } from "react-router-dom";
import AuthContext from "../store/auth-context";
import { useContext, useEffect, useState } from "react";

const MyAccommodations = () => {
  const authCtx = useContext(AuthContext);
  const [properties, setProperties] = useState();
  const navigate = useNavigate();
  const [selectedAccommodation, setSelecetedAccommodation] = useState();

  const selectPropertyHandler = (accommodation) => {
    setSelecetedAccommodation(accommodation);
    navigate("/my-accommodations/" + accommodation.id);
  };

  const addAccommodationHandler = () => {
    navigate("/new-accommodation/");
  };

  useEffect(() => {
    fetch("http://localhost:8000/accommodation/host/" + authCtx.id, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: authCtx.token,
      },
    })
      .then((response) => response.json())
      .then((actualData) => {
        console.log(actualData);
        setProperties(actualData.accommodations);
      });
  }, []);

  return (
    <div className={classes.body}>
      <br></br>
      <div className={classes.home}>
        <div className={classes.myAccContainer}>
          <h1>My accomodations</h1>
          <div className={utils.buttonContainerRight}>
            <button
              className={utils.lightBlueButton}
              onClick={addAccommodationHandler}
            >
              Add
            </button>
          </div>
          {properties?.map((property) => (
            <div className={classes.accommodationContainer}>
              <div className={classes.imgTitle}>
                <div className={classes.image}></div>
                <div>
                  <h2>{property?.name}</h2>
                  <h3>
                    {property?.location.street}, {property?.location.city},{" "}
                    {property?.location.country}
                  </h3>
                </div>
              </div>
              <div className={classes.property}>
                <div className={classes.propertyCont}>
                  <div className={classes.checkButtonContainer}>
                    <button
                      className={utils.blueButton}
                      onClick={() => {
                        selectPropertyHandler(property);
                      }}
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
