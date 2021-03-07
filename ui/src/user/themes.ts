import { Color, debugTheme } from "@/user/colors";

debugTheme();

export interface ThemeColors {
  readonly key: string;
  readonly mode: string;
  readonly navB: Color;
  readonly navF: Color;
  readonly menuB: Color;
  readonly menuF: Color;
  readonly menuL: Color;
  readonly bodyB: Color;
  readonly bodyL: Color;
}

const Default = {
  key: "Default",
  mode: "light",
  navB: "#193441",
  navF: "#dddddd",
  menuB: "#3e606f",
  menuF: "#cccccc",
  menuL: "#91aa9d",
  bodyB: "#fcfff5",
  bodyL: "#2f657f"
};
const Darkfault = {
  key: "Darkfault",
  mode: "dark",
  bodyB: "#000",
  bodyL: "#749eb1",
  navB: "#193441",
  navF: "#dddddd",
  menuB: "#3e606f",
  menuF: "#cccccc",
  menuL: "#91aa9d"
};

const AfternoonChai = { key: "AfternoonChai", mode: "light", navB: "#fff6c5", navF: "#695d46", menuB: "#a1e8d9", menuF: "#ff712c", menuL: "#695d46", bodyB: "#cfc291", bodyL: "#695d46" };
const AspirinC = { key: "AspirinC", mode: "light", navB: "#1695a3", navF: "#eb7f00", menuB: "#acf0f2", menuF: "#f3ffe2", menuL: "#eb7f00", bodyB: "#225378", bodyL: "#eb7f00" };
const BeachTime = { key: "BeachTime", mode: "light", navB: "#ffeead", navF: "#aad8b0", menuB: "#ff6f69", menuF: "#ffcc5c", menuL: "#aad8b0", bodyB: "#96ceb4", bodyL: "#aad8b0" };
const BirdfolioBlues = { key: "BirdfolioBlues", mode: "light", navB: "#3f5765", navF: "#ff530d", menuB: "#bdd4de", menuF: "#efefef", menuL: "#ff530d", bodyB: "#2b3a42", bodyL: "#ff530d" };
const BloggyGradientBlues = { key: "BloggyGradientBlues", mode: "light", navB: "#c4d7ed", navF: "#183152", menuB: "#abc8e2", menuF: "#375d81", menuL: "#183152", bodyB: "#e1e6fa", bodyL: "#183152" };
const BlueSky = { key: "BlueSky", mode: "light", navB: "#35478c", navF: "#add5f7", menuB: "#4e7ac7", menuF: "#7fb2f0", menuL: "#add5f7", bodyB: "#16193b", bodyL: "#add5f7" };
const Bogart = { key: "Bogart", mode: "light", navB: "#958976", navF: "#6a6a61", menuB: "#611427", menuF: "#1d2326", menuL: "#6a6a61", bodyB: "#dddcc5", bodyL: "#6a6a61" };
const CS04 = { key: "CS04", mode: "light", navB: "#333745", navF: "#ea2e49", menuB: "#77c4d3", menuF: "#daede2", menuL: "#ea2e49", bodyB: "#f6f792", bodyL: "#ea2e49" };
const Campfire = { key: "Campfire", mode: "light", navB: "#f2e394", navF: "#8c4646", menuB: "#f2ae72", menuF: "#d96459", menuL: "#8c4646", bodyB: "#588c7e", bodyL: "#8c4646" };
const CherryCheesecake = { key: "CherryCheesecake", mode: "light", navB: "#4c1b1b", navF: "#bd8d46", menuB: "#f6e497", menuF: "#fcfae1", menuL: "#bd8d46", bodyB: "#b9121b", bodyL: "#bd8d46" };
const CircusIII = { key: "CircusIII", mode: "light", navB: "#d90000", navF: "#04756f", menuB: "#ff2d00", menuF: "#ff8c00", menuL: "#04756f", bodyB: "#2e0927", bodyL: "#04756f" };
const CorporateOrangeBlue = { key: "CorporateOrangeBlue", mode: "light", navB: "#91bed4", navF: "#f26101", menuB: "#d9e8f5", menuF: "#ffffff", menuL: "#f26101", bodyB: "#304269", bodyL: "#f26101" };
const CoteAzur = { key: "CoteAzur", mode: "light", navB: "#009393", navF: "#ff3800", menuB: "#fffcc4", menuF: "#f0edbb", menuL: "#ff3800", bodyB: "#00585f", bodyL: "#ff3800" };
const Firenze = { key: "Firenze", mode: "light", navB: "#fff0a5", navF: "#8e2800", menuB: "#ffb03b", menuF: "#b64926", menuL: "#8e2800", bodyB: "#468966", bodyL: "#8e2800" };
const FlatUI = { key: "FlatUI", mode: "light", navB: "#e74c3c", navF: "#2980b9", menuB: "#ecf0f1", menuF: "#3498db", menuL: "#2980b9", bodyB: "#2c3e50", bodyL: "#2980b9" };
const FriendsAndFoes = { key: "FriendsAndFoes", mode: "light", navB: "#01a2a6", navF: "#ffffa6", menuB: "#29d9c2", menuF: "#bdf271", menuL: "#ffffa6", bodyB: "#2f2933", bodyL: "#ffffa6" };
const GardenSwimmingPool = { key: "GardenSwimmingPool", mode: "light", navB: "#b5e655", navF: "#7fc6bc", menuB: "#edf7f2", menuF: "#4bb5c1", menuL: "#7fc6bc", bodyB: "#96ca2d", bodyL: "#7fc6bc" };
const Gettysburg = { key: "Gettysburg", mode: "light", navB: "#343642", navF: "#348899", menuB: "#979c9c", menuF: "#f2ebc7", menuL: "#348899", bodyB: "#962d3e", bodyL: "#348899" };
const GrannySmithApple = { key: "GrannySmithApple", mode: "light", navB: "#cde855", navF: "#493f0b", menuB: "#f5f6d4", menuF: "#a7c520", menuL: "#493f0b", bodyB: "#85db18", bodyL: "#493f0b" };
const Harbor = { key: "Harbor", mode: "light", navB: "#9fbcbf", navF: "#59d8e6", menuB: "#647678", menuF: "#2f3738", menuL: "#59d8e6", bodyB: "#d5fbff", bodyL: "#59d8e6" };
const HerbsAndSpice = { key: "HerbsAndSpice", mode: "light", navB: "#d1570d", navF: "#a9cc66", menuB: "#fde792", menuF: "#477725", menuL: "#a9cc66", bodyB: "#5a1f00", bodyL: "#a9cc66" };
const HoneyPot = { key: "HoneyPot", mode: "light", navB: "#fffad5", navF: "#bd4932", menuB: "#ffd34e", menuF: "#db9e36", menuL: "#bd4932", bodyB: "#105b63", bodyL: "#bd4932" };
const JapaneseGarden = { key: "JapaneseGarden", mode: "light", navB: "#5c832f", navF: "#363942", menuB: "#284907", menuF: "#382513", menuL: "#363942", bodyB: "#d8caa8", bodyL: "#363942" };
const Kayak = {key: "Kayak", mode: "dark", bodyB: "#36362c", bodyL: "#825534", navB: "#5d917d", navF: "#825534", menuB: "#a8ad80", menuF: "#e6d4a7", menuL: "#825534"};
const KeepTheChange = { key: "KeepTheChange", mode: "light", navB: "#d9042b", navF: "#011c26", menuB: "#f4cb89", menuF: "#588c8c", menuL: "#011c26", bodyB: "#6b0c22", bodyL: "#011c26" };
const KnotJustNautical = { key: "KnotJustNautical", mode: "light", navB: "#fc4349", navF: "#ffffff", menuB: "#d7dadb", menuF: "#6dbcdb", menuL: "#ffffff", bodyB: "#2c3e50", bodyL: "#ffffff" };
const LifeIsBeautiful = { key: "LifeIsBeautiful", mode: "light", navB: "#047878", navF: "#c22121", menuB: "#ffb733", menuF: "#f57336", menuL: "#c22121", bodyB: "#801637", bodyL: "#c22121" };
const Lollapalooza = { key: "Lollapalooza", mode: "light", navB: "#013440", navF: "#efe7be", menuB: "#ab1a25", menuF: "#d97925", menuL: "#efe7be", bodyB: "#002635", bodyL: "#efe7be" };
const MountainsOfBurma = {
  key: "MountainsOfBurma",
  mode: "dark",
  bodyB: "#354242",
  bodyL: "#7d9100",
  navB: "#acebae",
  navF: "#7d9100",
  menuB: "#ffff9d",
  menuF: "#c9de55",
  menuL: "#7d9100"
};
const NeutralBlue = { key: "NeutralBlue", mode: "light", navB: "#d1dbbd", navF: "#193441", menuB: "#91aa9d", menuF: "#3e606f", menuL: "#193441", bodyB: "#fcfff5", bodyL: "#193441" };
const OceanSunset = { key: "OceanSunset", mode: "light", navB: "#9c9b7a", navF: "#f54f29", menuB: "#ffd393", menuF: "#ff974f", menuL: "#f54f29", bodyB: "#405952", bodyL: "#f54f29" };
const Optimist = { key: "Optimist", mode: "light", navB: "#3e423a", navF: "#f4f7d9", menuB: "#417378", menuF: "#a4cfbe", menuL: "#f4f7d9", bodyB: "#6c6e58", bodyL: "#f4f7d9" };
const PearLemonFizz = { key: "PearLemonFizz", mode: "light", navB: "#cafcd8", navF: "#588f27", menuB: "#f7e967", menuF: "#a9cf54", menuL: "#588f27", bodyB: "#04bfbf", bodyL: "#588f27" };
const Phaedra = { key: "Phaedra", mode: "light", navB: "#ffff9d", navF: "#00a388", menuB: "#beeb9f", menuF: "#79bd8f", menuL: "#00a388", bodyB: "#ff6138", bodyL: "#00a388" };
const PomegranateExplosion = { key: "PomegranateExplosion", mode: "light", navB: "#f2e1ac", navF: "#cd2c24", menuB: "#f2836b", menuF: "#f2594b", menuL: "#cd2c24", bodyB: "#63a69f", bodyL: "#cd2c24" };
const QuietCry = { key: "QuietCry", mode: "light", navB: "#31353d", navF: "#eeeff7", menuB: "#445878", menuF: "#92cdcf", menuL: "#eeeff7", bodyB: "#1c1d21", bodyL: "#eeeff7" };
const SalmonOnIce = { key: "SalmonOnIce", mode: "light", navB: "#2185c5", navF: "#ff7f66", menuB: "#7ecefd", menuF: "#fff6e5", menuL: "#ff7f66", bodyB: "#3e454c", bodyL: "#ff7f66" };
const SandyStoneBeach = { key: "SandyStoneBeach", mode: "light", navB: "#a7a37e", navF: "#002f2f", menuB: "#efecca", menuF: "#046380", menuL: "#002f2f", bodyB: "#e6e2af", bodyL: "#002f2f" };
const SeaWolf = { key: "SeaWolf", mode: "light", navB: "#d9cb9e", navF: "#1e1e20", menuB: "#374140", menuF: "#2a2c2b", menuL: "#1e1e20", bodyB: "#dc3522", bodyL: "#1e1e20" };
const SunshineOverGlacier = { key: "SunshineOverGlacier", mode: "light", navB: "#2c858d", navF: "#ffffcb", menuB: "#74ceb7", menuF: "#c9ffd5", menuL: "#ffffcb", bodyB: "#004056", bodyL: "#ffffcb" };
const TimesChanging = {
  key: "TimesChanging",
  mode: "dark",
  bodyB: "#332532",
  bodyL: "#a49a87",
  navB: "#644d52",
  navF: "#a49a87",
  menuB: "#f77a52",
  menuF: "#ff974f",
  menuL: "#a49a87"
};
const Unlike = { key: "Unlike", mode: "light", navB: "#35203b", navF: "#ed8c2b", menuB: "#911146", menuF: "#cf4a30", menuL: "#ed8c2b", bodyB: "#88a825", bodyL: "#ed8c2b" };
const VentanaAzul = { key: "VentanaAzul", mode: "light", navB: "#f2385a", navF: "#36b1bf", menuB: "#e9f1df", menuF: "#4ad9d9", menuL: "#36b1bf", bodyB: "#f5a503", bodyL: "#36b1bf" };
const ViewOverTheTown = { key: "ViewOverTheTown", mode: "light", navB: "#b39c85", navF: "#1d181f", menuB: "#306e73", menuF: "#3b424d", menuL: "#1d181f", bodyB: "#ff5335", bodyL: "#1d181f" };
const VintageCard = { key: "VintageCard", mode: "light", navB: "#8cbeb2", navF: "#f06060", menuB: "#f2ebbf", menuF: "#f3b562", menuL: "#f06060", bodyB: "#5c4b51", bodyL: "#f06060" };
const VintageRalphLauren = { key: "VintageRalphLauren", mode: "light", navB: "#2f343b", navF: "#c77966", menuB: "#7e827a", menuF: "#e3cda4", menuL: "#c77966", bodyB: "#703030", bodyL: "#c77966" };
const VitaminC = { key: "VitaminC", mode: "light", navB: "#1f8a70", navF: "#fd7400", menuB: "#bedb39", menuF: "#ffe11a", menuL: "#fd7400", bodyB: "#004358", bodyL: "#fd7400" };
const WomanInPurpleDress = { key: "WomanInPurpleDress", mode: "light", navB: "#e6b098", navF: "#31152b", menuB: "#cc4452", menuF: "#723147", menuL: "#31152b", bodyB: "#f9e4ad", bodyL: "#31152b" };
const XDustyPetrol = { key: "XDustyPetrol", mode: "light", navB: "#5b7876", navF: "#412a22", menuB: "#8f9e8b", menuF: "#f2e6b6", menuL: "#412a22", bodyB: "#292929", bodyL: "#412a22" };
const ZenAndTea = { key: "ZenAndTea", mode: "light", navB: "#95ab63", navF: "#f6ffe0", menuB: "#bdd684", menuF: "#e2f0d6", menuL: "#f6ffe0", bodyB: "#10222b", bodyL: "#f6ffe0" };

