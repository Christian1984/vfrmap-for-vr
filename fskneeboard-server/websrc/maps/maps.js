import "@fortawesome/fontawesome-free/css/all.min.css";
import "@fortawesome/fontawesome-free/js/all";

import { onDomContentLoaded } from "./mapsModule";
import Waypoints from "../freemium/waypoints/waypoints";

document.addEventListener("DOMContentLoaded", function() {
    onDomContentLoaded(Waypoints);
});