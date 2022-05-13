import "./FSKneeboardPanel.scss"

let MyTemplateElement;
let myCheckAutoload;

try {
    MyTemplateElement = TemplateElement;
}
catch (e) {
    if (parent && parent.window && parent.window.test_environment) {
        MyTemplateElement = class extends HTMLElement {
            constructor(args) { super(args); }
            connectedCallback() {}
            disconnectedCallback() {}
        }
    }
}

try {
    myCheckAutoload = checkAutoload;
}
catch (e) {
    if (parent && parent.window && parent.window.test_environment) {
        myCheckAutoload = function() {}
    }
}

class IngamePanelFSKneeboardPanel extends MyTemplateElement {
    constructor() {
        super(...arguments);

        this.panelActive = false;
        this.ingameUi = null;

        this.collapsed = false;
        this.collapse_hotkey = -1;
        this.collapse_altKey = false;
        this.collapse_ctrlKey = false;
        this.collapse_shiftKey = false;
    }

    collapse(collapsed) {
        if (collapsed) {
            this.classList.add("collapsed");
        }
        else {
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
        overlay_red.style.opacity = - 3 / 1000 * brightness + 1;
    }

    set_red_light(red) {
        const panel = document.querySelector("#FSKneeboardPanel");

        if (panel) {
            if (red) {
                panel.classList.add("red");
            }
            else {
                panel.classList.remove("red");
            }
        }
    }

    connectedCallback() {
        super.connectedCallback();

        var self = this;

        window.addEventListener("message", (e) => {
            try {
                const data = JSON.parse(e.data);

                switch (data.type) {
                    case "KeyboardEvent":
                        if (self.collapse_hotkey == -1) return;

                        if (data.data.type == "keydown"
                            && data.data.keyCode == self.collapse_hotkey
                            && data.data.altKey == self.collapse_altKey
                            && data.data.ctrlKey == self.collapse_ctrlKey
                            && data.data.shiftKey == self.collapse_shiftKey) {
                            self.toggle_collapse();
                        }

                        break;

                    case "HotkeyConfiguration":
                        self.collapse_hotkey = data.data.keyCode;
                        self.collapse_altKey = data.data.altKey;
                        self.collapse_ctrlKey = data.data.ctrlKey;
                        self.collapse_shiftKey = data.data.shiftKey;
                        break;

                    case "SetBrighness":
                        self.set_brightness(data.data.brightness);
                        break;
                        
                    case "SetRedLight":
                        self.set_red_light(data.data.red);
                        break;
                }
            }
            catch (e) {
                /* ignore silently */
            }
        });

        window.addEventListener("keydown", (e) => {
            if (self.collapse_hotkey == -1) return;

            if (e.keyCode == self.collapse_hotkey
                && e.altKey == self.collapse_altKey
                && e.ctrlKey == self.collapse_ctrlKey
                && e.shiftKey == self.collapse_shiftKey) {
                self.toggle_collapse();
            }
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
                        self.content_iframe.src = 'http://localhost:9000/index.html';
                    }

                    self.collapse(false);
                });
    
                this.ingameUi.addEventListener("panelInactive", (e) => {
                    self.panelActive = false;
                    self.warning_message.classList.remove("show");
    
                    if (self.content_iframe) {
                        self.content_iframe.src = '';
                    }
                });
            }
        } , 0);
    }

    disconnectedCallback() {
        super.disconnectedCallback();
    }

    updateImage() {

    }
}

window.customElements.define("ingamepanel-custom", IngamePanelFSKneeboardPanel);
myCheckAutoload();

if (parent && parent.window && parent.window.test_environment) {
    parent.document.addEventListener('testReady', function () {
        uis = document.querySelectorAll("ingame-ui");

        for (let ui of uis) {
            const event = new Event('panelActive');
            setTimeout(() => {
                ui.dispatchEvent(event);
            }, 250);
        }
    }, false);
}


