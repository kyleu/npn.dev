namespace call {
  export interface Result {
    readonly collection: string;
    readonly request: string;
    readonly requestHeaders: header.Header[];
    readonly status: string;
    readonly response: string;
    readonly timing: string;
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
