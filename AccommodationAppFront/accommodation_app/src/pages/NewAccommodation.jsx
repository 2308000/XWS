import React from "react";
import classes from "./NewAccomodation.module.css";
import utils from "./Utils.module.css";
import AuthContext from "../store/auth-context";
import { useContext, useEffect, useState, useRef } from "react";
import { useNavigate } from "react-router-dom";

const NewAccomodation = () => {
  const authCtx = useContext(AuthContext);
  const navigate = useNavigate();
  const nameRef = useRef();

  const cityRef = useRef();
  const countryRef = useRef();
  const addressRef = useRef();
  const minGuestRef = useRef();
  const maxGuestRef = useRef();
  const wifiRef = useRef();
  const parkingRef = useRef();
  const ACRef = useRef();
  const KFRef = useRef();
  const WMRef = useRef();
  const balconyRef = useRef();
  const bathtubRef = useRef();
  const MARef = useRef();

  const [image, setImage] = useState();

  const addHandler = () => {
    event.preventDefault();
    const fd = new FormData();
    fd.append("file", image);
    fetch("http://localhost:8000/accommodation", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: authCtx.token,
      },
      body: JSON.stringify({
        hostId: authCtx.id,
        hasWifi: wifiRef.current.checked,
        hasAirConditioning: ACRef.current.checked,
        hasFreeParking: parkingRef.current.checked,
        hasKitchen: KFRef.current.checked,
        hasWashingMachine: WMRef.current.checked,
        hasBathtub: bathtubRef.current.checked,
        hasBalcony: balconyRef.current.checked,
        photos: image,
        name: nameRef.current.value,
        minGuest: minGuestRef.current.value,
        maxGuest: maxGuestRef.current.value,
        location: {
          city: cityRef.current.value,
          country: countryRef.current.value,
          street: addressRef.current.value,
        },
      }),
    })
      .then((res) => {
        if (res.ok) {
          return res.json();
        }
      })
      .then((data) => {
        if (data) console.log(data);
        navigate("/my-accommodations");
      })
      .catch((error) => {
        alert(error);
      });
  };

  const fileSelectedHandler = (event) => {
    let reader = new FileReader();
    reader.readAsDataURL(event.target.files[0]);
    console.log(event.target.files[0]);
    reader.onload = () => {
      setImage(reader.result);
    };
  };

  return (
    <div className={classes.body}>
      <div className={classes.newAccomodation}>
        <div className={utils.title}>New accommodation </div>
        <div className={classes.newAccContainer}>
          <div className={classes.inputContainer}>
            <div className={classes.basics}>
              <div className={classes.span}>
                <span>Name: </span>
                <input className={classes.inputAcc} ref={nameRef}></input>
              </div>
              <div className={classes.span}>
                <span>Country: </span>
                <input className={classes.inputAcc} ref={countryRef}></input>
              </div>
              <div className={classes.span}>
                <span>City: </span>
                <input className={classes.inputAcc} ref={cityRef}></input>
              </div>
              <div className={classes.span}>
                <span>Address: </span>
                <input className={classes.inputAcc} ref={addressRef}></input>
              </div>

              <div className={classes.span}>
                <span>Min guest: </span>
                <input className={classes.inputAcc} ref={minGuestRef}></input>
              </div>
              <div className={classes.span}>
                <span>Max guest: </span>
                <input className={classes.inputAcc} ref={maxGuestRef}></input>
              </div>
            </div>
            <div className={classes.checks}>
              <div className={classes.spanCheck}>
                <span>Free parking: </span>
                <input
                  type="checkbox"
                  className={classes.check}
                  ref={parkingRef}
                ></input>
              </div>
              <div className={classes.spanCheck}>
                <span>Wifi: </span>
                <input
                  type="checkbox"
                  className={classes.check}
                  ref={wifiRef}
                ></input>
              </div>
              <div className={classes.spanCheck}>
                <span>Air Conditioning: </span>
                <input
                  type="checkbox"
                  className={classes.check}
                  ref={ACRef}
                ></input>
              </div>
              <div className={classes.spanCheck}>
                <span>Kitchen: </span>
                <input
                  type="checkbox"
                  className={classes.check}
                  ref={KFRef}
                ></input>
              </div>
              <div className={classes.spanCheck}>
                <span>Washing Machine: </span>
                <input
                  type="checkbox"
                  className={classes.check}
                  ref={WMRef}
                ></input>
              </div>
              <div className={classes.spanCheck}>
                <span>Bathtub: </span>
                <input
                  type="checkbox"
                  className={classes.check}
                  ref={bathtubRef}
                ></input>
              </div>
              <div className={classes.spanCheck}>
                <span>Balcony: </span>
                <input
                  type="checkbox"
                  className={classes.check}
                  ref={balconyRef}
                ></input>
              </div>
              <div className={classes.spanCheck}>
                <span>Manual acceptance: </span>
                <input
                  type="checkbox"
                  className={classes.check}
                  ref={MARef}
                ></input>
              </div>
            </div>
            <div>Images: </div>
            <input type="file" onChange={fileSelectedHandler}></input>
          </div>
        </div>
        <div className={utils.buttonContainerCenter}>
          <button className={utils.greenButton} onClick={addHandler}>
            Create
          </button>
        </div>
      </div>
    </div>
  );
};

export default NewAccomodation;
