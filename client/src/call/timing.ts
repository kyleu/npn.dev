namespace call {
  export interface Timing {
    readonly dnsStart: number,
    readonly dnsEnd: number,
    readonly connectStart: number,
    readonly connectEnd: number,
    readonly tlsStart?: number,
    readonly tlsEnd?: number,
    readonly wroteHeaders: number,
    readonly wroteRequest: number,
    readonly firstResponseByte: number,
    readonly responseHeaders: number,
    readonly completed: number
  }

  export interface TimingSection {
    readonly key: string;
    readonly start: number;
    readonly end: number;
  }

  export function timingSections(t: Timing): TimingSection[] {
    const ret: TimingSection[] = [];

    const add = function(k: string, s: number, e: number) {
      ret.push({key: k, start: s, end: e});
    }

    add("dns", t.dnsStart, t.dnsEnd);
    add("connect", t.connectStart, t.connectEnd);

    let cc = t.connectEnd;
    if ((t.tlsEnd || 0) > 0) {
      cc = t.tlsEnd || 0;
      add("tls", t.tlsStart || 0, cc);
    }

    add("reqheaders", cc, t.wroteHeaders);
    add("reqbody", t.wroteHeaders, t.wroteRequest);
    add("rspwait", t.wroteRequest, t.firstResponseByte);
    add("rspheaders", t.firstResponseByte, t.responseHeaders);
    add("rspbody", t.responseHeaders, t.completed);
    return ret;
  }
}
