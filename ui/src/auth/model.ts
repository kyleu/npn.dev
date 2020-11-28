export interface Auth {
  readonly type: string;
  readonly config: object;
}

export interface AuthType {
  readonly key: string;
  readonly title: string;
  readonly desc: string;
  readonly hidden: boolean;
}

export const AllTypes: AuthType[] = [
  {key: "basic", title: "Basic", desc: "a simple username and password", hidden: false},
  {key: "header", title: "Header", desc: "an API key passed as a request header", hidden: false},
  {key: "queryparam", title: "Query Param", desc: "an API key passed as a query parameter", hidden: false},
  {key: "bearertoken", title: "Bearer Token", desc: "uses an OAuth bearer token", hidden: false},
  {key: "digest", title: "Digest", desc: "", hidden: true},
  {key: "oauth", title: "OAuth", desc: "", hidden: true},
];
