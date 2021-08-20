class Waypoints {
    constructor() {
        //intentionally left empty
    }

    is_mode_available() {
        return false;
    }

    activate_mode_failed(show_infobox_callback) {
        show_infobox_callback(false);
    }

    load_trackdata() {
        //intentionally left empty
    }

    update_planepos() {
        //intentionally left empty
    }

    set_plane_visibility() {
        //intentionally left empty
    }
    
    set_rubberband_visibility() {
        //intentionally left empty
    }

    add_marker() {
        //intentionally left empty
    }

    has_waypoints() {
        return false;
    }

    update_track() {
        //intentionally left empty
    }
    
    load_flightplan() {
        //intentionally left empty
    }
    
    search_map() {
        //intentionally left empty
    }
}