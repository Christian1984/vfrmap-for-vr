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

function array_to_object(array, key) {
    const init = {};
    return array.reduce((acc, el) => {
        return {
            ...acc,
            [el[key]]: el.value,
        };
    }, init);
};

function store_data_set(key_value_array) {
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/data/", true);
    xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                var json = JSON.parse(xhr.responseText);
                console.log("POST to /data responded with:", json)
            }
        }
    };
    xhr.send(JSON.stringify({ data: key_value_array }));
}

function retrieve_data_set(key_array, callback) {
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/dataSet/?keys=" + JSON.stringify(key_array), true);
    xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                var json = JSON.parse(xhr.responseText);
                console.log("GET to /data responded with:", json)

                if (json != null && callback != null) {
                    const result_object = array_to_object(json.data, "key");
                    callback(result_object);
                }
            }
        }
    };

    xhr.send();
}

function store_data(key, string_data) {
    let xhr = new XMLHttpRequest();

    let payload = string_data;
    if (string_data.toString != null) {
        payload = payload.toString();
    }

    let body = {
        key: key,
        value: payload
    }
    xhr.open("POST", "/data/", true);
    xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                var json = JSON.parse(xhr.responseText);
                console.log("POST to /data responded with:", json)
            }
        }
    };
    xhr.send(JSON.stringify(body));
}

function retrieve_data(key, callback) {
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/data/?key=" + key, true);
    xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                var json = JSON.parse(xhr.responseText);
                console.log("GET to /data responded with:", json)

                if (json != null && callback != null) {
                    const result_object = array_to_object([json.data], "key");
                    callback(result_object);
                }
            }
        }
    };

    xhr.send();
}