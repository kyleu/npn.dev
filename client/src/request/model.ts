namespace request {
  export interface QueryParam {
    readonly k: string;
    readonly v: string;
    readonly desc?: string;
  }

  export interface Header {
    readonly k: string;
    readonly v: string;
    readonly desc?: string;
  }

  export interface Options {
    readonly timeout?: number;
    readonly ignoreRedirects?: boolean;
    readonly ignoreReferrer?: boolean;
    readonly ignoreCerts?: boolean;
    readonly excludeDefaultHeaders?: string[];
    readonly readCookieJars?: string[];
    readonly writeCookieJar?: string;
    readonly sslCert?: string;
    readonly userAgentOverride?: string;
  }

  export interface Prototype {
    readonly method: string;
    readonly protocol: string;
    readonly domain: string;
    readonly port?: number;
    readonly path?: string;
    readonly query?: QueryParam[];
    readonly fragment?: string;
    readonly headers?: Header[];
    readonly auth?: auth.Auth[];
    readonly body?: body.Body;
    readonly options?: Options;
  }

  export interface Request {
    readonly key: string;
    readonly title: string;
    readonly description: string;
    readonly prototype: Prototype;
  }
}
