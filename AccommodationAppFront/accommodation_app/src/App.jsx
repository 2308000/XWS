import { useState } from "react";
import "./App.css";
import { useRef, useContext } from "react";
import {
  createBrowserRouter,
  RouterProvider,
  Route,
  Link,
  createRoutesFromElements,
} from "react-router-dom";
import RootLayout from "./pages/RootLayout";
import Register from "./pages/Register";
import Login from "./pages/Login";
import Home from "./pages/Home";
import MyProfile from "./pages/MyProfile";
import AuthContext from "./store/auth-context";
import NewAccomodation from "./pages/NewAccommodation";
import Accomodations from "./pages/Accommodations";
import MyAccommodations from "./pages/MyAccommodations";
import MyAccommodation from "./pages/MyAccommodation";
import MyReservations from "./pages/MyReservations";
import ReservationRequests from "./pages/ReservationRequests";
import Accommodation from "./pages/Accommodation";

const router = createBrowserRouter(
  createRoutesFromElements(
    <Route>
      <Route path="/register" element={<Register></Register>} />
      <Route path="/login" element={<Login></Login>} />
      <Route path="/" element={<RootLayout></RootLayout>}>
        <Route path="/" index element={<Home></Home>}></Route>
        <Route path="/profile" element={<MyProfile></MyProfile>}></Route>
        <Route
          path="/accommodations"
          element={<Accomodations></Accomodations>}
        ></Route>
        <Route
          path="/accommodations/:id"
          element={<Accommodation></Accommodation>}
        ></Route>
        <Route
          path="/new-accommodation"
          element={<NewAccomodation></NewAccomodation>}
        ></Route>
        <Route
          path="/my-reservations"
          element={<MyReservations></MyReservations>}
        ></Route>
        <Route
          path="/reservation-requests"
          element={<ReservationRequests></ReservationRequests>}
        ></Route>
        <Route
          path="/my-accommodations"
          element={<MyAccommodations></MyAccommodations>}
        ></Route>
        <Route
          path="/my-accommodations/:id"
          element={<MyAccommodation></MyAccommodation>}
        ></Route>
      </Route>
    </Route>
  )
);

function App() {
  return (
    <>
      <RouterProvider router={router}></RouterProvider>
    </>
  );
}

export default App;
