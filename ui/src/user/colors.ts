import {socketRef} from "@/socket/socket";
import {systemService} from "@/util/services";
import {clientCommands} from "@/util/command";

export type Color = string;

// Paste from color.adobe.com's CSS export
const wip = `

/* Color Theme Swatches in Hex */
.Herbs-and-Spice-1-hex { color: #5A1F00; }
.Herbs-and-Spice-2-hex { color: #D1570D; }
.Herbs-and-Spice-3-hex { color: #FDE792; }
.Herbs-and-Spice-4-hex { color: #477725; }
.Herbs-and-Spice-5-hex { color: #A9CC66; }

/* Color Theme Swatches in RGBA */
.Herbs-and-Spice-1-rgba { color: rgba(89, 31, 0, 1); }
.Herbs-and-Spice-2-rgba { color: rgba(209, 86, 13, 1); }
.Herbs-and-Spice-3-rgba { color: rgba(253, 230, 145, 1); }
.Herbs-and-Spice-4-rgba { color: rgba(70, 119, 36, 1); }
.Herbs-and-Spice-5-rgba { color: rgba(168, 204, 102, 1); }

/* Color Theme Swatches in HSLA */
.Herbs-and-Spice-1-hsla { color: hsla(20, 100, 17, 1); }
.Herbs-and-Spice-2-hsla { color: hsla(22, 88, 43, 1); }
.Herbs-and-Spice-3-hsla { color: hsla(47, 96, 78, 1); }
.Herbs-and-Spice-4-hsla { color: hsla(95, 52, 30, 1); }
.Herbs-and-Spice-5-hsla { color: hsla(80, 50, 60, 1); }




`;

export function debugTheme(): void {
  const w = wip.trim();
  if (w.length > 0) {
    const hexes: string[] = [];

    const keyIdx = w.indexOf(".");
    const keyLen = w.indexOf(" ", keyIdx) - keyIdx - 1;
    let key = w.substr(keyIdx + 1, keyLen)
    key = key.split('-').slice(0, -2).map(x => x[0].toUpperCase() + x.substr(1)).join('')

    let idx = w.indexOf("#");
    while (idx > -1) {
      hexes.push(w.substr(idx, 7).toLowerCase());
      idx = w.indexOf("#", idx + 1);
    }
    if (hexes.length === 5) {
      const content = `const ${key} = { key: "${key}", mode: "light", navB: "${hexes[1]}", navF: "${hexes[4]}", menuB: "${hexes[2]}", menuF: "${hexes[3]}", menuL: "${hexes[4]}", bodyB: "${hexes[0]}", bodyL: "${hexes[4]}" }`;
      console.log(content);
      // @ts-ignore
      setTimeout(() => {
        if (socketRef.value) {
          socketRef.value.send({svc: systemService.key, cmd: clientCommands.testbed, param: {t: "theme", k: key, v: content}});
        } else {
          console.log("no socket");
        }
      }, 1000);
    } else {
      console.log("unhandled", hexes);
    }
  }
}
