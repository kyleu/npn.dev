namespace request {
  export interface QueryParam {
    readonly k: string;
    readonly v: string;
    readonly desc?: string;
  }

  export interface Options {
    readonly timeout?: number;
    readonly ignoreRedirects?: boolean;
    readonly ignoreReferrer?: boolean;
    readonly ignoreCerts?: boolean;
    readonly ignoreCookies?: boolean;
    readonly excludeDefaultHeaders?: string[];
    readonly readCookieJars?: string[];
    readonly writeCookieJar?: string;
    readonly sslCert?: string;
    readonly userAgentOverride?: string;
  }

  export interface Method {
    readonly key: string;
    readonly description?: string;
  }

  export const MethodGet     = {"key": "GET", "description": ""}
  export const MethodHead    = {"key": "HEAD", "description": ""}
  export const MethodPost    = {"key": "POST", "description": ""}
  export const MethodPut     = {"key": "PUT", "description": ""}
  export const MethodPatch   = {"key": "PATCH", "description": ""}
  export const MethodDelete  = {"key": "DELETE", "description": ""}
  export const MethodConnect = {"key": "CONNECT", "description": ""}
  export const MethodOptions = {"key": "OPTIONS", "description": ""}
  export const MethodTrace   = {"key": "TRACE", "description": ""}

  export const allMethods: Method[] = [MethodGet, MethodHead, MethodPost, MethodPut, MethodPatch, MethodDelete, MethodConnect, MethodOptions, MethodTrace];

  export interface Prototype {
    method: string;
    readonly protocol: string;
    readonly domain: string;
    readonly port?: number;
    readonly path?: string;
    query?: QueryParam[];
    readonly fragment?: string;
    headers?: header.Header[];
    auth?: auth.Auth[];
    body?: rbody.Body;
    options?: Options;
  }

  export interface Request {
    readonly key: string;
    readonly title: string;
    readonly description: string;
    readonly prototype: Prototype;
  }

  export interface Summary {
    readonly key: string;
    readonly title: string;
    readonly description: string;
    readonly url: string;
    readonly order: number;
  }
}
