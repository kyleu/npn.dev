import {QueryParam} from "@/request/model";

export interface BodyType {
  readonly key: string;
  readonly title: string;
  readonly hidden: boolean;
}

export const AllTypes: BodyType[] = [
  {key: "error", title: "Error", hidden: true},
  {key: "form", title: "Form", hidden: false},
  {key: "html", title: "HTML", hidden: false},
  {key: "xml", title: "XML", hidden: false},
  {key: "json", title: "JSON", hidden: false},
  {key: "large", title: "Large File", hidden: false},
  {key: "raw", title: "Raw", hidden: true}
];

export interface RBody {
  type: string;
  length?: number;
  config: object;
}

export interface FormConfig {
  data: QueryParam[];
}

export interface HTMLConfig {
  content: string;
}

export interface XMLConfig {
  content: string;
}

export interface JSONConfig {
  msg: unknown;
}

export interface ImageConfig {
  readonly type: string;
  readonly content: string;
}

export interface RawConfig {
  readonly type: string | undefined;
  readonly content: string;
  readonly length: number;
  readonly binary: boolean | undefined;
}

export interface ErrorConfig {
  readonly message: string;
}
