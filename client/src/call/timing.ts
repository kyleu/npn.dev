namespace call {
  export interface Timing {
    readonly began: number,
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
    readonly group: string;
    readonly start: number;
    readonly end: number;
  }

  export function timingSections(t: Timing): TimingSection[] {
    const ret: TimingSection[] = [];

    const add = function(k: string, g: string, s: number, e: number) {
      ret.push({key: k, group: g, start: s, end: e});
    }

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

  export function timingGraph(t: call.Timing) {
    const msg = "Still busted...";
    const rowHeight = 32
    const sections = timingSections(t);
    const h = (sections.length + 1) * rowHeight;

    let step = t.completed / 10;
    if (step > 1000000) {
      step = 1000000;
    } else if (step > 100000) {
      step = 100000;
    } else if (step > 10000) {
      step = 10000;
    } else if (step > 1000) {
      step = 1000;
    } else if (step > 100) {
      step = 100;
    }

    const secLines: string[] = [];
    for (let idx = step; idx < t.completed; idx += step) {
      secLines.push(`<line x1="${idx}" y1="0" x2="${idx}" y2="${h}" stroke="#666" />`);
    }
    secLines.push(`<line x1="${t.completed - 1}" y1="0" x2="${t.completed - 1}" y2="${h}" stroke="#666" />`);
    const secHTML: string[] = [];
    const f = (x: number) => (x / 1000) + "ms";
    for (let idx = 0; idx < sections.length; idx++) {
      const section = sections[idx]
      const cy = rowHeight * (idx + 1);
      const pc = Math.round(((section.end - section.start) / t.completed) * 10000) / 100;
      secHTML.push(`<rect x="0" y="${cy}" width="${t.completed}" height="${rowHeight}" fill="transparent" />`);
      secHTML.push(`<rect x="${section.start}" y="${cy}" width="${section.end - section.start}" height="${rowHeight}" class="${colorForSection(section.key)}-fill">
        <title>${section.key}: ${pc}%\n${f(section.start)} - ${f(section.end)}</title>
      </rect>`);
    }

    return `<svg height="${h}" width="100%" preserveAspectRatio="none" viewBox="0 0 ${t.completed} ${h}">
      ${secLines.join("\n")}
      ${secHTML.join("\n")}
      ${msg}
    </svg><div class="chart-tooltip"></div>`;
  }

  function colorForSection(key: string) {
    switch (key) {
      case "dns":
        return "bluegrey";
      case "connect":
        return "bluegrey";
      case "tls":
        return "orange";
      case "reqheaders":
        return "green";
      case "reqbody":
        return "green";
      case "rspwait":
        return "blue";
      case "rspheaders":
        return "blue";
      case "rspbody":
        return "blue";
      default:
        return "blue";
    }
  }
}
