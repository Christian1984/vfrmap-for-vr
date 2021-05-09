"use strict";

const MODES = {
    add_track_markers: 0,
    delete_track_markers: 1,
    teleport: 2,
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
let last_report = {};
const initial_pos = L.latLng(50.8694,7.1389);

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

let svgPlaneIconString = '<?xml version="1.0" encoding="UTF-8" standalone="no"?><svg xmlns="http://www.w3.org/2000/svg" height="249.84" width="248.25" version="1.0"><metadata id="metadata9"/><path id="path5724" d="M 247.51404,152.40266 139.05781,71.800946 c 0.80268,-12.451845 1.32473,-40.256266 0.85468,-45.417599 -3.94034,-43.266462 -31.23018,-24.6301193 -31.48335,-5.320367 -0.0693,5.281361 -1.01502,32.598388 -1.10471,50.836622 L 0.2842717,154.37562 0,180.19575 l 110.50058,-50.48239 3.99332,80.29163 -32.042567,22.93816 -0.203845,16.89693 42.271772,-11.59566 0.008,0.1395 42.71311,10.91879 -0.50929,-16.88213 -32.45374,-22.39903 2.61132,-80.35205 111.35995,48.50611 -0.73494,-25.77295 z" fill-rule="evenodd" fill="__COLOR__"/></svg>'

let planeIconBlack = L.icon({
    iconUrl: encodeURI("data:image/svg+xml," + svgPlaneIconString).replace("#","%23").replace("__COLOR__", "black"),
    iconSize: [map_resolution.icon_size, map_resolution.icon_size],
});

let planeIconWhite = L.icon({
    iconUrl: encodeURI("data:image/svg+xml," + svgPlaneIconString).replace("#","%23").replace("__COLOR__", "white"),
    iconSize: [map_resolution.icon_size, map_resolution.icon_size],
});

let planeIconGreen = L.icon({
    iconUrl: encodeURI("data:image/svg+xml," + svgPlaneIconString).replace("#","%23").replace("__COLOR__", "green"),
    iconSize: [map_resolution.icon_size, map_resolution.icon_size],
});

function open_in_google_maps() {
    var url = "https://www.google.com/maps/@" + last_report.latitude + "," + last_report.longitude + "," + map.getZoom() + "z"
    window.open(url,'_blank');
}

function updateMap(msg) {
    var pos = L.latLng(msg.latitude, msg.longitude);
    marker.setLatLng(pos);
    marker.setRotationAngle(msg.heading);

    waypoints.update_planepos(plane_visible, pos);

    plane_popup.pos.innerText = Number(pos.lat).toFixed(6) + "," + Number(pos.lng).toFixed(6);

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

    if (map !== undefined) {
        updateMap(msg);
    }
};

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
        layers: [ osm, openaip_cached_basemap ],
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
        icon: planeIconBlack,
        rotationAngle: 0,
        rotationOrigin: "center",
    });

    marker.addTo(map);
    marker.bindPopup(L.popup({autoPan: false}).setLatLng(pos).setContent(plane_popup.main));

    var markerPos = L.latLng(0,0);
    markerTeleport = L.marker(markerPos, {});
    markerTeleport.addTo(map);
    markerTeleport.bindPopup(L.popup({autoPan: false}).setContent(teleport_popup.main));
    set_teleport_marker(markerPos);
    hide_teleport_marker();

    waypoints = new Waypoints(map, pos, mode_options);

    map.on('dragstart', function(e) {
        set_follow(false);
    });

    map.on('dragend', function(e) {
        updateStoredCenter();
    });

    map.on('zoomend', function(e) {
        updateStoredCenter();
        localStorage.setItem("n_zoom", map.getZoom());
    });

    map.on('click', function(e) {
        switch(mode_options.mode) {
            case MODES.add_track_markers:
                waypoints.add_marker(e.latlng);
                break;
            case MODES.teleport:
                set_teleport_marker(e.latlng);
                break;
        }
    });

    map.on('baselayerchange', function(e) {
        if (e.name == "Carto Dark (Night Mode)") {
            marker.setIcon(planeIconWhite);
        }
        else if (e.name == "Stamen Toner") {
            marker.setIcon(planeIconGreen);
        }
        else {
            marker.setIcon(planeIconBlack);
        }
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
    localStorage.setItem('b_follow', follow);
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

    localStorage.setItem("b_show_airplane", visible);
}

function toggle_airplane_visibility(e) {
    set_airplane_marker_visibility(e.checked);
}

function center_airplane() {
    let cb = document.querySelector("#hud-controls-show-airplane");
    if (cb) {
        cb.checked = true;
    }

    set_airplane_marker_visibility(true);
    set_follow(true);

    let pos = initial_pos;

    if (last_report && last_report.latitude && last_report.longitude) {
        pos = L.latLng(last_report.latitude, last_report.longitude);
    }

    map.panTo(pos);
}

function updateStoredCenter() {
    const center = map.getCenter();
    localStorage.setItem("n_last_lat", center.lat);
    localStorage.setItem("n_last_long", center.lng);
}

function loadStoredState() {
    const show_airplane = localStorage.getItem('b_show_airplane');
    if (show_airplane !== undefined && show_airplane !== null) {
        set_airplane_marker_visibility(show_airplane == "true");

        const cb = document.querySelector("#hud-controls-show-airplane");
        if (cb) {
            cb.checked = show_airplane == "true";
        }
    }

    const follow = localStorage.getItem('b_follow');
    if (follow !== undefined && follow !== null) {
            set_follow(follow == "true" && show_airplane == "true");
    }
    else {
        set_follow(true);
    }

    const last_long = localStorage.getItem('n_last_long');
    const last_lat = localStorage.getItem('n_last_lat');
    if (!follow_plane
        && last_long !== undefined && last_long !== null
        && last_lat !== undefined && last_lat !== null) {
        map.panTo(L.latLng(last_lat, last_long));
    }

    const zoom = localStorage.getItem('n_zoom');
    if (zoom !== undefined && zoom !== null) {
        map.setZoom(zoom);
    }

    const nav_data = localStorage.getItem('b_nav_data');
    if (nav_data !== undefined && nav_data !== null) {
        const nav_data_cb = document.querySelector(".leaflet-control-layers-selector[type='checkbox']");
        if (nav_data_cb && nav_data == "false") {
            nav_data_cb.click();
        }
    }

    const active_map = localStorage.getItem('n_active_map');
    const nav_data_rbs = document.querySelectorAll(".leaflet-control-layers-selector[type='radio']");
    if (active_map !== undefined && active_map !== null && nav_data_rbs.length > 0 && active_map < nav_data_rbs.length) {
        nav_data_rbs[active_map].click();
    }

    waypoints.load_trackdata();
}

function registerHandlers() {
    const nav_data_cb = document.querySelector(".leaflet-control-layers-selector[type='checkbox']");
    if (nav_data_cb) {
        nav_data_cb.addEventListener("change", (e) => {
            localStorage.setItem("b_nav_data", nav_data_cb.checked)
        });
    }

    const nav_data_rbs = document.querySelectorAll(".leaflet-control-layers-selector[type='radio']");
    for (let i = 0; i < nav_data_rbs.length; i++) {
        nav_data_rbs[i].addEventListener("change", (e) => {
            localStorage.setItem("n_active_map", i)
        });
    }

    const mode_control_btns = document.querySelectorAll("#mode-controls > span");
    for (let i = 0; i < mode_control_btns.length; i++) {
        mode_control_btns[i].addEventListener("click", (e) => {
            for (let btn of mode_control_btns) {
                btn.classList.remove("selected");
            }

            mode_control_btns[i].classList.add("selected");
            mode_options.mode = i;

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
        const teleport = document.querySelector("#mode-controls-teleport");
        if (teleport) {
            teleport.click();
        }
    }
}

function hide_premium_info(hide = true) {
    const infobox = document.querySelector("#premium-info");
    if (!infobox) return;

    if (hide) {
        infobox.classList.add("hidden");
    }
    else {
        infobox.classList.remove("hidden");
    }
}


document.addEventListener("DOMContentLoaded", function(event) {
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

    initMap();
    loadStoredState();
    registerHandlers();
    activate_default_mode();
});