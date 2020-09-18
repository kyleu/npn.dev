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
        nav.init(function (p) {
            let parts = p.split("/");
            parts = parts.filter(x => x.length > 0);
            console.info("nav handler called, check it out: " + parts.join(" -> "));
            if (parts.length === 0) {
                return; // index
            }
            const svc = parts[0];
            switch (svc) {
                case "c":
                    const collName = parts[1];
                    if (collName !== collection.cache.active) {
                        collection.cache.setActiveCollection(collName);
                        socket.send({ svc: services.collection.key, cmd: command.client.getCollection, param: collName });
                    }
                    if (parts.length > 2) {
                        const reqName = parts[2];
                        if (reqName !== request.cache.active) {
                            request.cache.setActiveRequest(reqName);
                        }
                    }
                    break;
                default:
                    console.info("unhandled svc [" + svc + "]");
            }
        });
        socket.socketConnect(svc, id);
    }
    npn.init = init;
})(npn || (npn = {}));
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
        return JSX("ul", { class: "uk-list uk-list-divider" }, cs.map(renderCollection));
    }
    collection.renderCollections = renderCollections;
    function renderCollection(c) {
        let title = c.title;
        if (!title || c.title.length === 0) {
            title = c.key;
        }
        return JSX("li", null, nav.link("/c/" + c.key, title));
    }
    collection.renderCollection = renderCollection;
})(collection || (collection = {}));
var collection;
(function (collection) {
    function onCollectionMessage(cmd, param) {
        switch (cmd) {
            case command.server.collections:
                collection.cache.collections = param;
                log.info(`processing [${collection.cache.collections.length}] collections`);
                dom.setContent("#collection-list", collection.renderCollections(collection.cache.collections));
                break;
            case command.server.detail:
                const d = param;
                log.info(`processing [${d.requests.length}] requests for collection [${d.collection.key}]`);
                collection.cache.updateCollection(d.collection);
                request.cache.setCollectionRequests(d.collection.key, d.requests);
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
var request;
(function (request) {
    class Cache {
        constructor() {
            this.requests = new Map();
        }
        setCollectionRequests(key, requests) {
            this.requests.set(key, requests);
            if (key === collection.cache.active) {
                dom.setContent("#request-list", request.renderRequests(key, requests));
                for (let req of requests) {
                    if (this.active == req.key) {
                        renderActiveRequest(key, req);
                    }
                }
            }
        }
        setActiveRequest(key) {
            if (!collection.cache.active) {
                console.warn("no active collection");
                return;
            }
            const coll = collection.cache.active;
            const reqs = this.requests.get(coll) || [];
            this.active = key;
            for (let req of reqs) {
                if (req.key == key) {
                    renderActiveRequest(coll, req);
                }
            }
        }
    }
    function renderActiveRequest(key, req) {
        dom.setContent("#active-request", request.renderRequest(key, req));
    }
    request.cache = new Cache();
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
    function renderRequests(coll, rs) {
        return JSX("ul", { class: "uk-list uk-list-divider" }, rs.map(r => renderRequestLink(coll, r)));
    }
    request.renderRequests = renderRequests;
    function renderRequestLink(coll, r) {
        let title = r.title;
        if (!title || r.title.length === 0) {
            title = r.key;
        }
        return JSX("div", null, nav.link("/c/" + coll + "/" + r.key, title));
    }
    request.renderRequestLink = renderRequestLink;
    function renderRequest(coll, r) {
        return renderPrototype(r.prototype);
    }
    request.renderRequest = renderRequest;
    function renderPrototype(p) {
        return JSX("div", null, request.prototypeToURL(p));
    }
})(request || (request = {}));
var request;
(function (request) {
    function prototypeToURLParts(p) {
        const ret = [];
        let push = function (t, v) {
            ret.push({ t: t, v: v });
        };
        push("protocol", p.protocol);
        push("", "://");
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
    request.prototypeToURLParts = prototypeToURLParts;
    function prototypeToURL(p) {
        return prototypeToURLParts(p).map(x => x.v).join("");
    }
    request.prototypeToURL = prototypeToURL;
})(request || (request = {}));
var request;
(function (request) {
    var form;
    (function (form) {
        function initAuthEditor(el) {
        }
        form.initAuthEditor = initAuthEditor;
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
        form.setAuth = setAuth;
        function updateBasicAuth(cache, auth) {
            let currentAuth = [];
            try {
                currentAuth = JSON.parse(cache.auth.value);
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
            cache.auth.value = JSON.stringify(currentAuth, null, 2);
        }
        form.updateBasicAuth = updateBasicAuth;
    })(form = request.form || (request.form = {}));
})(request || (request = {}));
var request;
(function (request) {
    var form;
    (function (form) {
        function initBodyEditor(el) {
        }
        form.initBodyEditor = initBodyEditor;
        function setBody(cache, body) {
        }
        form.setBody = setBody;
    })(form = request.form || (request.form = {}));
})(request || (request = {}));
var request;
(function (request) {
    var form;
    (function (form) {
        function wireForm(prefix) {
            const id = function (k) {
                return "#" + prefix + "-" + k;
            };
            const cache = {
                url: dom.req(id("url")),
                auth: dom.req(id("auth")),
                qp: dom.req(id("queryparams")),
                headers: dom.req(id("headers")),
                body: dom.req(id("body"))
            };
            initEditors(prefix, cache);
            wireEvents(cache);
        }
        form.wireForm = wireForm;
        function initEditors(prefix, cache) {
            form.initURLEditor(cache.url);
            form.initAuthEditor(cache.auth);
            form.initQueryParamsEditor(cache.qp);
            form.initHeadersEditor(cache.headers);
            form.initBodyEditor(cache.body);
            form.initOptionsEditor(prefix);
        }
        function events(e, f) {
            e.onchange = f;
            e.onkeyup = f;
            e.onblur = f;
        }
        function wireEvents(cache) {
            events(cache.url, function () {
                form.setURL(cache, request.prototypeFromURL(cache.url.value));
            });
            events(cache.auth, function () {
                let auth;
                try {
                    auth = JSON.parse(cache.auth.value);
                }
                catch (e) {
                    console.log("invalid auth JSON [" + cache.auth.value + "]");
                    auth = [];
                }
                form.setAuth(cache, auth);
            });
            events(cache.qp, function () {
                let qp;
                try {
                    qp = JSON.parse(cache.qp.value);
                }
                catch (e) {
                    console.log("invalid qp JSON [" + cache.qp.value + "]");
                    qp = [];
                }
                form.setQueryParams(cache, qp);
            });
            events(cache.headers, function () {
                let h;
                try {
                    h = JSON.parse(cache.headers.value);
                }
                catch (e) {
                    console.log("invalid headers JSON [" + cache.headers.value + "]");
                    h = [];
                }
                form.setHeaders(cache, h);
            });
            events(cache.body, function () {
                let b;
                try {
                    b = JSON.parse(cache.body.value);
                }
                catch (e) {
                    console.log("invalid body JSON [" + cache.body.value + "]");
                }
                form.setBody(cache, b);
            });
        }
    })(form = request.form || (request.form = {}));
})(request || (request = {}));
var request;
(function (request) {
    var form;
    (function (form) {
        function createHeadersEditor(el) {
            const container = JSX("ul", { id: el.id + "-ul", class: "uk-list uk-list-divider" });
            const header = JSX("li", null,
                JSX("div", { "uk-grid": true },
                    JSX("div", { class: "uk-width-1-4" }, "Name"),
                    JSX("div", { class: "uk-width-1-4" }, "Value"),
                    JSX("div", { class: "uk-width-1-2" },
                        JSX("div", { class: "right" },
                            JSX("a", { class: style.linkColor, href: "", onclick: "request.form.addChild(dom.req('#" + el.id + "-ul" + "'), {k: '', v: ''});return false;", title: "new header" },
                                JSX("span", { "data-uk-icon": "icon: plus" }))),
                        "Description")));
            const updateFn = function () {
                const curr = JSON.parse(el.value);
                container.innerText = "";
                container.appendChild(header);
                for (let h of curr) {
                    addChild(container, h);
                }
            };
            updateFn();
            return container;
        }
        form.createHeadersEditor = createHeadersEditor;
        function addChild(container, h) {
            console.info(container);
            container.appendChild(JSX("li", null,
                JSX("div", { "uk-grid": true },
                    JSX("div", { class: "uk-width-1-4" }, h.k),
                    JSX("div", { class: "uk-width-1-4" }, h.v),
                    JSX("div", { class: "uk-width-1-2" },
                        JSX("div", { class: "right" },
                            JSX("a", { class: style.linkColor, href: "", onclick: "return false;", title: "new header" },
                                JSX("span", { "data-uk-icon": "icon: close" }))),
                        h.desc ? h.desc : ""))));
        }
        form.addChild = addChild;
    })(form = request.form || (request.form = {}));
})(request || (request = {}));
var request;
(function (request) {
    var form;
    (function (form) {
        function initHeadersEditor(el) {
            const parent = el.parentElement;
            parent.appendChild(form.createHeadersEditor(el));
        }
        form.initHeadersEditor = initHeadersEditor;
        function setHeaders(cache, headers) {
        }
        form.setHeaders = setHeaders;
    })(form = request.form || (request.form = {}));
})(request || (request = {}));
var request;
(function (request) {
    var form;
    (function (form) {
        function initOptionsEditor(prefix) {
        }
        form.initOptionsEditor = initOptionsEditor;
    })(form = request.form || (request.form = {}));
})(request || (request = {}));
var request;
(function (request) {
    var form;
    (function (form) {
        function initQueryParamsEditor(el) {
        }
        form.initQueryParamsEditor = initQueryParamsEditor;
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
        form.setQueryParams = setQueryParams;
        function updateQueryParams(cache, qp) {
            cache.qp.value = JSON.stringify(qp, null, 2);
        }
        form.updateQueryParams = updateQueryParams;
    })(form = request.form || (request.form = {}));
})(request || (request = {}));
var request;
(function (request) {
    var form;
    (function (form) {
        function initURLEditor(el) {
        }
        form.initURLEditor = initURLEditor;
        function setURL(cache, u) {
            if (!u) {
                cache.qp.value = "[]";
                return;
            }
            form.updateQueryParams(cache, u.query);
            form.updateBasicAuth(cache, u.auth);
        }
        form.setURL = setURL;
    })(form = request.form || (request.form = {}));
})(request || (request = {}));
var socket;
(function (socket) {
    const debug = true;
    let sock;
    let connected = false;
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
        sock.onmessage = (event) => onSocketMessage(JSON.parse(event.data));
        sock.onerror = (event) => npn.onError("socket", event.type);
        sock.onclose = onSocketClose;
    }
    socket.socketConnect = socketConnect;
    function send(msg) {
        if (connected) {
            if (debug) {
                console.debug("out", msg);
            }
            const m = JSON.stringify(msg, null, 2);
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
        pendingMessages.forEach(send);
        pendingMessages = [];
        // send({ svc: services.system.key, cmd: command.client.connect, param: currentID });
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
            default:
                console.warn(`unhandled message for service [${msg.svc}]`);
        }
    }
    socket.onSocketMessage = onSocketMessage;
    function onSocketClose() {
        function disconnect(seconds) {
            connected = false;
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
            case command.server.connected:
                system.cache.apply(param);
                break;
            default:
                console.warn(`unhandled system command [${cmd}]`);
        }
    }
    system.onSystemMessage = onSystemMessage;
})(system || (system = {}));
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
        getCollection: "getCollection"
    };
    command.server = {
        pong: "pong",
        connected: "connected",
        collections: "collections",
        detail: "detail",
        error: "error"
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
var log;
(function (log) {
    const started = Date.now();
    function info(msg) {
        const el = l("info", msg);
        const container = dom.req("#log-panel");
        container.appendChild(el);
    }
    log.info = info;
    function l(level, msg) {
        const n = Date.now() - started;
        return JSX("li", null,
            JSX("div", { class: "right" },
                n,
                "ms"),
            msg);
    }
    log.l = l;
})(log || (log = {}));
var nav;
(function (nav) {
    let handler = function (p) {
        console.info("default nav handler called: " + p);
    };
    function init(f) {
        handler = f;
        window.onpopstate = function (event) {
            f(event.state === null ? "" : event.state);
        };
        let path = location.pathname;
        if (path.startsWith("/w")) {
            path = path.substr(2);
        }
        navigate(path);
    }
    nav.init = init;
    function navigate(path) {
        if (path.startsWith("/")) {
            path = path.substr(1);
        }
        let fullpath = "/w";
        if (path.length > 0) {
            fullpath = fullpath + "/" + path;
        }
        if (location.pathname !== fullpath) {
            history.pushState(path, "", fullpath);
        }
        handler(path);
    }
    nav.navigate = navigate;
    function link(path, title) {
        return JSX("a", { class: style.linkColor, href: path, onclick: "nav.navigate('" + path + "', '" + title + "');return false;" }, title);
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