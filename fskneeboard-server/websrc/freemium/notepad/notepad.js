import "@fortawesome/fontawesome-free/css/all.min.css";
import "@fortawesome/fontawesome-free/js/all";

import "../../common/common.scss";
import "./notepad.scss";

import Logger from "./../../common/logger";

let loaded = false;

window.addEventListener("message", (m) => {
    const iframe = document.querySelector("iframe");
    if (iframe) {
        if (m.data == "load" && !loaded) {
            iframe.src = "https://fskneeboard.com/notes-ingame/";
            loaded = true;
        }
    }
});

document.addEventListener("DOMContentLoaded", function() {
    Logger.logDebug("notepad.js (FREE) => DOMContentLoaded fired!");
});