import React from "react";
import { useEffect, useState, useRef, useContext } from "react";
import AuthContext from "../store/auth-context";
import classes from "./Profile.module.css";
import utils from "./Utils.module.css";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import dayjs, { Dayjs } from "dayjs";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Modal from "@mui/material/Modal";

const user = { name: "petar" };

const MyProfile = () => {
  const style = {
    position: "absolute",
    top: "50%",
    left: "50%",
    transform: "translate(-50%, -50%)",
    bgcolor: "background.paper",
    boxShadow: 24,
    borderRadius: 3,
  };
  const styleQuestion = {
    position: "absolute",
    top: "50%",
    left: "50%",
    transform: "translate(-50%, -50%)",
    bgcolor: "background.paper",
    boxShadow: 24,
    borderRadius: 3,
  };
  const authCtx = useContext(AuthContext);
  const [open, setOpen] = useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);
  const [openDelete, setOpenDelete] = useState(false);
  const handleOpenDelete = () => setOpenDelete(true);
  const handleCloseDelete = () => setOpenDelete(false);
  const [toggleEdit, setToggleEdit] = useState();
  const [saved, setSaved] = useState();
  const nameInputRef = useRef();
  const surnameInputRef = useRef();
  const emailInputRef = useRef();
  const usernameInputRef = useRef();
  const birthDateInputRef = useRef();
  const countryInputRef = useRef();
  const cityInputRef = useRef();
  const addressInputRef = useRef();
  const phoneNumberInputRef = useRef();
  const genderInputRef = useRef();
  const [validation, setValidation] = useState(true);

  const [value, setValue] = useState(dayjs("2023-05-14"));
  const genders = ["Female", "Male"];
  const [newPw, setNewPw] = useState();
  const [newRePw, setNewRePw] = useState();
  const handleToggleEdit = () => {
    setToggleEdit(true);
  };
  const handleCloseEdit = () => {
    setSaved(false);
    setToggleEdit(false);
  };
  const handleSaveChanges = () => {
    event?.preventDefault();
    const enteredName = nameInputRef.current.value;
    const enteredSurname = surnameInputRef.current.value;
    const enteredGender = genderRef.current.value;

    fetch("http://localhost:5041/api/profiles/update/own", {
      method: "PUT",
      body: JSON.stringify({
        name: enteredName,
        gender: enteredGender,
        surname: enteredSurname,
        birthDate: "2023-04-12T21:01:31.611Z",
      }),
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + authCtx.token,
      },
    })
      .then((res) => {
        if (res.ok) {
          console.log(res);
          return res;
        }
      })
      .then((data) => {
        console.log(data);
        setSaved(true);
      });
  };

  const changePwHandler = () => {
    setNewPw(event.target.value);
  };
  const changeRePwHandler = () => {
    setNewRePw(event.target.value);
  };

  useEffect(() => {
    if (newRePw !== newPw) {
      setValidation(false);
    } else {
      setValidation(true);
    }
  }, [newRePw, newPw]);

  return (
    <>
      {!toggleEdit && (
        <div className={classes.myProfile}>
          <div className={utils.title}>My Profile</div>
          <div className={classes.inputContainerVertical}>
            <div className={classes.inputContainer}>
              <div>
                <div className={classes.span}>
                  <span>Name: </span>
                  <span>{user.name}</span>
                </div>
                <div className={classes.span}>
                  <span>Surname: </span>
                  <span>{user.surname}</span>
                </div>
                <div className={classes.span}>
                  <span>Email: </span>
                  <span>{user.email}</span>
                </div>
                <div className={classes.span}>
                  <span>Gender: </span>
                  <span>{user.gender}</span>
                </div>
                <div className={classes.span}>
                  <span>Birthdate: </span>
                  <span>{user.birthDate}</span>
                </div>
              </div>
              <div>
                <div className={classes.span}>
                  <span>Username: </span>
                  <span>{user.birthDate}</span>
                </div>

                <div className={classes.span}>
                  <span>Country: </span>
                  <span>{user.birthDate}</span>
                </div>

                <div className={classes.span}>
                  <span>City: </span>
                  <span>{user.birthDate}</span>
                </div>
                <div className={classes.span}>
                  <span>Address: </span>
                  <span>{user.birthDate}</span>
                </div>
                <div className={classes.span}>
                  <span>Phone number: </span>
                  <span>{user.birthDate}</span>
                </div>
              </div>
            </div>
            <div className={utils.buttonContainer}>
              <button className={utils.whiteButton} onClick={handleToggleEdit}>
                Edit Profile
              </button>
              <button className={utils.greenButton} onClick={handleOpen}>
                Change password
              </button>
              <button className={utils.redButton} onClick={handleOpenDelete}>
                Delete profile
              </button>
            </div>
          </div>
          <Modal
            open={open}
            onClose={handleClose}
            aria-labelledby="modal-modal-title"
            aria-describedby="modal-modal-description"
          >
            <Box sx={style}>
              <div>
                <div className={classes.modalTitle}>Change password</div>
                <form className={classes.changePW}>
                  <div className={classes.changePWCont}>
                    <label>New password:</label>
                    <input
                      className={classes.inputChangePW}
                      value={newPw}
                      onChange={changePwHandler}
                    ></input>
                  </div>
                  <div className={classes.changePWCont}>
                    <label>Re enter new password:</label>
                    <input
                      className={classes.inputChangePW}
                      value={newRePw}
                      onChange={changeRePwHandler}
                    ></input>
                  </div>

                  <div className={classes.buttonContainerCenter}>
                    {validation ? (
                      <button className={utils.greenButton}>Save</button>
                    ) : (
                      <button className={utils.greenButton} disabled>
                        Save
                      </button>
                    )}
                  </div>
                </form>
              </div>
            </Box>
          </Modal>
          <Modal
            open={openDelete}
            onClose={handleCloseDelete}
            aria-labelledby="modal-modal-title"
            aria-describedby="modal-modal-description"
          >
            <Box sx={styleQuestion}>
              <div className={utils.modalQuestion}>
                <div className={classes.modalTitle}>Delete account</div>
                <div className={utils.modalQuestionContainer}>
                  <div className={classes.spanReserve}>
                    <p>Are you sure you want to delete your account?</p>
                  </div>
                  <div className={utils.buttonContainerRight}>
                    <button className={utils.redButton}>Delete</button>
                  </div>
                </div>
              </div>
            </Box>
          </Modal>
        </div>
      )}
      {toggleEdit && (
        <div className={classes.myProfile}>
          <div className={utils.title}>Edit Profile</div>
          <div className={classes.inputContainerVerticalEdit}>
            <div>
              <div className={classes.spanEdit}>
                <span>Name: </span>
                <input
                  ref={nameInputRef}
                  defaultValue={user.name}
                  className={classes.input}
                ></input>
              </div>
              <div className={classes.spanEdit}>
                <span>Surname: </span>
                <input
                  ref={surnameInputRef}
                  defaultValue={user.surname}
                  className={classes.input}
                ></input>
              </div>
              <div className={classes.spanEdit}>
                <span>Email: </span>
                <input
                  ref={emailInputRef}
                  defaultValue={user.surname}
                  className={classes.input}
                ></input>
              </div>

              <div className={classes.spanEdit}>
                <span>Username: </span>
                <input
                  ref={usernameInputRef}
                  defaultValue={user.surname}
                  className={classes.input}
                ></input>
              </div>
              <div className={classes.spanEdit}>
                <span>Birthdate: </span>
                <div className={classes.container}>
                  <LocalizationProvider dateAdapter={AdapterDayjs}>
                    <DatePicker
                      renderInput={(props) => <TextField {...props} />}
                      value={value}
                      className={classes.timePicker}
                      onChange={(newValue) => {
                        setValue(newValue);
                      }}
                    />
                  </LocalizationProvider>
                </div>
              </div>
            </div>
            <div>
              <div className={classes.spanEdit}>
                <span>Country: </span>
                <input
                  ref={countryInputRef}
                  defaultValue={user.surname}
                  className={classes.input}
                ></input>
              </div>
              <div className={classes.spanEdit}>
                <span>City: </span>
                <input
                  ref={cityInputRef}
                  defaultValue={user.surname}
                  className={classes.input}
                ></input>
              </div>
              <div className={classes.spanEdit}>
                <span>Address: </span>
                <input
                  ref={addressInputRef}
                  defaultValue={user.surname}
                  className={classes.input}
                ></input>
              </div>
              <div className={classes.spanEdit}>
                <span>Phone number: </span>
                <input
                  ref={phoneNumberInputRef}
                  defaultValue={user.surname}
                  className={classes.input}
                ></input>
              </div>
              <div className={classes.spanEdit}>
                <span>Gender: </span>
                <select
                  defaultValue={user.gender}
                  ref={genderInputRef}
                  className={classes.select}
                >
                  {genders.map((gender, index) => {
                    return <option key={index}>{gender}</option>;
                  })}
                </select>
              </div>
            </div>
          </div>
          <div className={utils.buttonContainerEdit}>
            <button onClick={handleCloseEdit} className={utils.redButton}>
              Close
            </button>
            <button className={utils.greenButton} onClick={handleSaveChanges}>
              Save
            </button>
          </div>
        </div>
      )}
    </>
  );
};

export default MyProfile;
