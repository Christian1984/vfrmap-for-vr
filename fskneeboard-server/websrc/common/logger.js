import SENDER_ID from "./sender";

const LogLevel = {
    DEBUG: "debug",
    INFO: "info",
    WARN: "warn",
    ERROR: "error",
    OFF: "off"
}

let LoggerLevel = undefined;
let LoggerInitCalled = false;
let LoggerQueue = [];

export default class Logger {
    static init() {
        LoggerInitCalled = true;

        let xhr = new XMLHttpRequest();
        xhr.open("GET", "/loglevel/", true);
        xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
        xhr.onreadystatechange = function () {
            if (xhr.readyState === 4) {
                if (xhr.status === 200) {
                    const levelToLower = xhr.responseText.toLowerCase();

                    if (levelToLower == LogLevel.DEBUG || levelToLower == LogLevel.INFO || levelToLower == LogLevel.WARN || levelToLower == LogLevel.ERROR || levelToLower == LogLevel.OFF) {
                        LoggerLevel = levelToLower;
                        Logger.logInfo("Client LogLevel received and set to [" + levelToLower + "]");

                        // mitigate concurrency at least for current frame
                        setTimeout(() => Logger.processQueue(), 0);

                        /*
                        Logger.logMessage("OFF-Test", LogLevel.OFF);
                        Logger.logDebug("DEBUG-Test");
                        Logger.logInfo("INFO-Test");
                        Logger.logWarn("WARN-Test");
                        Logger.logError("ERROR-Test");
                        */
                    }
                    else {
                        Logger.logWarn("Received invalid client LogLevel: [" + levelToLower + "]; Logger turned off!");
                        LoggerLevel = LogLevel.OFF;
                    }
                }
            }
        };

        xhr.send();
    }

    static shouldLog(level) {
        if (LoggerLevel == undefined) {
            return true;
        }

        if (LoggerLevel == LogLevel.DEBUG && (level.toLowerCase() == LogLevel.DEBUG || level.toLowerCase() == LogLevel.INFO || level.toLowerCase() == LogLevel.WARN || level.toLowerCase() == LogLevel.ERROR)) {
            return true;
        }

        if (LoggerLevel == LogLevel.INFO && (level.toLowerCase() == LogLevel.INFO || level.toLowerCase() == LogLevel.WARN || level.toLowerCase() == LogLevel.ERROR)) {
            return true;
        }

        if (LoggerLevel == LogLevel.WARN && (level.toLowerCase() == LogLevel.WARN || level.toLowerCase() == LogLevel.ERROR)) {
            return true;
        }

        if (LoggerLevel == LogLevel.ERROR && level.toLowerCase() == LogLevel.ERROR) {
            return true;
        }

        return false;
    }

    static getSenderId() {
        return SENDER_ID;
    }

    static enqueueLog(message, level) {
        LoggerQueue.push({message: message, level: level});
    }

    static processQueue() {
        Logger.logDebug("Processing Logger queue... Queued items: " + LoggerQueue.length);

        if (LoggerLevel != undefined) {
            while (LoggerQueue.length > 0) {
                const log = LoggerQueue.shift();
                Logger.logMessage(log.message, log.level);
            }
        }
    }

    static logRemote(message, level) {
        let xhr = new XMLHttpRequest();

        //ensure that message and level are strings
        message = String(message);
        level = String(level);
        
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
        const logString = "[" + level + "] " + message;

        console.log(logString);

        const logDiv = document.querySelector("div#log");
    
        if (logDiv) {
            logDiv.innerHTML = "<p>" + logString + "</p>" + logDiv.innerHTML
        }
    }

    static logMessage(message, level) {
        if (LoggerLevel != undefined) {
            if (Logger.shouldLog(level)) {
                // run async
                setTimeout(() => {
                    Logger.logLocal(message, level);
                    Logger.logRemote(message, level);
                }, 0);
            }
        }
        else {
            if (!LoggerInitCalled) {
                Logger.init();
            }

            Logger.enqueueLog(message, level);
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