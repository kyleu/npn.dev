namespace ui {
  export function setPanels(coll?: string, req?: string, act?: string) {
    dom.setDisplay("#collection-list-panel", coll === undefined);
    dom.setDisplay("#collection-panel", coll !== undefined && coll.length > 0 && req === undefined);
    dom.setDisplay("#request-panel", req !== undefined && req.length > 0 && act === undefined);
    dom.setDisplay("#action-panel", act !== undefined && act.length > 0);
    setBreadcrumbs(coll, req, act);
  }
}
