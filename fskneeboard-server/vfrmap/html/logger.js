// depends on sender.js
const LogLevel = {
    DEBUG: "debug",
    INFO: "info",
    WARN: "warn",
    ERROR: "error",
    OFF: "off"
}

class Logger {
    static level = undefined;

    static init(callback) {
        //if (Logger.level == undefined) Logger.level = LogLevel.DEBUG;

        let xhr = new XMLHttpRequest();
        xhr.open("GET", "/loglevel/", true);
        xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
        xhr.onreadystatechange = function () {
            if (xhr.readyState === 4) {
                if (xhr.status === 200) {
                    console.log(xhr.responseText);

                    const levelToLower = xhr.responseText.toLowerCase();
                    if (levelToLower == LogLevel.DEBUG || levelToLower == LogLevel.INFO || levelToLower == LogLevel.WARN || levelToLower == LogLevel.ERROR || levelToLower == LogLevel.OFF) {
                        Logger.level = xhr.responseText.toLowerCase();
                        Logger.logDebug("Client LogLevel received and set to [" + levelToLower + "]");

                        if (callback) {
                            callback();
                        }

                        Logger.logMessage("OFF-Test", LogLevel.OFF);
                        Logger.logDebug("DEBUG-Test");
                        Logger.logInfo("INFO-Test");
                        Logger.logWarn("WARN-Test");
                        Logger.logError("ERROR-Test");
                    }
                    else {
                        Logger.logWarn("Received invalid client LogLevel: [" + levelToLower + "]; Logger turned off!");
                        Logger.level = LogLevel.OFF;
                    }
                }
            }
        };

        xhr.send();
    }

    static shouldLog(level) {
        if (Logger.level == undefined) {
            return true;
        }

        if (Logger.level == LogLevel.DEBUG && (level.toLowerCase() == LogLevel.DEBUG || level.toLowerCase() == LogLevel.INFO || level.toLowerCase() == LogLevel.WARN || level.toLowerCase() == LogLevel.ERROR)) {
            return true;
        }

        if (Logger.level == LogLevel.INFO && (level.toLowerCase() == LogLevel.INFO || level.toLowerCase() == LogLevel.WARN || level.toLowerCase() == LogLevel.ERROR)) {
            return true;
        }

        if (Logger.level == LogLevel.WARN && (level.toLowerCase() == LogLevel.WARN || level.toLowerCase() == LogLevel.ERROR)) {
            return true;
        }

        if (Logger.level == LogLevel.ERROR && level.toLowerCase() == LogLevel.ERROR) {
            return true;
        }

        return false;
    }

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

    static logMessage(message, level, retry = true) {
        if (Logger.level == undefined && retry === true) {
            Logger.init(() => Logger.logMessage(message, level, false));
            return;
        }

        if (Logger.shouldLog(level)) {
            Logger.logLocal(message, level);
            Logger.logRemote(message, level);
        }
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