function dispatch_keyevent(event) {
    //catch backspace and prevent navigation
    if (e.keyCode == 8) {
        e.preventDefault();
        return;
    }

    if (event instanceof KeyboardEvent) {
        window.parent.window.document.dispatchEvent(new KeyboardEvent(event.type, event));
    }
}