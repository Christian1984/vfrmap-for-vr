import "../../common/common.scss";
import "./notepad.scss";

let loaded = false;

window.addEventListener("message", (m) => {
    iframe = document.querySelector("iframe");
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