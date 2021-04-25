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

    add_marker() {
        //intentionally left empty
    }

    update_track() {
        //intentionally left empty
    }
}