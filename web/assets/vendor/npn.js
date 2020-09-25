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
        log.init();
        window.onbeforeunload = () => {
            socket.setAppUnloading();
        };
        nav.init(socket.route);
        socket.socketConnect(svc, id);
    }
    npn.init = init;
    function debug() {
        const dump = (k, v) => {
            console.warn(`${k}: ${v}`);
        };
        dump("Active Collection", collection.cache.active);
        dump("Active Request", request.cache.active);
        dump("Active Action", `${request.cache.action} [${request.cache.extra}]`);
    }
    npn.debug = debug;
    function testbed() {
        log.info("Testbed!");
    }
    npn.testbed = testbed;
})(npn || (npn = {}));
var body;
(function (body) {
    function renderBody(b) {
        if (!b) {
            return JSX("div", null, "No body");
        }
        switch (b.type) {
            case "json":
                return renderJSON(b.config);
            case "html":
                return renderHTML(b.config);
            default:
                return JSX("div", null,
                    "TODO: ",
                    b.type);
        }
    }
    body.renderBody = renderBody;
    function renderHTML(h) {
        return JSX("div", null,
            JSX("em", null, "HTML"),
            JSX("pre", { style: "overflow: auto; max-height: 720px;" }, h.content));
    }
    function renderJSON(j) {
        return JSX("div", null,
            JSX("em", null, "JSON"),
            JSX("pre", null, json.str(j.msg)));
    }
})(body || (body = {}));
var body;
(function (body) {
    body.AllTypes = [
        { key: "error", title: "Error", hidden: true },
        { key: "form", title: "Form", hidden: false },
        { key: "html", title: "HTML", hidden: false },
        { key: "json", title: "JSON", hidden: false },
        { key: "large", title: "Large File", hidden: false },
        { key: "raw", title: "Raw", hidden: true }
    ];
})(body || (body = {}));
var call;
(function (call) {
    function prepare(coll, r) {
        const param = { "coll": coll, "req": r.key };
        socket.send({ svc: services.request.key, cmd: command.client.requestCall, param: param });
    }
    call.prepare = prepare;
    function setResult(param) {
        const result = param;
        const container = dom.req(`#${result.collection}--${result.request}-call`);
        dom.setContent(container, call.renderResult(result));
    }
    call.setResult = setResult;
    function help(x) {
        console.log(x);
        let title = "";
        switch (x.label) {
            default:
                title = x.label;
        }
        return title + ": " + (x.labelVal / 1000) + "ms";
    }
})(call || (call = {}));
var call;
(function (call) {
    function renderResult(r) {
        var _a;
        const statusEl = JSX("div", null,
            r.status,
            ": ",
            (((_a = r.timing) === null || _a === void 0 ? void 0 : _a.completed) || 0) / 1000,
            "ms");
        return [
            JSX("div", { class: "right" },
                JSX("a", { class: "theme uk-icon", "data-uk-icon": "close", href: "", onclick: "nav.pop();return false;", title: "close result" })),
            section("Result", statusEl),
            JSX("hr", null),
            JSX("div", null, renderHeaders("Final Request Headers", r.requestHeaders)),
            JSX("hr", null),
            ...renderResponse(r.response),
            JSX("hr", null),
            section("Timing", renderTiming(r.timing))
        ];
    }
    call.renderResult = renderResult;
    function renderResponse(r) {
        if (!r) {
            return [JSX("div", null, "No response")];
        }
        return [
            section("Status", r.status),
            JSX("hr", null),
            section("Protocol", r.proto),
            JSX("hr", null),
            JSX("div", null, renderHeaders("Response Headers", r.headers)),
            JSX("hr", null),
            section("Content", `${r.contentType} (${r.contentLength} bytes)`),
            JSX("hr", null),
            section("Body", body.renderBody(r.body)),
        ];
    }
    function renderHeaders(title, headers) {
        if (!headers) {
            return section(title, "No headers");
        }
        return JSX("div", { class: "uk-overflow-auto" },
            JSX("h4", null, title),
            JSX("table", { class: "uk-table uk-table-divider uk-text-left uk-table-small uk-table-justify" },
                JSX("tbody", null, headers.map(h => JSX("tr", { title: h.desc },
                    JSX("td", { class: "uk-text-nowrap" }, h.k),
                    JSX("td", { class: "uk-text-nowrap" }, h.v))))));
    }
    function renderTiming(t) {
        if (!t) {
            return JSX("div", null, "No timing");
        }
        const sections = call.timingSections(t);
        return JSX("div", { class: "timing-panel" }, sections.map(sc => JSX("div", null,
            sc.key,
            ": ",
            sc.start,
            " - ",
            sc.end)));
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
        const ret = [];
        const add = function (k, g, s, e) {
            ret.push({ key: k, group: g, start: s, end: e });
        };
        add("dns", "connect", t.dnsStart, t.dnsEnd);
        add("connect", "connect", t.connectStart, t.connectEnd);
        let cc = t.connectEnd;
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
})(call || (call = {}));
var collection;
(function (collection_1) {
    class Cache {
        updateCollection(collection) {
            // TODO
        }
        setActiveCollection(key) {
            this.active = key;
        }
    }
    collection_1.cache = new Cache();
})(collection || (collection = {}));
var collection;
(function (collection) {
    function renderCollections(cs) {
        return cs.map(renderCollectionLink);
    }
    collection.renderCollections = renderCollections;
    function renderCollectionLink(c) {
        let title = c.title;
        if (!title || c.title.length === 0) {
            title = c.key;
        }
        return JSX("div", { class: "nav-item" }, nav.link("/c/" + c.key, title));
    }
    function renderCollection(coll, requests) {
        const cn = coll.title ? coll.title : coll.key;
        return JSX("div", null,
            JSX("div", { class: "uk-card uk-card-body uk-card-default uk-margin-top" },
                JSX("div", { class: "right" },
                    JSX("a", { class: "theme uk-icon", "data-uk-icon": "close", href: "", onclick: "nav.pop();return false;", title: "close collection" })),
                JSX("h3", { class: "uk-card-title" }, cn),
                JSX("p", null, coll.description || "")),
            JSX("div", { class: "uk-card uk-card-body uk-card-default uk-margin-top" },
                JSX("h3", { class: "uk-card-title" }, "Requests"),
                JSX("form", null,
                    JSX("input", { class: "uk-input", placeholder: "add a request by url" })),
                JSX("div", { id: "request-list", class: "uk-margin-top" }, renderRequests(coll.key, requests))));
    }
    collection.renderCollection = renderCollection;
    function renderRequests(coll, rs) {
        return JSX("ul", { class: "uk-list uk-list-divider" }, rs.map(r => renderRequestLink(coll, r)));
    }
    function renderRequestLink(coll, r) {
        let title = r.title;
        if (!title || r.title.length === 0) {
            title = r.key;
        }
        return JSX("li", null, nav.link("/c/" + coll + "/" + r.key, title));
    }
})(collection || (collection = {}));
var collection;
(function (collection) {
    function onCollectionMessage(cmd, param) {
        switch (cmd) {
            case command.server.collections:
                collection.cache.collections = param;
                if (collection.cache.collections) {
                    log.info(`processing [${collection.cache.collections.length}] collections`);
                    dom.els(".collection-list").forEach(el => {
                        dom.setContent(el, collection.renderCollections(collection.cache.collections));
                    });
                }
                break;
            case command.server.collectionDetail:
                const d = param;
                log.info(`processing [${d.requests.length}] requests for collection [${d.collection.key}]`);
                collection.cache.updateCollection(d.collection);
                request.cache.setCollectionRequests(d.collection, d.requests);
                break;
            default:
                console.warn(`unhandled collection command [${cmd}]`);
        }
    }
    collection.onCollectionMessage = onCollectionMessage;
})(collection || (collection = {}));
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
        if (Array.isArray(e)) {
            e.forEach(x => el.appendChild(x));
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
                if (child === undefined || child === null) {
                    throw `child array for tag [${tag}] is ${child}\n${e.outerHTML}`;
                }
                if (c === undefined || c === null) {
                    throw `child for tag [${tag}] is ${c}\n${e.outerHTML}`;
                }
                if (typeof c === "string") {
                    c = document.createTextNode(c);
                }
                e.appendChild(c);
            });
        }
        else if (child === undefined || child === null) {
            throw `child for tag [${tag}] is ${child}\n${e.outerHTML}`;
            // debugger;
            // child = document.createTextNode("NULL!");
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
        style.linkColor = `${color}-fg`;
        dom.els(".theme").forEach(el => {
            el.classList.add(style.linkColor);
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
var flash;
(function (flash) {
    function wire() {
        setTimeout(fadeOut, 4000);
    }
    flash.wire = wire;
    function fadeOut() {
        let matched = false;
        dom.els(".alert-top").forEach(el => {
            matched = true;
            el.classList.add("uk-animation-fade", "uk-animation-reverse");
        });
        if (matched) {
            setTimeout(remove, 1000);
        }
    }
    function remove() {
        dom.els(".alert-top").forEach(el => {
            el.remove();
        });
    }
})(flash || (flash = {}));
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
        input.onblur = () => {
            valueEl.innerText = input.value;
            dom.setDisplay(valueEl, true);
            dom.setDisplay(editorEl, false);
            updateEditor(el.parentElement);
        };
        input.onkeypress = (e) => {
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
    header.commonHeadersByName = new Map();
    for (const ch of header.commonHeaders) {
        header.commonHeadersByName.set(ch.key, ch);
    }
    function dumpCommonHeaders() {
        const dump = (title, req, rsp) => {
            let matched = false;
            console.log("\n::: " + title + " Headers");
            header.commonHeaders.forEach(ch => {
                if (ch.req == req && ch.rsp == rsp) {
                    matched = true;
                    console.log(`${ch.key}: ${ch.link}`);
                    console.log(`  - ${ch.description}`);
                }
            });
            if (!matched) {
                console.log("none");
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
    function renderActionEmpty(r) {
        return JSX("div", null);
    }
    request.renderActionEmpty = renderActionEmpty;
    function renderActionUnknown(key, extra, r) {
        return JSX("div", null,
            renderClose(r),
            "unknown action: ",
            key,
            " (",
            extra,
            ")");
    }
    request.renderActionUnknown = renderActionUnknown;
    function renderActionCall(coll, r) {
        return JSX("div", { id: coll + "--" + r.key + "-call" },
            renderClose(r),
            JSX("div", { class: "call-title" }, "Loading..."),
            JSX("div", { class: "call-result" }));
    }
    request.renderActionCall = renderActionCall;
    function renderClose(r) {
        return JSX("div", { class: "right" },
            JSX("a", { class: "theme uk-icon", "data-uk-icon": "close", href: "", onclick: "nav.navigate(`/c/${collection.cache.active}/${request.cache.active}`);return false;", title: "close collection" }));
    }
})(request || (request = {}));
var request;
(function (request) {
    class Cache {
        constructor() {
            this.requests = new Map();
            this.extra = [];
        }
        setCollectionRequests(coll, requests) {
            this.requests.set(coll.key, requests);
            if (coll.key === collection.cache.active) {
                dom.setContent("#collection-panel", collection.renderCollection(coll, requests));
                for (let req of requests) {
                    if (this.active === req.key) {
                        renderActiveRequest(collection.cache.active, req);
                        if (this.action) {
                            renderActiveAction(collection.cache.active, req, this.action, this.extra);
                        }
                    }
                }
            }
        }
        setActiveRequest(key) {
            if (!collection.cache.active) {
                console.warn("no active collection");
                return;
            }
            if (this.active !== key) {
                this.active = key;
                if (this.active) {
                    const r = getActiveRequest();
                    if (r) {
                        renderActiveRequest(collection.cache.active, r);
                    }
                }
                else {
                }
            }
        }
        setActiveAction(act, extra) {
            if (!collection.cache.active) {
                console.warn("no active collection");
                return;
            }
            const sameExtra = this.extra.length === extra.length && this.extra.every(function (value, index) { return value === extra[index]; });
            if (this.active && (this.action !== act || !sameExtra)) {
                this.action = act;
                this.extra = extra;
                const r = getActiveRequest();
                if (r) {
                    renderActiveAction(collection.cache.active, r, this.action, this.extra);
                }
            }
        }
    }
    function renderActiveRequest(coll, req) {
        dom.setContent("#request-panel", request.form.renderFormPanel(coll, req));
        request.editor.wireForm(req.key);
    }
    function renderActiveAction(coll, req, action, extra) {
        log.info(`new action: ${action} (${extra})`);
        const re = dom.req(".request-editor");
        const ra = dom.req(".request-action");
        switch (action) {
            case undefined:
                dom.setContent(ra, request.renderActionEmpty(req));
                break;
            case "call":
                call.prepare(coll, req);
                dom.setContent(ra, request.renderActionCall(coll, req));
                break;
            default:
                console.warn("unhandled request action [" + action + "]");
                dom.setContent(ra, request.renderActionUnknown(action, extra, req));
        }
        dom.setDisplay(re, action === undefined);
        dom.setDisplay(ra, action !== undefined);
    }
    function getActiveRequest() {
        const coll = collection.cache.active;
        if (!coll) {
            return undefined;
        }
        for (let req of request.cache.requests.get(coll) || []) {
            if (req.key === request.cache.active) {
                return req;
            }
        }
        return undefined;
    }
    request.cache = new Cache();
})(request || (request = {}));
var request;
(function (request) {
    const MethodGet = { "key": "GET", "description": "" };
    const MethodHead = { "key": "HEAD", "description": "" };
    const MethodPost = { "key": "POST", "description": "" };
    const MethodPut = { "key": "PUT", "description": "" };
    const MethodPatch = { "key": "PATCH", "description": "" };
    const MethodDelete = { "key": "DELETE", "description": "" };
    const MethodConnect = { "key": "CONNECT", "description": "" };
    const MethodOptions = { "key": "OPTIONS", "description": "" };
    const MethodTrace = { "key": "TRACE", "description": "" };
    request.allMethods = [MethodGet, MethodHead, MethodPost, MethodPut, MethodPatch, MethodDelete, MethodConnect, MethodOptions, MethodTrace];
})(request || (request = {}));
var request;
(function (request) {
    function newPrototype(protocol, hostname, port, path, qp, fragment, auth) {
        if (protocol.endsWith(":")) {
            protocol = protocol.substr(0, protocol.length - 1);
        }
        if (fragment.startsWith("#")) {
            fragment = fragment.substr(1);
        }
        return { method: "get", protocol: protocol, domain: hostname, port: port, path: path, query: qp, fragment: fragment, auth: auth };
    }
    function prototypeFromURL(u) {
        const url = new URL(u);
        const qp = [];
        for (const [k, v] of url.searchParams) {
            qp.push({ k: k, v: v });
        }
        const auth = [];
        if (url.username.length > 0) {
            auth.push({ type: "basic", config: { "username": url.username, "password": url.password, "showPassword": true } });
        }
        let port;
        if (url.port.length > 0) {
            port = parseInt(url.port);
        }
        return newPrototype(url.protocol, url.hostname, port, url.pathname, qp, url.hash, auth);
    }
    request.prototypeFromURL = prototypeFromURL;
})(request || (request = {}));
var request;
(function (request) {
    function prototypeToURL(p) {
        return prototypeToURLParts(p).map(x => x.v).join("");
    }
    request.prototypeToURL = prototypeToURL;
    function prototypeToHTML(p) {
        return JSX("span", null, prototypeToURLParts(p).map(x => JSX("span", { title: x.t, class: urlColor(x.t) }, x.v)));
    }
    request.prototypeToHTML = prototypeToHTML;
    function prototypeToURLParts(p) {
        const ret = [];
        let push = (t, v) => {
            ret.push({ t: t, v: v });
        };
        push("protocol", p.protocol);
        push("", "://");
        if (p.auth) {
            for (let a of p.auth) {
                if (a.type === "basic") {
                    const cfg = a.config;
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
            var query = p.query.map(k => encodeURIComponent(k.k) + '=' + encodeURIComponent(k.v)).join('&');
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
    function onRequestMessage(cmd, param) {
        switch (cmd) {
            case command.server.callResult:
                call.setResult(param);
                break;
            default:
                console.warn(`unhandled request command [${cmd}]`);
        }
    }
    request.onRequestMessage = onRequestMessage;
})(request || (request = {}));
var request;
(function (request) {
    var editor;
    (function (editor) {
        function initAuthEditor(el) {
        }
        editor.initAuthEditor = initAuthEditor;
        function setAuth(cache, auth) {
            const url = new URL(cache.url.value);
            let u = "";
            let p = "";
            if (auth) {
                for (let a of auth) {
                    if (a.type === "basic") {
                        const basic = a.config;
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
            let currentAuth = [];
            try {
                currentAuth = json.parse(cache.auth.value);
            }
            catch (e) {
                console.log("invalid auth JSON [" + cache.auth.value + "]");
            }
            let matched = -1;
            if (!currentAuth) {
                currentAuth = [];
            }
            for (let i = 0; i < currentAuth.length; i++) {
                const x = currentAuth[i];
                if (x.type === "basic") {
                    matched = i;
                }
            }
            let basic;
            if (auth) {
                for (let i = 0; i < auth.length; i++) {
                    const x = auth[i];
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
                    let curr = currentAuth[matched].config;
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
            const parent = el.parentElement;
            parent.appendChild(createBodyEditor(el));
        }
        editor.initBodyEditor = initBodyEditor;
        function createBodyEditor(el) {
            const b = json.parse(el.value);
            return JSX("div", { class: "uk-margin-top" },
                JSX("select", { class: "uk-select" },
                    JSX("option", { value: "" }, "No body"),
                    body.AllTypes.filter(t => !t.hidden).map(t => {
                        if (b && b.type === t.key) {
                            return JSX("option", { value: t.key, selected: "selected" }, t.title);
                        }
                        else {
                            return JSX("option", { value: t.key }, t.title);
                        }
                    }),
                    "\u02D9"),
                body.AllTypes.filter(t => !t.hidden).map(t => {
                    let cfg = (b && b.type == t.key) ? b.config : null;
                    return configEditor(t.key, cfg, t.key === (b ? b.type : ""));
                }));
        }
        function configEditor(key, config, active) {
            let cls = "uk-margin-top body-editor-" + key;
            if (!active) {
                cls += " hidden";
            }
            switch (key) {
                case "json":
                    const j = config;
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
            const id = (k) => {
                return "#" + prefix + "-" + k;
            };
            const cache = {
                url: dom.req(id("url")),
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
            e.onchange = f;
            e.onkeyup = f;
            e.onblur = f;
        }
        function wireEvents(cache) {
            events(cache.url, function () {
                editor.setURL(cache, request.prototypeFromURL(cache.url.value));
            });
            events(cache.auth, function () {
                let auth;
                try {
                    auth = json.parse(cache.auth.value);
                }
                catch (e) {
                    console.log("invalid auth JSON [" + cache.auth.value + "]");
                    auth = [];
                }
                editor.setAuth(cache, auth);
            });
            events(cache.qp, function () {
                let qp;
                try {
                    qp = json.parse(cache.qp.value);
                }
                catch (e) {
                    console.log("invalid qp JSON [" + cache.qp.value + "]");
                    qp = [];
                }
                editor.setQueryParams(cache, qp);
            });
            events(cache.headers, function () {
                let h;
                try {
                    h = json.parse(cache.headers.value);
                }
                catch (e) {
                    console.log("invalid headers JSON [" + cache.headers.value + "]");
                    h = [];
                }
                editor.setHeaders(cache, h);
            });
            events(cache.body, function () {
                let b;
                try {
                    b = json.parse(cache.body.value);
                }
                catch (e) {
                    console.log("invalid body JSON [" + cache.body.value + "]");
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
            const parent = el.parentElement;
            parent.appendChild(createHeadersEditor(el));
        }
        editor.initHeadersEditor = initHeadersEditor;
        function setHeaders(cache, headers) {
        }
        editor.setHeaders = setHeaders;
        function createHeadersEditor(el) {
            const container = JSX("ul", { id: el.id + "-ul", class: "uk-list uk-list-divider" });
            const header = JSX("li", null,
                JSX("div", { "data-uk-grid": "" },
                    JSX("div", { class: "uk-width-1-4" }, "Name"),
                    JSX("div", { class: "uk-width-1-4" }, "Value"),
                    JSX("div", { class: "uk-width-1-2" },
                        JSX("div", { class: "right" },
                            JSX("a", { class: style.linkColor, href: "", onclick: "request.editor.addChild(dom.req('#" + el.id + "-ul" + "'), {k: '', v: ''});return false;", title: "new header" },
                                JSX("span", { "data-uk-icon": "icon: plus" }))),
                        "Description")));
            const updateFn = () => {
                const curr = json.parse(el.value);
                container.innerText = "";
                container.appendChild(header);
                if (curr) {
                    for (let h of curr) {
                        addChild(container, h);
                    }
                }
            };
            updateFn();
            return container;
        }
        function addChild(container, h) {
            container.appendChild(JSX("li", null,
                JSX("div", { "data-uk-grid": "" },
                    JSX("div", { class: "uk-width-1-4" }, h.k),
                    JSX("div", { class: "uk-width-1-4" }, h.v),
                    JSX("div", { class: "uk-width-1-2" },
                        JSX("div", { class: "right" },
                            JSX("a", { class: style.linkColor, href: "", onclick: "return false;", title: "new header" },
                                JSX("span", { "data-uk-icon": "icon: close" }))),
                        h.desc ? h.desc : ""))));
        }
        editor.addChild = addChild;
    })(editor = request.editor || (request.editor = {}));
})(request || (request = {}));
var request;
(function (request) {
    var editor;
    (function (editor) {
        function initOptionsEditor(el) {
            const parent = el.parentElement;
            parent.appendChild(createOptionsEditor(el));
        }
        editor.initOptionsEditor = initOptionsEditor;
        function createOptionsEditor(el) {
            let opts = json.parse(el.value);
            if (!opts) {
                opts = {};
            }
            return JSX("div", null,
                JSX("div", { class: "uk-margin-top" },
                    JSX("label", { class: "uk-form-label", for: el.id + "-timeout" }, "Timeout"),
                    JSX("input", { class: "uk-input", id: el.id + "-timeout", name: "opt-timeout", type: "number", value: opts.timeout })),
                JSX("div", { class: "uk-margin-top" },
                    JSX("label", { class: "uk-form-label", for: el.id + "-ignoreRedirects" }, "Ignore Redirects"),
                    inputBool(el.id, "ignoreRedirects", opts.ignoreRedirects || false)),
                JSX("div", { class: "uk-margin-top" },
                    JSX("label", { class: "uk-form-label", for: "<%= key %>-ignoreReferrer" }, "Ignore Referrer"),
                    inputBool(el.id, "ignoreReferrer", opts.ignoreReferrer || false)),
                JSX("div", { class: "uk-margin-top" },
                    JSX("label", { class: "uk-form-label", for: el.id + "-ignoreCerts" }, "Ignore Certs"),
                    inputBool(el.id, "ignoreCerts", opts.ignoreCerts || false)),
                JSX("div", { class: "uk-margin-top" },
                    JSX("label", { class: "uk-form-label", for: el.id + "-ignoreCookies" }, "Ignore Cookies"),
                    inputBool(el.id, "ignoreCookies", opts.ignoreCookies || false)),
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
        function inputBool(key, prop, v) {
            const n = "opt-" + prop;
            const id = key + "-" + prop;
            if (v) {
                return JSX("div", null,
                    JSX("label", { class: "uk-margin-small-right" },
                        JSX("input", { class: "uk-radio", type: "radio", name: n, value: "true", checked: true }),
                        " True"),
                    JSX("label", null,
                        JSX("input", { class: "uk-radio", type: "radio", name: n, value: "false" }),
                        " False"));
            }
            else {
                return JSX("div", null,
                    JSX("label", { class: "uk-margin-small-right" },
                        JSX("input", { class: "uk-radio", type: "radio", name: n, value: "true" }),
                        " True"),
                    JSX("label", null,
                        JSX("input", { class: "uk-radio", type: "radio", name: n, value: "false", checked: true }),
                        " False"));
            }
        }
    })(editor = request.editor || (request.editor = {}));
})(request || (request = {}));
var request;
(function (request) {
    var editor;
    (function (editor) {
        function initQueryParamsEditor(el) {
        }
        editor.initQueryParamsEditor = initQueryParamsEditor;
        function setQueryParams(cache, qp) {
            let ret = [];
            if (qp) {
                for (let p of qp) {
                    ret.push(encodeURIComponent(p.k) + '=' + encodeURIComponent(p.v));
                }
            }
            const url = new URL(cache.url.value);
            url.search = ret.join("&");
            cache.url.value = url.toString();
        }
        editor.setQueryParams = setQueryParams;
        function updateQueryParams(cache, qp) {
            cache.qp.value = json.str(qp);
        }
        editor.updateQueryParams = updateQueryParams;
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
        function renderFormPanel(coll, r) {
            return JSX("div", null,
                JSX("div", { class: "uk-card uk-card-body uk-card-default" },
                    JSX("div", { class: "right" },
                        JSX("a", { class: "theme uk-icon", "data-uk-icon": "close", href: "", onclick: "nav.navigate('/c/" + coll + "');return false;", title: "close request" })),
                    JSX("h3", { class: "uk-card-title" }, r.title ? r.title : r.key),
                    form.renderURL(r),
                    renderActions(coll, r)),
                JSX("div", { class: "request-editor uk-card uk-card-body uk-card-default uk-margin-top" },
                    form.renderSwitcher(r),
                    JSX("div", { class: "uk-margin-top hidden" },
                        JSX("button", { class: "right uk-button uk-button-default uk-margin-top", type: "submit" }, "Save Changes"))),
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
                    JSX("input", { class: "uk-input", id: r.key + "-description", name: "description", type: "text", value: r.description || "", "data-lpignore": "true" })));
        }
        form.renderDetails = renderDetails;
        const transforms = {
            "http": "HTTP",
            "json": "JSON",
            "curl": "curl"
        };
        function renderActions(coll, r) {
            const path = "/c/" + coll + "/" + r.key;
            return JSX("div", { class: "uk-margin-top" },
                nav.link(path + "/call", "Call", "uk-button uk-button-default uk-margin-small-right", "", true),
                JSX("div", { class: "uk-inline" },
                    JSX("button", { type: "button", class: "uk-button uk-button-default uk-margin-small-right" }, "Export"),
                    JSX("div", { id: "export-dropdown", "uk-dropdown": "mode: click" },
                        JSX("ul", { class: "uk-list uk-list-divider", style: "margin-bottom: 0;" }, Object.keys(transforms).map(k => JSX("li", null, nav.link(path + "/transform/" + k, transforms[k], "", "UIkit.dropdown(dom.req('#export-dropdown')).hide(false);")))))),
                nav.link(path + "/delete", "Delete", "uk-button uk-button-default uk-margin-small-right", "if (!confirm('Are you sure you want to delete request [" + r.key + "]?')) { return false; }", true));
        }
    })(form = request.form || (request.form = {}));
})(request || (request = {}));
var request;
(function (request) {
    var form;
    (function (form) {
        function renderSwitcher(r) {
            const key = r.key;
            const p = r.prototype;
            return JSX("div", null,
                JSX("ul", { "data-uk-tab": "" },
                    JSX("li", null,
                        JSX("a", { href: "#" }, "Details")),
                    JSX("li", null,
                        JSX("a", { href: "#" }, "Query")),
                    JSX("li", null,
                        JSX("a", { href: "#" }, "Auth")),
                    JSX("li", null,
                        JSX("a", { href: "#" }, "Headers")),
                    JSX("li", null,
                        JSX("a", { href: "#" }, "Body")),
                    JSX("li", null,
                        JSX("a", { href: "#" }, "Options"))),
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
                    JSX("label", { class: "uk-form-label", for: key + "-queryparams" }, "Query Params"),
                    JSX("textarea", { class: "uk-textarea", id: key + "-queryparams", name: "queryparams" }, json.str(qp))));
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
                    JSX("label", { class: "uk-form-label", for: key + "-headers" }, "Headers"),
                    JSX("textarea", { class: "uk-textarea", id: key + "-headers", name: "headers" }, json.str(hs))));
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
                    JSX("label", { class: "uk-form-label", for: key + "-options" }, "Options"),
                    JSX("textarea", { class: "uk-textarea", id: key + "-options", name: "options" }, json.str(opts))));
        }
    })(form = request.form || (request.form = {}));
})(request || (request = {}));
var request;
(function (request) {
    var form;
    (function (form) {
        function renderURL(r) {
            const click = "nav.navigate(`/c/" + collection.cache.active + "/" + r.key + "/call`);return false;";
            return JSX("div", { class: "uk-margin-top uk-panel" },
                JSX("div", { class: "left", style: "width:120px;" },
                    JSX("select", { class: "uk-select", id: r.key + "-method", name: "method" }, request.allMethods.map(m => {
                        if (m.key === r.prototype.method) {
                            return JSX("option", { selected: "selected" }, m.key);
                        }
                        else {
                            return JSX("option", null, m.key);
                        }
                    }))),
                JSX("div", { class: "uk-inline right", style: "width:calc(100% - 120px);" },
                    JSX("a", { class: "uk-form-icon uk-form-icon-flip", href: "", onclick: click, "uk-icon": "icon: refresh" }),
                    JSX("input", { class: "uk-input", id: r.key + "-url", name: "url", type: "text", value: request.prototypeToURL(r.prototype), "data-lpignore": "true" })));
        }
        form.renderURL = renderURL;
    })(form = request.form || (request.form = {}));
})(request || (request = {}));
var socket;
(function (socket) {
    function route(p) {
        let parts = p.split("/");
        parts = parts.filter(x => x.length > 0);
        console.info("nav: " + parts.join(" -> "));
        if (parts.length === 0 || parts[0].length === 0) {
            ui.setPanels(undefined, undefined, undefined, []);
            return; // index
        }
        const svc = parts[0];
        switch (svc) {
            case "c":
                let coll = (parts.length > 1 && parts[1].length > 0) ? parts[1] : undefined;
                let req = (parts.length > 2 && parts[2].length > 0) ? parts[2] : undefined;
                let act = (parts.length > 3 && parts[3].length > 0) ? parts[3] : undefined;
                let extra = (parts.length > 4) ? parts.slice(4) : [];
                if (coll !== collection.cache.active) {
                    collection.cache.setActiveCollection(coll);
                    socket.send({ svc: services.collection.key, cmd: command.client.getCollection, param: coll });
                }
                request.cache.setActiveRequest(req);
                request.cache.setActiveAction(act, extra);
                ui.setPanels(coll, req, act, extra);
                break;
            default:
                console.info("unhandled svc [" + svc + "]");
        }
    }
    socket.route = route;
})(socket || (socket = {}));
var socket;
(function (socket) {
    const debug = true;
    let sock;
    let connected = false;
    let pauseSeconds = 0;
    let appUnloading = false;
    let pendingMessages = [];
    let currentService = "";
    let currentID = "";
    let connectTime;
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
    socket.setAppUnloading = setAppUnloading;
    function socketConnect(svc, id) {
        currentService = svc;
        currentID = id;
        connectTime = Date.now();
        sock = new WebSocket(socketUrl());
        sock.onopen = onSocketOpen;
        sock.onmessage = (event) => onSocketMessage(json.parse(event.data));
        sock.onerror = (event) => npn.onError("socket", event.type);
        sock.onclose = onSocketClose;
    }
    socket.socketConnect = socketConnect;
    function send(msg) {
        if (connected) {
            if (debug) {
                console.debug("out", msg);
            }
            const m = json.str(msg);
            sock.send(m);
        }
        else {
            pendingMessages.push(msg);
        }
    }
    socket.send = send;
    function onSocketOpen() {
        log.info("socket connected");
        connected = true;
        pauseSeconds = 1;
        pendingMessages.forEach(send);
        pendingMessages = [];
    }
    function onSocketMessage(msg) {
        if (debug) {
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
                console.warn(`unhandled message for service [${msg.svc}]`);
        }
    }
    socket.onSocketMessage = onSocketMessage;
    function onSocketClose() {
        function disconnect() {
            connected = false;
            const elapsed = Date.now() - connectTime;
            if (elapsed < 2000) {
                pauseSeconds = pauseSeconds * 2;
                if (debug) {
                    console.info(`socket closed immediately, reconnecting in ${pauseSeconds} seconds`);
                }
                setTimeout(() => {
                    socketConnect(currentService, currentID);
                }, pauseSeconds * 1000);
            }
            else {
                log.info("socket closed after [" + elapsed + "ms]");
                socketConnect(currentService, currentID);
            }
        }
        if (!appUnloading) {
            disconnect();
        }
    }
})(socket || (socket = {}));
var system;
(function (system) {
    class Cache {
        getProfile() {
            if (!this.profile) {
                throw "no active profile";
            }
            return this.profile;
        }
        apply(sj) {
            system.cache.profile = sj.profile;
        }
    }
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
                console.warn(`unhandled system command [${cmd}]`);
        }
    }
    system.onSystemMessage = onSystemMessage;
})(system || (system = {}));
var ui;
(function (ui) {
    function setBreadcrumbs(coll, req, act, extra) {
        const el = dom.req("#breadcrumbs");
        reset(el);
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
                        for (let i = 0; i < extra.length; i++) {
                            el.appendChild(sep());
                            const ret = [coll, req, act];
                            ret.push(...extra.slice(0, i));
                            el.appendChild(bcFor(extra[i], ...ret));
                        }
                    }
                }
            }
        }
    }
    ui.setBreadcrumbs = setBreadcrumbs;
    function reset(el) {
        for (let i = el.childElementCount - 1; i >= 0; i--) {
            const e = el.children[i];
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
    function bcFor(title, ...parts) {
        const path = parts.map(s => "/" + s).join("");
        return nav.link(path, title, "uk-navbar-item uk-logo uk-margin-remove uk-padding-remove dynamic");
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
        let title = "";
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
    function setPicture(p) {
        dom.setValue("#self-picture-input", p);
        return false;
    }
    profile.setPicture = setPicture;
})(profile || (profile = {}));
var command;
(function (command) {
    command.client = {
        ping: "ping",
        connect: "connect",
        getCollections: "getCollections",
        getCollection: "getCollection",
        requestCall: "requestCall"
    };
    command.server = {
        error: "error",
        pong: "pong",
        connected: "connected",
        collections: "collections",
        collectionDetail: "collectionDetail",
        callResult: "callResult"
    };
})(command || (command = {}));
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
var group;
(function (group_1) {
    class Group {
        constructor(key) {
            this.members = [];
            this.key = key;
        }
    }
    group_1.Group = Group;
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
    group_1.GroupSet = GroupSet;
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
    group_1.groupBy = groupBy;
    function findGroup(groups, key) {
        for (const g of groups) {
            if (g.key === key) {
                return g.members;
            }
        }
        return [];
    }
    group_1.findGroup = findGroup;
    function flatten(a) {
        const ret = [];
        a.forEach(v => ret.push(...v));
        return ret;
    }
    group_1.flatten = flatten;
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
    let started = 0;
    let content;
    let list;
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
        const n = Date.now() - started;
        const el = JSX("li", { class: color(level) },
            JSX("div", { class: "right" },
                n,
                "ms"),
            msg);
        if (!list) {
            list = dom.req("#log-list");
        }
        list.appendChild(el);
        if (!content) {
            content = dom.req("#log-content");
        }
        content.scrollTo(0, content.scrollHeight);
    }
    log.l = l;
    function toggle() {
        const wsc = dom.req("#workspace-content");
        const lp = dom.req("#log-container");
        const curr = (lp.style.display !== "") && (lp.style.display !== "none");
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
var nav;
(function (nav) {
    let handler = (p) => {
        console.info("default nav handler called: " + p);
    };
    function init(f) {
        handler = f;
        window.onpopstate = (event) => {
            f(event.state === null ? "" : event.state);
        };
        let path = location.pathname;
        navigate(path);
    }
    nav.init = init;
    function navigate(path) {
        if (path.startsWith("/")) {
            path = path.substr(1);
        }
        let locPath = location.pathname;
        if (locPath.startsWith("/")) {
            locPath = locPath.substr(1);
        }
        if (locPath !== path) {
            history.pushState(path, "", "/" + path);
        }
        handler(path);
    }
    nav.navigate = navigate;
    function pop() {
        let p = location.pathname.substr(0, location.pathname.lastIndexOf("/"));
        if (p === '/c') {
            p = "";
        }
        navigate(p);
    }
    nav.pop = pop;
    function navActiveRequest() {
        navigate(`/c/${collection.cache.active}/${request.cache.active}`);
    }
    nav.navActiveRequest = navActiveRequest;
    function link(path, title, cls, onclk, isButton) {
        let href = path;
        if (!href.startsWith("/")) {
            href = "/" + href;
        }
        if (cls) {
            cls = " " + cls.trim();
        }
        else {
            cls = "";
        }
        if (onclk) {
            if (!onclk.endsWith(";")) {
                onclk += ";";
            }
        }
        else {
            onclk = "";
        }
        if (!isButton) {
            cls = style.linkColor + cls;
        }
        return JSX("a", { class: cls, href: href, onclick: onclk + "nav.navigate('" + path + "', '" + title + "');return false;" }, title);
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
        const m = UIkit.modal(key);
        if (!m) {
            console.warn(`no modal available with key [${key}]`);
        }
        return m;
    }
    notify_1.modal = modal;
})(notify || (notify = {}));
var services;
(function (services) {
    services.system = { key: "system", title: "System", plural: "systems", icon: "close" };
    services.collection = { key: "collection", title: "Collection", plural: "Collections", icon: "folder" };
    services.request = { key: "request", title: "Request", plural: "Requests", icon: "file-text" };
    const allServices = [services.system, services.collection];
    function fromKey(key) {
        const ret = allServices.find(s => s.key === key);
        if (!ret) {
            throw `invalid service [${key}]`;
        }
        return ret;
    }
    services.fromKey = fromKey;
})(services || (services = {}));
//# sourceMappingURL=npn.js.map