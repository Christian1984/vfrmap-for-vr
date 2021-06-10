let MyTemplateElement;
let myCheckAutoload;

try {
    MyTemplateElement = TemplateElement;
}
catch (e) {
    if (parent && parent.window && parent.window.test_environment) {
        MyTemplateElement = class extends HTMLElement {
            constructor(args) {
                super(args);
            }
    
            connectedCallback() {
    
            }
    
            disconnectedCallback() {
                
            }
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
        this.started = false;
        this.ingameUi = null;
        this.busy = false;
        this.debugEnabled = false;

        this.collapsed = false;
        this.collapse_hotkey = -1;

        if (this.debugEnabled) {
            var self = this;
            setTimeout(() => {
                self.isDebugEnabled();
            }, 1000);
        } else {
            this.initialize();
        }
    }
    isDebugEnabled() {
        var self = this;
        if (typeof g_modDebugMgr != "undefined") {
            g_modDebugMgr.AddConsole(null);
            g_modDebugMgr.AddDebugButton("Identifier", function() {
                //console.log('Identifier');
                //console.log(self.instrumentIdentifier);
            });
            g_modDebugMgr.AddDebugButton("TemplateID", function() {
                //console.log('TemplateID');
                //console.log(self.templateID);
            });
            g_modDebugMgr.AddDebugButton("Source", function() {
                //console.log('Source');
                //console.log(window.document.documentElement.outerHTML);
            });
            g_modDebugMgr.AddDebugButton("close", function() {
                //console.log('close');
                if (self.ingameUi) {
                    //console.log('ingameUi');
                    self.ingameUi.closePanel();
                }
            });
            this.initialize();
        } else {
            Include.addScript("/JS/debug.js", function () {
                if (typeof g_modDebugMgr != "undefined") {
                    g_modDebugMgr.AddConsole(null);
                    g_modDebugMgr.AddDebugButton("Identifier", function() {
                        //console.log('Identifier');
                        //console.log(self.instrumentIdentifier);
                    });
                    g_modDebugMgr.AddDebugButton("TemplateID", function() {
                        //console.log('TemplateID');
                        //console.log(self.templateID);
                    });
                    g_modDebugMgr.AddDebugButton("Source", function() {
                        //console.log('Source');
                        //console.log(window.document.documentElement.outerHTML);
                    });
                    g_modDebugMgr.AddDebugButton("close", function() {
                        //console.log('close');
                        if (self.ingameUi) {
                            //console.log('ingameUi');
                            self.ingameUi.closePanel();
                        }
                    });
                    self.initialize();
                } else {
                    setTimeout(() => {
                        self.isDebugEnabled();
                    }, 2000);
                }
            });
        }
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
    
                        if (data.data.type == "keydown" && data.data.keyCode == self.collapse_hotkey && data.data.altKey) {
                            self.toggle_collapse();
                        }

                        break;

                    case "HotkeyConfiguration":
                        self.collapse_hotkey = data.data.keyCode;
                        break;

                    case "SetBrighness":
                        self.set_brightness(data.data.brightness);
                        break;
                        
                    case "SetRedLight":
                        self.set_red_light(data.data.red);
                        break;
                }

                if (data.type == "KeyboardEvent" && data.data != null) {
                }
            }
            catch (e) {
                /* ignore silently */
            }
        });

        window.addEventListener("keydown", (e) => {
            if (self.collapse_hotkey == -1) return;

            if (e.keyCode == self.collapse_hotkey && e.altKey) {
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

    initialize() {
        if (this.started) {
            return;
        }
        this.started = true;
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
    parent.document.addEventListener('testReady', function (e) { 
        //console.log("iframe => testReady");
        uis = document.querySelectorAll("ingame-ui");

        for (let ui of uis) {
            //console.log("ui", typeof(ui));

            const event = new Event('panelActive');
            setTimeout(() => {
                ui.dispatchEvent(event);
                //console.log("Event dispatched...");
            }, 250);
        }
    }, false);
}


