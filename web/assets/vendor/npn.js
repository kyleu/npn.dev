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
    var onOpen;
    var onMessage;
    var onError;
    function init(open, recv, err) {
        onOpen = open;
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
        onOpen(socket.currentID);
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
var command;
(function (command) {
    command.client = {
        ping: "ping",
        connect: "connect",
        // Collection
        getCollection: "getCollection",
        addCollection: "addCollection",
        addRequestURL: "addRequestURL",
        // Request
        getRequest: "getRequest",
        saveRequest: "saveRequest",
        call: "call",
        transform: "transform"
    };
    command.server = {
        error: "error",
        pong: "pong",
        connected: "connected",
        collections: "collections",
        collectionDetail: "collectionDetail",
        collectionAdded: "collectionAdded",
        requestDetail: "requestDetail",
        requestAdded: "requestAdded",
        callResult: "callResult",
        transformResult: "transformResult"
    };
})(command || (command = {}));
var npn;
(function (npn) {
    function onError(svc, err) {
        console.error(svc + ": " + err);
        var idx = err.lastIndexOf(":");
        if (idx > -1) {
            err = err.substr(idx + 1);
        }
        notify.notify(svc + " error: " + err, false);
    }
    npn.onError = onError;
    function init(svc, id) {
        if (inIframe()) {
            document.body.innerHTML = "";
            document.body.appendChild(rbody.iframeError());
            return;
        }
        log.init();
        window.onbeforeunload = function () {
            socket.setAppUnloading();
        };
        nav.init(routing.route);
        socket.init(function (id) { }, routing.recv, onError);
        socket.socketConnect(svc, id, svc === "wasm");
    }
    npn.init = init;
    function debug() {
        var dump = function (k, v) {
            if (v === void 0) { v = ""; }
            console.warn(k + ": " + v);
        };
        dump("Active Collection", collection.cache.active);
        dump("Active Request", request.cache.active);
        dump("Active Action", request.cache.action + " [" + request.cache.extra + "]");
    }
    npn.debug = debug;
    function testbed() {
        log.info("Testbed!");
    }
    npn.testbed = testbed;
    function inIframe() {
        try {
            return window.self !== window.top;
        }
        catch (e) {
            return true;
        }
    }
})(npn || (npn = {}));
var routing;
(function (routing) {
    function recv(msg) {
        if (socket.debug) {
            console.debug("in", msg);
        }
        switch (msg.svc) {
            case services.system.key:
                system.onSystemMessage(msg.cmd, msg.param);
                break;
            case services.collection.key:
                collection.onCollectionMessage(msg.cmd, msg.param);
                break;
            case services.request.key:
                request.onRequestMessage(msg.cmd, msg.param);
                break;
            default:
                console.warn("unhandled message for service [" + msg.svc + "]");
        }
    }
    routing.recv = recv;
    function route(p) {
        var parts = p.split("/");
        parts = parts.filter(function (x) { return x.length > 0; });
        console.debug("nav: " + parts.join(" -> "));
        var svc = (parts.length > 0) ? parts[0] : "c";
        switch (svc) {
            case "c":
                var coll = (parts.length > 1 && parts[1].length > 0) ? parts[1] : undefined;
                var req = (parts.length > 2 && parts[2].length > 0) ? parts[2] : undefined;
                var act = (parts.length > 3 && parts[3].length > 0) ? parts[3] : undefined;
                var extra = (parts.length > 4) ? parts.slice(4) : [];
                var currColl = collection.cache.active;
                collection.cache.setActiveCollection(coll);
                if (coll !== currColl && coll) {
                    socket.send({ svc: services.collection.key, cmd: command.client.getCollection, param: coll });
                }
                request.cache.setActiveRequest(req);
                request.cache.setActiveAction(act, extra);
                ui.setPanels(coll, req, act, extra);
                break;
            default:
                console.warn("unhandled svc [" + svc + "]");
        }
    }
    routing.route = route;
})(routing || (routing = {}));
var services;
(function (services) {
    services.system = { key: "system", title: "System", plural: "systems", icon: "close" };
    services.collection = { key: "collection", title: "Collection", plural: "Collections", icon: "folder" };
    services.request = { key: "request", title: "Request", plural: "Requests", icon: "file-text" };
    var allServices = [services.system, services.collection];
    function fromKey(key) {
        var ret = arr.find(allServices, function (_, s) { return s.key === key; });
        if (!ret) {
            throw "invalid service [" + key + "]";
        }
        return ret;
    }
    services.fromKey = fromKey;
})(services || (services = {}));
var rbody;
(function (rbody) {
    function renderBody(url, b) {
        if (!b) {
            return JSX("div", null, "No body");
        }
        switch (b.type) {
            case "json":
                return renderJSON(b.config);
            case "html":
                var baseURL = request.baseURL(url);
                return renderHTML(b.config, baseURL);
            case "image":
                return renderImage(b.config);
            case "raw":
                return renderRaw(b.config);
            case "error":
                return renderError(b.config);
            default:
                return JSX("div", null,
                    "unhandled body type [",
                    b.type,
                    "]");
        }
    }
    rbody.renderBody = renderBody;
    function renderHTML(h, baseURL) {
        return JSX("div", { class: "html-body" },
            JSX("span", { class: "base-url hidden" }, baseURL),
            JSX("span", { class: "preview-link right" },
                "(",
                JSX("a", { class: style.linkColor, href: "", onclick: "rbody.renderHTMLPreview(this);return false" }, "preview"),
                ")"),
            JSX("span", { class: "text-link right hidden" },
                "(",
                JSX("a", { class: style.linkColor, href: "", onclick: "rbody.renderHTMLText(this);return false" }, "text"),
                ")"),
            JSX("em", null, "HTML"),
            JSX("pre", { class: "text-content", style: "overflow: auto; max-height: 720px;" }, h.content),
            JSX("div", { class: "preview-content uk-margin-top hidden", style: "overflow: auto; max-height: 720px; border: 1px solid #666;" }));
    }
    function renderJSON(j) {
        return JSX("div", null,
            JSX("em", null, "JSON"),
            JSX("pre", null, json.str(j.msg)));
    }
    function renderImage(i) {
        var dataURL = "data:" + i.type + ";base64," + i.content;
        return JSX("img", { alt: "response image", src: dataURL });
    }
    function renderRaw(r) {
        return JSX("div", null,
            JSX("em", null, r.type ? r.type : "Unknown Type"),
            JSX("pre", null, json.str(r)));
    }
    function renderError(err) {
        return JSX("div", null,
            JSX("em", null, "Error"),
            JSX("pre", null, err.message));
    }
})(rbody || (rbody = {}));
var rbody;
(function (rbody) {
    function renderHTMLPreview(el) {
        var container = editorContent(el, true);
        var iframe = document.createElement("iframe");
        iframe.style.width = "100%";
        iframe.style.minHeight = "720px";
        var html = previewHTMLFor(container[1], container[0]);
        // iframe.src = "data:text/html;charset=utf-8," + encodeURI(html);
        container[2].innerHTML = "";
        container[2].appendChild(iframe);
        var idoc = iframe.contentDocument || iframe.contentWindow.document;
        idoc.open();
        idoc.write(html);
        idoc.close();
    }
    rbody.renderHTMLPreview = renderHTMLPreview;
    function renderHTMLText(el) {
        editorContent(el, false);
    }
    rbody.renderHTMLText = renderHTMLText;
    function editorContent(el, preview) {
        var container = el.parentElement.parentElement;
        if (!container.classList.contains("html-body")) {
            throw "container is not class [html-body]";
        }
        var baseURLEl = dom.req(".base-url", container);
        var tLink = dom.req(".text-link", container);
        var tContent = dom.req(".text-content", container);
        var pLink = dom.req(".preview-link", container);
        var pContent = dom.req(".preview-content", container);
        dom.setDisplay(tLink, preview);
        dom.setDisplay(tContent, !preview);
        dom.setDisplay(pLink, !preview);
        dom.setDisplay(pContent, preview);
        return [baseURLEl.innerText, tContent, pContent];
    }
    function previewHTMLFor(e, baseURL) {
        var ret = e.innerText;
        var headIdx = ret.indexOf("<head");
        if (headIdx > -1) {
            var headEnd = ret.indexOf(">", headIdx);
            if (headEnd > -1) {
                var base = "<base href=\"" + baseURL + "\" target=\"_blank\">";
                ret = ret.substr(0, headEnd + 1) + base + ret.substr(headEnd + 1);
            }
        }
        return ret;
    }
    function iframeError() {
        return JSX("div", { class: "uk-container" },
            JSX("div", { class: "uk-section uk-section-small" },
                JSX("h3", null, "Rendering Error"),
                JSX("p", null, "This page indicates that the HTML preview was unable to render.")));
    }
    rbody.iframeError = iframeError;
})(rbody || (rbody = {}));
var rbody;
(function (rbody) {
    rbody.AllTypes = [
        { key: "error", title: "Error", hidden: true },
        { key: "form", title: "Form", hidden: false },
        { key: "html", title: "HTML", hidden: false },
        { key: "json", title: "JSON", hidden: false },
        { key: "large", title: "Large File", hidden: false },
        { key: "raw", title: "Raw", hidden: true }
    ];
})(rbody || (rbody = {}));
var call;
(function (call) {
    function prepare(coll, r) {
        var param = { coll: coll, req: r.key, proto: r.prototype };
        socket.send({ svc: services.request.key, cmd: command.client.call, param: param });
        log.info("calling [" + request.prototypeToURL(r.prototype) + "]");
    }
    call.prepare = prepare;
    function setResult(result) {
        var container = dom.req("#" + result.collection + "--" + result.request + "-call");
        dom.setContent(container, call.renderResult(result));
        log.info("call result [" + result.id + "] received");
    }
    call.setResult = setResult;
})(call || (call = {}));
var call;
(function (call) {
    function renderResponse(rsp) {
        var _a;
        if (!rsp) {
            return JSX("div", null, "no response");
        }
        var ct = rsp.contentType || "";
        var cl = (rsp.contentLength && rsp.contentLength > -1) ? "(" + rsp.contentLength + " bytes)" : ((rsp.body && rsp.body.length > -1) ? "(" + rsp.body.length + " bytes)" : "");
        var ret = JSX("div", null,
            JSX("h3", null, rsp ? rsp.status : "Unknown"),
            JSX("em", null,
                rsp.method,
                " ",
                rsp.url),
            JSX("div", { class: "mt" },
                JSX("ul", { "data-uk-tab": "" },
                    JSX("li", null,
                        JSX("a", { href: "#result" }, "Result")),
                    JSX("li", null,
                        JSX("a", { href: "#request" }, "Request")),
                    JSX("li", null,
                        JSX("a", { href: "#headers" }, "Response")),
                    JSX("li", null,
                        JSX("a", { href: "#body" }, "Body")),
                    JSX("li", null,
                        JSX("a", { href: "#timing" }, "Timing"))),
                JSX("ul", { class: "uk-switcher uk-margin" },
                    JSX("li", null,
                        JSX("div", null,
                            (((_a = rsp.timing) === null || _a === void 0 ? void 0 : _a.completed) || 0) / 1000,
                            "ms"),
                        JSX("div", null,
                            rsp.proto,
                            " ",
                            JSX("em", null, rsp.status),
                            JSX("div", null,
                                ct,
                                " ",
                                cl))),
                    JSX("li", null, renderHeaders("Final Request Headers", rsp.requestHeaders)),
                    JSX("li", null, renderHeaders("Response Headers", rsp.headers)),
                    JSX("li", null, rbody.renderBody(rsp.url, rsp.body)),
                    JSX("li", null, renderTiming(rsp.timing)))));
        if (rsp.prior) {
            return JSX("div", null,
                renderResponse(rsp.prior),
                JSX("hr", null),
                ret);
        }
        return ret;
    }
    function renderResult(r) {
        var ret = [
            JSX("div", { class: "right" },
                JSX("a", { class: "theme uk-icon", "data-uk-icon": "close", href: "", onclick: "nav.pop();return false;", title: "close result" })),
            r.error ? JSX("div", null,
                JSX("div", { class: "red-fg" },
                    "error: ",
                    r.error)) : JSX("div", null),
            renderResponse(r.response)
        ];
        return ret;
    }
    call.renderResult = renderResult;
    function renderHeaders(title, headers) {
        if (headers === void 0) { headers = []; }
        if (headers.length === 0) {
            return section(title, "No headers");
        }
        return JSX("div", { class: "uk-overflow-auto" },
            JSX("h4", null, title),
            JSX("table", { class: "uk-table uk-table-divider uk-text-left uk-table-small uk-table-justify" },
                JSX("tbody", null, headers.map(function (h) { return JSX("tr", { title: h.desc },
                    JSX("td", { class: "uk-text-nowrap" }, h.k),
                    JSX("td", { class: "uk-text-nowrap" }, h.v)); }))));
    }
    function renderTiming(t) {
        if (!t) {
            return JSX("div", null, "No timing");
        }
        var sections = call.timingSections(t);
        return JSX("div", { class: "timing-panel" },
            JSX("div", { class: "result-timing-graph" },
                JSX("div", null,
                    JSX("div", { class: "timing-start" }, "0ms"),
                    JSX("div", { class: "timing-end" },
                        t.completed / 1000,
                        "ms")),
                JSX("object", { type: "image/svg+xml", style: "width: 100%; height: " + (sections.length * 24) + "px", data: call.timingGraph(sections) }, "SVG not supported")),
            JSX("hr", null),
            sections.map(function (sc) { return JSX("div", null,
                sc.key,
                ": ",
                sc.start,
                " - ",
                sc.end); }));
    }
    function section(k, v) {
        if (!v) {
            v = "undefined";
        }
        return JSX("div", null,
            JSX("h4", null, k),
            " ",
            v);
    }
})(call || (call = {}));
var call;
(function (call) {
    function timingSections(t) {
        var ret = [];
        var add = function (k, g, s, e) {
            if (s && e) {
                ret.push({ key: k, group: g, start: s, end: e });
            }
        };
        add("dns", "connect", t.dnsStart, t.dnsEnd);
        add("connect", "connect", t.connectStart, t.connectEnd);
        var cc = t.connectEnd;
        if ((t.tlsEnd || 0) > 0) {
            cc = t.tlsEnd || 0;
            add("tls", "connect", t.tlsStart || 0, cc);
        }
        add("reqheaders", "request", cc, t.wroteHeaders);
        if ((t.wroteRequest - t.wroteHeaders) > 2) {
            add("reqbody", "request", t.wroteHeaders, t.wroteRequest);
        }
        add("rspwait", "response", t.wroteRequest, t.firstResponseByte);
        add("rspheaders", "response", t.firstResponseByte, t.responseHeaders);
        add("rspbody", "response", t.responseHeaders, t.completed);
        return ret;
    }
    call.timingSections = timingSections;
    function timingGraph(ts) {
        var _a;
        var ret = [];
        for (var _i = 0, ts_1 = ts; _i < ts_1.length; _i++) {
            var t = ts_1[_i];
            if (t.group.length > 0) {
                ret.push(encodeURIComponent(t.key + ".g") + '=' + encodeURIComponent(t.group));
            }
            ret.push(encodeURIComponent(t.key + ".s") + '=' + encodeURIComponent(t.start));
            ret.push(encodeURIComponent(t.key + ".e") + '=' + encodeURIComponent(t.end));
        }
        ret.push("t=" + ((_a = system.cache.profile) === null || _a === void 0 ? void 0 : _a.theme) || "light");
        return "/svg/gantt?" + ret.join("&");
    }
    call.timingGraph = timingGraph;
})(call || (call = {}));
var collection;
(function (collection_1) {
    var Cache = /** @class */ (function () {
        function Cache() {
            this.collections = [];
        }
        Cache.prototype.updateCollection = function (collection) {
            this.collections = group.updateAndSort(this.collections, collection, function (t) { return t.key; });
            collection_1.renderCollections(this.collections);
        };
        Cache.prototype.setActiveCollection = function (key) {
            if (this.active !== key) {
                this.active = key;
                collection_1.renderCollections(this.collections);
            }
        };
        Cache.prototype.getActiveCollection = function () {
            for (var _i = 0, _a = this.collections; _i < _a.length; _i++) {
                var x = _a[_i];
                if (x.key == this.active) {
                    return x;
                }
            }
            return undefined;
        };
        return Cache;
    }());
    collection_1.cache = new Cache();
})(collection || (collection = {}));
var collection;
(function (collection) {
    function renderCollections(cs) {
        return dom.els(".collection-list").forEach(function (el) {
            dom.setContent(el, cs.map(function (c) { return renderCollectionLink(c); }));
        });
    }
    collection.renderCollections = renderCollections;
    function renderCollectionLink(c) {
        var title = c.title;
        if (!title || title.length === 0) {
            title = c.key;
        }
        var link = nav.link({ path: "/c/" + c.key, title: title, icon: "folder" });
        if (collection.cache.active === c.key) {
            var activeReq = request.cache.active;
            var summs = request.cache.summaries.get(c.key);
            if (summs) {
                var collLink = void 0;
                if (activeReq) {
                    collLink = nav.link({ path: "/c/" + c.key, title: title, icon: "album" });
                }
                else {
                    collLink = JSX("strong", null, nav.link({ path: "/", title: title, icon: "album" }));
                }
                link = JSX("div", null,
                    collLink,
                    summs.map(function (s) {
                        var l = nav.link({ path: "/c/" + c.key + "/" + s.key, title: (s.title && s.title.length > 0) ? s.title : s.key, icon: "link" });
                        return JSX("div", { class: "uk-margin-small-left" }, request.cache.active === s.key ? JSX("strong", null, l) : l);
                    }));
            }
        }
        return JSX("div", { class: "nav-item collection-link collection-link-" + c.key }, link);
    }
    function renderCollection(coll, requests) {
        var cn = coll.title ? coll.title : coll.key;
        return JSX("div", null,
            JSX("div", { class: "uk-card uk-card-body uk-card-default" },
                JSX("div", { class: "right" },
                    JSX("a", { class: "theme uk-icon", "data-uk-icon": "close", href: "", onclick: "nav.pop();return false;", title: "close collection" })),
                JSX("h3", { class: "uk-card-title" },
                    JSX("span", { class: "nav-icon-h3", "data-uk-icon": "icon: album" }),
                    cn),
                JSX("p", null, coll.description || "")),
            JSX("div", { class: "uk-card uk-card-body uk-card-default uk-margin-top" },
                JSX("h3", { class: "uk-card-title" }, "Requests"),
                JSX("form", { onsubmit: "collection.addRequestURL();return false;" },
                    JSX("div", { class: "uk-margin-top uk-inline uk-width-expand" },
                        JSX("button", { class: "uk-form-icon uk-form-icon-flip", type: "submit", title: "add request", "uk-icon": "icon: plus" }),
                        JSX("input", { id: "coll-request-add-url", class: "uk-input", placeholder: "add a request by url", "data-lpignore": "true" }))),
                JSX("div", { id: "request-list", class: "uk-margin-top" }, renderRequests(coll.key, requests))));
    }
    collection.renderCollection = renderCollection;
    function addFromInput() {
        var input = dom.req("#coll-add-input");
        var name = input.value.trim();
        if (name && name.length > 0) {
            input.value = "";
            socket.send({ svc: services.collection.key, cmd: command.client.addCollection, param: name });
            log.info("adding request [" + name + "]");
        }
    }
    collection.addFromInput = addFromInput;
    function addRequestURL() {
        var input = dom.req("#coll-request-add-url");
        var url = input.value.trim();
        if (url && url.length > 0) {
            input.value = "";
            var param = { "coll": collection.cache.active, "url": url };
            socket.send({ svc: services.collection.key, cmd: command.client.addRequestURL, param: param });
            log.info("adding request [" + url + "]");
        }
    }
    collection.addRequestURL = addRequestURL;
    function renderRequests(coll, rs) {
        rs = group.sort(rs, function (x) { return x.order; });
        return JSX("ul", { class: "uk-list uk-list-divider" }, rs.map(function (r) { return renderRequestLink(coll, r); }));
    }
    function renderRequestLink(coll, r) {
        var title = r.title;
        if (!title || r.title.length === 0) {
            title = r.key;
        }
        return JSX("li", null,
            nav.link({ path: "/c/" + coll + "/" + r.key, title: title }),
            r.description && r.description.length ? JSX("div", null,
                JSX("em", null, r.description)) : JSX("span", null));
    }
})(collection || (collection = {}));
var collection;
(function (collection) {
    function onCollectionMessage(cmd, param) {
        switch (cmd) {
            case command.server.collections:
                collection.cache.collections = group.sort(param, function (c) { return c.key; });
                log.info("processing [" + collection.cache.collections.length + "] collections");
                collection.renderCollections(collection.cache.collections);
                break;
            case command.server.collectionDetail:
                var d = param;
                log.info("processing [" + d.requests.length + "] requests for collection [" + d.collection.key + "]");
                collection.cache.updateCollection(d.collection);
                request.cache.setCollectionRequests(d.collection, d.requests);
                collection.renderCollections(collection.cache.collections);
                break;
            case command.server.collectionAdded:
                var a = param;
                log.info("processing new collection [" + a.active + "]");
                collection.cache.collections = a.collections;
                nav.navigate("/c/" + a.active);
                break;
            default:
                console.warn("unhandled collection command [" + cmd + "]");
        }
    }
    collection.onCollectionMessage = onCollectionMessage;
})(collection || (collection = {}));
var header;
(function (header) {
    function nch(key, description, req, rsp, link) {
        return { "key": key, "description": description, "req": req, "rsp": rsp, "link": link };
    }
    function snch(key, description, req, rsp) {
        return nch(key, description, req, rsp, mdnLink(key));
    }
    function mdnLink(s) {
        return "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/" + s;
    }
    header.commonHeaders = [
        snch("Accept", "Informs the server about the types of data that can be sent back.", true, false),
        snch("Access-Control-Allow-Headers", "Used in response to a preflight request to indicate which HTTP headers can be used when making the actual request.", false, true),
        snch("Access-Control-Allow-Methods", "Specifies the methods allowed when accessing the resource in response to a preflight request.", false, true),
        snch("Access-Control-Allow-Origin", "Indicates whether the response can be shared.", false, true),
        snch("Authorization", "Contains the credentials to authenticate a user-agent with a server.", true, false),
        snch("Connection", "Controls whether the network connection stays open after the current transaction finishes.", true, false),
        snch("Content-Encoding", "Used to specify the compression algorithm.", true, true),
        snch("Content-Length", "The size of the resource, in decimal number of bytes.", true, true),
        snch("Content-Type", "Indicates the media type of the resource.", true, true),
        snch("Cookie", "Contains stored HTTP cookies previously sent by the server with the Set-Cookie header.", true, false),
        snch("Date", "The Date general HTTP header contains the date and time at which the message was originated.", false, true),
        snch("ETag", "A unique string identifying the version of the resource.", false, true),
        snch("Expires", "The date/time after which the response is considered stale.", false, true),
        snch("Host", "Specifies the domain name of the server (for virtual hosting), and (optionally) the TCP port number on which the server is listening.", true, false),
        snch("Last-Modified", "The last modification date of the resource, used to compare several versions of the same resource.", false, true),
        snch("Location", "Indicates the URL to redirect a page to. ", false, true),
        snch("Origin", "Indicates where a fetch originates from.", true, false),
        snch("Referer", "The address of the previous web page from which a link to the currently requested page was followed.", true, false),
        snch("Server", "Contains information about the software used by the origin server to handle the request.", false, true),
        snch("Set-Cookie", "Send cookies from the server to the user-agent.", false, true),
        snch("User-Agent", "Contains a characteristic string that allows the network protocol peers to identify the application", true, false)
    ];
    var commonHeadersByName;
    function getCommonHeaderByName(key) {
        if (!commonHeadersByName) {
            commonHeadersByName = new map.Map();
            for (var _i = 0, commonHeaders_1 = header.commonHeaders; _i < commonHeaders_1.length; _i++) {
                var ch = commonHeaders_1[_i];
                commonHeadersByName.set(ch.key, ch);
            }
        }
        return commonHeadersByName.get(key);
    }
    header.getCommonHeaderByName = getCommonHeaderByName;
    function dumpCommonHeaders() {
        var dump = function (title, req, rsp) {
            var matched = false;
            console.debug("\n::: " + title + " Headers");
            header.commonHeaders.forEach(function (ch) {
                if (ch.req == req && ch.rsp == rsp) {
                    matched = true;
                    console.debug(ch.key + ": " + ch.link);
                    console.debug("  - " + ch.description);
                }
            });
            if (!matched) {
                console.debug("none");
            }
        };
        dump("Common", true, true);
        dump("Request", true, false);
        dump("Response", false, true);
        dump("Invalid", false, false);
    }
    header.dumpCommonHeaders = dumpCommonHeaders;
})(header || (header = {}));
var request;
(function (request) {
    function renderActionEmpty() {
        return JSX("div", null);
    }
    request.renderActionEmpty = renderActionEmpty;
    function renderActionUnknown(key, extra) {
        return JSX("div", null,
            renderActionClose(),
            "unknown action: ",
            key,
            " (",
            extra,
            ")");
    }
    request.renderActionUnknown = renderActionUnknown;
    function renderActionCall(coll, req) {
        return JSX("div", { id: coll + "--" + req + "-call" },
            renderActionClose(),
            JSX("div", { class: "call-title" }, "Loading..."),
            JSX("div", { class: "call-result" }));
    }
    request.renderActionCall = renderActionCall;
    function renderActionClose() {
        return JSX("div", { class: "right" },
            JSX("a", { class: "theme uk-icon", "data-uk-icon": "close", href: "", onclick: "nav.navigate(`/c/${collection.cache.active}/${request.cache.active}`);return false;", title: "close collection" }));
    }
    request.renderActionClose = renderActionClose;
})(request || (request = {}));
var request;
(function (request) {
    var Cache = /** @class */ (function () {
        function Cache() {
            this.summaries = new map.Map();
            this.requests = new map.Map();
            this.extra = [];
        }
        Cache.prototype.setCollectionRequests = function (coll, summs) {
            this.summaries.set(coll.key, summs);
            if (coll.key === collection.cache.active) {
                dom.setContent("#collection-panel", collection.renderCollection(coll, summs));
                for (var _i = 0, summs_1 = summs; _i < summs_1.length; _i++) {
                    var req = summs_1[_i];
                    if (this.active === req.key) {
                        request.renderActiveRequest(collection.cache.active);
                        if (this.action) {
                            request.renderAction(collection.cache.active, req.key, this.action, this.extra);
                        }
                    }
                }
            }
        };
        Cache.prototype.setActiveRequest = function (key) {
            if (!collection.cache.active) {
                return;
            }
            if (this.active !== key) {
                this.active = key;
                if (this.active) {
                    request.renderActiveRequest(collection.cache.active);
                }
                collection.renderCollections(collection.cache.collections);
            }
        };
        Cache.prototype.setActiveAction = function (act, extra) {
            if (!collection.cache.active) {
                return;
            }
            var sameExtra = this.extra.length === extra.length && this.extra.every(function (value, index) { return value === extra[index]; });
            if (this.active /* && (this.action !== act || !sameExtra) */) {
                this.action = act;
                this.extra = extra;
                request.renderAction(collection.cache.active, this.active, this.action, this.extra);
            }
        };
        Cache.prototype.updateRequest = function (r) {
            if (!collection.cache.active) {
                return;
            }
            var curr = this.requests.get(collection.cache.active);
            var updated = group.update(curr, r, function (x) { return x.key; });
            this.requests.set(collection.cache.active, updated);
        };
        return Cache;
    }());
    request.cache = new Cache();
})(request || (request = {}));
var request;
(function (request) {
    function diff(l, r) {
        var ret = [];
        var p = function (k, lv, rv) { return ret.push({ k: k, l: lv, r: rv }); };
        var comp = function (k, lv, rv) {
            if (lv === undefined || lv === null) {
                lv = "";
            }
            if (rv === undefined || rv === null) {
                rv = "";
            }
            if (typeof lv === "object" && typeof rv === "object") {
                for (var f in lv) {
                    if (lv.hasOwnProperty(f)) {
                        comp(k + "." + f, lv[f], rv[f]);
                    }
                }
                for (var f in rv) {
                    if (rv.hasOwnProperty(f) && !lv.hasOwnProperty(f)) {
                        comp(k + "." + f, lv[f], rv[f]);
                    }
                }
            }
            else {
                if (lv !== rv) {
                    p(k, lv, rv);
                }
            }
        };
        var compArray = function (k, lv, rv) {
            if (lv === undefined || lv === null) {
                lv = [];
            }
            if (rv === undefined || rv === null) {
                rv = [];
            }
            if (lv.length !== rv.length) {
                p(k + ".length", lv.length, rv.length);
            }
            for (var i = 0; i < lv.length; i++) {
                comp(k + "[" + i + "]", lv[i], rv[i]);
            }
        };
        var checkNull = function (k, lv, rv) {
            if (!l) {
                if (r) {
                    p(k, null, "(defined)");
                }
                return true;
            }
            if (!r) {
                p(k, "(defined)", null);
                return true;
            }
            return false;
        };
        if (checkNull("request", l, r)) {
            return ret;
        }
        comp("key", l.key, r.key);
        comp("title", l.title, r.title);
        comp("description", l.description, r.description);
        var lp = l.prototype;
        var rp = r.prototype;
        comp("method", lp.method, rp.method);
        comp("protocol", lp.protocol, rp.protocol);
        comp("domain", lp.domain, rp.domain);
        comp("port", lp.port, rp.port);
        comp("path", lp.path, rp.path);
        compArray("query", lp.query, rp.query);
        comp("fragment", lp.fragment, rp.fragment);
        compArray("headers", lp.headers, rp.headers);
        compArray("auth", lp.auth, rp.auth);
        if (!checkNull("body", lp.body, rp.body)) {
            if (lp.body && rp.body) {
                comp("body.type", lp.body.type, rp.body.type);
                comp("body.config", lp.body.config, rp.body.config);
            }
        }
        var lpo = lp.options;
        var rpo = rp.options;
        if (checkNull("options", lpo, rpo)) {
            return ret;
        }
        if ((!lpo) || (!rpo)) {
            return ret;
        }
        comp("timeout", lpo.timeout, rpo.timeout);
        comp("ignoreRedirects", lpo.ignoreRedirects, rpo.ignoreRedirects);
        comp("ignoreReferrer", lpo.ignoreReferrer, rpo.ignoreReferrer);
        comp("ignoreCerts", lpo.ignoreCerts, rpo.ignoreCerts);
        comp("ignoreCookies", lpo.ignoreCookies, rpo.ignoreCookies);
        compArray("excludeDefaultHeaders", lpo.excludeDefaultHeaders, rpo.excludeDefaultHeaders);
        comp("readCookieJars", lpo.readCookieJars, rpo.readCookieJars);
        comp("writeCookieJar", lpo.writeCookieJar, rpo.writeCookieJar);
        comp("sslCert", lpo.sslCert, rpo.sslCert);
        comp("userAgentOverride", lpo.userAgentOverride, rpo.userAgentOverride);
        return ret;
    }
    request.diff = diff;
})(request || (request = {}));
var request;
(function (request) {
    function getActiveRequest() {
        return getRequest(collection.cache.active, request.cache.active);
    }
    request.getActiveRequest = getActiveRequest;
    function getSummary(coll, key) {
        for (var _i = 0, _a = request.cache.summaries.get(coll) || []; _i < _a.length; _i++) {
            var req = _a[_i];
            if (req.key === key) {
                return req;
            }
        }
        return undefined;
    }
    request.getSummary = getSummary;
    function getRequest(coll, key) {
        for (var _i = 0, _a = request.cache.requests.get(coll) || []; _i < _a.length; _i++) {
            var req = _a[_i];
            if (req.key === key) {
                return req;
            }
        }
        return undefined;
    }
    request.getRequest = getRequest;
    function onRequestMessage(cmd, param) {
        switch (cmd) {
            case command.server.requestDetail:
                var req = param;
                log.info("received details for request [" + req.key + "]");
                request.cache.updateRequest(req);
                if (request.cache.active === req.key) {
                    request.renderActiveRequest(collection.cache.active);
                    request.renderAction(collection.cache.active, request.cache.active, request.cache.action, request.cache.extra);
                }
                break;
            case command.server.requestAdded:
                var ra = param;
                log.info("received details for new request [" + ra.key + "]");
                request.cache.updateRequest(ra);
                request.cache.setActiveRequest(ra.key);
                nav.navigate("/c/" + collection.cache.active + "/" + ra.key);
                break;
            case command.server.callResult:
                var result = param;
                call.setResult(result);
                var path = "r/" + result.id;
                // TODO history.replaceState(path, "", "/" + path);
                break;
            case command.server.transformResult:
                transform.setResult(param);
                break;
            default:
                console.warn("unhandled request command [" + cmd + "]");
        }
    }
    request.onRequestMessage = onRequestMessage;
})(request || (request = {}));
var request;
(function (request) {
    request.MethodGet = { "key": "GET", "description": "" };
    request.MethodHead = { "key": "HEAD", "description": "" };
    request.MethodPost = { "key": "POST", "description": "" };
    request.MethodPut = { "key": "PUT", "description": "" };
    request.MethodPatch = { "key": "PATCH", "description": "" };
    request.MethodDelete = { "key": "DELETE", "description": "" };
    request.MethodConnect = { "key": "CONNECT", "description": "" };
    request.MethodOptions = { "key": "OPTIONS", "description": "" };
    request.MethodTrace = { "key": "TRACE", "description": "" };
    request.allMethods = [request.MethodGet, request.MethodHead, request.MethodPost, request.MethodPut, request.MethodPatch, request.MethodDelete, request.MethodConnect, request.MethodOptions, request.MethodTrace];
})(request || (request = {}));
var request;
(function (request) {
    function newPrototype(protocol, hostname, port, path, qp, fragment, auth) {
        if (str.endsWith(protocol, ":")) {
            protocol = protocol.substr(0, protocol.length - 1);
        }
        if (str.startsWith(fragment, "#")) {
            fragment = fragment.substr(1);
        }
        return { method: "get", protocol: protocol, domain: hostname, port: port, path: path, query: qp, fragment: fragment, auth: auth };
    }
    function prototypeFromURL(u) {
        var url = new URL(u);
        var qp = [];
        url.searchParams.forEach(function (v, k) { return qp.push({ k: k, v: v }); });
        var auth = [];
        if (url.username.length > 0) {
            auth.push({ type: "basic", config: { "username": url.username, "password": url.password, "showPassword": true } });
        }
        var port;
        if (url.port && url.port.length > 0) {
            port = parseInt(url.port, 10);
        }
        return newPrototype(url.protocol, url.hostname, port, url.pathname, qp, url.hash, auth);
    }
    request.prototypeFromURL = prototypeFromURL;
})(request || (request = {}));
var request;
(function (request) {
    function renderActiveRequest(coll) {
        if (request.cache.active) {
            render(coll, request.cache.active);
        }
        else {
            console.warn("no active request");
        }
    }
    request.renderActiveRequest = renderActiveRequest;
    function render(coll, reqKey) {
        var req = request.getRequest(coll, reqKey);
        if (req) {
            dom.setContent("#request-panel", request.form.renderFormPanel(coll, req));
            request.editor.wireForm(req.key);
        }
        else {
            var summ = request.getSummary(coll, reqKey);
            if (summ) {
                dom.setContent("#request-panel", request.renderSummaryPanel(coll, summ));
                var param = { coll: coll, req: summ.key };
                socket.send({ svc: services.request.key, cmd: command.client.getRequest, param: param });
            }
        }
    }
    request.render = render;
    function renderAction(coll, reqKey, action, extra) {
        var re = dom.opt(".request-editor");
        var ra = dom.opt(".request-action");
        if (!re || !ra) {
            return;
        }
        switch (action) {
            case undefined:
                dom.setContent(ra, request.renderActionEmpty());
                break;
            case "call":
                // call.prepare(coll, getRequest(coll, reqKey));
                call.prepare(coll, request.form.extractRequest(request.cache.active));
                dom.setContent(ra, request.renderActionCall(coll, reqKey));
                break;
            case "transform":
                var req = request.form.extractRequest(request.cache.active);
                dom.setContent(ra, transform.renderRequest(coll, reqKey, extra[0]));
                var param = { coll: coll, req: reqKey, fmt: extra[0], proto: req.prototype };
                socket.send({ svc: services.request.key, cmd: command.client.transform, param: param });
                break;
            default:
                console.warn("unhandled request action [" + action + "]");
                dom.setContent(ra, request.renderActionUnknown(action, extra));
        }
        dom.setDisplay(re, action === undefined);
        dom.setDisplay(ra, action !== undefined);
    }
    request.renderAction = renderAction;
})(request || (request = {}));
var request;
(function (request) {
    function renderSummaryPanel(coll, r) {
        return JSX("div", null,
            JSX("div", { class: "uk-card uk-card-body uk-card-default" },
                JSX("div", { class: "right" },
                    JSX("a", { class: "theme uk-icon", "data-uk-icon": "close", href: "", onclick: "nav.navigate('/c/" + coll + "');return false;", title: "close request" })),
                JSX("h3", { class: "uk-card-title" }, r.title ? r.title : r.key),
                JSX("p", null, "Loading...")));
    }
    request.renderSummaryPanel = renderSummaryPanel;
})(request || (request = {}));
var request;
(function (request) {
    function urlToPrototype(url) {
        var u = new URL(url);
        return {
            method: request.MethodGet.key,
            protocol: str.trimSuffix(u.protocol, ":"),
            domain: u.hostname,
            port: u.port ? parseInt(u.port, 10) : undefined,
            path: str.trimPrefix(u.pathname, "/"),
            fragment: str.trimPrefix(u.hash, "#")
        };
    }
    request.urlToPrototype = urlToPrototype;
    function prototypeToURL(p) {
        return prototypeToURLParts(p).map(function (x) { return x.v; }).join("");
    }
    request.prototypeToURL = prototypeToURL;
    function prototypeToHTML(p) {
        return JSX("span", null, prototypeToURLParts(p).map(function (x) { return JSX("span", { title: x.t, class: urlColor(x.t) }, x.v); }));
    }
    request.prototypeToHTML = prototypeToHTML;
    function baseURL(s) {
        return prototypeBaseURL(urlToPrototype(s));
    }
    request.baseURL = baseURL;
    function prototypeBaseURL(p) {
        if (!p) {
            return "invalid";
        }
        var d = p.domain;
        if (p.port && p.port > 0) {
            d += ":" + p.port;
        }
        return p.protocol + "://" + d + "/";
    }
    request.prototypeBaseURL = prototypeBaseURL;
    function prototypeToURLParts(p) {
        var ret = [];
        var push = function (t, v) {
            ret.push({ t: t, v: v });
        };
        push("protocol", p.protocol);
        push("", "://");
        if (p.auth) {
            for (var _i = 0, _a = p.auth; _i < _a.length; _i++) {
                var a = _a[_i];
                if (a.type === "basic") {
                    var cfg = a.config;
                    push("username", cfg.username);
                    push("", ":");
                    if (cfg.showPassword) {
                        push("password", cfg.password);
                    }
                    else {
                        push("password", "****");
                    }
                    push("", "@");
                    break;
                }
            }
        }
        push("domain", p.domain);
        if (p.port) {
            push("", ":");
            push("port", p.port.toString());
        }
        if (p.path && p.path.length > 0) {
            push("", "/");
            push("path", p.path);
        }
        if (p.query && p.query.length > 0) {
            push("", "?");
            var query = p.query.map(function (k) { return encodeURIComponent(k.k) + '=' + encodeURIComponent(k.v); }).join('&');
            push("query", query);
        }
        if (p.fragment && p.fragment.length > 0) {
            push("", "#");
            push("fragment", encodeURIComponent(p.fragment));
        }
        return ret;
    }
    function urlColor(key) {
        switch (key) {
            case "username":
            case "password":
            case "protocol":
            case "auth":
                return "green-fg";
            case "domain":
            case "port":
                return "blue-fg";
            case "path":
                return "bluegrey-fg";
            case "query":
                return "purple-fg";
            default:
                return "";
        }
    }
})(request || (request = {}));
var request;
(function (request) {
    var editor;
    (function (editor) {
        function initAuthEditor(el) {
        }
        editor.initAuthEditor = initAuthEditor;
        function setAuth(cache, auth) {
            var url = new URL(cache.url.value);
            var u = "";
            var p = "";
            if (auth) {
                for (var _i = 0, auth_1 = auth; _i < auth_1.length; _i++) {
                    var a = auth_1[_i];
                    if (a.type === "basic") {
                        var basic = a.config;
                        u = encodeURIComponent(basic.username);
                        p = encodeURIComponent(basic.password);
                    }
                }
            }
            url.username = u;
            url.password = p;
            cache.url.value = url.toString();
        }
        editor.setAuth = setAuth;
        function updateBasicAuth(cache, auth) {
            var currentAuth = [];
            try {
                currentAuth = json.parse(cache.auth.value);
            }
            catch (e) {
                console.warn("invalid auth JSON [" + cache.auth.value + "]");
            }
            var matched = -1;
            if (!currentAuth) {
                currentAuth = [];
            }
            for (var i = 0; i < currentAuth.length; i++) {
                var x = currentAuth[i];
                if (x.type === "basic") {
                    matched = i;
                }
            }
            var basic;
            if (auth) {
                for (var i = 0; i < auth.length; i++) {
                    var x = auth[i];
                    if (x.type === "basic") {
                        basic = x.config;
                    }
                }
            }
            if (matched === -1) {
                if (basic) {
                    currentAuth.push({ type: "basic", config: basic });
                }
            }
            else {
                if (basic) {
                    var curr = currentAuth[matched].config;
                    if (curr) {
                        curr = {
                            username: basic.username,
                            password: basic.password,
                            showPassword: curr.showPassword
                        };
                    }
                    else {
                        curr = basic;
                    }
                    currentAuth[matched] = { type: "basic", config: curr };
                }
                else {
                    currentAuth.splice(matched, 1);
                }
            }
            cache.auth.value = json.str(currentAuth);
        }
        editor.updateBasicAuth = updateBasicAuth;
    })(editor = request.editor || (request.editor = {}));
})(request || (request = {}));
var request;
(function (request) {
    var editor;
    (function (editor) {
        function initBodyEditor(el) {
            var parent = el.parentElement;
            parent.appendChild(createBodyEditor(el));
        }
        editor.initBodyEditor = initBodyEditor;
        function createBodyEditor(el) {
            var b = json.parse(el.value);
            return JSX("div", { class: "uk-margin-top" },
                JSX("select", { class: "uk-select" },
                    JSX("option", { value: "" }, "No body"),
                    rbody.AllTypes.filter(function (t) { return !t.hidden; }).map(function (t) {
                        if (b && b.type === t.key) {
                            return JSX("option", { value: t.key, selected: "selected" }, t.title);
                        }
                        else {
                            return JSX("option", { value: t.key }, t.title);
                        }
                    }),
                    "\u02D9"),
                rbody.AllTypes.filter(function (t) { return !t.hidden; }).map(function (t) {
                    var cfg = (b && b.type == t.key) ? b.config : null;
                    return configEditor(t.key, cfg, t.key === (b ? b.type : ""));
                }));
        }
        function configEditor(key, config, active) {
            var cls = "uk-margin-top body-editor-" + key;
            if (!active) {
                cls += " hidden";
            }
            switch (key) {
                case "json":
                    var j = config;
                    return JSX("div", { class: cls },
                        JSX("textarea", { class: "uk-textarea" }, json.str(j ? j.msg : null)));
                default:
                    return JSX("div", { class: cls },
                        "Unimplemented [",
                        key,
                        "] editor");
            }
        }
        function setBody(cache, body) {
        }
        editor.setBody = setBody;
    })(editor = request.editor || (request.editor = {}));
})(request || (request = {}));
var request;
(function (request) {
    var editor;
    (function (editor) {
        function wireForm(prefix) {
            var id = function (k) {
                return "#" + prefix + "-" + k;
            };
            var cache = {
                key: dom.req(id("key")),
                title: dom.req(id("title")),
                desc: dom.req(id("description")),
                url: dom.req(id("url")),
                method: dom.req(id("method")),
                auth: dom.req(id("auth")),
                qp: dom.req(id("queryparams")),
                headers: dom.req(id("headers")),
                body: dom.req(id("body")),
                options: dom.req(id("options"))
            };
            initEditors(prefix, cache);
            wireEvents(cache);
        }
        editor.wireForm = wireForm;
        function initEditors(prefix, cache) {
            editor.initURLEditor(cache.url);
            editor.initAuthEditor(cache.auth);
            editor.initQueryParamsEditor(cache.qp);
            editor.initHeadersEditor(cache.headers);
            editor.initBodyEditor(cache.body);
            editor.initOptionsEditor(cache.options);
        }
        function events(e, f) {
            var x = function () {
                f();
                return true;
            };
            e.onchange = x;
            e.onkeyup = x;
            e.onblur = x;
        }
        editor.events = events;
        function wireEvents(cache) {
            events(cache.key, function () {
                request.form.checkEditor(request.cache.active);
            });
            events(cache.title, function () {
                request.form.checkEditor(request.cache.active);
            });
            events(cache.desc, function () {
                request.form.checkEditor(request.cache.active);
            });
            events(cache.method, function () {
                request.form.checkEditor(request.cache.active);
            });
            events(cache.url, function () {
                var p = request.prototypeFromURL(cache.url.value);
                editor.setURL(cache, p);
                request.form.checkEditor(request.cache.active);
            });
            events(cache.auth, function () {
                var auth;
                try {
                    auth = json.parse(cache.auth.value);
                }
                catch (e) {
                    console.warn("invalid auth JSON [" + cache.auth.value + "]");
                    auth = [];
                }
                editor.setAuth(cache, auth);
            });
            events(cache.qp, function () {
                var qp;
                try {
                    qp = json.parse(cache.qp.value);
                }
                catch (e) {
                    console.warn("invalid qp JSON [" + cache.qp.value + "]");
                    qp = [];
                }
                editor.setQueryParams(cache.url, qp);
            });
            events(cache.headers, function () {
                var h;
                try {
                    h = json.parse(cache.headers.value);
                }
                catch (e) {
                    console.warn("invalid headers JSON [" + cache.headers.value + "]");
                    h = [];
                }
                editor.setHeaders(cache, h);
            });
            events(cache.body, function () {
                var b;
                try {
                    b = json.parse(cache.body.value);
                }
                catch (e) {
                    console.warn("invalid body JSON [" + cache.body.value + "]");
                }
                editor.setBody(cache, b);
            });
        }
    })(editor = request.editor || (request.editor = {}));
})(request || (request = {}));
var request;
(function (request) {
    var editor;
    (function (editor) {
        function initHeadersEditor(el) {
            var parent = el.parentElement;
            parent.appendChild(createHeadersEditor(el));
        }
        editor.initHeadersEditor = initHeadersEditor;
        function setHeaders(cache, headers) {
        }
        editor.setHeaders = setHeaders;
        function createHeadersEditor(el) {
            var container = JSX("ul", { id: el.id + "-ul", class: "uk-list uk-list-divider" });
            var header = JSX("li", null,
                JSX("div", { "data-uk-grid": "" },
                    JSX("div", { class: "uk-width-1-4" }, "Name"),
                    JSX("div", { class: "uk-width-1-4" }, "Value"),
                    JSX("div", { class: "uk-width-1-2" },
                        JSX("div", { class: "right" },
                            JSX("a", { class: style.linkColor, href: "", onclick: "return request.editor.addHeaderRow('" + el.id + "')", title: "new header" },
                                JSX("span", { "data-uk-icon": "icon: plus" }))),
                        "Description")));
            var curr = json.parse(el.value);
            container.innerText = "";
            container.appendChild(header);
            if (curr) {
                for (var idx = 0; idx < curr.length; idx++) {
                    addChild(el.id, idx, container, curr[idx]);
                }
            }
            return container;
        }
        function addHeaderRow(id) {
            var ul = dom.req("#" + id + "-ul");
            var idx = ul.children.length - 1;
            addChild(id, idx, ul, { k: '', v: '' });
            return false;
        }
        editor.addHeaderRow = addHeaderRow;
        function removeHeaderRow(id, el) {
            el.parentElement.parentElement.parentElement.parentElement.remove();
            parseHeaders(id);
            return false;
        }
        editor.removeHeaderRow = removeHeaderRow;
        function addChild(elID, idx, container, h) {
            var ret = JSX("li", null,
                JSX("div", { "data-uk-grid": "" },
                    JSX("div", { class: "uk-width-1-4" },
                        JSX("input", { class: "uk-input", "data-field": idx + "-key", type: "text", value: h.k })),
                    JSX("div", { class: "uk-width-1-4" },
                        JSX("input", { class: "uk-input", "data-field": idx + "-value", type: "text", value: h.v })),
                    JSX("div", { class: "uk-width-1-2" },
                        JSX("div", { class: "right", style: "margin-top: 6px;" },
                            JSX("a", { class: style.linkColor, href: "", onclick: "return request.editor.removeHeaderRow('" + elID + "', this);", title: "new header" },
                                JSX("span", { "data-uk-icon": "icon: close" }))),
                        JSX("input", { style: "width: calc(100% - 48px);", class: "uk-input", "data-field": idx + "-desc", type: "text", value: h.desc }))));
            editor.events(ret, function () { return parseHeaders(elID); });
            container.appendChild(ret);
        }
        function parseHeaders(elID) {
            var ta = dom.req("#" + elID);
            var ul = dom.req("#" + elID + "-ul");
            var inputs = dom.els("input", ul);
            var ret = [];
            for (var _i = 0, inputs_1 = inputs; _i < inputs_1.length; _i++) {
                var i = inputs_1[_i];
                var field = i.dataset["field"] || "";
                var dash = field.lastIndexOf("-");
                var idx = parseInt(field.substring(0, dash), 10);
                var key = field.substring(dash + 1);
                if (!ret[idx]) {
                    ret[idx] = { k: "", v: "" };
                }
                switch (key) {
                    case "key":
                        ret[idx].k = i.value.trim();
                        break;
                    case "value":
                        ret[idx].v = i.value.trim();
                        break;
                    case "desc":
                        var desc = i.value.trim();
                        if (desc.length > 0) {
                            ret[idx].desc = desc;
                        }
                        break;
                    default:
                        throw "unknown key [" + key + "]";
                }
            }
            ret = ret.filter(function (x) { return x.k.length > 0; });
            ta.value = json.str(ret);
            request.form.checkEditor(elID.substr(0, elID.lastIndexOf("-")));
            return ret;
        }
    })(editor = request.editor || (request.editor = {}));
})(request || (request = {}));
var request;
(function (request) {
    var editor;
    (function (editor) {
        function initOptionsEditor(el) {
            var parent = el.parentElement;
            parent.appendChild(createOptionsEditor(el));
        }
        editor.initOptionsEditor = initOptionsEditor;
        function createOptionsEditor(el) {
            var opts = json.parse(el.value);
            if (!opts) {
                opts = {};
            }
            return JSX("div", null,
                JSX("div", { class: "uk-margin-top" },
                    JSX("label", { class: "uk-form-label", for: el.id + "-timeout" }, "Timeout"),
                    JSX("input", { class: "uk-input", id: el.id + "-timeout", name: "opt-timeout", type: "number", value: opts.timeout })),
                JSX("div", { class: "uk-margin-top" },
                    JSX("label", { class: "uk-form-label" }, "Ignore"),
                    JSX("div", null,
                        inputCheckbox(el.id, "ignoreRedirects", "Redirects", opts.ignoreRedirects || false),
                        inputCheckbox(el.id, "ignoreReferrer", "Referrer", opts.ignoreReferrer || false),
                        inputCheckbox(el.id, "ignoreCerts", "Certs", opts.ignoreCerts || false),
                        inputCheckbox(el.id, "ignoreCookies", "Cookies", opts.ignoreCookies || false))),
                JSX("div", { class: "uk-margin-top" },
                    JSX("label", { class: "uk-form-label", for: el.id + "-excludeDefaultHeaders" }, "Exclude Default Headers"),
                    JSX("input", { class: "uk-input", id: el.id + "-excludeDefaultHeaders", name: "opt-excludeDefaultHeaders", type: "text", value: opts.excludeDefaultHeaders })),
                JSX("div", { class: "uk-margin-top" },
                    JSX("label", { class: "uk-form-label", for: el.id + "-readCookieJars" }, "Read Cookie Jars"),
                    JSX("input", { class: "uk-input", id: el.id + "-readCookieJars", name: "opt-readCookieJars", type: "text", value: opts.readCookieJars })),
                JSX("div", { class: "uk-margin-top" },
                    JSX("label", { class: "uk-form-label", for: el.id + "writeCookieJar" }, "Write Cookie Jar"),
                    JSX("input", { class: "uk-input", id: el.id + "-writeCookieJar", name: "opt-writeCookieJar", type: "text", value: opts.writeCookieJar })),
                JSX("div", { class: "uk-margin-top" },
                    JSX("label", { class: "uk-form-label", for: el.id + "-sslCert" }, "SSL Cert"),
                    JSX("input", { class: "uk-input", id: el.id + "-sslCert", name: "opt-sslCert", type: "text", value: opts.sslCert })),
                JSX("div", { class: "uk-margin-top" },
                    JSX("label", { class: "uk-form-label", for: el.id + "-userAgentOverride" }, "User Agent Override"),
                    JSX("input", { class: "uk-input", id: el.id + "-userAgentOverride", name: "opt-userAgentOverride", type: "text", value: opts.userAgentOverride })));
        }
        function inputCheckbox(key, prop, title, v) {
            var n = "opt-" + prop;
            var id = key + "-" + prop;
            if (v) {
                return JSX("label", { class: "uk-margin-right" },
                    JSX("input", { type: "checkbox", name: n, value: "true", checked: true }),
                    " ",
                    title);
            }
            else {
                return JSX("label", { class: "uk-margin-right" },
                    JSX("input", { type: "checkbox", name: n, value: "true" }),
                    " ",
                    title);
            }
        }
    })(editor = request.editor || (request.editor = {}));
})(request || (request = {}));
var request;
(function (request) {
    var editor;
    (function (editor) {
        function initQueryParamsEditor(el) {
            var parent = el.parentElement;
            parent.appendChild(createQueryParamsEditor(el));
        }
        editor.initQueryParamsEditor = initQueryParamsEditor;
        function setQueryParams(el, qp) {
            var ret = [];
            if (qp) {
                for (var _i = 0, qp_1 = qp; _i < qp_1.length; _i++) {
                    var p = qp_1[_i];
                    ret.push(encodeURIComponent(p.k) + '=' + encodeURIComponent(p.v));
                }
            }
            var url = new URL(el.value);
            url.search = ret.join("&");
            el.value = url.toString();
        }
        editor.setQueryParams = setQueryParams;
        function updateQueryParams(cache, qp) {
            cache.qp.value = json.str(qp);
            updateFn(cache.qp, dom.req("#" + cache.qp.id + "-ul"));
        }
        editor.updateQueryParams = updateQueryParams;
        function header(id) {
            return JSX("li", null,
                JSX("div", { "data-uk-grid": "" },
                    JSX("div", { class: "uk-width-1-4" }, "Name"),
                    JSX("div", { class: "uk-width-1-4" }, "Value"),
                    JSX("div", { class: "uk-width-1-2" },
                        JSX("div", { class: "right" },
                            JSX("a", { class: style.linkColor, href: "", onclick: "return request.editor.addHeaderRow('" + id + "')", title: "new header" },
                                JSX("span", { "data-uk-icon": "icon: plus" }))),
                        "Description")));
        }
        function updateFn(el, container) {
            var curr = json.parse(el.value);
            container.innerText = "";
            container.appendChild(header(el.id));
            if (curr) {
                for (var idx = 0; idx < curr.length; idx++) {
                    addChild(el.id, idx, container, curr[idx]);
                }
            }
        }
        function createQueryParamsEditor(el) {
            var container = JSX("ul", { id: el.id + "-ul", class: "uk-list uk-list-divider" });
            updateFn(el, container);
            return container;
        }
        function addQueryParamRow(id) {
            var ul = dom.req("#" + id + "-ul");
            var idx = ul.children.length - 1;
            addChild(id, idx, ul, { k: '', v: '' });
            return false;
        }
        editor.addQueryParamRow = addQueryParamRow;
        function removeQueryParamRow(id, el) {
            el.parentElement.parentElement.parentElement.parentElement.remove();
            parseQueryParams(id);
            return false;
        }
        editor.removeQueryParamRow = removeQueryParamRow;
        function addChild(elID, idx, container, h) {
            var ret = JSX("li", null,
                JSX("div", { "data-uk-grid": "" },
                    JSX("div", { class: "uk-width-1-4" },
                        JSX("input", { class: "uk-input", "data-field": idx + "-key", type: "text", value: h.k })),
                    JSX("div", { class: "uk-width-1-4" },
                        JSX("input", { class: "uk-input", "data-field": idx + "-value", type: "text", value: h.v })),
                    JSX("div", { class: "uk-width-1-2" },
                        JSX("div", { class: "right", style: "margin-top: 6px;" },
                            JSX("a", { class: style.linkColor, href: "", onclick: "return request.editor.removeHeaderRow('" + elID + "', this);", title: "new header" },
                                JSX("span", { "data-uk-icon": "icon: close" }))),
                        JSX("input", { style: "width: calc(100% - 48px);", class: "uk-input", "data-field": idx + "-desc", type: "text", value: h.desc }))));
            editor.events(ret, function () { return parseQueryParams(elID); });
            container.appendChild(ret);
        }
        function parseQueryParams(elID) {
            var ta = dom.req("#" + elID);
            var ul = dom.req("#" + elID + "-ul");
            var inputs = dom.els("input", ul);
            var ret = [];
            for (var _i = 0, inputs_2 = inputs; _i < inputs_2.length; _i++) {
                var i = inputs_2[_i];
                var field = i.dataset["field"] || "";
                var dash = field.lastIndexOf("-");
                var idx = parseInt(field.substring(0, dash), 10);
                var key = field.substring(dash + 1);
                if (!ret[idx]) {
                    ret[idx] = { k: "", v: "" };
                }
                switch (key) {
                    case "key":
                        ret[idx].k = i.value.trim();
                        break;
                    case "value":
                        ret[idx].v = i.value.trim();
                        break;
                    case "desc":
                        var desc = i.value.trim();
                        if (desc.length > 0) {
                            ret[idx].desc = desc;
                        }
                        break;
                    default:
                        throw "unknown key [" + key + "]";
                }
            }
            ret = ret.filter(function (x) { return x.k.length > 0; });
            ta.value = json.str(ret);
            setQueryParams(dom.req("#" + elID.replace("queryparams", "url")), ret);
            request.form.checkEditor(elID.substr(0, elID.lastIndexOf("-")));
            return ret;
        }
    })(editor = request.editor || (request.editor = {}));
})(request || (request = {}));
var request;
(function (request) {
    var editor;
    (function (editor) {
        function initURLEditor(el) {
        }
        editor.initURLEditor = initURLEditor;
        function setURL(cache, u) {
            if (!u) {
                cache.qp.value = "[]";
                return;
            }
            editor.updateQueryParams(cache, u.query);
            editor.updateBasicAuth(cache, u.auth);
        }
        editor.setURL = setURL;
    })(editor = request.editor || (request.editor = {}));
})(request || (request = {}));
var request;
(function (request) {
    var form;
    (function (form) {
        function extractRequest(reqID) {
            var key = gv(reqID, "key");
            var title = gv(reqID, "title");
            var desc = gv(reqID, "description");
            var url = gv(reqID, "url");
            var proto = request.urlToPrototype(url);
            proto.method = gv(reqID, "method");
            proto.query = json.parse(gv(reqID, "queryparams"));
            proto.headers = json.parse(gv(reqID, "headers"));
            proto.auth = json.parse(gv(reqID, "auth"));
            proto.body = json.parse(gv(reqID, "body"));
            proto.options = json.parse(gv(reqID, "options"));
            return { key: key, title: title, description: desc, prototype: proto };
        }
        form.extractRequest = extractRequest;
        function checkEditor(reqID) {
            var o = request.getRequest(collection.cache.active, reqID);
            var changed = false;
            if (o) {
                var n = extractRequest(reqID);
                var diff_1 = request.diff(o, n);
                // console.log(o, n, diff);
                changed = diff_1.length > 0;
            }
            else {
                changed = true;
            }
            dom.setDisplay("#save-panel", changed);
        }
        form.checkEditor = checkEditor;
        function saveCurrentRequest(reqID) {
            var req = extractRequest(reqID);
            var msg = { "coll": collection.cache.active, "orig": reqID, "req": req };
            socket.send({ svc: services.request.key, cmd: command.client.saveRequest, param: msg });
        }
        form.saveCurrentRequest = saveCurrentRequest;
        function gv(r, k) {
            return dom.req("#" + r + "-" + k).value;
        }
    })(form = request.form || (request.form = {}));
})(request || (request = {}));
var request;
(function (request) {
    var form;
    (function (form) {
        function renderFormPanel(coll, r) {
            return JSX("div", null,
                JSX("div", { class: "uk-card uk-card-body uk-card-default" },
                    JSX("div", { class: "right" },
                        JSX("a", { class: "theme uk-icon", "data-uk-icon": "close", href: "", onclick: "nav.navigate('/c/" + coll + "');return false;", title: "close request" })),
                    JSX("h3", { class: "uk-card-title" }, r.title ? r.title : r.key),
                    form.renderURL(r),
                    renderSavePanel(r),
                    renderActions(coll, r)),
                JSX("div", { class: "request-editor uk-card uk-card-body uk-card-default uk-margin-top" },
                    JSX("form", { action: "", method: "post", onsubmit: "console.log('XXXXXXX');return false;" }, form.renderSwitcher(r))),
                JSX("div", { class: "request-action uk-card uk-card-body uk-card-default uk-margin-top hidden" }));
        }
        form.renderFormPanel = renderFormPanel;
        function renderDetails(r) {
            return JSX("li", { class: "request-details-panel" },
                JSX("div", { class: "uk-margin-top" },
                    JSX("label", { class: "uk-form-label", for: r.key + "-key" }, "Key"),
                    JSX("input", { class: "uk-input", id: r.key + "-key", name: "key", type: "text", value: r.key || "", "data-lpignore": "true" })),
                JSX("div", { class: "uk-margin-top" },
                    JSX("label", { class: "uk-form-label", for: r.key + "-title" }, "Title"),
                    JSX("input", { class: "uk-input", id: r.key + "-title", name: "title", type: "text", value: r.title || "", "data-lpignore": "true" })),
                JSX("div", { class: "uk-margin-top" },
                    JSX("label", { class: "uk-form-label", for: r.key + "-description" }, "Description"),
                    JSX("textarea", { class: "uk-textarea", id: r.key + "-description", name: "description", "data-lpignore": "true" }, r.description || "")));
        }
        form.renderDetails = renderDetails;
        function reset(r) {
            request.render(collection.cache.active, r);
        }
        form.reset = reset;
        var transforms = {
            "http": "HTTP",
            "json": "JSON",
            "curl": "curl"
        };
        function renderSavePanel(r) {
            return JSX("div", { id: "save-panel", class: "right hidden" },
                JSX("button", { class: "uk-button uk-button-default uk-margin-small-right uk-margin-top", onclick: "request.form.reset('" + r.key + "');" }, "Reset"),
                JSX("button", { class: "uk-button uk-button-default uk-margin-top", onclick: "request.form.saveCurrentRequest('" + r.key + "');" }, "Save Changes"));
        }
        function renderActions(coll, r) {
            var path = "/c/" + coll + "/" + r.key;
            var btnClass = "uk-button uk-button-default uk-margin-small-right uk-margin-top";
            var delWarn = "if (!confirm('Are you sure you want to delete request [" + r.key + "]?')) { return false; }";
            return JSX("div", null,
                nav.link({ path: path + "/call", title: "Call", cls: btnClass, isButton: true }),
                JSX("div", { class: "uk-inline" },
                    JSX("button", { type: "button", class: btnClass }, "Export"),
                    JSX("div", { id: "export-dropdown", "uk-dropdown": "mode: click" },
                        JSX("ul", { class: "uk-list uk-list-divider", style: "margin-bottom: 0;" }, Object.keys(transforms).map(function (k) { return JSX("li", null, nav.link({ path: path + "/transform/" + k, title: transforms[k], onclk: "UIkit.dropdown(dom.req('#export-dropdown')).hide(false);" })); })))),
                nav.link({ path: path + "/delete", title: "Delete", cls: btnClass, onclk: delWarn, isButton: true }));
        }
    })(form = request.form || (request.form = {}));
})(request || (request = {}));
var request;
(function (request) {
    var form;
    (function (form) {
        function renderSwitcher(r) {
            var key = r.key;
            var p = r.prototype;
            return JSX("div", null,
                JSX("ul", { "data-uk-tab": "" },
                    JSX("li", null,
                        JSX("a", { href: "#details" }, "Details")),
                    JSX("li", null,
                        JSX("a", { href: "#query" }, "Query")),
                    JSX("li", null,
                        JSX("a", { href: "#auth" }, "Auth")),
                    JSX("li", null,
                        JSX("a", { href: "#headers" }, "Headers")),
                    JSX("li", null,
                        JSX("a", { href: "#body" }, "Body")),
                    JSX("li", null,
                        JSX("a", { href: "#options" }, "Options"))),
                JSX("ul", { class: "uk-switcher uk-margin" },
                    form.renderDetails(r),
                    renderQueryParams(key, p.query),
                    renderAuth(key, p.auth),
                    renderHeaders(key, p.headers),
                    renderBody(key, p.body),
                    renderOptions(key, p.options)));
        }
        form.renderSwitcher = renderSwitcher;
        function renderQueryParams(key, qp) {
            return JSX("li", { class: "request-queryparams-panel" },
                JSX("div", { class: "uk-margin-top" },
                    JSX("textarea", { class: "uk-textarea hidden", id: key + "-queryparams", name: "queryparams" }, json.str(qp))));
        }
        function renderAuth(key, as) {
            return JSX("li", { class: "request-auth-panel" },
                JSX("div", { class: "uk-margin-top" },
                    JSX("label", { class: "uk-form-label", for: key + "-auth" }, "Auth"),
                    JSX("textarea", { class: "uk-textarea", id: key + "-auth", name: "auth" }, json.str(as))));
        }
        function renderHeaders(key, hs) {
            return JSX("li", { class: "request-headers-panel" },
                JSX("div", { class: "uk-margin-top" },
                    JSX("textarea", { class: "uk-textarea hidden", id: key + "-headers", name: "headers" }, json.str(hs))));
        }
        function renderBody(key, b) {
            return JSX("li", { class: "request-body-panel" },
                JSX("div", { class: "uk-margin-top" },
                    JSX("label", { class: "uk-form-label", for: key + "-body" }, "Body"),
                    JSX("textarea", { class: "uk-textarea", id: key + "-body", name: "body" }, json.str(b))));
        }
        function renderOptions(key, opts) {
            return JSX("li", { class: "request-options-panel" },
                JSX("div", { class: "uk-margin-top" },
                    JSX("textarea", { class: "uk-textarea hidden", id: key + "-options", name: "options" }, json.str(opts))));
        }
    })(form = request.form || (request.form = {}));
})(request || (request = {}));
var request;
(function (request) {
    var form;
    (function (form) {
        function renderURL(r) {
            var call = "nav.navigate(`/c/" + collection.cache.active + "/" + r.key + "/call`);return false;";
            return JSX("div", { class: "uk-margin-top uk-panel" },
                JSX("div", { class: "left", style: "width:120px;" },
                    JSX("select", { class: "uk-select", id: r.key + "-method", name: "method" }, request.allMethods.map(function (m) {
                        if (m.key === r.prototype.method) {
                            return JSX("option", { selected: "selected" }, m.key);
                        }
                        else {
                            return JSX("option", null, m.key);
                        }
                    }))),
                JSX("div", { class: "uk-inline right", style: "width:calc(100% - 120px);" },
                    JSX("a", { class: "uk-form-icon uk-form-icon-flip", href: "", onclick: call, title: "send request", "uk-icon": "icon: play" }),
                    JSX("form", { onsubmit: call },
                        JSX("input", { class: "uk-input", id: r.key + "-url", name: "url", type: "text", value: request.prototypeToURL(r.prototype), "data-lpignore": "true" }))));
        }
        form.renderURL = renderURL;
    })(form = request.form || (request.form = {}));
})(request || (request = {}));
var transform;
(function (transform) {
    function renderRequest(coll, req, format) {
        return JSX("div", { id: coll + "--" + req + "-transform" },
            request.renderActionClose(),
            JSX("div", { class: "transform-title" }, format),
            JSX("div", { class: "transform-result" }));
    }
    transform.renderRequest = renderRequest;
    function setResult(result) {
        var container = dom.req("#" + result.coll + "--" + result.req + "-transform .transform-result");
        dom.setContent(container, render(result));
        log.info("call result [" + result.coll + "/" + result.req + ": " + result.fmt + "] received");
    }
    transform.setResult = setResult;
    function render(r) {
        return JSX("div", { class: "uk-margin-top" },
            JSX("pre", null, r.out));
    }
})(transform || (transform = {}));
var system;
(function (system) {
    var Cache = /** @class */ (function () {
        function Cache() {
        }
        Cache.prototype.getProfile = function () {
            if (!this.profile) {
                throw "no active profile";
            }
            return this.profile;
        };
        Cache.prototype.apply = function (sj) {
            system.cache.profile = sj.profile;
        };
        return Cache;
    }());
    system.cache = new Cache();
})(system || (system = {}));
var system;
(function (system) {
    function onSystemMessage(cmd, param) {
        switch (cmd) {
            case command.server.error:
                console.warn("error from server: " + param);
                break;
            case command.server.connected:
                system.cache.apply(param);
                break;
            default:
                console.warn("unhandled system command [" + cmd + "]");
        }
    }
    system.onSystemMessage = onSystemMessage;
})(system || (system = {}));
var __spreadArrays = (this && this.__spreadArrays) || function () {
    for (var s = 0, i = 0, il = arguments.length; i < il; i++) s += arguments[i].length;
    for (var r = Array(s), k = 0, i = 0; i < il; i++)
        for (var a = arguments[i], j = 0, jl = a.length; j < jl; j++, k++)
            r[k] = a[j];
    return r;
};
var ui;
(function (ui) {
    var bcCls = "uk-navbar-item uk-logo uk-margin-remove uk-padding-remove dynamic";
    function setBreadcrumbs(coll, req, act, extra) {
        var el = dom.req("#breadcrumbs");
        reset(el);
        el.appendChild(nav.link({ path: "/", title: "npn", cls: bcCls }));
        if (coll) {
            el.appendChild(sep());
            el.appendChild(bcFor(coll, "c", coll));
            if (req) {
                el.appendChild(sep());
                el.appendChild(bcFor(req, "c", coll, req));
                if (act) {
                    el.appendChild(sep());
                    el.appendChild(bcFor(act, "c", coll, req, act));
                    if (extra && extra.length > 0) {
                        for (var i = 0; i < extra.length; i++) {
                            el.appendChild(sep());
                            var ret = [coll, req, act];
                            ret.push.apply(ret, extra.slice(0, i));
                            el.appendChild(bcFor.apply(void 0, __spreadArrays([extra[i]], ret)));
                        }
                    }
                }
            }
        }
    }
    ui.setBreadcrumbs = setBreadcrumbs;
    function reset(el) {
        for (var i = el.childElementCount - 1; i >= 0; i--) {
            var e = el.children[i];
            if (e.classList.contains("dynamic")) {
                el.removeChild(e);
            }
        }
    }
    ui.reset = reset;
    function sep() {
        return JSX("span", { class: "uk-navbar-item dynamic", style: "padding: 0 8px;" }, " / ");
    }
    function bcForExtra(coll, req, act, extra) {
        return bcFor(act, "c", coll, req, act);
    }
    function bcFor(title) {
        var parts = [];
        for (var _i = 1; _i < arguments.length; _i++) {
            parts[_i - 1] = arguments[_i];
        }
        var path = parts.map(function (s) { return "/" + s; }).join("");
        return nav.link({ path: path, title: title, cls: bcCls });
    }
})(ui || (ui = {}));
var ui;
(function (ui) {
    function setPanels(coll, req, act, extra) {
        dom.setDisplay("#welcome-panel", coll === undefined);
        dom.setDisplay("#collection-panel", coll !== undefined && coll.length > 0 && req === undefined);
        dom.setDisplay("#request-panel", req !== undefined && req.length > 0);
        ui.setBreadcrumbs(coll, req, act, extra);
        setTitle(coll, req, act);
    }
    ui.setPanels = setPanels;
    function setTitle(coll, req, act) {
        var title = "";
        if (act) {
            title += act + " ";
        }
        if (coll) {
            title += coll;
        }
        if (req) {
            title += "/" + req;
        }
        if (title.length == 0) {
            title = "npn";
        }
        else {
            title = "npn: " + title;
        }
        document.title = title;
    }
})(ui || (ui = {}));
//# sourceMappingURL=npn.js.map