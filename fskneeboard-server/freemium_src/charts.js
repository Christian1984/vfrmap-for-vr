let loaded = false;

window.addEventListener("message", (m) => {
    iframe = document.querySelector("iframe");
    if (iframe) {
        if (m.data == "load" && !loaded) {
            iframe.src = "https://fskneeboard.com/charts-ingame/";
            loaded = true;
        }
    }
});