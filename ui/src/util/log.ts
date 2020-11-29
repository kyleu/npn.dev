import {profileRef} from "@/user/profile";

const maxLogs = 32;
let debug = false;
let logList: HTMLElement | null;

export interface LogMessage {
  level: string;
  message: string;
  context: {[key: string]: string};
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

// @ts-ignore
// eslint-disable-next-line
function elementFor(level: string, s: string, ...params: any[]): HTMLElement {
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

  return el;
}

// @ts-ignore
// eslint-disable-next-line
function log(level: string, msg: string, ...params: any[]): void {
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

// @ts-ignore
// eslint-disable-next-line
export function logDebug(msg: string, ...params: any[]): void {
  if (debug) {
    log("debug", msg, ...params);
  }
}

// @ts-ignore
// eslint-disable-next-line
export function logInfo(msg: string, ...params: any[]): void {
  log("info", msg, ...params);
}

// @ts-ignore
// eslint-disable-next-line
export function logWarn(msg: string, ...params: any[]): void {
  log("warn", msg, ...params);
}

// @ts-ignore
// eslint-disable-next-line
export function logError(msg: string, ...params: any[]): void {
  log("error", msg, ...params);
}
