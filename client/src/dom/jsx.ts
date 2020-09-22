declare namespace JSX {
  interface Element extends HTMLElement {}

  // noinspection JSUnusedGlobalSymbols
  interface IntrinsicElements {
    [elemName: string]: any;
  }
}

// noinspection JSUnusedGlobalSymbols
function JSX(tag: string, attrs: any) {
  const e = document.createElement(tag);
  for (const name in attrs) {
    if (name && attrs.hasOwnProperty(name)) {
      const v = attrs[name];
      if (name === "dangerouslySetInnerHTML") {
        dom.setHTML(e, v["__html"]);
      } else if (v === true) {
        e.setAttribute(name, name);
      } else if (v !== false && v !== null && v !== undefined) {
        e.setAttribute(name, v.toString());
      }
    }
  }
  for (let i = 2; i < arguments.length; i++) {
    let child = arguments[i];
    if (Array.isArray(child)) {
      child.forEach(c => {
        if (child === undefined || child === null) {
          throw `child array for tag [${tag}] is ${child}\n${e.outerHTML}`;
        }
        if (c === undefined || c === null) {
          throw `child for tag [${tag}] is ${c}\n${e.outerHTML}`;
        }
        if (typeof c === "string") {
          c = document.createTextNode(c);
        }
        e.appendChild(c);
      });
    } else if (child === undefined || child === null) {
      throw `child for tag [${tag}] is ${child}\n${e.outerHTML}`;
      // debugger;
      // child = document.createTextNode("NULL!");
    } else {
      if (!child.nodeType) {
        child = document.createTextNode(child.toString());
      }
      e.appendChild(child);
    }
  }
  return e;
}
