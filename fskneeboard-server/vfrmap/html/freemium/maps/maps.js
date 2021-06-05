"use strict";

const MODES = {
    add_track_markers: 0,
    delete_track_markers: 1,
    teleport: 2,
}

const AC_TYPE = {
    airplane: 0,
    helicopter: 1,
}

const AC_COLOR = {
    black: 0,
    white: 1,
    green: 2,
}

let map;
let marker;
let markerTeleport;
let markerIcon;
let popup;
let ws;
let plane_popup;
let teleport_popup;
let waypoints;
let follow_plane = false;
let plane_visible = true;
let mode_options = { mode: MODES.add_track_markers };
let ac_visibility_options = {
    ac_visibility: true,
    ac_type: AC_TYPE.airplane, 
    ac_color: AC_COLOR.black
};
let last_report = {};
const initial_pos = L.latLng(50.8694,7.1389);
const autoremoval_proximity_threshold = 0.5; //miles

let wind_indicator;
let wind_indicator_arrow;
let wind_indicator_direction;
let wind_indicator_velocity;

const map_resolutions = {
    high: {
        tile_size: 256,
        zoom_offset: 0,
        icon_size: 32
    },
    medium: {
        tile_size: 512,
        zoom_offset: -1,
        icon_size: 64
    },
    low: {
        tile_size: 1024,
        zoom_offset: -2,
        icon_size: 128
    }
};

let map_resolution = map_resolutions.high;

const svgPlaneIconString = '<?xml version="1.0" encoding="UTF-8" standalone="no"?><svg xmlns="http://www.w3.org/2000/svg" height="249.84" width="248.25" version="1.0"><metadata id="metadata9"/><path id="path5724" d="M 247.51404,152.40266 139.05781,71.800946 c 0.80268,-12.451845 1.32473,-40.256266 0.85468,-45.417599 -3.94034,-43.266462 -31.23018,-24.6301193 -31.48335,-5.320367 -0.0693,5.281361 -1.01502,32.598388 -1.10471,50.836622 L 0.2842717,154.37562 0,180.19575 l 110.50058,-50.48239 3.99332,80.29163 -32.042567,22.93816 -0.203845,16.89693 42.271772,-11.59566 0.008,0.1395 42.71311,10.91879 -0.50929,-16.88213 -32.45374,-22.39903 2.61132,-80.35205 111.35995,48.50611 -0.73494,-25.77295 z" fill-rule="evenodd" fill="__COLOR__"/></svg>';
const svgHelicopterIconString = '<?xml version="1.0" encoding="iso-8859-1"?><svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="250px" height="250px" viewBox="0 0 478.874 478.873"><g><g><path d="M463.096,252.605l-133.38-52.861V78.503V47.101c0-4.338-3.519-7.851-7.851-7.851s-7.851,3.513-7.851,7.851v31.402h-11.569C293.433,32.987,266.884,0,235.512,0c-31.37,0-57.919,32.987-66.938,78.503h-19.416V47.101c0-4.338-3.519-7.851-7.851-7.851s-7.85,3.513-7.85,7.851v31.402v43.46l-109-43.2c-6.987-2.771-14.597-0.112-16.99,5.933c-2.395,6.045,1.327,13.187,8.312,15.961l117.678,46.639v80.363v23.551c0,4.341,3.518,7.851,7.85,7.851s7.851-3.51,7.851-7.851V227.66h48.1c7.64,25.239,14.703,58.196,14.703,94.207v78.502h7.851v39.528c0,8.079,7.027,14.644,15.701,14.644c8.674,0,15.699-6.564,15.699-14.644v-39.528h7.851v-78.502c0-35.618,6.984-68.655,14.606-94.207h40.347v23.551c0,4.341,3.519,7.851,7.851,7.851s7.851-3.51,7.851-7.851V227.66v-2.583l124.703,49.425c6.981,2.773,14.596,0.121,16.987-5.935C473.799,262.512,470.081,255.383,463.096,252.605z M314.015,94.204v99.322l-24.132-9.567c9.91-19.424,15.877-44.248,15.877-71.307c0-6.297-0.409-12.435-1.03-18.448H314.015z M149.158,94.204h17.132c-0.621,6.014-1.023,12.151-1.023,18.448c0,7.694,0.486,15.207,1.406,22.468l-17.515-6.939V94.204z M149.158,211.958v-58.436l23.536,9.327c1.775,5.688,3.829,11.093,6.155,16.186l-0.433-0.148c0,0,6.476,12.457,13.74,33.071H149.158z M278.714,211.958c0.749-2.18,1.479-4.208,2.22-6.215l15.682,6.215H278.714z" fill="__COLOR__"/><path d="M266.913,408.219c-4.328,0-7.851,3.518-7.851,7.85v54.954c0,4.332,3.522,7.851,7.851,7.851c4.332,0,7.85-3.519,7.85-7.851v-54.954C274.762,411.736,271.245,408.219,266.913,408.219z" fill="__COLOR__"/></g></g></svg>';

