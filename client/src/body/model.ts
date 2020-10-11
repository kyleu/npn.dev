namespace rbody {
  export interface BodyType {
    readonly key: string;
    readonly title: string;
    readonly hidden: boolean;
  }

  export const AllTypes: BodyType[] = [
    {key: "error", title: "Error", hidden: true},
    {key: "form", title: "Form", hidden: false},
    {key: "html", title: "HTML", hidden: false},
    {key: "json", title: "JSON", hidden: false},
    {key: "large", title: "Large File", hidden: false},
    {key: "raw", title: "Raw", hidden: true}
  ]

  export interface Body {
    readonly type: string;
    readonly length?: number;
    readonly config: any;
  }

  export interface JSONConfig {
    readonly msg: any;
  }

  export interface HTMLConfig {
    readonly content: string;
  }

  export interface RawConfig {
    readonly type: string | undefined;
    readonly content: string;
    readonly length: number;
    readonly binary: boolean | undefined;
  }
}
