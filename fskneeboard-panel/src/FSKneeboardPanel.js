import "./FSKneeboardPanel.scss";

let MyTemplateElement;
let myCheckAutoload;

try {
    MyTemplateElement = TemplateElement;
} catch (e) {
    if (parent && parent.window && parent.window.test_environment) {
        MyTemplateElement = class extends HTMLElement {
            constructor(args) {
                super(args);
            }
            connectedCallback() {}
            disconnectedCallback() {}
        };
    }
}

try {
    myCheckAutoload = checkAutoload;
} catch (e) {
    if (parent && parent.window && parent.window.test_environment) {
        myCheckAutoload = function () {};
    }
}

const emptyHotkey = () => {
    return {
        key: -1,
        keycode: -1,
        altkey: false,
        ctrlkey: false,
        shiftkey: false,
    };
};

const equalsHotkey = (event, hotkey) => {
    const res =
        hotkey.key !== -1 &&
        event.keyCode == hotkey.keycode &&
        event.altKey == hotkey.altkey &&
        event.ctrlKey == hotkey.ctrlkey &&
        event.shiftKey == hotkey.shiftkey;

    // console.log("evaluated hotkey equals:", res);

    return res;
};

class IngamePanelFSKneeboardPanel extends MyTemplateElement {
    constructor() {
        super(...arguments);

        this.panelActive = false;
        this.ingameUi = null;
        this.child = null;

        this.collapsed = false;

        this.collapse_hotkey = emptyHotkey();
        this.maps_hotkey = emptyHotkey();
        this.charts_hotkey = emptyHotkey();
        this.notepad_hotkey = emptyHotkey();

        //this.posInterval = -1;

        /*
        this.logDiv = document.querySelector("#log");
        if (this.logDiv) {
            this.logDiv.innerHTML = "<p>log initialized</p>";
        }
        */
    }

    /*
    log(msg) {
        if (this.logDiv) {
            this.logDiv.innerHTML = "<p>" + msg + "</p>" + this.logDiv.innerHTML;
        }
    }
    */

    collapse(collapsed) {
        if (collapsed) {
            this.classList.add("collapsed");
        } else {
            this.classList.remove("collapsed");
        }

        this.collapsed = collapsed;
    }

    toggle_collapse() {
        this.collapse(!this.collapsed);
    }

    set_brightness(brightness) {
        const overlay_brightness = document.querySelector("#overlay_brightness");
        overlay_brightness.style.opacity = (100 - brightness) / 100;

        const overlay_red = document.querySelector("#overlay_red");
        overlay_red.style.filter = "brightness(" + brightness + "%)";
        overlay_red.style.opacity = (-3 / 1000) * brightness + 1;
    }

    set_red_light(red) {
        const panel = document.querySelector("#FSKneeboardPanel");

        if (panel) {
            if (red) {
                panel.classList.add("red");
            } else {
                panel.classList.remove("red");
            }
        }
    }

    sendNavigationIntent(target) {
        if (this.child) {
            const msg = JSON.stringify({
                type: "NavigationIntent",
                data: { target: target },
            });

            this.child.postMessage(msg, "*");
        }
    }

    processHotkeyEvent(event) {
        if (equalsHotkey(event, this.collapse_hotkey)) {
            this.toggle_collapse();
        } else if (equalsHotkey(event, this.maps_hotkey)) {
            this.sendNavigationIntent("maps");
        } else if (equalsHotkey(event, this.charts_hotkey)) {
            this.sendNavigationIntent("charts");
        } else if (equalsHotkey(event, this.notepad_hotkey)) {
            this.sendNavigationIntent("notepad");
        }
    }

    connectedCallback() {
        super.connectedCallback();

        var self = this;

        window.addEventListener("message", (e) => {
            try {
                const data = JSON.parse(e.data);

                switch (data.type) {
                    case "RegisterChild":
                        self.child = e.source;
                        break;

                    case "KeyboardEvent":
                        // console.log("received KeyboardEvent");
                        if (data.data.type == "keydown") {
                            self.processHotkeyEvent(data.data);
                        }
                        break;

                    case "HotkeyConfiguration":
                        self.collapse_hotkey = data.data.masterHotkey ?? emptyHotkey();
                        self.maps_hotkey = data.data.mapsHotkey ?? emptyHotkey();
                        self.charts_hotkey = data.data.chartsHotkey ?? emptyHotkey();
                        self.notepad_hotkey = data.data.notepadHotkey ?? emptyHotkey();
                        break;

                    case "SetBrighness":
                        self.set_brightness(data.data.brightness);
                        break;

                    case "SetRedLight":
                        self.set_red_light(data.data.red);
                        break;
                }
            } catch (e) {
                /* ignore silently */
            }
        });

        window.addEventListener("keydown", (e) => {
            self.processHotkeyEvent(e);
        });

        setTimeout(() => {
            this.ingameUi = this.querySelector("ingame-ui");
            this.content_iframe = document.getElementById("content_iframe");
            this.warning_message = document.getElementById("warning_message");

            if (this.ingameUi) {
                this.ingameUi.addEventListener("panelActive", (e) => {
                    self.panelActive = true;
                    self.warning_message.classList.add("show");

                    if (self.content_iframe) {
                        self.content_iframe.src = "http://localhost:9000/index.html";
                    }

                    /*
                    if (this.posInterval == -1) {
                        this.posInterval = setInterval(() => {
                            try {
                                const lat = SimVar.GetSimVarValue("PLANE LATITUDE", "degrees");
                                const lng = SimVar.GetSimVarValue("PLANE LONGITUDE", "degrees");
                                const alt = SimVar.GetSimVarValue("A:INDICATED ALTITUDE", "meters");
                                this.log("Lat: " + lat + ", Lng: " + lng + ", alt: " + alt + "m");
                            } catch (e) {
                                this.log(e);
                            }
                        }, 5000);
                    }
                    */

                    self.collapse(false);
                });

                this.ingameUi.addEventListener("panelInactive", (e) => {
                    self.panelActive = false;
                    self.warning_message.classList.remove("show");
                    //clearInterval(this.posInterval);
                    //this.posInterval = -1;

                    if (self.content_iframe) {
                        self.content_iframe.src = "";
                    }
                });
            }
        }, 0);
    }

    disconnectedCallback() {
        super.disconnectedCallback();
    }

    updateImage() {}
}

window.customElements.define("ingamepanel-custom", IngamePanelFSKneeboardPanel);
myCheckAutoload();

if (parent && parent.window && parent.window.test_environment) {
    parent.document.addEventListener(
        "testReady",
        function () {
            const uis = document.querySelectorAll("ingame-ui");

            for (let ui of uis) {
                const event = new Event("panelActive");
                setTimeout(() => {
                    ui.dispatchEvent(event);
                }, 250);
            }
        },
        false
    );
}
