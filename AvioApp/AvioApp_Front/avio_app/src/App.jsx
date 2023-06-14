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
import AdminFlights from "./pages/AdminFlights";
import NewFlight from "./pages/NewFlight";
import Flights from "./pages/Flights";
import Reservations from "./pages/Reservations";
import AuthContext from "./store/auth-context";

const router = createBrowserRouter(
  createRoutesFromElements(
    <Route>
      <Route path="/register" element={<Register></Register>} />
      <Route path="/login" element={<Login></Login>} />
      <Route path="/" element={<RootLayout></RootLayout>}>
        <Route path="/" index element={<Home></Home>}></Route>
        <Route
          path="/admin-flights"
          index
          element={<AdminFlights></AdminFlights>}
        ></Route>
        <Route path="/flights" index element={<Flights></Flights>}></Route>
        <Route
          path="/reservations"
          index
          element={<Reservations></Reservations>}
        ></Route>
        <Route
          path="/new-flight"
          index
          element={<NewFlight></NewFlight>}
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
