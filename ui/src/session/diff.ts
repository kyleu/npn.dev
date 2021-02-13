import {jsonClone} from "@/util/json";
import {checkNull, comp, compArray, Diff} from "@/util/diff";
import {Session} from "@/session/model";

const debugDiff = false;

function debug(lpo: Session, rpo: Session, ret: Diff[]): void {
  if (debugDiff) {
    console.debug(jsonClone(lpo), jsonClone(rpo));
    if (ret.length > 0) {
      console.debug(ret);
    }
  }
}

export function diffSessions(l: Session | undefined, r: Session | undefined): Diff[] {
  const ret: Diff[] = [];
  const p = (k: string, lv: unknown, rv: unknown): number => ret.push({k: k, l: lv, r: rv});

  if ((!l || !r)) {
    checkNull("request", l, r, p);
    return ret;
  }

  comp("key", l.key, r.key, p);
  comp("title", l.title, r.title, p);

  compArray("cookies", l.cookies, r.cookies, p);
  compArray("variables", l.variables, r.variables, p);

  debug(l, r, ret);

  return ret;
}
