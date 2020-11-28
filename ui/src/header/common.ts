import {logDebug} from "@/util/log";

export interface CommonHeader {
  readonly key: string;
  readonly description: string;
  readonly req?: boolean;
  readonly rsp?: boolean;
  readonly link?: string;
}

function nch(key: string, description: string, req: boolean, rsp: boolean, link: string): CommonHeader {
  return {"key": key, "description": description, "req": req, "rsp": rsp, "link": link};
}

function mdnLink(s: string): string {
  return "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/" + s;
}

function snch(key: string, description: string, req: boolean, rsp: boolean): CommonHeader {
  return nch(key, description, req, rsp, mdnLink(key))
}

export const commonHeaders: CommonHeader[] = [
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

let commonHeadersByName: Map<string, CommonHeader>;

export function getCommonHeaderByName(key: string): CommonHeader | undefined {
  if(!commonHeadersByName) {
    commonHeadersByName = new Map<string, CommonHeader>();
    for (const ch of commonHeaders) {
      commonHeadersByName.set(ch.key, ch);
    }
  }
  return commonHeadersByName.get(key);
}
