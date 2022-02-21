// depends on sender.js
const keycode_map = {
    8: "Backspace",
    13: "Enter",
    17: "Control",
    32: " ",
    33: "!",
    34: "\"",
    35: "#",
    36: "$",
    37: "%",
    38: "&",
    39: "'",
    40: "(",
    41: ")",
    42: "*",
    43: "+",
    44: ",",
    45: "-",
    46: ".",
    47: "/",
    48: "0",
    49: "1",
    50: "2",
    51: "3",
    52: "4",
    53: "5",
    54: "6",
    55: "7",
    56: "8",
    57: "9",
    58: ":",
    59: ";",
    60: "<",
    61: "=",
    62: ">",
    63: "?",
    64: "@",
    65: "A",
    66: "B",
    67: "C",
    68: "D",
    69: "E",
    70: "F",
    71: "G",
    72: "H",
    73: "I",
    74: "J",
    75: "K",
    76: "L",
    77: "M",
    78: "N",
    79: "O",
    80: "P",
    81: "Q",
    82: "R",
    83: "S",
    84: "T",
    85: "U",
    86: "V",
    87: "W",
    88: "X",
    89: "Y",
    90: "Z",
    91: "[",
    92: "\\",
    93: "]",
    94: "^",
    95: "_",
    96: "`",
    97: "a",
    98: "b",
    99: "c",
    100: "d",
    101: "e",
    102: "f",
    103: "g",
    104: "h",
    105: "i",
    106: "j",
    107: "k",
    108: "l",
    109: "m",
    110: "n",
    111: "o",
    112: "p",
    113: "q",
    114: "r",
    115: "s",
    116: "t",
    117: "u",
    118: "v",
    119: "w",
    120: "x",
    121: "y",
    122: "z",
    123: "{",
    124: "|",
    125: "}",
    126: "~",
    167: "§",
    180: "´"
};

function get_key_from_keycode(keycode) {
    let res = "";

    try {
        res = keycode_map[keycode];
    }
    catch (e) {
        Logger.logError("Key for keycode " + keycode + " not found, details: " + e.message);
    }

    return res != undefined ? res : "";
}

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
        acc[el[key]] = el.value
        return acc;
    }, init);
}

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
    
        xhr.send(JSON.stringify({ data: key_string_value_array, sender: SENDER_ID }));
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
            sender: SENDER_ID
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

document.addEventListener("DOMContentLoaded", function() {
    Logger.logDebug("common.js => DOMContentLoaded fired!");
});