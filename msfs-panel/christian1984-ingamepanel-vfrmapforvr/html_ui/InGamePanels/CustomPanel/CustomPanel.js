class IngamePanelCustomPanel extends TemplateElement {
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
        this.ingameUi = this.querySelector('ingame-ui');
        
        this.warningMessage = document.getElementById("WarningMessage");
        this.iframeElement = document.getElementById("CustomPanelIframe");

        this.m_MainDisplay = document.querySelector("#MainDisplay");
        this.m_MainDisplay.classList.add("hidden");

        this.m_Footer = document.querySelector("#Footer");
        this.m_Footer.classList.add("hidden");

        if (this.ingameUi) {
            this.ingameUi.addEventListener("panelActive", (e) => {
                console.log('panelActive');
                self.panelActive = true;
                self.warningMessage.classList.add("show");
                if (self.iframeElement) {
                    self.iframeElement.src = 'http://localhost:9000';
                }
            });
            this.ingameUi.addEventListener("panelInactive", (e) => {
                console.log('panelInactive');
                self.panelActive = false;
                self.warningMessage.classList.remove("show");
                if (self.iframeElement) {
                    self.iframeElement.src = '';
                }
            });
        }
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
checkAutoload();