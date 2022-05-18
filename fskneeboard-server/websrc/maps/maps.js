import "@fortawesome/fontawesome-free/css/all.min.css";

import { onDomContentLoaded } from "./mapsModule";
import Waypoints from "../freemium/waypoints/waypoints";

document.addEventListener("DOMContentLoaded", function() {
    onDomContentLoaded(Waypoints);
});