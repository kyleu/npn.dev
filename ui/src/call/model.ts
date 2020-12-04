import {Header} from "@/header/model";
import {RBody} from "@/body/model";
import {Timing} from "@/call/timing";

export interface NPNResponse {
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
  readonly timing?: Timing;
  readonly error?: string;
}

export interface RequestStarted {
  readonly coll: string;
  readonly req: string;
  readonly id: string;
  readonly idx: number;
  readonly method: string;
  readonly url: string;
  readonly started: string;
}

export interface RequestCompleted {
  readonly coll: string;
  readonly req: string;
  readonly id: string;
  readonly idx: number;
  readonly status: string;
  readonly rsp: NPNResponse;
  readonly error: string;
  readonly duration: number;
}

export interface RequestResultCycle {
  readonly idx: number;
  readonly method: string;
  readonly url: string;
  readonly started: string;

  status: string;
  rsp: NPNResponse | null;
  error: string;
  duration: number;
}

export interface RequestResults {
  readonly id: string;
  readonly coll: string;
  readonly req: string;
  readonly cycles: RequestResultCycle[];
}
