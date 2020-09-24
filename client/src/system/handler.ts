namespace system {
  export interface MsgConnected {
    readonly profile: profile.Profile;
  }

  export function onSystemMessage(cmd: string, param: any) {
    switch (cmd) {
      case command.server.error:
        console.warn("error from server: " + param)
        break;
      case command.server.connected:
        cache.apply(param as MsgConnected);
        break;
      default:
        console.warn(`unhandled system command [${cmd}]`);
    }
  }
}
