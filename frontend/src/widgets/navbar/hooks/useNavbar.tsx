import { useContext } from "react";
import { NavbarContext } from "../context/navbar.context";

export const useNavbar = () => useContext(NavbarContext);
