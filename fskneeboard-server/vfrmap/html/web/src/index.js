import "./index.scss";
import "./mainmenu.scss";

import "./md5.min.js";
import "./sender.js";
import "./logger.js";
import "./common.js";

const zoom_modification_factor = 1.05;
const brightness_modification = 10;

const content_div = document.getElementById("content");

const iframe_map = document.getElementById("iframe_map");
const iframe_charts = document.getElementById("iframe_charts");
const iframe_notepad = document.getElementById("iframe_notepad");

const switch_map = document.getElementById("switch_map");
const switch_charts = document.getElementById("switch_charts");
const switch_notepad = document.getElementById("switch_notepad");

const temp = document.getElementById("temp");

const red_light = document.getElementById("red_light");

const brightness_down = document.getElementById("brightness_down");
const brightness_up = document.getElementById("brightness_up");
const brightness_reset = document.getElementById("brightness_reset");

const zoom_in = document.getElementById("zoom_in");
const zoom_out = document.getElementById("zoom_out");
const zoom_reset = document.getElementById("zoom_reset");


const stretch = document.getElementById("stretch");
const unstretch = document.getElementById("unstretch");
const reset_stretch_button = document.getElementById("reset-stretch");

let current_zoom = { x: 1, y: 1 };
let current_brightness = 100;

function dispatch_keyevent_top(event) {
    const msg = JSON.stringify({
        type: "KeyboardEvent",
        data: {
            type: event.type,
            keyCode: event.keyCode,
            altKey: event.altKey,
            shiftKey: event.shiftKey,
            ctrlKey: event.ctrlKey
        }
    });

    Logger.logDebug("index.js => dispatching key event to parent: msg=" + msg);
    
    window.parent.window.postMessage(msg , "*");
}

function hide_all_iframes() {
    iframe_map.classList.add("hidden");
    iframe_charts.classList.add("hidden");
    iframe_notepad.classList.add("hidden");
}

function unselect_all_buttons() {
    switch_map.classList.remove("active");
    switch_charts.classList.remove("active");
    switch_notepad.classList.remove("active");
}

function switch_to_map() {
    hide_all_iframes();
    unselect_all_buttons();
    iframe_map.classList.remove("hidden");
    switch_map.classList.add("active");
    save_tab(0);
}

function switch_to_charts() {
    hide_all_iframes();
    unselect_all_buttons();
    iframe_charts.classList.remove("hidden");
    switch_charts.classList.add("active");

    setTimeout(() => {
        iframe_charts.contentWindow.postMessage("load", "*");
    }, 1000);

    save_tab(1);
}

function switch_to_notepad() {
    hide_all_iframes();
    unselect_all_buttons();
    iframe_notepad.classList.remove("hidden");
    switch_notepad.classList.add("active");

    setTimeout(() => {
        iframe_notepad.contentWindow.postMessage("load", "*");
    }, 1000);
    
    save_tab(2);
}

function save_tab(tab_id) {
    // store the active tab only on session-level/locally
    // to make sure that map loads properly upon first page load
    store_data("active_tab", tab_id, false);
}

function save_red(red) {
    if (red == null) return;
    store_data("red", red);
}

function save_brightness() {
    store_data("brightness", current_brightness);
}

function save_zoom() {
    store_data("zoom", JSON.stringify(current_zoom));
}

function load_state() {
    retrieve_data("active_tab", data => {
        if (data.active_tab != null && data.active_tab !== "") {
            switch(data.active_tab) {
                case "1":
                    switch_to_charts();
                    break;
                case "2":
                    switch_to_notepad();
                    break;
                default:
                    switch_to_map();
                    break;
            }
        }
    }, false);

    retrieve_data_set(["zoom", "red", "brightness"], data => {
        if (data.zoom != null && data.zoom !== "") {
            try {
                current_zoom = JSON.parse(data.zoom);
            }
            catch(e) { /* ignore silently */ }
    
            apply_zoom();
        }

        if (data.red != null && data.red !== "") {
            set_red_light(data.red == "true");
    
            if (red_light) {
                red_light.checked = data.red == "true";
            }
        }

        if (data.brightness != null && data.brightness !== "") {
            set_brightness(data.brightness);
        }
    });
}

function set_red_light(red) {
    const msg = JSON.stringify({
        type: "SetRedLight",
        data: {
            red: red
        }
    });
    
    window.parent.window.postMessage(msg , "*");
}

function reset_red_light() {
    red_light.checked = false;
    set_red_light(false);
    save_red(false);
}

function set_brightness(brightness) {
    try {
        current_brightness = Number.parseInt(brightness);
    }
    catch (e) {
        current_brightness = 100;
    }

    if (current_brightness > 100) {
        current_brightness = 100;
    }
    else if (current_brightness <= 0) {
        current_brightness = brightness_modification;
    }

    const msg = JSON.stringify({
        type: "SetBrighness",
        data: {
            brightness: current_brightness
        }
    });
    
    window.parent.window.postMessage(msg , "*");

}

function brightness_increase() {
    set_brightness(current_brightness + brightness_modification);
    save_brightness();
}

function brightness_decrease() {
    set_brightness(current_brightness - brightness_modification);
    save_brightness();
}

function reset_brightness() {
    set_brightness(100);
    save_brightness();
}

