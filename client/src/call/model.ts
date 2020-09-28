namespace call {
  export interface Response {
    readonly status: string;
    readonly statusCode?: number;
    readonly proto?: string;
    readonly protoMajor?: number;
    readonly protoMinor?: number;
    readonly headers?: header.Header[];
    readonly contentLength?: number;
    readonly contentType?: string;
    readonly charset?: string;
    readonly transferEncoding?: string[];
    readonly close?: boolean;
    readonly uncompressed?: boolean;
    readonly body?: rbody.Body;
    readonly error?: string;
  }

  export interface Result {
    readonly id: string;
    readonly collection: string;
    readonly request: string;
    readonly requestHeaders?: header.Header[];
    readonly status: string;
    readonly response?: Response;
    readonly timing?: Timing;
    readonly error?: string;
  }

  export function prepare(coll: string, r: request.Request) {
    const param = {"coll": coll, "req": r.key};
    socket.send({svc: services.request.key, cmd: command.client.requestCall, param: param});
  }

  export function setResult(param: any) {
    const result = param as Result;
    const container = dom.req(`#${result.collection}--${result.request}-call`);
    dom.setContent(container, renderResult(result));
  }
}
