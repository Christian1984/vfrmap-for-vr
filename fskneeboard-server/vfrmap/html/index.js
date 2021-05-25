const zoom_modification_factor = 1.05;

const content_div = document.getElementById("content");

const iframe_map = document.getElementById("iframe_map");
const iframe_charts = document.getElementById("iframe_charts");
const iframe_notepad = document.getElementById("iframe_notepad");

const switch_map = document.getElementById("switch_map");
const switch_charts = document.getElementById("switch_charts");
const switch_notepad = document.getElementById("switch_notepad");

const temp = document.getElementById("temp");

const zoom_in = document.getElementById("zoom_in");
const zoom_out = document.getElementById("zoom_out");
const stretch = document.getElementById("stretch");
const unstretch = document.getElementById("unstretch");
const reset = document.getElementById("reset");

const current_zoom = { x: 1, y: 1 };

function dispatch_keyevent(event) {
    if (event instanceof KeyboardEvent) {
        const msg = JSON.stringify({
            type: "KeyboardEvent",
            data: {
                type: event.type,
                keyCode: event.keyCode,
                altKey: event.altKey
            }
        });
        
        window.parent.window.postMessage(msg , "*")
    }
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
            reset_zoom();
        });
    }

    window.document.addEventListener("keydown", (e) => {
        dispatch_keyevent(e);
    });

    load_state();
}

init();