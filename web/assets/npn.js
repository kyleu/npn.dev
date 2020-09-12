"use strict";
var npn;
(function (npn) {
    function onError(svc, err) {
        console.error(`${svc}: ${err}`);
        const idx = err.lastIndexOf(":");
        if (idx > -1) {
            err = err.substr(idx + 1);
        }
        notify.notify(`${svc} error: ${err}`, false);
    }
    npn.onError = onError;
    function init(svc, id) {
        window.onbeforeunload = () => {
            socket.setAppUnloading();
        };
        socket.socketConnect(svc, id);
    }
    npn.init = init;
})(npn || (npn = {}));
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
    }
    dom.initDom = initDom;
    function els(selector, context) {
        let result;
        if (context) {
            result = context.querySelectorAll(selector);
        }
        else {
            result = document.querySelectorAll(selector);
        }
        const ret = [];
        result.forEach(v => {
            ret.push(v);
        });
        return ret;
    }
    dom.els = els;
    function opt(selector, context) {
        const e = els(selector, context);
        switch (e.length) {
            case 0:
                return undefined;
            case 1:
                return e[0];
            default:
                console.warn(`found [${e.length}] elements with selector [${selector}], wanted zero or one`);
        }
    }
    dom.opt = opt;
    function req(selector, context) {
        const res = opt(selector, context);
        if (!res) {
            console.warn(`no element found for selector [${selector}]`);
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
    function setDisplay(el, condition, v = "block") {
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
        el.appendChild(e);
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
            text.style.height = `${text.scrollHeight < 64 ? 64 : text.scrollHeight + 6}px`;
        }
        function delayedResize() {
            window.setTimeout(resize, 0);
        }
        const x = text.dataset["autoresize"];
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
        categories.forEach(c => {
            const opt = document.createElement("option");
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
        for (let i = 0; i < el.children.length; i++) {
            const e = el.children.item(i);
            e.selected = e.value === o;
        }
    }
    dom.setSelectOption = setSelectOption;
    function insertAtCaret(e, text) {
        if (e.selectionStart || e.selectionStart === 0) {
            let startPos = e.selectionStart;
            let endPos = e.selectionEnd;
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
    const e = document.createElement(tag);
    for (const name in attrs) {
        if (name && attrs.hasOwnProperty(name)) {
            const v = attrs[name];
            if (name === "dangerouslySetInnerHTML") {
                dom.setHTML(e, v["__html"]);
            }
            else if (v === true) {
                e.setAttribute(name, name);
            }
            else if (v !== false && v !== null && v !== undefined) {
                e.setAttribute(name, v.toString());
            }
        }
    }
    for (let i = 2; i < arguments.length; i++) {
        let child = arguments[i];
        if (Array.isArray(child)) {
            child.forEach(c => {
                e.appendChild(c);
            });
        }
        else if (child === undefined || child === null) {
            throw `child for tag [${tag}] is ${child}`;
        }
        else {
            if (!child.nodeType) {
                child = document.createTextNode(child.toString());
            }
            e.appendChild(child);
        }
    }
    return e;
}
var style;
(function (style) {
    function setTheme(theme) {
        wireEmoji(theme);
        const card = dom.els(".uk-card");
        switch (theme) {
            case "auto":
                let t = "light";
                if (window.matchMedia && window.matchMedia("(prefers-color-scheme: dark)").matches) {
                    t = "dark";
                }
                setTheme(t);
                fetch("/profile/theme/" + t).then(r => r.text()).then(() => {
                    // console.log(`Set theme to [${t}]`);
                });
                break;
            case "light":
                document.documentElement.classList.remove("uk-light");
                document.body.classList.remove("uk-light");
                document.documentElement.classList.add("uk-dark");
                document.body.classList.add("uk-dark");
                card.forEach(x => {
                    x.classList.add("uk-card-default");
                    x.classList.remove("uk-card-secondary");
                });
                break;
            case "dark":
                document.documentElement.classList.add("uk-light");
                document.body.classList.add("uk-light");
                document.documentElement.classList.remove("uk-dark");
                document.body.classList.remove("uk-dark");
                card.forEach(x => {
                    x.classList.remove("uk-card-default");
                    x.classList.add("uk-card-secondary");
                });
                break;
            default:
                console.warn("invalid theme");
                break;
        }
    }
    style.setTheme = setTheme;
    function themeLinks(color) {
        dom.els(".theme").forEach(el => {
            el.classList.add(`${color}-fg`);
        });
    }
    style.themeLinks = themeLinks;
    function wireEmoji(t) {
        if (typeof EmojiButton === "undefined") {
            dom.els(".picker-toggle").forEach(el => dom.setDisplay(el, false));
            return;
        }
        const opts = { position: "bottom-end", theme: t, zIndex: 1021 };
        dom.els(".textarea-emoji").forEach(el => {
            const toggle = dom.req(".picker-toggle", el);
            toggle.addEventListener("click", () => {
                const textarea = dom.req(".uk-textarea", el);
                const picker = new EmojiButton(opts);
                picker.on("emoji", (emoji) => {
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
        dom.els(".drop").forEach(el => {
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
        const el = e.target;
        const key = el.dataset["key"] || "";
        let t = el.dataset["t"] || "";
        const f = events.getOpenEvent(key);
        if (f) {
            f(t);
        }
        else {
            console.warn(`no drop open handler registered for [${key}]`);
        }
    }
    function onDropHide(e) {
        if (!e.target) {
            return;
        }
        const el = e.target;
        if (el.classList.contains("uk-open")) {
            const key = el.dataset["key"] || "";
            const t = el.dataset["t"] || "";
            const f = events.getCloseEvent(key);
            if (f) {
                f(t);
            }
        }
    }
    let emojiPicked = false;
    function onEmojiPicked() {
        emojiPicked = true;
        setTimeout(() => (emojiPicked = false), 200);
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
    let openEvents = new Map();
    let closeEvents = new Map();
    function register(key, o, c) {
        if (!o) {
            o = () => { };
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
var modal;
(function (modal) {
    let activeParam;
    function wire() {
        dom.els(".modal").forEach(el => {
            el.addEventListener("show", onModalOpen);
            el.addEventListener("hide", onModalHide);
        });
    }
    modal.wire = wire;
    function open(key, param) {
        activeParam = param;
        const m = notify.modal(`#modal-${key}`);
        m.show();
        return false;
    }
    modal.open = open;
    function openSoon(key) {
        setTimeout(() => open(key), 0);
    }
    modal.openSoon = openSoon;
    function hide(key) {
        const m = notify.modal(`#modal-${key}`);
        const el = m.$el;
        if (el.classList.contains("uk-open")) {
            m.hide();
        }
    }
    modal.hide = hide;
    function onModalOpen(e) {
        if (!e.target) {
            return;
        }
        const el = e.target;
        if (el.id.indexOf("modal") !== 0) {
            return;
        }
        const key = el.id.substr("modal-".length);
        const f = events.getOpenEvent(key);
        if (f) {
            f(activeParam);
        }
        else {
            console.warn(`no modal open handler registered for [${key}]`);
        }
        activeParam = undefined;
    }
    function onModalHide(e) {
        if (!e.target) {
            return;
        }
        const el = e.target;
        if (el.classList.contains("uk-open")) {
            const key = el.id.substr("modal-".length);
            const f = events.getCloseEvent(key);
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
        dom.els(".tag-editor").forEach(el => {
            el.addEventListener("moved", onTagEditorUpdate);
            el.addEventListener("added", onTagEditorUpdate);
            el.addEventListener("removed", onTagEditorUpdate);
        });
    }
    tags.wire = wire;
    function removeTag(el) {
        const itemEl = el.parentElement;
        const editorEl = itemEl.parentElement;
        itemEl.remove();
        updateEditor(editorEl);
    }
    tags.removeTag = removeTag;
    function addTag(el) {
        const editorEl = el.parentElement;
        if (!editorEl) {
            return;
        }
        const itemEl = tags.renderItem();
        editorEl.insertBefore(itemEl, dom.req(".add-item", editorEl));
        editTag(itemEl);
    }
    tags.addTag = addTag;
    function editTag(el) {
        const valueEl = dom.req(".value", el);
        const editorEl = dom.req(".editor", el);
        dom.setDisplay(valueEl, false);
        dom.setDisplay(editorEl, true);
        const input = tags.renderInput(valueEl.innerText);
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
        const el = e.target;
        updateEditor(el);
    }
    function updateEditor(el) {
        const key = el.dataset["key"] || "";
        const f = events.getOpenEvent(key);
        if (f) {
            f();
        }
        else {
            console.warn(`no tag open handler registered for [${key}]`);
        }
        const ret = dom.els(".item", el).map(e => e.innerText);
        dom.setValue(`#model-${key}-input`, ret.join(","));
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
            a.map(s => JSX("span", { class: "item" }, s)),
            JSX("div", { class: "clear" }));
    }
    tags.renderTagsView = renderTagsView;
})(tags || (tags = {}));
var socket;
(function (socket_1) {
    const debug = true;
    let socket;
    let appUnloading = false;
    let currentService = "";
    let currentID = "";
    function socketUrl() {
        const l = document.location;
        let protocol = "ws";
        if (l.protocol === "https:") {
            protocol = "wss";
        }
        return protocol + `://${l.host}/s`;
    }
    function setAppUnloading() {
        appUnloading = true;
    }
    socket_1.setAppUnloading = setAppUnloading;
    function socketConnect(svc, id) {
        // system.cache.currentService = svc;
        // system.cache.currentID = id;
        // system.cache.connectTime = Date.now();
        socket = new WebSocket(socketUrl());
        socket.onopen = () => {
            send({ svc: svc, cmd: "connect", param: id });
        };
        socket.onmessage = (event) => {
            const msg = JSON.parse(event.data);
            onSocketMessage(msg);
        };
        socket.onerror = (event) => {
            // rituals.onError(services.system, event.type);
        };
        socket.onclose = () => {
            onSocketClose();
        };
    }
    socket_1.socketConnect = socketConnect;
    function send(msg) {
        if (debug) {
            console.debug("out", msg);
        }
        socket.send(JSON.stringify(msg));
    }
    socket_1.send = send;
    function onSocketMessage(msg) {
        if (debug) {
            console.debug("in", msg);
        }
        switch (msg.svc) {
            default:
                console.warn(`unhandled message for service [${msg.svc}]`);
        }
    }
    socket_1.onSocketMessage = onSocketMessage;
    function onSocketClose() {
        function disconnect(seconds) {
            if (debug) {
                console.info(`socket closed, reconnecting in ${seconds} seconds`);
            }
            setTimeout(() => {
                socketConnect(currentService, currentID);
            }, seconds * 1000);
        }
        if (!appUnloading) {
            disconnect(10);
        }
    }
})(socket || (socket = {}));
var profile;
(function (profile) {
    // noinspection JSUnusedGlobalSymbols
    function setNavColor(el, c) {
        dom.setValue("#nav-color", c);
        const nb = dom.req("#navbar");
        nb.className = `${c}-bg uk-navbar-container uk-navbar`;
        const colors = document.querySelectorAll(".nav_swatch");
        colors.forEach(function (i) {
            i.classList.remove("active");
        });
        el.classList.add("active");
    }
    profile.setNavColor = setNavColor;
    // noinspection JSUnusedGlobalSymbols
    function setLinkColor(el, c) {
        dom.setValue("#link-color", c);
        const links = dom.els(".profile-link");
        links.forEach(l => {
            l.classList.forEach(x => {
                if (x.indexOf("-fg") > -1) {
                    l.classList.remove(x);
                }
                l.classList.add(`${c}-fg`);
            });
        });
        const colors = document.querySelectorAll(".link_swatch");
        colors.forEach(function (i) {
            i.classList.remove("active");
        });
        el.classList.add("active");
    }
    profile.setLinkColor = setLinkColor;
})(profile || (profile = {}));
var collection;
(function (collection) {
    class Group {
        constructor(key) {
            this.members = [];
            this.key = key;
        }
    }
    collection.Group = Group;
    class GroupSet {
        constructor() {
            this.groups = [];
        }
        findOrInsert(key) {
            const ret = this.groups.find(x => x.key === key);
            if (ret) {
                return ret;
            }
            const n = new Group(key);
            this.groups.push(n);
            return n;
        }
    }
    collection.GroupSet = GroupSet;
    function groupBy(list, func) {
        const res = new GroupSet();
        if (list) {
            list.forEach(o => {
                const group = res.findOrInsert(func(o));
                group.members.push(o);
            });
        }
        return res;
    }
    collection.groupBy = groupBy;
    function findGroup(groups, key) {
        for (const g of groups) {
            if (g.key === key) {
                return g.members;
            }
        }
        return [];
    }
    collection.findGroup = findGroup;
    function flatten(a) {
        const ret = [];
        a.forEach(v => ret.push(...v));
        return ret;
    }
    collection.flatten = flatten;
})(collection || (collection = {}));
var date;
(function (date) {
    function dateToYMD(dt) {
        const d = dt.getDate();
        const m = dt.getMonth() + 1;
        const y = dt.getFullYear();
        return `${y}-${m <= 9 ? `0${m}` : m}-${d <= 9 ? `0${d}` : d}`;
    }
    date.dateToYMD = dateToYMD;
    function dateFromYMD(s) {
        const d = new Date(s);
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
        return `${toDateString(d)} ${toTimeString(d)}`;
    }
    date.toDateTimeString = toDateTimeString;
    const tzOffset = new Date().getTimezoneOffset() * 60000;
    function utcDate(s) {
        return new Date(Date.parse(s) + tzOffset);
    }
    date.utcDate = utcDate;
})(date || (date = {}));
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
        const m = UIkit.modal(key);
        if (!m) {
            console.warn(`no modal available with key [${key}]`);
        }
        return m;
    }
    notify_1.modal = modal;
})(notify || (notify = {}));
//# sourceMappingURL=npn.js.map