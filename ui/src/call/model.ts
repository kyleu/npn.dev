import {Header} from "@/header/model";
import {RBody} from "@/body/model";
import {Timing} from "@/call/timing";

export interface Response {
  readonly method: string;
  readonly url: string;
  readonly requestHeaders: Header[];
  readonly status: string;
  readonly statusCode?: number;
  readonly proto?: string;
  readonly protoMajor?: number;
  readonly protoMinor?: number;
  readonly headers?: Header[];
  readonly contentLength?: number;
  readonly contentType?: string;
  readonly charset?: string;
  readonly transferEncoding?: string[];
  readonly close?: boolean;
  readonly uncompressed?: boolean;
  readonly body?: RBody;
  readonly prior?: Response;
  readonly timing?: Timing;
  readonly error?: string;
}

export interface CallResult {
  readonly id: string;
  readonly collection: string;
  readonly request: string;
  readonly status: string;
  readonly response?: Response;
  readonly error?: string;
}
