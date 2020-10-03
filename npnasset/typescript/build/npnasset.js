"use strict";
var dom;
(function (dom) {
    function initDom(t, color) {
        try {
            style.themeLinks(color);
            style.setTheme(t);
        }
        catch (e) {
            console.warn("error setting style", e);
        }
        try {
            modal.wire();
        }
        catch (e) {
            console.warn("error wiring modals", e);
        }
        try {
            drop.wire();
        }
        catch (e) {
            console.warn("error wiring drops", e);
        }
        try {
            tags.wire();
        }
        catch (e) {
            console.warn("error wiring tag editors", e);
        }
        try {
            flash.wire();
        }
        catch (e) {
            console.warn("error wiring tag editors", e);
        }
    }
    dom.initDom = initDom;
    function els(selector, context) {
        var result;
        if (context) {
            result = context.querySelectorAll(selector);
        }
        else {
            result = document.querySelectorAll(selector);
        }
        var ret = [];
        result.forEach(function (v) {
            ret.push(v);
        });
        return ret;
    }
    dom.els = els;
    function opt(selector, context) {
        var e = els(selector, context);
        switch (e.length) {
            case 0:
                return undefined;
            case 1:
                return e[0];
            default:
                console.warn("found [" + e.length + "] elements with selector [" + selector + "], wanted zero or one");
        }
    }
    dom.opt = opt;
    function req(selector, context) {
        var res = opt(selector, context);
        if (!res) {
            console.warn("no element found for selector [" + selector + "]");
        }
        return res;
    }
    dom.req = req;
    function setHTML(el, html) {
        if (typeof el === "string") {
            el = req(el);
        }
        el.innerHTML = html;
        return el;
    }
    dom.setHTML = setHTML;
    function setDisplay(el, condition, v) {
        if (v === void 0) { v = "block"; }
        if (typeof el === "string") {
            el = req(el);
        }
        el.style.display = condition ? v : "none";
        return el;
    }
    dom.setDisplay = setDisplay;
    function setContent(el, e) {
        if (typeof el === "string") {
            el = req(el);
        }
        dom.clear(el);
        if (Array.isArray(e)) {
            e.forEach(function (x) { return el.appendChild(x); });
        }
        else {
            el.appendChild(e);
        }
        return el;
    }
    dom.setContent = setContent;
    function setText(el, text) {
        if (typeof el === "string") {
            el = req(el);
        }
        el.innerText = text;
        return el;
    }
    dom.setText = setText;
    function switchElements(el, tgt) {
        setDisplay(el, false);
        setDisplay(tgt, true);
        return false;
    }
    dom.switchElements = switchElements;
    function clear(el) {
        return setHTML(el, "");
    }
    dom.clear = clear;
})(dom || (dom = {}));
var dom;
(function (dom) {
    function setValue(el, text) {
        if (typeof el === "string") {
            el = dom.req(el);
        }
        el.value = text;
        return el;
    }
    dom.setValue = setValue;
    function wireTextarea(text) {
        function resize() {
            text.style.height = "auto";
            text.style.height = (text.scrollHeight < 64 ? 64 : text.scrollHeight + 6) + "px";
        }
        function delayedResize() {
            window.setTimeout(resize, 0);
        }
        var x = text.dataset["autoresize"];
        if (!x) {
            text.dataset["autoresize"] = "true";
            text.addEventListener("change", resize, false);
            text.addEventListener("cut", delayedResize, false);
            text.addEventListener("paste", delayedResize, false);
            text.addEventListener("drop", delayedResize, false);
            text.addEventListener("keydown", delayedResize, false);
            text.focus();
            text.select();
        }
        resize();
    }
    dom.wireTextarea = wireTextarea;
    function setOptions(el, categories) {
        if (typeof el === "string") {
            el = dom.req(el);
        }
        dom.clear(el);
        categories.forEach(function (c) {
            var opt = document.createElement("option");
            opt.value = c;
            dom.setText(opt, c);
            el.appendChild(opt);
        });
    }
    dom.setOptions = setOptions;
    function setSelectOption(el, o) {
        if (typeof el === "string") {
            el = dom.req(el);
        }
        for (var i = 0; i < el.children.length; i++) {
            var e = el.children.item(i);
            e.selected = e.value === o;
        }
    }
    dom.setSelectOption = setSelectOption;
    function insertAtCaret(e, text) {
        if (e.selectionStart || e.selectionStart === 0) {
            var startPos = e.selectionStart;
            var endPos = e.selectionEnd;
            e.value = e.value.substring(0, startPos) + text + e.value.substring(endPos, e.value.length);
            e.selectionStart = startPos + text.length;
            e.selectionEnd = startPos + text.length;
        }
        else {
            e.value += text;
        }
    }
    dom.insertAtCaret = insertAtCaret;
})(dom || (dom = {}));
// noinspection JSUnusedGlobalSymbols
function JSX(tag, attrs) {
    var e = document.createElement(tag);
    for (var name_1 in attrs) {
        if (name_1 && attrs.hasOwnProperty(name_1)) {
            var v = attrs[name_1];
            if (name_1 === "dangerouslySetInnerHTML") {
                dom.setHTML(e, v["__html"]);
            }
            else if (v === true) {
                e.setAttribute(name_1, name_1);
            }
            else if (v !== false && v !== null && v !== undefined) {
                e.setAttribute(name_1, v.toString());
            }
        }
    }
    var _loop_1 = function (i) {
        var child = arguments_1[i];
        if (Array.isArray(child)) {
            child.forEach(function (c) {
                if (child === undefined || child === null) {
                    throw "child array for tag [" + tag + "] is " + child + "\n" + e.outerHTML;
                }
                if (c === undefined || c === null) {
                    throw "child for tag [" + tag + "] is " + c + "\n" + e.outerHTML;
                }
                if (typeof c === "string") {
                    c = document.createTextNode(c);
                }
                e.appendChild(c);
            });
        }
        else if (child === undefined || child === null) {
            throw "child for tag [" + tag + "] is " + child + "\n" + e.outerHTML;
            // debugger;
            // child = document.createTextNode("NULL!");
        }
        else {
            if (!child.nodeType) {
                child = document.createTextNode(child.toString());
            }
            e.appendChild(child);
        }
    };
    var arguments_1 = arguments;
    for (var i = 2; i < arguments.length; i++) {
        _loop_1(i);
    }
    return e;
}
var style;
(function (style) {
    function setTheme(theme) {
        wireEmoji(theme);
        switch (theme) {
            case "auto":
                var t = "light";
                if (window.matchMedia && window.matchMedia("(prefers-color-scheme: dark)").matches) {
                    t = "dark";
                }
                setTheme(t);
                fetch("/profile/theme/" + t).then(function (r) { return r.text(); }).then(function () {
                    // console.log(`Set theme to [${t}]`);
                });
                break;
            case "light":
                document.documentElement.classList.remove("uk-light");
                document.body.classList.remove("uk-light");
                document.documentElement.classList.add("uk-dark");
                document.body.classList.add("uk-dark");
                break;
            case "dark":
                document.documentElement.classList.add("uk-light");
                document.body.classList.add("uk-light");
                document.documentElement.classList.remove("uk-dark");
                document.body.classList.remove("uk-dark");
                break;
            default:
                console.warn("invalid theme");
                break;
        }
    }
    style.setTheme = setTheme;
    style.linkColor = "";
    function themeLinks(color) {
        style.linkColor = color + "-fg";
        dom.els(".theme").forEach(function (el) {
            el.classList.add(style.linkColor);
        });
    }
    style.themeLinks = themeLinks;
    function wireEmoji(t) {
        if (typeof EmojiButton === "undefined") {
            dom.els(".picker-toggle").forEach(function (el) { return dom.setDisplay(el, false); });
            return;
        }
        var opts = { position: "bottom-end", theme: t, zIndex: 1021 };
        dom.els(".textarea-emoji").forEach(function (el) {
            var toggle = dom.req(".picker-toggle", el);
            toggle.addEventListener("click", function () {
                var textarea = dom.req(".uk-textarea", el);
                var picker = new EmojiButton(opts);
                picker.on("emoji", function (emoji) {
                    drop.onEmojiPicked();
                    dom.insertAtCaret(textarea, emoji);
                });
                picker.togglePicker(toggle);
            }, false);
        });
    }
})(style || (style = {}));
var drop;
(function (drop) {
    function wire() {
        dom.els(".drop").forEach(function (el) {
            el.addEventListener("show", onDropOpen);
            el.addEventListener("beforehide", onDropBeforeHide);
            el.addEventListener("hide", onDropHide);
        });
    }
    drop.wire = wire;
    function onDropOpen(e) {
        if (!e.target) {
            return;
        }
        var el = e.target;
        var key = el.dataset["key"] || "";
        var t = el.dataset["t"] || "";
        var f = events.getOpenEvent(key);
        if (f) {
            f(t);
        }
        else {
            console.warn("no drop open handler registered for [" + key + "]");
        }
    }
    function onDropHide(e) {
        if (!e.target) {
            return;
        }
        var el = e.target;
        if (el.classList.contains("uk-open")) {
            var key = el.dataset["key"] || "";
            var t = el.dataset["t"] || "";
            var f = events.getCloseEvent(key);
            if (f) {
                f(t);
            }
        }
    }
    var emojiPicked = false;
    function onEmojiPicked() {
        emojiPicked = true;
        setTimeout(function () { return (emojiPicked = false); }, 200);
    }
    drop.onEmojiPicked = onEmojiPicked;
    function onDropBeforeHide(e) {
        if (emojiPicked) {
            e.preventDefault();
        }
    }
})(drop || (drop = {}));
var events;
(function (events) {
    var openEvents;
    var closeEvents;
    function register(key, o, c) {
        if (!openEvents) {
            openEvents = new map.Map();
        }
        if (!closeEvents) {
            closeEvents = new map.Map();
        }
        if (!o) {
            o = function () { };
        }
        openEvents.set(key, o);
        if (c) {
            closeEvents.set(key, c);
        }
    }
    events.register = register;
    function getOpenEvent(key) {
        return openEvents.get(key);
    }
    events.getOpenEvent = getOpenEvent;
    function getCloseEvent(key) {
        return closeEvents.get(key);
    }
    events.getCloseEvent = getCloseEvent;
})(events || (events = {}));
var flash;
(function (flash) {
    function wire() {
        setTimeout(fadeOut, 4000);
    }
    flash.wire = wire;
    function fadeOut() {
        var matched = false;
        dom.els(".alert-top").forEach(function (el) {
            matched = true;
            el.classList.add("uk-animation-fade", "uk-animation-reverse");
        });
        if (matched) {
            setTimeout(remove, 1000);
        }
    }
    function remove() {
        dom.els(".alert-top").forEach(function (el) {
            el.remove();
        });
    }
})(flash || (flash = {}));
var modal;
(function (modal) {
    var activeParam;
    function wire() {
        dom.els(".modal").forEach(function (el) {
            el.addEventListener("show", onModalOpen);
            el.addEventListener("hide", onModalHide);
        });
    }
    modal.wire = wire;
    function open(key, param) {
        activeParam = param;
        var m = notify.modal("#modal-" + key);
        m.show();
        return false;
    }
    modal.open = open;
    function openSoon(key) {
        setTimeout(function () { return open(key); }, 0);
    }
    modal.openSoon = openSoon;
    function hide(key) {
        var m = notify.modal("#modal-" + key);
        var el = m.$el;
        if (el.classList.contains("uk-open")) {
            m.hide();
        }
    }
    modal.hide = hide;
    function onModalOpen(e) {
        if (!e.target) {
            return;
        }
        var el = e.target;
        if (el.id.indexOf("modal") !== 0) {
            return;
        }
        var key = el.id.substr("modal-".length);
        var f = events.getOpenEvent(key);
        if (f) {
            f(activeParam);
        }
        else {
            console.warn("no modal open handler registered for [" + key + "]");
        }
        activeParam = undefined;
    }
    function onModalHide(e) {
        if (!e.target) {
            return;
        }
        var el = e.target;
        if (el.classList.contains("uk-open")) {
            var key = el.id.substr("modal-".length);
            var f = events.getCloseEvent(key);
            if (f) {
                f(activeParam);
            }
            activeParam = undefined;
        }
    }
})(modal || (modal = {}));
var tags;
(function (tags) {
    function wire() {
        dom.els(".tag-editor").forEach(function (el) {
            el.addEventListener("moved", onTagEditorUpdate);
            el.addEventListener("added", onTagEditorUpdate);
            el.addEventListener("removed", onTagEditorUpdate);
        });
    }
    tags.wire = wire;
    function removeTag(el) {
        var itemEl = el.parentElement;
        var editorEl = itemEl.parentElement;
        itemEl.remove();
        updateEditor(editorEl);
    }
    tags.removeTag = removeTag;
    function addTag(el) {
        var editorEl = el.parentElement;
        if (!editorEl) {
            return;
        }
        var itemEl = tags.renderItem();
        editorEl.insertBefore(itemEl, dom.req(".add-item", editorEl));
        editTag(itemEl);
    }
    tags.addTag = addTag;
    function editTag(el) {
        var valueEl = dom.req(".value", el);
        var editorEl = dom.req(".editor", el);
        dom.setDisplay(valueEl, false);
        dom.setDisplay(editorEl, true);
        var input = tags.renderInput(valueEl.innerText);
        input.onblur = function () {
            valueEl.innerText = input.value;
            dom.setDisplay(valueEl, true);
            dom.setDisplay(editorEl, false);
            updateEditor(el.parentElement);
        };
        input.onkeypress = function (e) {
            if (e.key === "Enter") {
                input.blur();
                return false;
            }
            return true;
        };
        dom.setContent(editorEl, input);
        input.focus();
    }
    tags.editTag = editTag;
    function onTagEditorUpdate(e) {
        if (!e.target) {
            console.warn("no event target");
            return;
        }
        var el = e.target;
        updateEditor(el);
    }
    function updateEditor(el) {
        var key = el.dataset["key"] || "";
        var f = events.getOpenEvent(key);
        if (f) {
            f();
        }
        else {
            console.warn("no tag open handler registered for [" + key + "]");
        }
        var ret = dom.els(".item", el).map(function (e) { return e.innerText; });
        dom.setValue("#model-" + key + "-input", ret.join(","));
    }
})(tags || (tags = {}));
var tags;
(function (tags) {
    function renderInput(v) {
        return JSX("input", { type: "text", class: "uk-input", value: v });
    }
    tags.renderInput = renderInput;
    function renderItem() {
        return JSX("span", { class: "item" },
            JSX("span", { class: "value", onclick: "tags.editTag(this.parentElement);" }),
            JSX("span", { class: "editor" }),
            JSX("span", { class: "close", "data-uk-icon": "icon: close; ratio: 0.6;", onclick: "tags.removeTag(this);" }));
    }
    tags.renderItem = renderItem;
    function renderTagsView(a) {
        return JSX("div", { class: "tag-view" },
            a.map(function (s) { return JSX("span", { class: "item" }, s); }),
            JSX("div", { class: "clear" }));
    }
    tags.renderTagsView = renderTagsView;
})(tags || (tags = {}));
var socket;
(function (socket) {
    function initBypass() {
        socket.bypass = true;
        socket.connected = true;
        nav.enabled = false;
    }
    socket.initBypass = initBypass;
    function bypassSend(msg) {
        if (socket.debug) {
            console.debug("out", msg);
        }
        if (npn_handler) {
            npn_handler(JSON.stringify(msg, null, 2));
        }
        else {
            console.warn("no bypass handler configured");
        }
    }
    socket.bypassSend = bypassSend;
})(socket || (socket = {}));
var socket;
(function (socket) {
    socket.debug = true;
    socket.appUnloading = false;
    socket.currentService = "";
    socket.currentID = "";
    socket.bypass = false;
    function setAppUnloading() {
        socket.appUnloading = true;
    }
    socket.setAppUnloading = setAppUnloading;
    function send(msg) {
        if (socket.bypass) {
            socket.bypassSend(msg);
        }
        else {
            socket.socketSend(msg);
        }
    }
    socket.send = send;
})(socket || (socket = {}));
var socket;
(function (socket) {
    var sock;
    socket.connected = false;
    var pauseSeconds = 0;
    var pendingMessages = [];
    var onMessage;
    var onError;
    function init(recv, err) {
        onMessage = recv;
        onError = err;
    }
    socket.init = init;
    function socketUrl() {
        var l = document.location;
        var protocol = "ws";
        if (l.protocol === "https:") {
            protocol = "wss";
        }
        return protocol + ("://" + l.host + "/s");
    }
    function initSocket() {
        sock = new WebSocket(socketUrl());
        sock.onopen = onSocketOpen;
        sock.onmessage = function (event) { return onMessage(json.parse(event.data)); };
        sock.onerror = function (event) { return onError("socket", event.type); };
        sock.onclose = onSocketClose;
    }
    socket.initSocket = initSocket;
    function socketConnect(svc, id, useBypass) {
        socket.currentService = svc;
        socket.currentID = id;
        socket.connectTime = Date.now();
        if (!onMessage) {
            throw "onMessage not initialized";
        }
        if (useBypass) {
            socket.initBypass();
        }
        else {
            initSocket();
        }
    }
    socket.socketConnect = socketConnect;
    function onSocketOpen() {
        log.info("socket connected");
        socket.connected = true;
        pauseSeconds = 1;
        pendingMessages.forEach(socket.send);
        pendingMessages = [];
    }
    function onSocketClose() {
        function disconnect() {
            socket.connected = false;
            var elapsed = Date.now() - socket.connectTime;
            if (elapsed < 2000) {
                pauseSeconds = pauseSeconds * 2;
                if (socket.debug) {
                    console.debug("socket closed immediately, reconnecting in " + pauseSeconds + " seconds");
                }
                setTimeout(function () {
                    socketConnect(socket.currentService, socket.currentID);
                }, pauseSeconds * 1000);
            }
            else {
                log.info("socket closed after [" + elapsed + "ms]");
                socketConnect(socket.currentService, socket.currentID);
            }
        }
        if (!socket.appUnloading) {
            disconnect();
        }
    }
    function socketSend(msg) {
        if (socket.debug) {
            console.debug("out", msg);
        }
        if (socket.connected) {
            var m = json.str(msg);
            sock.send(m);
        }
        else {
            pendingMessages.push(msg);
        }
    }
    socket.socketSend = socketSend;
})(socket || (socket = {}));
var profile;
(function (profile) {
    // noinspection JSUnusedGlobalSymbols
    function setNavColor(el, c) {
        dom.setValue("#nav-color", c);
        var nb = dom.req("#navbar");
        nb.className = c + "-bg uk-navbar-container uk-navbar";
        var colors = document.querySelectorAll(".nav_swatch");
        colors.forEach(function (i) {
            i.classList.remove("active");
        });
        el.classList.add("active");
    }
    profile.setNavColor = setNavColor;
    // noinspection JSUnusedGlobalSymbols
    function setLinkColor(el, c) {
        dom.setValue("#link-color", c);
        var links = dom.els(".profile-link");
        links.forEach(function (l) {
            l.classList.forEach(function (x) {
                if (x.indexOf("-fg") > -1) {
                    l.classList.remove(x);
                }
                l.classList.add(c + "-fg");
            });
        });
        var colors = document.querySelectorAll(".link_swatch");
        colors.forEach(function (i) {
            i.classList.remove("active");
        });
        el.classList.add("active");
    }
    profile.setLinkColor = setLinkColor;
    function setPicture(p) {
        dom.setValue("#self-picture-input", p);
        return false;
    }
    profile.setPicture = setPicture;
})(profile || (profile = {}));
var arr;
(function (arr) {
    function find(a, predicate) {
        var len = a.length >>> 0;
        var k = 0;
        while (k < len) {
            var kValue = a[k];
            if (predicate(k, kValue)) {
                return kValue;
            }
            k++;
        }
        return undefined;
    }
    arr.find = find;
})(arr || (arr = {}));
var date;
(function (date) {
    function dateToYMD(dt) {
        var d = dt.getDate();
        var m = dt.getMonth() + 1;
        var y = dt.getFullYear();
        return y + "-" + (m <= 9 ? "0" + m : m) + "-" + (d <= 9 ? "0" + d : d);
    }
    date.dateToYMD = dateToYMD;
    function dateFromYMD(s) {
        var d = new Date(s);
        return new Date(d.getTime() + d.getTimezoneOffset() * 60000);
    }
    date.dateFromYMD = dateFromYMD;
    function dow(i) {
        switch (i) {
            case 0:
                return "Sun";
            case 1:
                return "Mon";
            case 2:
                return "Tue";
            case 3:
                return "Wed";
            case 4:
                return "Thu";
            case 5:
                return "Fri";
            case 6:
                return "Sat";
            default:
                return "???";
        }
    }
    date.dow = dow;
    function toDateString(d) {
        return d.toLocaleDateString();
    }
    date.toDateString = toDateString;
    function toTimeString(d) {
        return d.toLocaleTimeString().slice(0, 8);
    }
    date.toTimeString = toTimeString;
    function toDateTimeString(d) {
        return toDateString(d) + " " + toTimeString(d);
    }
    date.toDateTimeString = toDateTimeString;
    var tzOffset = new Date().getTimezoneOffset() * 60000;
    function utcDate(s) {
        return new Date(Date.parse(s) + tzOffset);
    }
    date.utcDate = utcDate;
})(date || (date = {}));
var group;
(function (group_1) {
    var Group = /** @class */ (function () {
        function Group(key) {
            this.members = [];
            this.key = key;
        }
        return Group;
    }());
    group_1.Group = Group;
    var GroupSet = /** @class */ (function () {
        function GroupSet() {
            this.groups = [];
        }
        GroupSet.prototype.findOrInsert = function (key) {
            var ret = arr.find(this.groups, function (_, x) { return x.key === key; });
            if (ret) {
                return ret;
            }
            var n = new Group(key);
            this.groups.push(n);
            return n;
        };
        return GroupSet;
    }());
    group_1.GroupSet = GroupSet;
    function groupBy(list, func) {
        var res = new GroupSet();
        if (list) {
            list.forEach(function (o) {
                var group = res.findOrInsert(func(o));
                group.members.push(o);
            });
        }
        return res;
    }
    group_1.groupBy = groupBy;
    function findGroup(groups, key) {
        for (var _i = 0, groups_1 = groups; _i < groups_1.length; _i++) {
            var g = groups_1[_i];
            if (g.key === key) {
                return g.members;
            }
        }
        return [];
    }
    group_1.findGroup = findGroup;
    function flatten(a) {
        var ret = [];
        a.forEach(function (v) { return ret.push.apply(ret, v); });
        return ret;
    }
    group_1.flatten = flatten;
    function sort(a, matchFn) {
        if (!a) {
            return [];
        }
        a.sort(function (l, r) {
            var lv = matchFn(l);
            var rv = matchFn(r);
            if (lv > rv) {
                return 1;
            }
            if (lv < rv) {
                return -1;
            }
            return 0;
        });
        return a;
    }
    group_1.sort = sort;
    function update(a, v, matchFn) {
        if (!a) {
            return [v];
        }
        var matched = false;
        var key = matchFn(v);
        for (var idx in a) {
            var c = a[idx];
            if (matchFn(c) == key) {
                matched = true;
                a[idx] = v;
            }
        }
        if (!matched) {
            a.push(v);
        }
        return a;
    }
    group_1.update = update;
    function updateAndSort(a, v, matchFn) {
        return sort(update(a, v, matchFn), matchFn);
    }
    group_1.updateAndSort = updateAndSort;
})(group || (group = {}));
var json;
(function (json) {
    function str(x) {
        if (x === undefined) {
            return "null";
        }
        return JSON.stringify(x, null, 2);
    }
    json.str = str;
    function parse(s) {
        return JSON.parse(s);
    }
    json.parse = parse;
})(json || (json = {}));
var log;
(function (log) {
    var started = 0;
    var content;
    var list;
    function init() {
        started = Date.now();
        l("debug", "npn started");
    }
    log.init = init;
    function info(msg) {
        l("info", msg);
    }
    log.info = info;
    function l(level, msg) {
        if (started === 0) {
            console.warn("call `log.init()` before attempting to log");
            return;
        }
        var n = Date.now() - started;
        var el = JSX("li", { class: color(level) },
            JSX("div", { class: "right" },
                n,
                "ms"),
            msg);
        if (!list) {
            list = dom.opt("#log-list");
            if (!list) {
                console.warn(level + ": " + msg);
                return;
            }
        }
        list.appendChild(el);
        if (!content) {
            content = dom.req("#log-content");
        }
        content.scrollTo(0, content.scrollHeight);
    }
    log.l = l;
    function toggle() {
        var wsc = dom.req("#workspace-content");
        var lp = dom.req("#log-container");
        var curr = (lp.style.display !== "") && (lp.style.display !== "none");
        if (curr) {
            wsc.classList.remove("log-visible");
        }
        else {
            wsc.classList.add("log-visible");
        }
        dom.setDisplay(lp, !curr);
        if (!content) {
            content = dom.req("#log-content");
        }
        content.scrollTo(0, content.scrollHeight);
    }
    log.toggle = toggle;
    function color(level) {
        switch (level) {
            case "debug":
                return "grey-fg";
            case "info":
                return "";
            case "warn":
                return "yellow-fg";
            case "error":
                return "red-fg";
            default:
                return "";
        }
    }
})(log || (log = {}));
var map;
(function (map) {
    var Map = /** @class */ (function () {
        function Map() {
            this.storage = Object.create(null);
        }
        Map.prototype.get = function (key) {
            return this.storage[key];
        };
        ;
        Map.prototype.set = function (key, v) {
            return this.storage[key] = v;
        };
        ;
        return Map;
    }());
    map.Map = Map;
})(map || (map = {}));
var nav;
(function (nav) {
    nav.enabled = true;
    var handler = function (p) {
        console.warn("default nav handler called: " + p);
    };
    function init(f) {
        handler = f;
        window.onpopstate = function (event) {
            if (event.state) {
                var s = event.state;
                handler(s);
            }
            else {
                handler("");
            }
        };
        var path = location.pathname;
        navigate(path);
    }
    nav.init = init;
    function pop() {
        var p = location.pathname.substr(0, location.pathname.lastIndexOf("/"));
        if (p === '/c') {
            p = "";
        }
        navigate(p);
    }
    nav.pop = pop;
    function navigate(path) {
        if (!nav.enabled) {
            handler(path);
            return "";
        }
        if (str.startsWith(path, "text/html;")) {
            return "";
        }
        if (str.startsWith(path, "/")) {
            path = path.substr(1);
        }
        var locPath = location.pathname;
        if (str.startsWith(locPath, "/")) {
            locPath = locPath.substr(1);
        }
        if (locPath !== path) {
            var final = path;
            history.pushState(final, "", "/" + final);
        }
        handler(path);
    }
    nav.navigate = navigate;
    function link(o) {
        var href = o.path;
        if (!str.startsWith(href, "/")) {
            href = "/" + href;
        }
        if (o.cls) {
            o.cls = " " + o.cls.trim();
        }
        else {
            o.cls = "";
        }
        var i = JSX("span", null);
        if (o.icon) {
            i = JSX("span", { class: "nav-icon", "data-uk-icon": "icon: " + o.icon });
        }
        if (o.onclk) {
            if (!str.endsWith(o.onclk, ";")) {
                o.onclk += ";";
            }
        }
        else {
            o.onclk = "";
        }
        if (!o.isButton) {
            o.cls = style.linkColor + o.cls;
        }
        return JSX("a", { class: o.cls, href: href, onclick: o.onclk + "nav.navigate('" + o.path + "', '" + o.title + "');return false;" },
            i,
            o.title);
    }
    nav.link = link;
})(nav || (nav = {}));
var notify;
(function (notify_1) {
    function notify(msg, status) {
        UIkit.notification(msg, { status: status ? "success" : "danger", pos: "top-right" });
    }
    notify_1.notify = notify;
    function confirm(msg, f) {
        UIkit.modal.confirm(msg).then(f);
    }
    notify_1.confirm = confirm;
    function modal(key) {
        var m = UIkit.modal(key);
        if (!m) {
            console.warn("no modal available with key [" + key + "]");
        }
        return m;
    }
    notify_1.modal = modal;
})(notify || (notify = {}));
var str;
(function (str_1) {
    function startsWith(str, search, pos) {
        pos = (pos && pos > 0) ? pos | 0 : 0;
        return str.substring(pos, pos + search.length) === search;
    }
    str_1.startsWith = startsWith;
    function endsWith(str, search, pos) {
        pos = (pos && pos < str.length) ? pos : str.length;
        return str.substring(pos - search.length, pos) === search;
    }
    str_1.endsWith = endsWith;
    function trimPrefix(s, prefix) {
        if (startsWith(s, prefix)) {
            return s.slice(prefix.length);
        }
        else {
            return s;
        }
    }
    str_1.trimPrefix = trimPrefix;
    function trimSuffix(s, suffix) {
        if (endsWith(s, suffix)) {
            return s.substring(0, s.lastIndexOf(suffix));
        }
        else {
            return s;
        }
    }
    str_1.trimSuffix = trimSuffix;
})(str || (str = {}));
//# sourceMappingURL=npnasset.js.map