const icons = {
    planes: {
        black: L.icon({
            iconUrl: encodeURI("data:image/svg+xml," + svgPlaneIconString).replace("__COLOR__", "black"),
            iconSize: [map_resolution.icon_size, map_resolution.icon_size],
        }),
        white: L.icon({
            iconUrl: encodeURI("data:image/svg+xml," + svgPlaneIconString).replace("__COLOR__", "white"),
            iconSize: [map_resolution.icon_size, map_resolution.icon_size],
        }),
        green: L.icon({
            iconUrl: encodeURI("data:image/svg+xml," + svgPlaneIconString).replace("__COLOR__", "green"),
            iconSize: [map_resolution.icon_size, map_resolution.icon_size],
        })
    },
    helicopters: {
        black: L.icon({
            iconUrl: encodeURI("data:image/svg+xml," + svgHelicopterIconString).replace("__COLOR__", "black"),
            iconSize: [map_resolution.icon_size, map_resolution.icon_size],
        }),
        white: L.icon({
            iconUrl: encodeURI("data:image/svg+xml," + svgHelicopterIconString).replace("__COLOR__", "white"),
            iconSize: [map_resolution.icon_size, map_resolution.icon_size],
        }),
        green: L.icon({
            iconUrl: encodeURI("data:image/svg+xml," + svgHelicopterIconString).replace("__COLOR__", "green"),
            iconSize: [map_resolution.icon_size, map_resolution.icon_size],
        })
    }
}

let currentIconGroup = icons.planes;
let currentIcon = currentIconGroup.black;

function open_in_google_maps() {
    var url = "https://www.google.com/maps/@" + last_report.latitude + "," + last_report.longitude + "," + map.getZoom() + "z"
    window.open(url,"_blank");
}

function hide_wind_indicator(hide = true) {
    if (wind_indicator == null) return;

    if (hide) {
        wind_indicator.classList.add("hidden");
    }
    else {
        wind_indicator.classList.remove("hidden");
    }
}

function update_wind_indicator(dir, vel) {
    if (dir == null || vel == null) {
        if (wind_indicator_arrow != null) {
            wind_indicator_arrow.classList.add("hidden");
        }

        if (wind_indicator_direction != null) {
            wind_indicator_direction.innerText = "N/A";
        }

        if (wind_indicator_velocity != null) {
            wind_indicator_velocity.innerText = "N/A";
        }
    }
    else {
        if (wind_indicator_arrow != null) {
            wind_indicator_arrow.classList.remove("hidden");
            wind_indicator_arrow.style.transform = "rotate(" + dir + "deg)";
        }

        if (wind_indicator_direction != null) {
            wind_indicator_direction.innerText = dir;
        }

        if (wind_indicator_velocity != null) {
            wind_indicator_velocity.innerText = vel;
        }
    }
}

function updateMap(msg) {
    var pos = L.latLng(msg.latitude, msg.longitude);
    marker.setLatLng(pos);
    marker.setRotationAngle(msg.heading);

    waypoints.update_planepos(plane_visible, pos);

    plane_popup.pos.innerText = Number(pos.lat).toFixed(6) + "," + Number(pos.lng).toFixed(6);

    update_wind_indicator(msg.wind_direction, msg.wind_velocity);

    if (follow_plane) {
        map.panTo(pos);
    }
}

ws = new WebSocket("ws://" + window.location.hostname + ":" + window.location.port + "/ws");
ws.onopen = function() {
    //console.log("ws open");
};
ws.onclose = function() {
    //console.log("ws close");
};
ws.onmessage = function(e) {
    var msg = JSON.parse(e.data);
    last_report = msg;

    if (map != null) {
        updateMap(msg);
    }
};

