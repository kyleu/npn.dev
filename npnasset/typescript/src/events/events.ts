namespace events {
  let openEvents: map.Map<string, Function>;
  let closeEvents: map.Map<string, Function>;

  export function register(key: string, o?: (param?: string) => void, c?: (param?: string) => void) {
    if (!openEvents) {
      openEvents = new map.Map();
    }
    if (!closeEvents) {
      closeEvents = new map.Map();
    }

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
