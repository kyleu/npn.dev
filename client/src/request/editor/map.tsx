namespace request.editor {
  export interface MapParam {
    k: string;
    v: string;
    desc?: string;
  }

  export function mapHeader(id: string, cb: string) {
    return <li>
      <div data-uk-grid="">
        <div class="uk-width-1-4">Name</div>
        <div class="uk-width-1-4">Value</div>
        <div class="uk-width-1-2">
          <div class="right">
            <a class={style.linkColor} href="" onclick={"return request.editor." + cb + "('" + id + "')"} title="new row">
              <span data-uk-icon="icon: plus" />
            </a>
          </div>
          Description
        </div>
      </div>
    </li>;
  }

  export function newChild(elID: string, idx: number, h: request.editor.MapParam, cb: string) {
    return <li>
      <div data-uk-grid="">
        <div class="uk-width-1-4">
          <input class="uk-input" data-field={idx + "-key"} type="text" value={h.k} />
        </div>
        <div class="uk-width-1-4">
          <input class="uk-input" data-field={idx + "-value"} type="text" value={h.v} />
        </div>
        <div class="uk-width-1-2">
          <div class="right" style="margin-top: 6px;">
            <a class={style.linkColor} href="" onclick={"return request.editor." + cb + "('" + elID + "', this);"} title="remove row"><span data-uk-icon="icon: close" /></a>
          </div>
          <input style="width: calc(100% - 48px);" class="uk-input" data-field={idx + "-desc"} type="text" value={h.desc} />
        </div>
      </div>
    </li>
  }

  export function parseMapParams(elID: string) {
    const ul = dom.req("#" + elID + "-ul");
    const inputs = dom.els<HTMLInputElement>("input", ul);
    let ret: MapParam[] = []
    for (const i of inputs) {
      const field = i.dataset["field"] || "";
      const dash = field.lastIndexOf("-");
      const idx = parseInt(field.substring(0, dash), 10);
      const key = field.substring(dash + 1);
      if (!ret[idx]) {
        ret[idx] = {k: "", v: ""};
      }
      switch (key) {
        case "key":
          ret[idx].k = i.value.trim();
          break;
        case "value":
          ret[idx].v = i.value.trim();
          break;
        case "desc":
          const desc = i.value.trim();
          if (desc.length > 0) {
            ret[idx].desc = desc;
          }
          break;
        default:
          throw "unknown key [" + key + "]";
      }
    }

    ret = ret.filter(x => x.k.length > 0);

    return ret;
  }
}
