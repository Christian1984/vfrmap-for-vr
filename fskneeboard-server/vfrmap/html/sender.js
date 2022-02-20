// depends on md5.min.js
const SENDER_ID = md5(Math.floor(Math.random() * Number.MAX_VALUE).toString());

document.addEventListener("DOMContentLoaded", function() {
    Logger.logDebug("sender.js => DOMContentLoaded fired!");
});