function updateIcon() {
    let iconType = icons.planes;
    
    if (ac_visibility_options.ac_type === AC_TYPE.helicopter) {
        iconType = icons.helicopters;
    }

    let currentIcon;

    switch (ac_visibility_options.ac_color) {
        case AC_COLOR.white:
            currentIcon = iconType.white;
            break;
        case AC_COLOR.green:
            currentIcon = iconType.green;
            break;
        default:
            currentIcon = iconType.black;
            break;
    }

    marker.setIcon(currentIcon);

    set_airplane_marker_visibility(ac_visibility_options.ac_visibility);
}

function initMap() {
    var pos = initial_pos;

    var osm = new L.TileLayer("http://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
        maxZoom: 18,
        minZoom: 2,
        tileSize: map_resolution.tile_size,
        zoomOffset: map_resolution.zoom_offset,
        format: "image/png",
        subdomains: ["a", "b", "c"]
    });

    var openaip_cached_basemap = new L.TileLayer("http://{s}.tile.maps.openaip.net/geowebcache/service/tms/1.0.0/openaip_basemap@EPSG%3A900913@png/{z}/{x}/{y}.png", {
        maxZoom: 14,
        minZoom: 4,
        tileSize: map_resolution.tile_size,
        zoomOffset: map_resolution.zoom_offset,
        tms: true,
        //detectRetina: true,
        subdomains: "12",
        format: "image/png",
        transparent: true
    });

    var stamen_black_white = new L.TileLayer("http://{s}.tile.stamen.com/toner/{z}/{x}/{y}.png", {
        maxZoom: 18,
        minZoom: 2,
        tileSize: map_resolution.tile_size,
        zoomOffset: map_resolution.zoom_offset,
        format: "image/png",
        subdomains: ["a", "b", "c"]
    });

    var stamen_terrain = new L.TileLayer("http://{s}.tile.stamen.com/terrain/{z}/{x}/{y}.png", {
        maxZoom: 18,
        minZoom: 2,
        tileSize: map_resolution.tile_size,
        zoomOffset: map_resolution.zoom_offset,
        format: "image/png",
        subdomains: ["a", "b", "c"]
    });

    var stamen_water = new L.TileLayer("http://{s}.tile.stamen.com/watercolor/{z}/{x}/{y}.png", {
        maxZoom: 18,
        minZoom: 2,
        tileSize: map_resolution.tile_size,
        zoomOffset: map_resolution.zoom_offset,
        format: "image/png",
        subdomains: ["a", "b", "c"]
    });

    var carto_dark = new L.TileLayer("https://cartodb-basemaps-{s}.global.ssl.fastly.net/dark_all/{z}/{x}/{y}.png", {
        maxZoom: 18,
        minZoom: 2,
        tileSize: map_resolution.tile_size,
        zoomOffset: map_resolution.zoom_offset,
        format: "image/png",
        subdomains: ["a", "b", "c"]
    });

    map = new L.Map("map", {
        layers: [ osm ],
        center: pos,
        zoom: 10,
        attributionControl: false,
        zoomControl: false
    });

    var attrib = L.control.attribution({position: "bottomleft"});
    attrib.addAttribution("<a href=\"https://www.openstreetmap.org/copyright\" target=\"_blank\" style=\"\">OpenStreetMap</a>");
    attrib.addAttribution("<a href=\"https://www.openaip.net\" target=\"_blank\" style=\"\">openAIP</a>");
    attrib.addAttribution("<a href=\"http://maps.stamen.com\" target=\"_blank\" style=\"\">Stamen</a>");
    attrib.addAttribution("<a href=\"https://carto.com/\" target=\"_blank\" style=\"\">Carto</a>");
    attrib.addAttribution("<a href=\"https://github.com/lian/msfs2020-go\" target=\"_blank\" style=\"\">lian</a>");
    attrib.addAttribution("<a href=\"https://github.com/bymaximus/msfs2020-toolbar-window-template\" target=\"_blank\" style=\"\">bymaximus</a>");
    attrib.addAttribution("<a href=\"https://github.com/Christian1984/msfs2020-go\" target=\"_blank\" style=\"\">Christian1984</a>");

    attrib.addTo(map);

    var baseMaps = {
        "OpenStreetMap": osm,
        "Stamen Terrain": stamen_terrain,
        "Stamen Toner": stamen_black_white,
        "Stamen Water": stamen_water,
        "Carto Dark (Night Mode)": carto_dark,
    };

    var overlayMaps = {
        "Navigational Data": openaip_cached_basemap,
    };

    L.control.layers(baseMaps, overlayMaps).addTo(map);

    marker = L.marker(pos, {
        icon: icons.planes.black,
        rotationAngle: 0,
        rotationOrigin: "center",
    });

    marker.addTo(map);
    //marker.bindPopup(L.popup({autoPan: false}).setLatLng(pos).setContent(plane_popup.main));

    var markerPos = L.latLng(0,0);
    markerTeleport = L.marker(markerPos, {});
    markerTeleport.addTo(map);
    markerTeleport.bindPopup(L.popup({autoPan: false}).setContent(teleport_popup.main));
    set_teleport_marker(markerPos);
    hide_teleport_marker();

    waypoints = new Waypoints(map, pos, mode_options, autoremoval_proximity_threshold);

    map.on("dragstart", function(e) {
        set_follow(false);
    });

    map.on("moveend", function(e) {
        updateStoredCenter();
    });

    map.on("zoomend", function(e) {
        updateStoredCenter();
        localStorage.setItem("n_zoom", map.getZoom());
    });

    map.on("click", function(e) {
        switch(mode_options.mode) {
            case MODES.add_track_markers:
                waypoints.add_marker(e.latlng);
                break;
            case MODES.teleport:
                set_teleport_marker(e.latlng);
                break;
        }
    });

    map.on("baselayerchange", function(e) {
        if (e.name == "Carto Dark (Night Mode)") {
            ac_visibility_options.ac_color = AC_COLOR.white;
        }
        else if (e.name == "Stamen Toner") {
            ac_visibility_options.ac_color = AC_COLOR.green;
        }
        else {
            ac_visibility_options.ac_color = AC_COLOR.black;
        }

        updateIcon();
    });

    map.whenReady(function() {
        registerHandlers();
        loadStoredState();
        activate_default_mode();
    });
}

