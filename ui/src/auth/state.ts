import {ref, watchEffect} from "@vue/composition-api";
import {Basic} from "@/auth/basic";
import {Auth} from "@/auth/model";
import {jsonClone} from "@/util/json";
import {requestEditingRef} from "@/request/state";

export interface AuthConfig {
  type: string;
  basicContent: Basic;
}

const def = {
  type: "",
  basicContent: { username: "", password: "", showPassword: true }
};

export const authConfigRef = ref<AuthConfig>(def);

export function toAuth(c: AuthConfig): Auth | undefined {
  switch (c.type) {
    case "basic":
      return { type: c.type, config: c.basicContent };
    default:
      return undefined;
  }
}

export function toAuthConfig(a: Auth | undefined): AuthConfig {
  if(!a) {
    return jsonClone(def);
  }
  switch (a.type) {
    case "basic":
      return { type: a.type, basicContent: jsonClone(a.config) as Basic };
    default:
      return jsonClone(def);
  }
}

function compareBasic(lb: Basic, rb: Basic): boolean {
  return lb.username !== rb.username || lb.password !== rb.password || lb.showPassword !== rb.showPassword;
}

function diff(t: AuthConfig, p: Auth | undefined): boolean {
  if (!p) {
    return t.type !== "";
  }
  if (t.type !== p.type) {
    return true;
  }

  switch (t.type) {
    case "":
      return false;
    case "basic":
      return compareBasic(t.basicContent, p.config as Basic);
    default:
      return false;
  }
}

watchEffect(() => {
  const t = authConfigRef.value;
  if (requestEditingRef) {
    const v = requestEditingRef.value;
    if (v) {
      const p = v.prototype.auth;
      if (diff(t, p)) {
        v.prototype.auth = toAuth(t);
      }
    }
  }
});
