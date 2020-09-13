namespace system {
  export function onSystemMessage(cmd: string, param: any) {
    switch (cmd) {
      case command.server.connected:
        console.info("Connected!!!!")
        break;
      default:
        console.warn(`unhandled system command [${cmd}]`);
    }
  }
}