export const allThemes: ThemeColors[] = [
  Default,
  Darkfault,
  AfternoonChai,
  AspirinC,
  BeachTime,
  BirdfolioBlues,
  BloggyGradientBlues,
  BlueSky,
  Bogart,
  CS04,
  Campfire,
  CherryCheesecake,
  CircusIII,
  CorporateOrangeBlue,
  CoteAzur,
  Firenze,
  FlatUI,
  FriendsAndFoes,
  GardenSwimmingPool,
  Gettysburg,
  GrannySmithApple,
  Harbor,
  HerbsAndSpice,
  HoneyPot,
  JapaneseGarden,
  Kayak,
  KeepTheChange,
  KnotJustNautical,
  LifeIsBeautiful,
  Lollapalooza,
  MountainsOfBurma,
  NeutralBlue,
  OceanSunset,
  Optimist,
  PearLemonFizz,
  Phaedra,
  PomegranateExplosion,
  QuietCry,
  SalmonOnIce,
  SandyStoneBeach,
  SeaWolf,
  SunshineOverGlacier,
  TimesChanging,
  Unlike,
  VentanaAzul,
  ViewOverTheTown,
  VintageCard,
  VintageRalphLauren,
  VitaminC,
  WomanInPurpleDress,
  XDustyPetrol,
  ZenAndTea
];
