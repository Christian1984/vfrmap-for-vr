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

const brightness_down = document.getElementById("brightness_down");
const red_light = document.getElementById("red_light");
const brightness_up = document.getElementById("brightness_up");
const zoom_in = document.getElementById("zoom_in");
const zoom_out = document.getElementById("zoom_out");
const stretch = document.getElementById("stretch");
const unstretch = document.getElementById("unstretch");
const reset = document.getElementById("reset");

let current_zoom = { x: 1, y: 1 };
let current_brightness = 100;

function dispatch_keyevent(event) {
    console.log(event);
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
    localStorage.setItem("active_tab", tab_id);
}

function save_red(red) {
    if (red == null) return;
    localStorage.setItem("red", red);
}

function save_brightness() {
    localStorage.setItem("brightness", current_brightness);
}

function save_zoom() {
    localStorage.setItem("zoom", JSON.stringify(current_zoom));
}

function load_state() {
    const zoom = localStorage.getItem("zoom");
    if (zoom != null) {
        try {
            current_zoom = JSON.parse(zoom);
        }
        catch(e) { /* ignore silently */ }

        apply_zoom();
    }

    const red = localStorage.getItem("red");
    if (red != null) {
        set_red_light(red == "true");
        red_light.checked = red == "true";
    }

    const brightness = localStorage.getItem("brightness");
    if (brightness != null) {
        set_brightness(brightness);
    }

    const active_tab = localStorage.getItem("active_tab");
    if (active_tab != null) {
        switch(active_tab) {
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
    current_brightness = brightness;

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
    current_zoom.x = 1;
    current_zoom.y = 1;

    apply_zoom();
    save_zoom();
}

function request_hotkey() {
    var xhr = new XMLHttpRequest();
    var url = "/hotkey";
    xhr.open("GET", url, true);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
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
                    
                    window.parent.window.postMessage(msg , "*");
                }
            }
        }
    };
    xhr.send();
}

function init() {
    if (iframe_map) {
        iframe_map.src = 'http://localhost:9000/freemium/maps.html';
    }

    if (iframe_charts) {
        iframe_charts.src = 'http://localhost:9000/premium/charts.html';
    }

    if (iframe_notepad) {
        iframe_notepad.src = 'http://localhost:9000/premium/notepad.html';
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

    if (reset) {
        reset.addEventListener("click", () => {
            reset_brightness();
            reset_zoom();
            reset_red_light();
        });
    }

    request_hotkey();

    window.document.addEventListener("keydown", (e) => {
        dispatch_keyevent(e);
    });

    load_state();
}

init();