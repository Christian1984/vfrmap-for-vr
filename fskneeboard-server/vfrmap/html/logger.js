// depends on sender.js
const LogLevel = {
    DEBUG: "debug",
    INFO: "info",
    WARN: "warn",
    ERROR: "error",
}

class Logger {
    static logRemote(message, level) {
        let xhr = new XMLHttpRequest();
        
        let body = {
            level: level,
            message: message,
            sender: SENDER_ID
        }
    
        xhr.open("POST", "/log/", true);
        xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
        xhr.send(JSON.stringify(body));
    }

    static logLocal(message, level) {
        const logDiv = document.querySelector("div#log");
    
        if (logDiv) {
            logDiv.innerHTML = "<p>[" + level + "] " + message + "</p>" + logDiv.innerHTML
        }
    }

    static logMessage(message, level) {
        Logger.logLocal(message, level);
        Logger.logRemote(message, level);
    }

    static logDebug(message) {
        Logger.logMessage(message, LogLevel.DEBUG);
    }

    static logInfo(message) {
        Logger.logMessage(message, LogLevel.INFO);
    }

    static logWarn(message) {
        Logger.logMessage(message, LogLevel.WARN);
    }

    static logError(message) {
        Logger.logMessage(message, LogLevel.ERROR);
    }
}

document.addEventListener("DOMContentLoaded", function() {
    Logger.logDebug("logger.js => DOMContentLoaded fired!");
});