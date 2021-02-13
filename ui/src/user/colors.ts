import {socketRef} from "@/socket/socket";
import {systemService} from "@/util/services";
import {clientCommands} from "@/util/command";
import {logError} from "@/util/log";

export type Color = string;

// Paste from color.adobe.com's CSS export
const wip = ``;

export function debugTheme(): void {
  const w = wip.trim();
  if (w.length > 0) {
    const hexes: string[] = [];

    const keyIdx = w.indexOf(".");
    const keyLen = w.indexOf(" ", keyIdx) - keyIdx - 1;
    let key = w.substr(keyIdx + 1, keyLen);
    key = key.split('-').slice(0, -2).map(x => x[0].toUpperCase() + x.substr(1)).join('');

    let idx = w.indexOf("#");
    while (idx > -1) {
      hexes.push(w.substr(idx, 7).toLowerCase());
      idx = w.indexOf("#", idx + 1);
    }
    if (hexes.length === 5) {
      const content = `const ${key} = { key: "${key}", mode: "light", navB: "${hexes[1]}", navF: "${hexes[4]}", menuB: "${hexes[2]}", menuF: "${hexes[3]}", menuL: "${hexes[4]}", bodyB: "${hexes[0]}", bodyL: "${hexes[4]}" }`;
      setTimeout(() => {
        if (socketRef.value) {
          socketRef.value.send({channel: systemService.key, cmd: clientCommands.testbed, param: {t: "theme", k: key, v: content}});
        }
      }, 1000);
    } else {
      logError("unhandled", hexes);
    }
  }
}
