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

class IngamePanelCustomPanel extends MyTemplateElement {
    constructor() {
        super(...arguments);

        this.panelActive = false;
        this.started = false;
        this.ingameUi = null;
        this.busy = false;
        this.debugEnabled = false;

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
                console.log('Identifier');
                console.log(self.instrumentIdentifier);
            });
            g_modDebugMgr.AddDebugButton("TemplateID", function() {
                console.log('TemplateID');
                console.log(self.templateID);
            });
            g_modDebugMgr.AddDebugButton("Source", function() {
                console.log('Source');
                console.log(window.document.documentElement.outerHTML);
            });
            g_modDebugMgr.AddDebugButton("close", function() {
                console.log('close');
                if (self.ingameUi) {
                    console.log('ingameUi');
                    self.ingameUi.closePanel();
                }
            });
            this.initialize();
        } else {
            Include.addScript("/JS/debug.js", function () {
                if (typeof g_modDebugMgr != "undefined") {
                    g_modDebugMgr.AddConsole(null);
                    g_modDebugMgr.AddDebugButton("Identifier", function() {
                        console.log('Identifier');
                        console.log(self.instrumentIdentifier);
                    });
                    g_modDebugMgr.AddDebugButton("TemplateID", function() {
                        console.log('TemplateID');
                        console.log(self.templateID);
                    });
                    g_modDebugMgr.AddDebugButton("Source", function() {
                        console.log('Source');
                        console.log(window.document.documentElement.outerHTML);
                    });
                    g_modDebugMgr.AddDebugButton("close", function() {
                        console.log('close');
                        if (self.ingameUi) {
                            console.log('ingameUi');
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

    connectedCallback() {
        super.connectedCallback();

        var self = this;

        setTimeout(() => {
            this.ingameUi = this.querySelector("ingame-ui");

            this.warningMessage = document.getElementById("WarningMessage");
    
            this.iframe_map = document.getElementById("iframe_map");
            this.iframe_charts = document.getElementById("iframe_charts");
            this.iframe_notepad = document.getElementById("iframe_notepad");
    
            this.switch_map = document.getElementById("switch_map");
            this.switch_charts = document.getElementById("switch_charts");
            this.switch_notepad = document.getElementById("switch_notepad");
    
            this.hide_all_iframes = function() {
                self.iframe_map.classList.add("hidden");
                self.iframe_charts.classList.add("hidden");
                self.iframe_notepad.classList.add("hidden");
            }
    
            this.unselect_all_buttons = function() {
                self.switch_map.classList.remove("active");
                self.switch_charts.classList.remove("active");
                self.switch_notepad.classList.remove("active");
            }
    
            this.switch_to_map = function() {
                self.hide_all_iframes();
                self.unselect_all_buttons();
                self.iframe_map.classList.remove("hidden");
                self.switch_map.classList.add("active");
            }
    
            this.switch_to_charts = function() {
                self.hide_all_iframes();
                self.unselect_all_buttons();
                self.iframe_charts.classList.remove("hidden");
                self.switch_charts.classList.add("active");
            }
    
            this.switch_to_notepad = function() {
                self.hide_all_iframes();
                self.unselect_all_buttons();
                self.iframe_notepad.classList.remove("hidden");
                self.switch_notepad.classList.add("active");
            }
    
            if (this.ingameUi) {
    
                this.ingameUi.addEventListener("panelActive", (e) => {
    
                    self.panelActive = true;
                    self.warningMessage.classList.add("show");
                    if (self.iframe_map) {
                        self.iframe_map.src = 'http://localhost:9000';
                    }
    
                    if (self.iframe_charts) {
                        self.iframe_charts.src = 'http://localhost:9000/premium/charts.html';
                    }
                    if (self.iframe_notepad) {
                        self.iframe_notepad.src = 'http://localhost:9000/premium/notepad.html';
                    }
    
                    if(self.switch_map) {
                        self.switch_map.addEventListener("click", () => {
                            self.switch_to_map();
                        });
                    }
    
                    if(self.switch_charts) {
                        self.switch_charts.addEventListener("click", () => {
                            self.switch_to_charts();
                        });
                    }
    
                    if(self.switch_notepad) {
                        self.switch_notepad.addEventListener("click", () => {
                            self.switch_to_notepad();
                        });
                    }
                });
    
                this.ingameUi.addEventListener("panelInactive", (e) => {
                    self.panelActive = false;
                    self.warningMessage.classList.remove("show");
    
                    if (self.iframe_map) {
                        self.iframe_map.src = '';
                    }
                    if (self.iframe_charts) {
                        self.iframe_charts.src = '';
                    }
                    if (self.iframe_notepad) {
                        self.iframe_notepad.src = '';
                    }
    
                    if(self.switch_map) {
                        self.switch_map.removeEventListener("click", this.switch_to_map);
                    }
    
                    if(self.switch_charts) {
                        self.switch_charts.removeEventListener("click", this.switch_charts);
                    }
    
                    if(self.switch_notepad) {
                        self.switch_charts.removeEventListener("click", this.switch_notepad);
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

window.customElements.define("ingamepanel-custom", IngamePanelCustomPanel);
myCheckAutoload();

if (parent && parent.window && parent.window.test_environment) {
    parent.document.addEventListener('testReady', function (e) { 
        console.log("iframe => testReady");
        uis = document.querySelectorAll("ingame-ui");

        for (let ui of uis) {
            console.log("ui", typeof(ui));

            const event = new Event('panelActive');
            setTimeout(() => {
                ui.dispatchEvent(event);
                console.log("Event dispatched...");
            }, 250);
        }
    }, false);
}


