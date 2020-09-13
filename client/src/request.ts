namespace request {
  export interface QueryParam {
    readonly key: string;
    readonly value: string;
  }

  export interface Header {
    readonly key: string;
    readonly value: string;
  }

  export interface Auth {
    readonly type: string;
    readonly config: object;
  }

  export interface Body {
    readonly type: string;
    readonly config: object;
  }

  export interface Options {
    readonly todo: string;
  }

  export interface Prototype {
    readonly method: string;
    readonly protocol: string;
    readonly domain: string;
    readonly port: number;
    readonly path: string;
    readonly query: [QueryParam] | undefined;
    readonly fragment: string;
    readonly headers: [Header] | undefined;
    readonly auth: [Auth] | undefined;
    readonly body: Body | undefined;
    readonly options: Options | undefined;
  }

  export interface Request {
    readonly title: string;
    readonly description: string;
    readonly prototype: string;
  }
}
