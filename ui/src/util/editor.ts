export interface Editor {
  on(event: string, f: () => void): void;
  refresh(): void;
  setSize(h: string, w: string): void;

  getValue(): string;
  setValue(v: string): void;

  getCursor(): unknown;
  setCursor(u: unknown): void;
}

declare const CodeMirror: (el: HTMLElement, x: object) => Editor;

export function editorFor(el: HTMLElement, lineNumbers: boolean, mode: string, value: string, readOnly: boolean): Editor {
  const args: {[key: string]: unknown} = {lineNumbers, mode, value};
  if (readOnly) {
    args["readOnly"] = "nocursor";
  }
  return CodeMirror(el, args);
}
