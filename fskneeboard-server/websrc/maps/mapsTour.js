import "shepherd.js/dist/css/shepherd.css";
import "../common/tour.scss";

import Shepherd from "shepherd.js";

Shepherd.on("complete", () => {
    console.log("complete");
    // TODO: trigger sub panel tours
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
        title: "VR Zoom Adjustment",
        text: "Let's start with the most important feature for VR pilots: The <b>VR Zoom</b> or <b>Dynamic Resolution Scaling</b>.<br /><br />If anything inside FSKneeboard appears too small in VR, use the <i class='fas fa-plus-square'></i> button here to improve the readability of the ingame panel's contents.<br /><br />Clicking the <i class='fas fa-undo'></i> button will reset the VR Zoom.",
        attachTo: {
            element: "#tour-vr-zoom",
            on: "top"
        }
    });

    tour.start();
}

export default tour;