function set_teleport_marker(latlng) {
    if (!markerTeleport || !markerTeleport._icon) return;

    markerTeleport.setLatLng(latlng);
    teleport_popup.gps.value = latlng.lat.toFixed(8) + "," + latlng.lng.toFixed(8);

    if (last_report.altitude) {
        teleport_popup.altitude.value = last_report.altitude;
    }

    markerTeleport._icon.style.display = "block";
    markerTeleport._shadow.style.display = "block";
}

function hide_teleport_marker() {
    if (!markerTeleport || !markerTeleport._icon) return;

    markerTeleport._icon.style.display = "none";
    markerTeleport._shadow.style.display = "none";
}

function teleport_here() {
    var msg = JSON.stringify(
        {
            "type": "teleport",
            "lat": parseFloat(teleport_popup.gps.value.split(",")[0]),
            "lng": parseFloat(teleport_popup.gps.value.split(",")[1]),
            "altitude": parseFloat(teleport_popup.altitude.value) + 0.5,
        }
    );

    ws.send(msg);
}

function set_follow(follow) {
    if (follow) {
        plane_popup.follow.innerText = "don't follow plane";
    }
    else {
        plane_popup.follow.innerText = "follow plane";
    }

    follow_plane = follow;
    localStorage.setItem("b_follow", follow);

    if (follow) {
        ac_visibility_options.ac_visibility = follow
        save_ac_visibility();
    }
}

function toggle_follow() {
    set_follow(!follow_plane);
}

function set_airplane_marker_visibility(visible) {
    plane_visible = visible;

    if (visible) {
        marker.setOpacity(1);
    }
    else {
        marker.setOpacity(0);
        set_follow(false);
    }

    waypoints.update_track();
}

function update_visibility_buttons() {
    let rb_hidden = document.querySelector("#ac-visibility-none");
    let rb_plane = document.querySelector("#ac-visibility-plane");
    let rb_helicopter = document.querySelector("#ac-visibility-helicopter");

    if (ac_visibility_options.ac_visibility == false) {
        if (rb_hidden) rb_hidden.click();
    }
    else {
        if (ac_visibility_options.ac_type == AC_TYPE.helicopter) {
            if (rb_helicopter) rb_helicopter.click();
        }
        else {
            if (rb_plane) rb_plane.click();
        }
    }
}

