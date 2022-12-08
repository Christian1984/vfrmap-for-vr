import initTour from "../common/initTour";

const mapsTour = () => {
    const tour = initTour();

    tour.addStep({
        title: "Welcome to the Map Module",
        text: "Currently, you are looking at <b>FSKneeboard's Map Module</b>.<br /><br />It offers a fully-fledged moving map with several, <b>Highly Detailed Map Modes</b> and even <b>Navigation Overlays</b> for all your VFR navigation needs.<br /><br />If you have opted in to support FSKneeboard's development by purchasing a <b>PRO License</b>, you have access to even more features such as <b>Waypoints and Tracks</b> or the wonderful <b>Location Finder</b>.<br /><br/>Let's take a look around and familiarize ourselves with everything the <b>Map Module</b> has to offer...",
        buttons: [
            {
                action() {
                    return this.cancel();
                },
                classes: "shepherd-button-secondary",
                text: "Skip Tour"
            },
            {
                action() {
                    return this.next();
                },
                text: "Next"
            }
        ]
    });

    tour.addStep({
        title: "Change Map Style & Overlays",
        text: "FSKneeboard offers a variety of different <b>Map Styles and Overlays</b> which you can choose from here.<br /><br />The first section allows you to change the appearance of the <b>Base Map</b> (i.e. different map styles), while the second section allows you to toggle a variety of <b>Navigation Data Overlays</b>.",
        attachTo: {
            element: ".leaflet-control-layers .leaflet-control-layers-list",
            on: "left"
        },
        when: {
            "before-hide": () => {
                const layersList = document.querySelector(".leaflet-control-layers");
                if (layersList) {
                    layersList.classList.remove("leaflet-control-layers-expanded")
                }
            },
            "before-show": () => {
                const layersList = document.querySelector(".leaflet-control-layers");
                if (layersList) {
                    layersList.classList.add("leaflet-control-layers-expanded")
                }
            }
        }
    });

    tour.addStep({
        title: "Center Aircraft",
        text: "You can move the map around by dragging it with your mouse cursor. Click this button to <b>Center the Map</b> on your aircraft and enable the <b>Follow Mode</b>.",
        attachTo: {
            element: "#tour-center-ac",
            on: "left"
        }
    });

    tour.addStep({
        title: "Manage Your Flight Plan (PRO Feature)",
        text: "Use these buttons to <ul><li>load the <b>Currently Active Flight Plan</b> from Flight Simulator directly into FSKneeboard,</li><li><b>Delete All Waypoints</b>,</li><li>and open the <b>Location Finder</b> to locate and navigate to any POI or address in the entire world!</li>",
        attachTo: {
            element: "#tour-flightplan",
            on: "left"
        }
    });

    tour.addStep({
        title: "Add Waypoints (PRO)",
        text: "Click here to activate the <b>Add Waypoint Mode</b>. Then, click anywhere on the map to append a new waypoint to your current flight plan. Individual waypoints will be connected by a purple <b>Track Line</b>.<br /><br />Clicking an existing track line will insert a new waypoint between the two existing waypoints.",
        attachTo: {
            element: "#tour-add-waypoints",
            on: "left"
        }
    });

    tour.addStep({
        title: "Rubberband (PRO)",
        text: "Click <b>Your Own Aircraft</b> to toggle the <b>Purple Rubberband</b> that indicates the direction from you Aircraft to the <b>First Waypoint</b> of your flight plan.<br/><br/>You can also click individual <b>Waypoint Icons</b> to toggle the <b>Navigation Data Flag</b> for this waypoint!",
        /*attachTo: {
            element: ".leaflet-marker-icon",
            on: "top"
        },*/
        when: {
            "before-show": () => {
                const btn = document.querySelector("#ac-toggle-follow");
                if (btn) {
                    btn.click();
                }
            }
        }
    });

    tour.addStep({
        title: "Add Measurement Tool Pins (PRO)",
        text: "Click here to activate the <b>Add Measurement Tool Pins Mode</b>. You can now click anywhere on the map to create up to 3 \"tool pins\".<br /><br />Distances and directions between two markers will automatically be drawn and calculated. Once you place the third pin, FSKneeboard will also calculate the angle between the \"rulers\".",
        attachTo: {
            element: "#tour-add-tool-pin",
            on: "left"
        }
    });


    tour.addStep({
        title: "Remove Waypoints and Tool Pins (PRO)",
        text: "Click this icon to activate the <b>Remove Waypoints and Tool Pins Mode</b>. Next, click any waypoint or tool pin to remove it from your current flight plan.<br /><br />FYI: The <b>Active Waypoint</b> of your current flight plan will get removed automatically once you arrive at a close proximity of that waypoint (i.e. less than .5 NM).",
        attachTo: {
            element: "#tour-remove-waypoints",
            on: "left"
        }
    });

    tour.addStep({
        title: "Teleport",
        text: "<b>Click This Button</b> and then <b>Anywhere on The Map</b> to bring up the <b>Teleport Dialog</b>. Set your desired altitude and confirm to get teleported to the selected location.",
        attachTo: {
            element: "#tour-teleport",
            on: "left"
        }
    });

    tour.addStep({
        title: "Deactivate Tools",
        text: "Click this 'Tool' to disable all map tools and avoid unintended changes on your current flight plan etc.",
        attachTo: {
            element: "#tour-deactivate-maptools",
            on: "left"
        }
    });

    tour.addStep({
        title: "Map Zoom",
        text: "Use the <i class='fas fa-search-plus'></i> and <i class='fas fa-search-minus'></i> buttons to set the <b>Map Zoom</b>. Alternatively, use your mouse wheel while hovering over the map itself to zoom in and out.",
        attachTo: {
            element: "#tour-zoom>span:last-child",
            on: "left"
        }
    });

    tour.addStep({
        title: "Visibility Options",
        text: "With these buttons you can set one of three <b>Visibility Options</b> for your aircraft's representation on the map <ul><li><i class='fas fa-eye-slash'></i> renders your own aircraft's position <b>invisible</b> for a real 'paper map on kneeboard' navigation experience</li><li><i class='fas fa-plane fa-flip-horizontal'></i> shows your aircraft's location represented by an <b>Aeroplane Icon</b></li><li><i class='fas fa-helicopter fa-flip-horizontal'></i> shows your aircraft's location represented by an <b>Helicopter Icon</b></li></ul>Below, you can see two toggles:<ul><li><i class='fas fa-wind'></i> allows you to show or hide the <b>Wind Indicator</b> in the upper left corner of the map module.</li><li><li><i class='fas fa-route'> toggles the visibility of the <b>Trail Line</b> that tracks the path you have travelled so far.</i></li></ul>",
        attachTo: {
            element: "#tour-visibility",
            on: "left"
        }
    });

    tour.addStep({
        title: "Map Rotation",
        text: "You're almost done :-)<br /><br />Here you can control how the <b>Map is Rotated</b>. Three options are available: <ul><li><i class='fas fa-arrow-alt-circle-up'></i> enables the default <b>North Up Mode</b>.</li><li><i class='fas fa-plane fa-rotate-270'></i> enables the <b>Track Up Mode</b> which will ensure that your aircraft's nose always faces up (similar to your car's navigation system).</li><li>Finally, <i class='fas fa-lock-open'></i> enables the <b>Manual Rotation Mode</b> which enables you to rotate the map freely in the direction of choice by dragging the little <b>Rotation Gizmo</b> that appears on the bottom right of the map when this mode is active.</li></ul>",
        attachTo: {
            element: "#tour-rotation",
            on: "left"
        }
    });

    tour.addStep({
        title: "You're All Set",
        text: "This concludes the tour of FSKneeboard's <b>Map Module</b>. You can now explore the <b>Map Module</b> on your own or visit the <b>Charts Viewer</b> or the <b>Notepad</b> to learn more about the other modules as well.<br /><br /><em>If you want to take the tour again, you can always restart it through the <b>Settings Dialog</b> in the <b>FSKneeboard Server GUI Window</b>.</em>",
        buttons: [
            {
                action() {
                    return this.back();
                },
                    classes: "shepherd-button-secondary",
                    text: "Back"
            },
            {
                action() {
                    return this.complete();
                },
                text: "Finish Tour"
            }
        ],
    });

    fetch("/tour/mapStarted")
        .then(resp => resp.text())
        .then(t => {
            if (t === "false") {
                tour.start();
            }
        })
        .catch(() => logger.logError("Could not load TourMapStatus!"));
}

export default mapsTour;