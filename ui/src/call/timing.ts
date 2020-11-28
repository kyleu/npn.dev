export interface Timing {
  readonly began: number;
  readonly dnsStart: number;
  readonly dnsEnd: number;
  readonly connectStart: number;
  readonly connectEnd: number;
  readonly tlsStart?: number;
  readonly tlsEnd?: number;
  readonly wroteHeaders: number;
  readonly wroteRequest: number;
  readonly firstResponseByte: number;
  readonly responseHeaders: number;
  readonly completed: number;
}

export interface TimingSection {
  readonly key: string;
  readonly group: string;
  readonly start: number;
  readonly end: number;
}

export function timingSections(t: Timing): TimingSection[] {
  const ret: TimingSection[] = [];

  const add = (k: string, g: string, s?: number, e?: number): void => {
    if (s && e) {
      ret.push({key: k, group: g, start: s, end: e});
    }
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

export function timingGraph(url: string, ts: TimingSection[]): string {
  const ret: string[] = [];
  for (const t of ts) {
    if (t.group.length > 0) {
      ret.push(encodeURIComponent(t.key + ".g") + '=' + encodeURIComponent(t.group));
    }
    ret.push(encodeURIComponent(t.key + ".s") + '=' + encodeURIComponent(t.start));
    ret.push(encodeURIComponent(t.key + ".e") + '=' + encodeURIComponent(t.end));
  }
  const TODO = undefined;
  ret.push("t=" + TODO || "light");

  if (url.endsWith("/")) {
    url = url.substr(0, url.length - 1);
  }

  return url + "/svg/gantt?" + ret.join("&");
}
