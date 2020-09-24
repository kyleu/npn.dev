namespace call {
  export function renderResult(r: Result) {
    return <div style="overflow: auto;max-width: 820px;"><pre>{json.str(r)}</pre></div>;
  }
}
