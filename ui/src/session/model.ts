export interface SessionSummary {
  readonly key: string;
  readonly title: string | undefined;
  readonly cookieCount: number;
  readonly variableCount: number;
}

export interface Cookie {
  name: string;
  value: string;
  domain?: string;
  path?: string;
  expires?: string;
  size?: number;
  httpOnly?: boolean;
  secure?: boolean;
  sameSite?: boolean;
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

export interface SessAdded {
  sessions: SessionSummary[];
  active: Session;
}
