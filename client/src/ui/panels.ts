namespace ui {
  export function setPanels(coll?: string, req?: string, act?: string) {
    dom.setDisplay("#welcome-panel", coll === undefined);
    dom.setDisplay("#collection-panel", coll !== undefined && coll.length > 0 && req === undefined);
    dom.setDisplay("#request-panel", req !== undefined && req.length > 0 && act === undefined);
    dom.setDisplay("#action-panel", act !== undefined && act.length > 0);
    setBreadcrumbs(coll, req, act);
    setTitle(coll, req, act);
  }

  function setTitle(coll: string | undefined, req: string | undefined, act: string | undefined) {
    let title = "";
    if (act) {
      title += act + " ";
    }
    if (coll) {
      title += coll;
    }
    if (req) {
      title += "/" + req;
    }
    if (title.length == 0) {
      title = "npn"
    } else {
      title = "npn: " + title
    }
    document.title = title;
  }
}