function apply_zoom() {
    if (!content_div) return;

    const offX = 100 * 0.5 * (1 - 1 / current_zoom.x);
    const offY = 100 * 0.5 * (1 - 1 / current_zoom.y);

    content_div.style.transform = `scale(${current_zoom.x}, ${current_zoom.y})`;
    content_div.style.left = `${offX}%`;  
    content_div.style.right = `${offX}%`;
    content_div.style.top = `${offY}%`;
    content_div.style.bottom = `${offY}%`;
}

function zoom_views(zoom_in) {
    if (zoom_in) {
        current_zoom.x *= zoom_modification_factor;
        current_zoom.y *= zoom_modification_factor;
    }
    else {
        current_zoom.x /= zoom_modification_factor;
        current_zoom.y /= zoom_modification_factor;
    }

    apply_zoom();
    save_zoom();
}

function stretch_views(stretch) {
    if (stretch) {
        current_zoom.x *= zoom_modification_factor;
    }
    else {
        current_zoom.x /= zoom_modification_factor;
    }

    apply_zoom();
    save_zoom();
}

function reset_zoom() {
    const ratio = current_zoom.y / current_zoom.x;

    current_zoom.x = 1;
    current_zoom.y = ratio;

    apply_zoom();
    save_zoom();
}

function reset_stretch() {
    current_zoom.x = current_zoom.y;

    apply_zoom();
    save_zoom();
}

function request_hotkey() {
    Logger.logDebug("index.js => requesting hotkey configuration...");

    var xhr = new XMLHttpRequest();
    var url = "/hotkey";
    xhr.open("GET", url, true);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                Logger.logDebug("index.js => received hotkey configuration: " + xhr.responseText);

                var json = JSON.parse(xhr.responseText);

                if (json && json.keycode != null) {
                    const msg = JSON.stringify({
                        type: "HotkeyConfiguration",
                        data: {
                            keyCode: json.keycode,
                            altKey: json.altkey,
                            ctrlKey: json.ctrlkey,
                            shiftKey: json.shiftkey
                        }
                    });

                    Logger.logDebug("index.js => propagating hotkey configuration to parent: msg=" + msg);
                    
                    window.parent.window.postMessage(msg , "*");
                }
            }
        }
    };
    xhr.send();
}

function init() {
    if (iframe_map) {
        iframe_map.src = '/freemium/maps.html';
    }

    if (iframe_charts) {
        iframe_charts.src = '/premium/charts.html';
    }

    if (iframe_notepad) {
        iframe_notepad.src = '/premium/notepad.html';
    }

    if(switch_map) {
        switch_map.addEventListener("click", () => {
            switch_to_map();
        });
    }

    if(switch_charts) {
        switch_charts.addEventListener("click", () => {
            switch_to_charts();
        });
    }

    if(switch_notepad) {
        switch_notepad.addEventListener("click", () => {
            switch_to_notepad();
        });
    }

    if (brightness_down) {
        brightness_down.addEventListener("click", () => {
            brightness_decrease();
        });
    }

    if (red_light) {
        red_light.addEventListener("change", () => {
            set_red_light(red_light.checked);
            save_red(red_light.checked);
        });
    }

    if (brightness_up) {
        brightness_up.addEventListener("click", () => {
            brightness_increase();
        });
    }

    if (zoom_in) {
        zoom_in.addEventListener("click", () => {
            zoom_views(true);
        });
    }

    if (zoom_out) {
        zoom_out.addEventListener("click", () => {
            zoom_views(false);
        });
    }

    if (stretch) {
        stretch.addEventListener("click", () => {
            stretch_views(true);
        });
    }

    if (unstretch) {
        unstretch.addEventListener("click", () => {
            stretch_views(false);
        });
    }

    if (brightness_reset) {
        brightness_reset.addEventListener("click", () => {
            reset_brightness();
        });
    }

    if (zoom_reset) {
        zoom_reset.addEventListener("click", () => {
            reset_zoom();
        });
    }

    if (stretch_reset) {
        stretch_reset.addEventListener("click", () => {
            reset_stretch();
        });
    }

    request_hotkey();

    // subscribe to websocket
    const ws = new WebSocket("ws://" + window.location.hostname + ":" + window.location.port + "/hotkeysWs");
    ws.onopen = function() {
        Logger.logDebug("index.js => connection to /hotkeysWs websocket opened");
    };
    ws.onclose = function() {
        Logger.logDebug("index.js => connection to /hotkeysWs websocket closed");
    };
    ws.onmessage = e => {
        Logger.logDebug("index.js => received message from /hotkeysWs websocket: " + e.data);
        const msg = JSON.parse(e.data);

        if (msg.msg) {
            request_hotkey();
        }
    };

    window.document.addEventListener("keydown", (e) => {
        Logger.logDebug("index.js => keydown event registered, e.key=[" + e.key + "], e.keyCode=[" + e.keyCode + "], e.code=[" + e.code + "]");
        dispatch_keyevent_top(e);
    });

    load_state();
}

document.addEventListener("DOMContentLoaded", function() {
    Logger.logDebug("index.js => DOMContentLoaded fired!");

    try {
        let txt = "Browser Information:\n";

        txt += "\tCodeName:          " + navigator.appCodeName + "\n";
        txt += "\tName:              " + navigator.appName + "\n";
        txt += "\tVersion:           " + navigator.appVersion + "\n";
        txt += "\tCookies Enabled:   " + navigator.cookieEnabled + "\n";
        txt += "\tPlatform:          " + navigator.platform + "\n";
        txt += "\tUser-agent header: " + navigator.userAgent + "\n";
    
        Logger.logDebug(txt);
    }
    catch (e) {
        Logger.logWarn("Could not retrieve browser information");
    }

    init();
});