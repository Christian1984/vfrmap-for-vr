import "shepherd.js/dist/css/shepherd.css";
import "./tour.scss";

import Shepherd from "shepherd.js";

const initTour = () => {
    return new Shepherd.Tour({
        useModalOverlay: true,
        defaultStepOptions: {
            cancelIcon: {
                enabled: false
            },
            scrollTo: {
                behavior: "smooth",
                block: "center" 
            },
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
            ],
            popperOptions: {
                modifiers: [{ name: 'offset', options: { offset: [0, 12] } }]
            }
        }
    });
}

export default initTour;