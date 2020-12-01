import {ref, watchEffect} from "@vue/composition-api";
import {jsonClone, jsonParseTry, jsonStr} from "@/util/json";
import {FormConfig, HTMLConfig, JSONConfig, RBody} from "@/body/model";
import {requestEditingRef} from "@/request/state";
import {QueryParam} from "@/request/model";

export interface BodyConfig {
  type: string;
  formContent: QueryParam[];
  htmlContent: string;
  jsonContent: string;
}

const def = {type: "", formContent: [], htmlContent: "", jsonContent: ""};

export const bodyConfigRef = ref<BodyConfig>(def);

export function toBody(bc: BodyConfig): RBody | undefined {
  switch (bc.type) {
    case "form":
      return { type: bc.type, length: bc.formContent.length, config: {data: bc.formContent} };
    case "html":
      return { type: bc.type, length: bc.htmlContent.length, config: {content: bc.htmlContent} };
    case "json":
      return { type: bc.type, length: bc.jsonContent.length, config: {msg: jsonParseTry(bc.jsonContent)} };
    default:
      return undefined;
  }
}

export function toBodyConfig(b: RBody | undefined): BodyConfig {
  if(!b) {
    return jsonClone(def);
  }
  switch (b.type) {
    case "form":
      return { type: b.type, formContent: (b.config as FormConfig).data, htmlContent: "", jsonContent: "" };
    case "html":
      return { type: b.type, formContent: [], htmlContent: (b.config as HTMLConfig).content, jsonContent: "" };
    case "json":
      if(typeof (b.config as JSONConfig).msg === "string") {
        return { type: b.type, formContent: [], htmlContent: "", jsonContent: (b.config as JSONConfig).msg };
      }
      return { type: b.type, formContent: [], htmlContent: "", jsonContent: jsonStr((b.config as JSONConfig).msg) };
    default:
      return jsonClone(def);
  }
}

function diff(t: BodyConfig, b: RBody | undefined): boolean {
  if (!b) {
    return t.type !== "";
  }
  if (t.type !== b.type) {
    return true;
  }

  switch (t.type) {
    case "":
      return false;
    case "form":
      return t.formContent !== (b.config as FormConfig).data;
    case "html":
      return t.htmlContent !== (b.config as HTMLConfig).content;
    case "json":
      if (typeof (b.config as JSONConfig).msg === "string") {
        return t.jsonContent !== (b.config as JSONConfig).msg;
      }
      return jsonStr(jsonParseTry(t.jsonContent)) !== jsonStr((b.config as JSONConfig).msg);
    default:
      return false;
  }
}

watchEffect(() => {
  const t = bodyConfigRef.value;
  if (requestEditingRef) {
    const v = requestEditingRef.value;
    if (v) {
      const p = v.prototype.body;
      if (diff(t, p)) {
        v.prototype.body = toBody(t);
      }
    }
  }
});