function center_airplane() {
    set_follow(true);

    let pos = initial_pos;

    if (last_report && last_report.latitude && last_report.longitude) {
        pos = L.latLng(last_report.latitude, last_report.longitude);
    }

    map.panTo(pos);
    update_visibility_buttons();
}

function updateStoredCenter() {
    const center = map.getCenter();
    localStorage.setItem("n_last_lat", center.lat);
    localStorage.setItem("n_last_long", center.lng);
}

function save_ac_visibility() {
    localStorage.setItem("ac_visibility_options", JSON.stringify(ac_visibility_options));
}

function loadStoredState() {
    const stored_vos = localStorage.getItem("ac_visibility_options");
    if (stored_vos != null) {
        try {
            ac_visibility_options = JSON.parse(stored_vos);
            updateIcon();
        }
        catch(e) {
            /* ignore silently */
        }
    }

    const follow = localStorage.getItem("b_follow");
    if (follow != null) {
        set_follow(follow == "true" && ac_visibility_options.ac_visibility);
    }
    else {
        set_follow(true);
    }

    const last_long = localStorage.getItem("n_last_long");
    const last_lat = localStorage.getItem("n_last_lat");

    if (!follow_plane && last_long != null && last_lat != null) {
        setTimeout(() => {
            map.panTo(L.latLng(last_lat, last_long));
        }, 500);
    }

    const zoom = localStorage.getItem("n_zoom");
    if (zoom != null) {
        map.setZoom(zoom);
    }
    
    const nav_data = localStorage.getItem("b_nav_data");
    if (nav_data != null) {
        const nav_data_cb = document.querySelector(".leaflet-control-layers-selector[type='checkbox']");
        if (nav_data_cb && nav_data == "true") {
            nav_data_cb.click();
        }
    }

    const active_map = localStorage.getItem("n_active_map");
    const nav_data_rbs = document.querySelectorAll(".leaflet-control-layers-selector[type='radio']");
    if (active_map != null && nav_data_rbs.length > 0 && active_map < nav_data_rbs.length) {
        nav_data_rbs[active_map].click();
    }

    const wind_indicator_visibility = localStorage.getItem("wind_indicator_visibility");
    const wind_indicator_toggle = document.querySelector("#wind-indicator-toggle");
    if (wind_indicator_visibility != null && wind_indicator_visibility == "false" && wind_indicator_toggle != null) {
        wind_indicator_toggle.click();
    }

    update_visibility_buttons();
    waypoints.load_trackdata();
}

