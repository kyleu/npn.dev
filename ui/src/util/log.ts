let debug = true;

// @ts-ignore
// eslint-disable-next-line
function log(level: string, ...msg: any[]): void {
  console.log(`[${level}]`, ...msg);
}

export function setDebug(d: boolean): void {
  debug = d
}

export function isDebug(): boolean {
  return debug;
}

// @ts-ignore
// eslint-disable-next-line
export function logDebug(...msg: any[]): void {
  if (debug) {
    log("debug", ...msg);
  }
}

// @ts-ignore
// eslint-disable-next-line
export function logInfo(...msg: any[]): void {
  log("info", ...msg);
}

// @ts-ignore
// eslint-disable-next-line
export function logWarn(...msg: any[]): void {
  log("warn", ...msg);
}

// @ts-ignore
// eslint-disable-next-line
export function logError(...msg: any[]): void {
  log("error", ...msg);
}
