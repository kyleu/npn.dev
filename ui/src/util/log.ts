import {profileRef} from "@/user/profile";

let debug = true;
let pub = false;

const maxLogs = 32;
let logList: HTMLElement | null;

export interface LogMessage {
  level: string;
  message: string;
  context: {[key: string]: unknown};
}

function colorForLevel(l: string): string {
  const isDark = profileRef.value?.settings.mode === "dark";
  switch(l) {
    case "debug":
      return "#999999";
    case "info":
      return "inherit";
    case "warn":
      return "#3e606f";
    case "error":
      return "#bc0022";
    default:
      return isDark ? "#330000": "#ff0000";
  }
}

function elementFor(level: string, s: string, ...params: unknown[]): HTMLElement {
  const el = document.createElement("li");
  el.style.color = colorForLevel(level);

  const lEl = document.createElement("div");
  lEl.innerText = `[${level}]`;
  lEl.classList.add("log-level");
  el.appendChild(lEl);

  const tEl = document.createElement("div");
  tEl.innerText = new Date().toLocaleTimeString();
  tEl.classList.add("log-timestamp");
  el.appendChild(tEl);

  const cEl = document.createElement("div");
  cEl.innerText = s;
  el.appendChild(cEl);

  if (params !== undefined && params.length > 0) {
    for (const p of params) {
      if (p !== undefined) {
        const pEl = document.createElement("div");
        pEl.innerText = JSON.stringify(p, null, 2);
        el.appendChild(pEl);
      }
    }
  }

  return el;
}

function log(level: string, msg: string, ...params: unknown[]): void {
  if (level !== "debug") {
    if (!logList) {
      logList = document.getElementById("log-list");
    }
    if (logList) {
      const li = elementFor(level, msg, ...params);
      while (logList.children.length > maxLogs) {
        logList.removeChild(logList.children[0]);
      }
      logList.appendChild(li);
      li.scrollIntoView();
    }
  }

  console.log(`[${level}]`, msg, ...params);
}

export function onLog(l: LogMessage): void {
  log(l.level, l.message, l.context);
}

export function setDebug(d: boolean): void {
  debug = d;
}

export function isDebug(): boolean {
  return debug;
}

export function setPublic(d: boolean): void {
  pub = d;
}

export function isPublic(): boolean {
  return pub;
}

export function logDebug(msg: string, ...params: unknown[]): void {
  if (debug) {
    log("debug", msg, ...params);
  }
}

export function logInfo(msg: string, ...params: unknown[]): void {
  log("info", msg, ...params);
}

export function logWarn(msg: string, ...params: unknown[]): void {
  log("warn", msg, ...params);
}

export function logError(msg: string, ...params: unknown[]): void {
  log("error", msg, ...params);
}

export function logToggle(): void {
  const el = document.getElementById("log-container");
  if (el) {
    el.style.display = el.style.display === "block" ? "none" : "block";
  }
}
