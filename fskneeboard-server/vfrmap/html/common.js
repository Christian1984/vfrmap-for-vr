const sender_id = Math.floor(Math.random() * Number.MAX_VALUE).toString();

function dispatch_keyevent(event) {
    //catch backspace and prevent navigation
    if (event.keyCode == 8) {
        event.preventDefault();
        return;
    }

    if (event instanceof KeyboardEvent) {
        window.parent.window.document.dispatchEvent(new KeyboardEvent(event.type, event));
    }
}

function array_to_object(arr, key) {
    const init = {};
    return arr.reduce((acc, el) => {
        return {
            ...acc,
            [el[key]]: el.value,
        };
    }, init);
};

function store_data_set(key_string_value_array, remote = true) {
    if (remote) {
        let xhr = new XMLHttpRequest();
        xhr.open("POST", "/dataSet/", true);
        xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
        xhr.onreadystatechange = function () {
            if (xhr.readyState === 4) {
                if (xhr.status === 200) {
                    //console.log("POST to /data responded with:", JSON.parse(xhr.responseText))
                }
            }
        };
    
        for (kv of key_string_value_array) {
            if (kv.value.toString != null) {
                kv.value = kv.value.toString();
            }
        }
    
        xhr.send(JSON.stringify({ data: key_string_value_array, sender: sender_id }));
    }
    else {
        for (const el of key_string_value_array) {
            localStorage.setItem(el.key, el.value);
        }
    }
}

function retrieve_data_set(key_array, callback, remote = true) {
    if (remote) {
        let xhr = new XMLHttpRequest();
        xhr.open("GET", "/dataSet/?keys=" + JSON.stringify(key_array), true);
        xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
        xhr.onreadystatechange = function () {
            if (xhr.readyState === 4) {
                if (xhr.status === 200) {
                    var json = JSON.parse(xhr.responseText);
                    //console.log("GET to /data responded with:", json)
    
                    if (json != null && callback != null) {
                        const result_object = array_to_object(json.data, "key");
                        callback(result_object);
                    }
                }
            }
        };
    
        xhr.send();
    }
    else {
        const result_object = {};

        for (const key of key_array) {
            result_object[key] = localStorage.getItem(key);
        }

        callback(result_object);
    }
}

function store_data(key, string_data, remote = true) {
    if (remote) {
        let xhr = new XMLHttpRequest();
    
        let payload = string_data;
        if (string_data.toString != null) {
            payload = payload.toString();
        }
    
        let body = {
            key: key,
            value: payload,
            sender: sender_id
        }
        xhr.open("POST", "/data/", true);
        xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
        xhr.onreadystatechange = function () {
            if (xhr.readyState === 4) {
                if (xhr.status === 200) {
                    //console.log("POST to /data responded with:", JSON.parse(xhr.responseText))
                }
            }
        };
        xhr.send(JSON.stringify(body));
    }
    else {
        localStorage.setItem(key, string_data);
    }
}

function retrieve_data(key, callback, remote = true) {
    if (remote) {
        let xhr = new XMLHttpRequest();
        xhr.open("GET", "/data/?key=" + key, true);
        xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
        xhr.onreadystatechange = function () {
            if (xhr.readyState === 4) {
                if (xhr.status === 200) {
                    var json = JSON.parse(xhr.responseText);
                    //console.log("GET to /data responded with:", json)
    
                    if (json != null && callback != null) {
                        const result_object = array_to_object([json], "key");
                        callback(result_object);
                    }
                }
            }
        };
    
        xhr.send();
    }
    else {
        const result_object = {};
        result_object[key] = localStorage.getItem(key);

        callback(result_object);
    }
}

function hide_confirm_dialog(wrapper_selector, hide) {
    const confirm_dialog_wrapper = document.querySelector(wrapper_selector);
    if (!confirm_dialog_wrapper) return;

    if (hide) {
        confirm_dialog_wrapper.classList.add("hidden");
    }
    else {
        confirm_dialog_wrapper.classList.remove("hidden");
    }
}