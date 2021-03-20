const mkdirp = require("mkdirp");
const puppeteer = require("puppeteer");

const path = "../../screenshots/";

const sizes = [
  { w: 1920, h: 1080, k: "1080p", s: 19, d: 1 },
  { w: 3840, h: 2160, k: "4K", s: 19, d: 1 },
  { w: 2880, h: 1800, k: "macOS 1440p", s: 24, d: 2 },

  { w: 2732, h: 2048, k: "iPad Pro (12.9)", s: 12.9, d: 2 },
  { w: 2388, h: 1668, k: "iPad Pro (11)", s: 11, d: 2 },
  { w: 2224, h: 1668, k: "iPad Pro (10.5)", s: 10.5, d: 2 },
  { w: 2008, h: 1536, k: "iPad Pro (9.7)", s: 9.7, d: 2 },

  { w: 2778, h: 1284, k: "iPhone 12 Pro Max", s: 6.7, d: 3 },
  { w: 2532, h: 1170, k: "iPhone 12 Pro", s: 6.1, d: 3 },
  { w: 2688, h: 1242, k: "iPhone XS Max", s: 6.5, d: 3 },
  { w: 2436, h: 1125, k: "iPhone XS", s: 5.8, d: 3 },
  { w: 2208, h: 1242, k: "iPhone 8 Plus", s: 5.5, d: 3 },
  { w: 1334, h: 750, k: "iPhone 8", s: 4.7, d: 2 },
  { w: 1136, h: 640, k: "iPhone SE", s: 4, d: 2 }
];

function getFilename(k, w, h, d, s, fn) {
  const dir = path + s + "-" + k;
  mkdirp(dir);
  console.log(" - [" + w + "x" + h + "@" + d + "] (" + k + ")");
  return dir + "/" + fn + ".png";
}

async function setViewLandscape(page, x, fn) {
  await page.setViewport({width: x.w / x.d, height: x.h / x.d, deviceScaleFactor: x.d});
  return getFilename(x.k, x.w, x.h, x.d, x.s, fn + "-landscape");
}
async function setViewPortrait(page, x, fn) {
  await page.setViewport({width: x.h / x.d, height: x.w / x.d, deviceScaleFactor: x.d});
  return getFilename(x.k, x.w, x.h, x.d, x.s, fn + "-portrait");
}

async function ss(page, fn) {
  console.log("starting capture of [" + fn + "]...");
  for (const s of sizes) {
    const fnl = await setViewLandscape(page, s, fn);
    await page.screenshot({ path: fnl });
    const fnp = await setViewPortrait(page, s, fn);
    await page.screenshot({ path: fnp });
  }
  console.log("completed capture of [" + fn + "]");
}

(async () => {
  const browser = await puppeteer.launch();
  const page = await browser.newPage();

  await page.goto("http://localhost:10101");
  await ss(page, "00-marketing");

  await page.evaluate(() => init(true, true));
  await ss(page, "01-home");

  await page.goto("http://localhost:10101/c/personal/amazon/call");
  console.log("pausing for load...");
  await page.waitForTimeout(1500);
  console.log("pause complete");
  await page.click("#tab-response-headers");
  await ss(page, "02-call");

  await page.goto("http://localhost:10101/s/_");
  await ss(page, "03-session");

  await browser.close();
})();

