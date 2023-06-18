import React from "react";
import classes from "./Flights.module.css";
import utils from "./Utils.module.css";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Modal from "@mui/material/Modal";
import { useState, useRef, useEffect } from "react";
import dayjs, { Dayjs } from "dayjs";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import { useNavigate, useParams } from "react-router-dom";
import AuthContext from "../store/auth-context";
import { useContext } from "react";
const style = {
  position: "absolute",
  top: "50%",
  left: "50%",
  transform: "translate(-50%, -50%)",
  bgcolor: "background.paper",
  boxShadow: 24,
  borderRadius: 3,
};
const MyAccommodation = () => {
  const authCtx = useContext(AuthContext);
  const [open, setOpen] = useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);
  const today = new Date();
  const tomorrow = new Date(today);
  tomorrow.setDate(tomorrow.getDate() + 1);
  const [value, setValue] = useState(dayjs(tomorrow));
  const [valueEnd, setValueEnd] = useState(dayjs(tomorrow));
  let { id } = useParams();
  const [acc, setAcc] = useState();
  const [accName, setAccName] = useState();
  const priceRef = useRef();
  const pricingRef = useRef();
  const navigate = useNavigate();
  const [refresh, setRefresh] = useState(false);
  const addAvailabilityHandler = () => {
    event.preventDefault();
    console.log(pricingRef.current.checked);
    console.log(priceRef.current.value);
    console.log(value);
    console.log(valueEnd);
    fetch("http://localhost:8000/accommodation/availability", {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: authCtx.token,
      },
      body: JSON.stringify({
        accommodationId: id,
        availableDate: {
          beginning: value,
          ending: valueEnd,
          price: priceRef.current.value,
          isPricePerGuest: pricingRef.current.checked,
        },
      }),
    })
      .then((res) => {
        if (res.ok) {
          return res.json();
        } else if (res.status == 500) {
          throw new Error(
            "Can not update because there are existing reservations!"
          );
        }
      })
      .then((data) => {
        if (data) console.log(data);
        setRefresh(true);
        //navigate("/my-accommodation/" + id);
      })
      .catch((error) => {
        alert(error);
      });
  };

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
        setAcc(actualData.accommodation.availability);
        setAccName(actualData.accommodation.name);
      });
  }, [refresh]);

  return (
    <div className={classes.body}>
      <div className={classes.home}>
        <br></br>
        <h1>{accName}</h1>

        <div className={classes.buttonContainerRight}>
          <button className={utils.greenButton} onClick={handleOpen}>
            Add
          </button>
        </div>
        <table className={classes.styledTable}>
          <thead>
            <tr>
              <th>Starting date</th>
              <th>Ending date</th>
              <th>Price</th>
              <th>Pricing</th>
            </tr>
          </thead>
          <tbody>
            {acc?.map((app, index) => (
              <tr key={index}>
                <td>{dayjs(app.beginning).format("DD-MM-YYYY")}</td>
                <td>{dayjs(app.ending).format("DD-MM-YYYY")}</td>
                <td>{app.price}</td>
                <td>{app.isPricePerGuest ? "Per guest" : "Per unit"}</td>
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
            <div className={classes.modalTitle}>Add pricing</div>
            <form className={classes.register}>
              <div className={classes.spanReserve}>
                <label>Starting date:</label>
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
              <div className={classes.spanReserve}>
                <label>Ending date:</label>
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
              <div className={classes.spanReserve}>
                <label>Price:</label>
                <input className={utils.inputPrice} ref={priceRef}></input>
              </div>
              <div className={classes.spanReserve}>
                <label>Pricing:</label>

                <input
                  type="radio"
                  id="unit"
                  name="pricing"
                  value="unit"
                  ref={pricingRef}
                ></input>
                <label>Per unit</label>
                <input
                  ref={pricingRef}
                  type="radio"
                  id="person"
                  name="pricing"
                  value="person"
                ></input>
                <label>Per person</label>
              </div>
              <button
                className={classes.reserveButton}
                onClick={addAvailabilityHandler}
              >
                Add
              </button>
            </form>
          </div>
        </Box>
      </Modal>
    </div>
  );
};

export default MyAccommodation;
