namespace events {
  let openEvents: Map<string, Function> = new Map<string, Function>();
  let closeEvents: Map<string, Function> = new Map<string, Function>();

  export function register(key: string, o?: (param?: string) => void, c?: (param?: string) => void) {
    if (!o) {
      o = () => {};
    }
    openEvents.set(key, o);
    if (c) {
      closeEvents.set(key, c);
    }
  }

  export function getOpenEvent(key: string) {
    return openEvents.get(key);
  }

  export function getCloseEvent(key: string) {
    return closeEvents.get(key);
  }
}
