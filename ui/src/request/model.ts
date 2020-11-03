import {Auth} from "@/auth/model";
import {RBody} from "@/body/model";
import {Header} from "@/header/model";

export interface QueryParam {
  k: string;
  v: string;
  desc?: string;
}

export interface Options {
  timeout?: number;
  ignoreRedirects?: boolean;
  ignoreReferrer?: boolean;
  ignoreCerts?: boolean;
  ignoreCookies?: boolean;
  excludeDefaultHeaders?: string[];
  readCookieJars?: string[];
  writeCookieJar?: string;
  sslCert?: string;
  userAgentOverride?: string;
}

export interface Method {
  readonly key: string;
  readonly description?: string;
}

export const MethodGet = {"key": "GET", "description": ""}
export const MethodHead = {"key": "HEAD", "description": ""}
export const MethodPost = {"key": "POST", "description": ""}
export const MethodPut = {"key": "PUT", "description": ""}
export const MethodPatch = {"key": "PATCH", "description": ""}
export const MethodDelete = {"key": "DELETE", "description": ""}
export const MethodConnect = {"key": "CONNECT", "description": ""}
export const MethodOptions = {"key": "OPTIONS", "description": ""}
export const MethodTrace = {"key": "TRACE", "description": ""}

export const allMethods: Method[] = [MethodGet, MethodHead, MethodPost, MethodPut, MethodPatch, MethodDelete, MethodConnect, MethodOptions, MethodTrace];

export interface Prototype {
  method: string;
  protocol: string;
  domain: string;
  port?: number;
  path?: string;
  query?: QueryParam[];
  fragment?: string;
  headers?: Header[];
  auth?: Auth;
  body?: RBody;
  options?: Options;
}

export interface NPNRequest {
  readonly key: string;
  readonly title: string;
  readonly description: string;
  readonly prototype: Prototype;
}

export function cloneRequest(r: NPNRequest | undefined): NPNRequest | undefined {
  return JSON.parse(JSON.stringify(r));
}

export interface Summary {
  readonly key: string;
  readonly title: string;
  readonly description: string;
  readonly url: string;
  readonly order: number;
}
