namespace collection {
  export function onCollectionMessage(cmd: string, param: any) {
    switch (cmd) {
      case command.server.collections:
        console.warn("Collections!");
        break;
      default:
        console.warn(`unhandled collection command [${cmd}]`);
    }
  }
}
