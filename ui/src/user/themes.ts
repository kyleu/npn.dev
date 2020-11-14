import {Color, debugTheme} from "@/user/colors";

debugTheme();

export interface ThemeColors {
  readonly key: string;
  readonly mode: string;
  readonly bodyB: Color;
  readonly bodyL: Color;
  readonly navB: Color;
  readonly navF: Color;
  readonly menuB: Color;
  readonly menuF: Color;
  readonly menuL: Color;
}

const Default = {
  key: "Default",
  mode: "light",
  bodyB: "#fcfff5",
  bodyL: "#193441",
  navB: "#193441",
  navF: "#fff",
  menuB: "#3e606f",
  menuF: "#ccc",
  menuL: "#91aa9d"
};

const Firenze = {
  key: "Firenze",
  mode: "light",
  bodyB: "#468966",
  bodyL: "#8e2800",
  navB: "#fff0a5",
  navF: "#8e2800",
  menuB: "#ffb03b",
  menuF: "#b64926",
  menuL: "#8e2800"
};

const NeutralBlue = {
  key: "NeutralBlue",
  mode: "light",
  bodyB: "#fcfff5",
  bodyL: "#193441",
  navB: "#d1dbbd",
  navF: "#193441",
  menuB: "#91aa9d",
  menuF: "#3e606f",
  menuL: "#193441"
};

const Phaedra = {
  key: "Phaedra",
  mode: "light",
  bodyB: "#ff6138",
  bodyL: "#00a388",
  navB: "#ffff9d",
  navF: "#00a388",
  menuB: "#beeb9f",
  menuF: "#79bd8f",
  menuL: "#00a388"
};

const SandyStoneBeach = {
  key: "SandyStoneBeach",
  mode: "light",
  bodyB: "#e6e2af",
  bodyL: "#002f2f",
  navB: "#a7a37e",
  navF: "#002f2f",
  menuB: "#efecca",
  menuF: "#046380",
  menuL: "#002f2f"
};

const FlatUI = {
  key: "FlatUI",
  mode: "light",
  bodyB: "#2c3e50",
  bodyL: "#2980b9",
  navB: "#e74c3c",
  navF: "#2980b9",
  menuB: "#ecf0f1",
  menuF: "#3498db",
  menuL: "#2980b9"
};

const AspirinC = {
  key: "AspirinC",
  mode: "light",
  bodyB: "#225378",
  bodyL: "#eb7f00",
  navB: "#1695a3",
  navF: "#eb7f00",
  menuB: "#acf0f2",
  menuF: "#f3ffe2",
  menuL: "#eb7f00"
};

const HoneyPot = {
  key: "HoneyPot",
  mode: "light",
  bodyB: "#105b63",
  bodyL: "#bd4932",
  navB: "#fffad5",
  navF: "#bd4932",
  menuB: "#ffd34e",
  menuF: "#db9e36",
  menuL: "#bd4932"
};

const VitaminC = {
  key: "VitaminC",
  mode: "light",
  bodyB: "#004358",
  bodyL: "#fd7400",
  navB: "#1f8a70",
  navF: "#fd7400",
  menuB: "#bedb39",
  menuF: "#ffe11a",
  menuL: "#fd7400"
};

const SeaWolf = {
  key: "SeaWolf",
  mode: "light",
  bodyB: "#dc3522",
  bodyL: "#1e1e20",
  navB: "#d9cb9e",
  navF: "#1e1e20",
  menuB: "#374140",
  menuF: "#2a2c2b",
  menuL: "#1e1e20"
};

const CircusIII = {
  key: "CircusIII",
  mode: "light",
  bodyB: "#2e0927",
  bodyL: "#04756f",
  navB: "#d90000",
  navF: "#04756f",
  menuB: "#ff2d00",
  menuF: "#ff8c00",
  menuL: "#04756f"
};

const VintageRalphLauren = {
  key: "VintageRalphLauren",
  mode: "light",
  bodyB: "#703030",
  bodyL: "#c77966",
  navB: "#2f343b",
  navF: "#c77966",
  menuB: "#7e827a",
  menuF: "#e3cda4",
  menuL: "#c77966"
};

const CherryCheesecake = {
  key: "CherryCheesecake",
  mode: "light",
  bodyB: "#b9121b",
  bodyL: "#bd8d46",
  navB: "#4c1b1b",
  navF: "#bd8d46",
  menuB: "#f6e497",
  menuF: "#fcfae1",
  menuL: "#bd8d46"
};

const FriendsAndFoes = {
  key: "FriendsAndFoes",
  mode: "light",
  bodyB: "#2f2933",
  bodyL: "#ffffa6",
  navB: "#01a2a6",
  navF: "#ffffa6",
  menuB: "#29d9c2",
  menuF: "#bdf271",
  menuL: "#ffffa6"
};

const CS04 = {
  key: "CS04",
  mode: "light",
  bodyB: "#f6f792",
  bodyL: "#ea2e49",
  navB: "#333745",
  navF: "#ea2e49",
  menuB: "#77c4d3",
  menuF: "#daede2",
  menuL: "#ea2e49"
};

const TimesChanging = {
  key: "TimesChanging",
  mode: "light",
  bodyB: "#332532",
  bodyL: "#a49a87",
  navB: "#644d52",
  navF: "#a49a87",
  menuB: "#f77a52",
  menuF: "#ff974f",
  menuL: "#a49a87"
};

const OceanSunset = {
  key: "OceanSunset",
  mode: "light",
  bodyB: "#405952",
  bodyL: "#f54f29",
  navB: "#9c9b7a",
  navF: "#f54f29",
  menuB: "#ffd393",
  menuF: "#ff974f",
  menuL: "#f54f29"
};

const VentanaAzul = {
  key: "VentanaAzul",
  mode: "light",
  bodyB: "#f5a503",
  bodyL: "#36b1bf",
  navB: "#f2385a",
  navF: "#36b1bf",
  menuB: "#e9f1df",
  menuF: "#4ad9d9",
  menuL: "#36b1bf"
};

const ZenAndTea = {
  key: "ZenAndTea",
  mode: "light",
  bodyB: "#10222b",
  bodyL: "#f6ffe0",
  navB: "#95ab63",
  navF: "#f6ffe0",
  menuB: "#bdd684",
  menuF: "#e2f0d6",
  menuL: "#f6ffe0"
};

const BirdfolioBlues = {
  key: "BirdfolioBlues",
  mode: "light",
  bodyB: "#2b3a42",
  bodyL: "#ff530d",
  navB: "#3f5765",
  navF: "#ff530d",
  menuB: "#bdd4de",
  menuF: "#efefef",
  menuL: "#ff530d"
};

const PearLemonFizz = {
  key: "PearLemonFizz",
  mode: "light",
  bodyB: "#04bfbf",
  bodyL: "#588f27",
  navB: "#cafcd8",
  navF: "#588f27",
  menuB: "#f7e967",
  menuF: "#a9cf54",
  menuL: "#588f27"
};

export const allThemes: ThemeColors[] = [
  Default, Firenze, NeutralBlue, Phaedra, SandyStoneBeach, FlatUI, AspirinC, HoneyPot, VitaminC, SeaWolf, CircusIII, VintageRalphLauren, CherryCheesecake, FriendsAndFoes, CS04, TimesChanging, OceanSunset, VentanaAzul, ZenAndTea, BirdfolioBlues, PearLemonFizz
];
