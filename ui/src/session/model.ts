import {ref} from "@vue/composition-api";

export interface SessionSummary {
  readonly key: string;
  readonly title: string | undefined;
  readonly cookieCount: number;
  readonly variableCount: number;
}

export interface Cookie {
  name: string;
  value: string;
  domain: string;
  path: string;
  expires: string;
  size: string;
  httpOnly: string;
  secure: string;
  sameSite: string;
}

export interface Variable {
  k: string;
  v: string;
  desc?: string;
}

export interface Session {
  key: string;
  title: string | undefined;
  cookies: Cookie[];
  variables: Variable[];
}
