import React from "react";
import classes from "./NewAccomodation.module.css";
import utils from "./Utils.module.css";

const NewAccomodation = () => {
  return (
    <div className={classes.body}>
      <div className={classes.newAccomodation}>
        <div className={utils.title}>New accommodation </div>
        <div className={classes.newAccContainer}>
          <div className={classes.inputContainer}>
            <div className={classes.basics}>
              <div className={classes.span}>
                <span>Name: </span>
                <input className={classes.inputAcc}></input>
              </div>
              <div className={classes.span}>
                <span>Country: </span>
                <input className={classes.inputAcc}></input>
              </div>
              <div className={classes.span}>
                <span>City: </span>
                <input className={classes.inputAcc}></input>
              </div>
              <div className={classes.span}>
                <span>Address: </span>
                <input className={classes.inputAcc}></input>
              </div>

              <div className={classes.span}>
                <span>Min guest: </span>
                <input className={classes.inputAcc}></input>
              </div>
              <div className={classes.span}>
                <span>Max guest: </span>
                <input className={classes.inputAcc}></input>
              </div>
            </div>
            <div className={classes.checks}>
              <div className={classes.spanCheck}>
                <span>Free parking: </span>
                <input type="checkbox" className={classes.check}></input>
              </div>
              <div className={classes.spanCheck}>
                <span>Wifi: </span>
                <input type="checkbox" className={classes.check}></input>
              </div>
              <div className={classes.spanCheck}>
                <span>Air Conditioning: </span>
                <input type="checkbox" className={classes.check}></input>
              </div>
              <div className={classes.spanCheck}>
                <span>Kitchen: </span>
                <input type="checkbox" className={classes.check}></input>
              </div>
              <div className={classes.spanCheck}>
                <span>Washing Machine: </span>
                <input type="checkbox" className={classes.check}></input>
              </div>
              <div className={classes.spanCheck}>
                <span>Bathtub: </span>
                <input type="checkbox" className={classes.check}></input>
              </div>
              <div className={classes.spanCheck}>
                <span>Balcony: </span>
                <input type="checkbox" className={classes.check}></input>
              </div>
              <div className={classes.spanCheck}>
                <span>Manual acceptance: </span>
                <input type="checkbox" className={classes.check}></input>
              </div>
            </div>
            <div>Images: </div>
            <input type="file"></input>
          </div>
        </div>
        <div className={utils.buttonContainerCenter}>
          <button className={utils.greenButton}>Create</button>
        </div>
      </div>
    </div>
  );
};

export default NewAccomodation;