function registerHandlers() {
    const nav_data_cb = document.querySelector(".leaflet-control-layers-selector[type='checkbox']");
    if (nav_data_cb) {
        nav_data_cb.addEventListener("change", () => {
            localStorage.setItem("b_nav_data", nav_data_cb.checked)
        });
    }

    const nav_data_rbs = document.querySelectorAll(".leaflet-control-layers-selector[type='radio']");
    for (let i = 0; i < nav_data_rbs.length; i++) {
        nav_data_rbs[i].addEventListener("change", () => {
            localStorage.setItem("n_active_map", i)
        });
    }

    const load_flightplan_btn = document.querySelector("#load-waypoints");
    if (load_flightplan_btn) {
        load_flightplan_btn.addEventListener("click", () => {
            if (!waypoints.is_mode_available()) {
                waypoints.activate_mode_failed(hide_premium_info);
                activate_default_mode();
            }
            else {
                if (waypoints.has_waypoints()) {
                    hide_waypoint_confirm_dialog(false);
                }
                else {
                    waypoints.load_flightplan();
                }
            }
        });
    }

    const confirm_load_flightplan_btn = document.querySelector("#waypoint-confirm-dialog-yes");
    if (confirm_load_flightplan_btn) {
        confirm_load_flightplan_btn.addEventListener("click", () => {
            if (waypoints.is_mode_available()) {
                waypoints.load_flightplan();
            }
            hide_waypoint_confirm_dialog();
        });
    }

    const cancel_load_flightplan_btn = document.querySelector("#waypoint-confirm-dialog-no");
    if (cancel_load_flightplan_btn) {
        cancel_load_flightplan_btn.addEventListener("click", () => {
            hide_waypoint_confirm_dialog();
        });
    }

    const mode_control_btns = document.querySelectorAll("#mode-controls > input");
    for (let i = 0; i < mode_control_btns.length; i++) {
        mode_control_btns[i].addEventListener("click", () => {
            switch (mode_control_btns[i].value) {
                case "add-marker":
                    mode_options.mode = MODES.add_track_markers;
                    break;
                case "remove-marker":
                    mode_options.mode = MODES.delete_track_markers;
                    break;
                default:
                    mode_options.mode = MODES.teleport;
            }

            if (mode_options.mode != MODES.teleport) {
                if (!waypoints.is_mode_available()) {
                    waypoints.activate_mode_failed(hide_premium_info);
                    activate_default_mode();
                }
                else {
                    hide_teleport_marker();
                }
            }
        });
    }

    const ac_visibility_control_btns = document.querySelectorAll("#hud-controls > input.ac-visibility");
    for (let i = 0; i < ac_visibility_control_btns.length; i++) {
        ac_visibility_control_btns[i].addEventListener("click", () => {
            ac_visibility_options.ac_visibility = true;

            switch (ac_visibility_control_btns[i].value) {
                case "none":
                    ac_visibility_options.ac_visibility = false;
                    break;
                case "helicopter":
                    ac_visibility_options.ac_type = AC_TYPE.helicopter;
                    break;
                default:
                    ac_visibility_options.ac_type = AC_TYPE.airplane;
            }

            updateIcon();
            save_ac_visibility();
        });
    }

    const wind_indicator_btn = document.querySelector("#wind-indicator-toggle");
    if (wind_indicator_btn) {
        wind_indicator_btn.addEventListener("change", () => {
            console.log(wind_indicator_btn.checked);
            hide_wind_indicator(!wind_indicator_btn.checked);
            localStorage.setItem("wind_indicator_visibility", wind_indicator_btn.checked);
        })
    }

    const center_ac_btn = document.querySelector("#ac-toggle-follow");
    if (center_ac_btn) {
        center_ac_btn.addEventListener("click", () => {
            center_airplane();
        });
    }

    const premium_info_close = document.querySelector("#premium-info-close");
    if (premium_info_close) {
        premium_info_close.addEventListener("click", (e) => {
            e.preventDefault();
            hide_premium_info();
        });
    }
}

function activate_default_mode() {
    if (!waypoints.is_mode_available()) {
        const teleport = document.querySelector("#mode-teleport");
        if (teleport) {
            teleport.click();
        }
    }
}

function hide_waypoint_confirm_dialog(hide = true) {
    const waypoint_confirm_dialog_wrapper = document.querySelector("#waypoint-confirm-dialog-wrapper");
    if (!waypoint_confirm_dialog_wrapper) return;

    if (hide) {
        waypoint_confirm_dialog_wrapper.classList.add("hidden");
    }
    else {
        waypoint_confirm_dialog_wrapper.classList.remove("hidden");
    }
}

function hide_premium_info(hide = true) {
    const infobox = document.querySelector("#premium-info");
    if (!infobox) return;

    const infobox_iframe = infobox.querySelector("iframe");
    if (!infobox_iframe) return;

    if (hide) {
        infobox.classList.add("hidden");
        infobox_iframe.src = "";
    }
    else {
        infobox.classList.remove("hidden");
        infobox_iframe.src = "https://fskneeboard.com/maps-ingame/";
    }
}

document.addEventListener("DOMContentLoaded", function() {
    plane_popup = {
        main: document.getElementById("plane-popup"),
        pos: document.getElementById("plane-popup-pos"),
        gmap: document.getElementById("plane-popup-gmap"),
        follow: document.getElementById("plane-popup-follow"),
    };

    teleport_popup = {
        main: document.getElementById("teleport-popup"),
        submit: document.getElementById("teleport-popup-submit"),
        gps: document.getElementById("teleport-popup-gps"),
        altitude: document.getElementById("teleport-popup-altitude"),
    };

    wind_indicator = document.getElementById("wind-indicator");
    wind_indicator_arrow = document.getElementById("wind-indicator-arrow");
    wind_indicator_direction = document.getElementById("wind-indicator-direction");
    wind_indicator_velocity = document.getElementById("wind-indicator-velocity");

    window.document.addEventListener("keydown", (e) => {
        dispatch_keyevent(e);
    });

    initMap();
});