import "shepherd.js/dist/css/shepherd.css";
import "../common/tour.scss";

import Shepherd from "shepherd.js";

const triggerTour = (selector) => {
    const target_iframe = document.getElementById(selector);
    if (target_iframe) {
        target_iframe.contentWindow.postMessage("tour", "*");
    }
}

Shepherd.on("complete", () => {
    console.log("complete");
    // TODO: trigger sub panel tours

    triggerTour("iframe_map");
    triggerTour("iframe_charts");
    triggerTour("iframe_notepad");
});

const tour = () => {
    const tour = new Shepherd.Tour({
        useModalOverlay: true,
        defaultStepOptions: {
            cancelIcon: {
                enabled: true
            },
            scrollTo: {
                behavior: "smooth",
                block: "center" 
            },
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
                        return this.next();
                    },
                    text: "Next"
                }
            ]
        }
    });

    tour.addStep({
        title: "Welcome to FSKneeboard",
        text: "<b>Great you're here!</b> Let's take a quick look around to get you up to speed...<br /><br /><em>If you decide to skip the tour for now, you can always restart it through the <b>Settings Dialog</b> in the <b>FSKneeboard Server GUI Window</b>.</em>",
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
        title: "VR Zoom Adjustment",
        text: "Let's start with the most important feature for VR pilots: The <b>VR Zoom</b> or <b>Dynamic Resolution Scaling</b>.<br /><br />If anything inside FSKneeboard appears too small in VR, use the <i class='fas fa-plus-square'></i> button here to improve the readability of the ingame panel's contents.<br /><br />Clicking the <i class='fas fa-undo'></i> button will reset the VR Zoom.",
        attachTo: {
            element: "#tour-vr-zoom",
            on: "top"
        }
    });

    tour.addStep({
        title: "Adjust Stretch",
        text: "If the map inside FSKneeboard appears <b>stretched</b> or <b>squashed</b>, adjust with the <i class='fas fa-plus-square'></i> and <i class='fas fa-minus-square'></i> buttons here.",
        attachTo: {
            element: "#tour-vr-stretch",
            on: "top"
        }
    });

    tour.addStep({
        title: "Brightness Controls",
        text: "Flying at night can be challenging in VR. Here you can adjust the <b>Brightness</b> of the FSKneeboard panel with the <i class='fas fa-plus-square'></i> and <i class='fas fa-minus-square'></i> buttons or ...",
        attachTo: {
            element: "#tour-brightness",
            on: "top"
        }
    });

    tour.addStep({
        title: "Red Flashlight Mode",
        text: "... even toggle a <b>Red Flashlight Mode</b> by flipping the <i class='checked fas fa-toggle-on'></i> switch.",
        attachTo: {
            element: "#tour-flashlight",
            on: "top"
        }
    });

    tour.addStep({
        title: "Main Navigation",
        text: "Finally, there is the <b>Main Navigation</b>. Here you can switch between the <b>Map Module</b>, the <b>Charts Viewer</b> and the <b>Notepad</b>.",
        attachTo: {
            element: "#switch_map",
            on: "right"
        }
    });

    tour.addStep({
        title: "What's Next?",
        text: "Next, let's take a look at the individual modules!"
    });

    tour.start();
}

export default tour;