import React from "react";
import classes from "./Accommodation.module.css";
import utils from "./Utils.module.css";
import dayjs, { Dayjs } from "dayjs";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import { useState, useRef, useEffect, useContext } from "react";
import Property from "../components/Property";
import { useNavigate } from "react-router-dom";

const Accommodations = () => {
  const [properties, setProperties] = useState();
  const today = new Date();
  const tomorrow = new Date(today);
  const start = new Date(today);
  const end = new Date(today);
  start.setDate(tomorrow.getDate() - 2);
  end.setDate(tomorrow.getDate() + 3);
  tomorrow.setDate(tomorrow.getDate() + 1);
  const [valueStart, setValueStart] = useState(dayjs(start));
  const [valueEnd, setValueEnd] = useState(dayjs(end));
  const cityRef = useRef();
  const minPriceRef = useRef();
  const maxPriceRef = useRef();
  const countryRef = useRef();
  const wifiRef = useRef();
  const parkingRef = useRef();
  const ACRef = useRef();
  const KFRef = useRef();
  const WMRef = useRef();
  const balconyRef = useRef();
  const bathtubRef = useRef();
  const OHRef = useRef();
  const numberOfGuestsRef = useRef();
  const navigate = useNavigate();
  const [nog, setNog] = useState(2);

  useEffect(() => {
    fetch("http://localhost:8000/accommodation/search", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        numberOfGuests: 2,
        location: {
          city: "",
          country: "",
          street: "",
        },
        beginning: start,
        ending: end,
      }),
    })
      .then((response) => response.json())
      .then((actualData) => {
        console.log(actualData);
        setProperties(actualData.accommodations);
      });
  }, []);

  const searchHandler = () => {
    if (!numberOfGuestsRef.current.value) {
      alert("Please enter number of guests!");
      return;
    }
    if (numberOfGuestsRef.current.value < 1) {
      alert("Please enter valid number of guests!");
      return;
    }
    fetch("http://localhost:8000/accommodation/search", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        numberOfGuests: numberOfGuestsRef.current.value,
        location: {
          city: cityRef.current.value,
          country: countryRef.current.value,
          street: "",
        },
        beginning: valueStart,
        ending: valueEnd,
      }),
    })
      .then((response) => response.json())
      .then((actualData) => {
        console.log(actualData);
        setProperties(actualData.accommodations);
      });
  };

  const checkHandler = (id) => {
    window.open(`/accommodations/${id}`, "_blank");
  };

  const filterHandler = () => {
    if (!minPriceRef.current.value || !maxPriceRef.current.value) {
      alert("Please enter price range!");
      return;
    }

    fetch("http://localhost:8000/accommodation/filter", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        searchQuery: {
          numberOfGuests: numberOfGuestsRef.current.value,
          location: {
            city: cityRef.current.value,
            country: countryRef.current.value,
            street: "",
          },
          beginning: valueStart,
          ending: valueEnd,
        },
        priceRangeLowerBound: minPriceRef.current.value,
        priceRangeUpperBound: maxPriceRef.current.value,
        benefits: {
          hasWifi: wifiRef.current.checked,
          hasAirConditioning: ACRef.current.checked,
          hasFreeParking: parkingRef.current.checked,
          hasKitchen: KFRef.current.checked,
          hasWashingMachine: WMRef.current.checked,
          hasBathtub: bathtubRef.current.checked,
          hasBalcony: balconyRef.current.checked,
        },
        isOutstandingHost: OHRef.current.checked,
      }),
    })
      .then((response) => response.json())
      .then((actualData) => {
        console.log(actualData);
        setProperties(actualData.accommodations);
      });
  };

  return (
    <div className={classes.body}>
      <br></br>
      <div className={classes.home}>
        <div className={classes.search}>
          <div className={classes.searchItem}>
            <label>City</label>
            <input className={classes.inputLoc} ref={cityRef}></input>
          </div>
          <div className={classes.searchItem}>
            <label>Country</label>
            <input className={classes.inputLoc} ref={countryRef}></input>
          </div>
          <div className={classes.searchItem}>
            <label>From</label>
            <LocalizationProvider dateAdapter={AdapterDayjs}>
              <DatePicker
                value={valueStart}
                onChange={(newValue) => {
                  setValueStart(newValue);
                }}
                className={classes.DatePicker}
              />
            </LocalizationProvider>
          </div>
          <div className={classes.searchItem}>
            <label>To</label>
            <LocalizationProvider dateAdapter={AdapterDayjs}>
              <DatePicker
                value={valueEnd}
                onChange={(newValue) => {
                  setValueEnd(newValue);
                }}
                className={classes.DatePicker}
              />
            </LocalizationProvider>
          </div>
          <div className={classes.searchItem}>
            <label>Number of guests</label>
            <input
              className={classes.inputNumber}
              ref={numberOfGuestsRef}
              defaultValue={nog}
            ></input>
          </div>
          <button className={classes.searchButton} onClick={searchHandler}>
            Search
          </button>
        </div>
        <div className={classes.accomodationsContainer}>
          <div className={classes.filters}>
            <h4>Filter by:</h4>
            <h5>Price range: </h5>
            <div className={classes.filterPrice}>
              <label>From</label>
              <input className={classes.inputPrice} ref={minPriceRef}></input>
            </div>
            <div className={classes.filterPrice}>
              <label>To</label>
              <input className={classes.inputPrice} ref={maxPriceRef}></input>
            </div>
            <h5>Benefits: </h5>
            <div className={classes.filter}>
              <input type="checkbox" ref={parkingRef}></input>
              <label>Free Parking</label>
            </div>
            <div className={classes.filter}>
              <input type="checkbox" ref={wifiRef}></input>
              <label>WiFi</label>
            </div>
            <div className={classes.filter}>
              <input type="checkbox" ref={ACRef}></input>
              <label>Air Conditioning</label>
            </div>
            <div className={classes.filter}>
              <input type="checkbox" ref={KFRef}></input>
              <label>Kitchen facilities</label>
            </div>
            <div className={classes.filter}>
              <input type="checkbox" ref={WMRef}></input>
              <label>Washing Machine</label>
            </div>
            <div className={classes.filter}>
              <input type="checkbox" ref={balconyRef}></input>
              <label>Balcony</label>
            </div>
            <div className={classes.filter}>
              <input type="checkbox" ref={bathtubRef}></input>
              <label>Bathtub</label>
            </div>
            <h5>Host: </h5>
            <div className={classes.filter}>
              <input type="checkbox" ref={OHRef}></input>
              <label>Outstanding host</label>
            </div>
            <div className={utils.buttonContainerCenter}>
              <button className={classes.searchButton} onClick={filterHandler}>
                Filter
              </button>
            </div>
          </div>
          <div className={classes.tableContainer}>
            {cityRef?.current?.value?.length > 0 ||
            countryRef?.current?.value?.length > 0 ? (
              <h2>
                Properties in {cityRef?.current.value} ,{" "}
                {countryRef?.current.value}
              </h2>
            ) : (
              <h2>Properties</h2>
            )}
            {properties?.map((property) => (
              <div
                className={
                  property?.accommodation.host.isOutstanding
                    ? classes.propertyContainerOutstanding
                    : classes.propertyContainer
                }
                key={property.accommodation.id}
              >
                <div className={classes.imgTitle}>
                  <div className={classes.image}></div>
                  <div>
                    <h2 className={classes.blueText}>
                      {property.accommodation.name}
                    </h2>
                    <h5 className={classes.underlinedBlueText}>
                      {property.accommodation.location.street},{" "}
                      {property.accommodation.location.city},{" "}
                      {property.accommodation.location.country}
                    </h5>
                    <br></br>
                    <h5>
                      Lorem ipsum dolor sit amet consectetur adipisicing elit.
                      Eaque suscipit autem odit amet!
                    </h5>
                    <br></br>
                    <h5>{property?.accommodation.hasWifi && "Wifi"}</h5>
                    <h5>
                      {property?.accommodation.hasFreeParking && "Free Parking"}
                    </h5>{" "}
                    <h5>
                      {property?.accommodation.hasAirConditioning &&
                        "Air Conditioning"}
                    </h5>{" "}
                    <h5>
                      {property?.accommodation.hasKitchen &&
                        "Kitchen facilities"}
                    </h5>
                    <h5>{property.Name}</h5>
                  </div>
                </div>
                <div className={classes.property}>
                  <div className={classes.propertyCont}>
                    <div className={classes.checkButtonContainer}>
                      <div className={classes.pricesContainer}>
                        <div className={classes.grade}>
                          <p>
                            {property?.accommodation
                              ?.averageAccommodationGrade !== "NaN"
                              ? (property?.accommodation?.averageAccommodationGrade).toFixed(
                                  2
                                )
                              : 0}
                          </p>
                        </div>
                      </div>
                      <div>
                        <div className={classes.propertyPriceContainer}>
                          {numberOfGuestsRef.current.value && (
                            <h4> {numberOfGuestsRef.current.value} guests </h4>
                          )}
                          <h4>Price per night: {property.pricePerNight}</h4>
                          <br></br>
                          <h4>Total price:</h4>
                          <h2>EUR {property.totalPrice}</h2>
                        </div>
                        <br></br>
                        <button
                          className={utils.blueButton}
                          onClick={() => {
                            checkHandler(property.accommodation.id);
                            console.log(valueEnd);
                            localStorage.setItem("startDate", valueStart);
                            localStorage.setItem("endDate", valueEnd);
                            localStorage.setItem(
                              "numberOfGuests",
                              numberOfGuestsRef.current.value
                            );
                            localStorage.setItem(
                              "pricePerNight",
                              property?.pricePerNight
                            );
                            localStorage.setItem(
                              "totalPrice",
                              property?.totalPrice
                            );
                          }}
                        >
                          Check
                        </button>
                      </div>
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

export default Accommodations;
