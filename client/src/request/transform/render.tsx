
namespace transform {
  export interface Result {
    readonly coll: string;
    readonly req: string;
    readonly fmt: string;
    readonly out: string;
  }

  export function renderRequest(coll: string, req: string, format: string) {
    return <div id={coll + "--" + req + "-transform"}>
      {request.renderActionClose()}
      <div class="transform-title">{format}</div>
      <div class="transform-result" />
    </div>;
  }

  export function setResult(result: Result) {
    const container = dom.req(`#${result.coll}--${result.req}-transform .transform-result`);
    dom.setContent(container, render(result));
    log.info("call result [" + result.coll + "/" + result.req + ": " + result.fmt + "] received");
  }

  function render(r: Result) {
    return <div class="uk-margin-top"><pre>{r.out}</pre></div>;
  }
}
