function dispatch_keyevent(event) {
    if (event instanceof KeyboardEvent) {
        window.parent.window.document.dispatchEvent(new KeyboardEvent(event.type, event));
    }
}