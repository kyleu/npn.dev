import {ref} from "@vue/composition-api";

export interface SessionSummary {
  readonly key: string;
  readonly title: string | undefined;
  readonly cookieCount: number;
  readonly variableCount: number;
}

export const activeSessionRef = ref<string>("_");
export const sessionsRef = ref<SessionSummary[]>([]);

export function getSession(key: string): SessionSummary | undefined {
  for (const s of sessionsRef.value) {
    if (s.key === key) {
      return s;
    }
  }
}
